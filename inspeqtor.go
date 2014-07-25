package inspeqtor

import (
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
		if svc.Manager == nil {
			for _, sm := range i.ServiceManagers {
				pid, status, err := (*sm).LookupService(svc.Name)
				if err != nil {
					util.Warn(err.Error())
					continue
				}
				if pid == -1 {
					util.Debug((*sm).Name() + " doesn't have " + svc.Name)
					continue
				}
				svc.PID = pid
				svc.Status = status
				svc.Manager = sm
				break
			}
		}
		if svc.Manager == nil {
			util.Warn("Could not find service for " + svc.Name)
			continue
		}
		if svc.Status == services.Down {
			(*svc.Manager).Start(svc.Name)
			svc.Status = services.Starting
		}
	}
}
