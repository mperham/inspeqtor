package inspeqtor

import (
	"bytes"
	"fmt"
	"inspeqtor/inq"
	"inspeqtor/metrics"
	"inspeqtor/services"
	"inspeqtor/util"
	"log"
	"net/smtp"
	"os"
	"os/signal"
	"time"
)

const (
	VERSION = "1.0.0"
)

type Inspeqtor struct {
	RootDir         string
	ServiceManagers []Init
	Checks          *inq.Checks
}

type Init interface {
	// Name of the init system: "upstart", "runit", etc.
	Name() string

	// Look up PID for the given service name, returns
	// positive integer if successful, -1 if the service
	// name was not found or error if there was an
	// unexpected failure.
	FindServicePID(name string) (int32, error)
}

func New(dir string) (*Inspeqtor, error) {
	return &Inspeqtor{RootDir: dir}, nil
}

func (i *Inspeqtor) DetectManagers() error {
	serviceMapping := make(map[string]int32)

	launchctl, err := services.DetectLaunchctl("/")
	if err != nil {
		return err
	}

	if launchctl != nil {
		services := []string{"homebrew.mxcl.memcached", "bob"}
		for _, service := range services {
			pid, err := launchctl.FindServicePID(service)
			if err != nil {
				util.Debug("Couldn't find service " + service + ", skipping...")
			} else {
				serviceMapping[service] = pid
			}
		}
	}

	upstart, err := services.DetectUpstart("/etc/init")
	if err != nil {
		return err
	}

	if upstart != nil {
		services := []string{"mysql", "pass", "bob"}
		for _, service := range services {
			pid, err := upstart.FindServicePID(service)
			if err != nil {
				return err
			} else {
				serviceMapping[service] = pid
			}
		}
	}

	log.Println(serviceMapping)

	return nil
}

func (i *Inspeqtor) Start() {
	shutdown := make(chan int)
	go i.pollSystem(shutdown)

	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt)
	<-signals

	util.Debug("Inspeqtor shutting down...")
	shutdown <- 1
}

func (i *Inspeqtor) Parse() error {
	checks, err := inq.Parse(i.RootDir)
	if err != nil {
		return err
	}
	i.Checks = checks
	util.DebugDebug("Checks: %+v\n", checks)
	return nil
}

func (i *Inspeqtor) pollSystem(shutdown chan int) {
	scanSystem()
	select {
	case <-shutdown:
		util.DebugDebug("Exiting poll loop")
		return
	case <-time.After(30 * time.Second):
		scanSystem()
	}
}

func scanSystem() {
	util.DebugDebug("Scanning...")
	metrics, err := metrics.CollectHostMetrics("/proc")
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println(metrics)
	}
}

func sendEmail(data interface{}) error {
	auth := smtp.PlainAuth("", "mperham", "", "smtp.gmail.com")
	err := smtp.SendMail("smtp.gmail.com:587", auth,
		"mperham@gmail.com",
		[]string{"mperham@gmail.com"},
		bytes.NewBufferString(fmt.Sprint(data)).Bytes())
	if err != nil {
		return err
	}
	return nil
}
