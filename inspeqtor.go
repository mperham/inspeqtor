package inspeqtor

import (
	"bytes"
	"fmt"
	"inspeqtor/conf"
	"inspeqtor/darwin"
	"inspeqtor/linux"
	"log"
	"net/smtp"
	"os"
	"os/signal"
	"time"
)

var (
	VERSION = "1.0.0"
)

type Inspeqtor struct {
	RootDir         string
	ServiceManagers []Init
	Checks          *conf.Checks
}

type Init interface {
	Name() string
	FindService(name string) (int32, error)
}

func New(dir string) (*Inspeqtor, error) {
	return &Inspeqtor{RootDir: dir}, nil
}

func (i *Inspeqtor) DetectManagers() error {
	serviceMapping := make(map[string]int)

	launchctl, err := darwin.DetectLaunchctl()
	if err != nil {
		return err
	}

	if launchctl != nil {
		services := []string{"homebrew.mxcl.memcached", "bob"}
		for _, service := range services {
			name, pid, err := launchctl.FindService(service)
			if err != nil {
				log.Println("Couldn't find service " + service + ", skipping...")
			} else {
				serviceMapping[name] = pid
			}
		}
	}

	upstart, err := linux.DetectUpstart("/etc/init")
	if err != nil {
		return err
	}

	if upstart != nil {
		services := []string{"mysql", "pass", "bob"}
		for _, service := range services {
			name, pid, err := upstart.FindService(service)
			if err != nil {
				return err
			} else {
				serviceMapping[name] = pid
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

	log.Println("Inspeqtor shutting down...")
	shutdown <- 1
}

func (i *Inspeqtor) Parse() error {
	checks, err := conf.Parse(i.RootDir)
	if err != nil {
		return err
	}
	i.Checks = checks
	return nil
}

func (i *Inspeqtor) pollSystem(shutdown chan int) {
	scanSystem()
	select {
	case <-shutdown:
		log.Println("Exiting poll loop")
		return
	case <-time.After(30 * time.Second):
		scanSystem()
	}
}

func scanSystem() {
	log.Println("Scanning...")
	metrics, err := linux.CollectHostMetrics("/proc")
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
