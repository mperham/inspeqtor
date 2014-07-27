package inspeqtor

import (
	"fmt"
	"inspeqtor/metrics"
	"inspeqtor/services"
	"inspeqtor/util"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	VERSION = "1.0.0"
)

type Inspeqtor struct {
	RootDir         string
	ServiceManagers []*services.InitSystem
	Host            *Host
	Services        []*Service
	GlobalConfig    *ConfigFile
}

func New(dir string) (*Inspeqtor, error) {
	return &Inspeqtor{RootDir: dir}, nil
}

var (
	Quit os.Signal = syscall.SIGQUIT
)

func (i *Inspeqtor) Start() {
	go i.runLoop()

	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt)
	signal.Notify(signals, Quit)
	<-signals

	util.Debug("Inspeqtor shutting down...")
	os.Exit(0)
}

func (i *Inspeqtor) Parse() error {
	config, err := ParseGlobal(i.RootDir)
	if err != nil {
		return err
	}
	i.GlobalConfig = config

	host, services, err := ParseInq(i.RootDir + "/conf.d")
	if err != nil {
		return err
	}
	i.Host = host
	i.Services = services

	util.DebugDebug("Config: %+v", config)
	util.DebugDebug("Host: %+v", host)
	util.DebugDebug("Services: %+v", services)
	return nil
}

func (i *Inspeqtor) runLoop() {
	i.scanSystem(true)
	for {
		select {
		case <-time.After(time.Duration(i.GlobalConfig.Top.CycleTime) * time.Second):
			i.scanSystem(false)
		}
	}
}

func (i *Inspeqtor) scanSystem(firstTime bool) {
	util.DebugDebug("Scanning...")
	metrics, err := metrics.CollectHostMetrics("/proc")
	if err != nil {
		util.Warn("%v", err)
	} else {
		util.DebugDebug("%+v", metrics)
	}

	for _, svc := range i.Services {
		go i.collectService(svc)
	}
}

func (i *Inspeqtor) collectService(svc *Service) {
	if svc.Manager == nil {
		for _, sm := range i.ServiceManagers {
			pid, status, err := (*sm).LookupService(svc.Name)
			if err != nil {
				util.Warn(err.Error())
				return
			}
			if pid == -1 {
				util.Debug((*sm).Name() + " doesn't have " + svc.Name)
				return
			}
			svc.PID = pid
			svc.Status = status
			svc.Manager = sm
			break
		}
	}
	if svc.Manager == nil {
		util.Warn("Could not find service for " + svc.Name)
		return
	}
	if svc.Status == services.Starting {
		status := (*svc.Manager).Status(svc.Name)
		svc.Status = status
	}
	if svc.Status == services.Down {
		util.Info("%s is Down, asking %s to start it", svc.Name, (*svc.Manager).Name())
		(*svc.Manager).Start(svc.Name)
		svc.Status = services.Starting
	}
	if svc.Status == services.Up {
		i.captureProcess(svc)
	}
}

func (i *Inspeqtor) captureProcess(svc *Service) error {
	insist(svc.PID > 0 && svc.Status == services.Up,
		fmt.Sprintf("%+v should be Up with valid PID\n", svc))

	m, err := metrics.CaptureProcess("/proc", int32(svc.PID))
	if err != nil {
		return err
	}
	svc.Metrics.Add(m)
	return nil
}

// assert is taken by testing helpers.
func insist(expr bool, msg string) {
	if !expr {
		panic(msg)
	}
}
