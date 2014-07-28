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
	ServiceManagers []services.InitSystem
	Host            *Host
	Services        []*Service
	GlobalConfig    *ConfigFile
}

func New(dir string) (*Inspeqtor, error) {
	i := &Inspeqtor{RootDir: dir}
	var sm services.InitSystem

	sm, err := services.DetectLaunchctl("/")
	if err != nil {
		util.Warn("Couldn't detect launchctl: " + err.Error())
	} else {
		i.ServiceManagers = append(i.ServiceManagers, sm)
	}

	sm, err = services.DetectUpstart("/etc/init")
	if err != nil {
		util.Warn("Couldn't detect upstart: " + err.Error())
	} else {
		i.ServiceManagers = append(i.ServiceManagers, sm)
	}

	sm, err = services.DetectRunit("/")
	if err != nil {
		util.Warn("Couldn't detect runit: " + err.Error())
	} else {
		i.ServiceManagers = append(i.ServiceManagers, sm)
	}

	return i, nil
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

	if firstTime {
		util.Info("Resolving services")
		i.resolveServices()
	}
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

/*
Resolve each defined service to its managing init system.
*/
func (i *Inspeqtor) resolveServices() {
	for _, svc := range i.Services {
		for _, sm := range i.ServiceManagers {
			pid, status, err := sm.LookupService(svc.Name)
			if err != nil {
				util.Warn(err.Error())
				return
			}
			if pid == -1 {
				util.Debug(sm.Name() + " doesn't have " + svc.Name)
				return
			}
			util.Info("Found " + sm.Name() + "/" + svc.Name + " with PID " + string(pid))
			svc.PID = pid
			svc.Status = status
			svc.Manager = sm
			break
		}
	}
}

/*
Called for each service each cycle, in parallel.  This
method must be thread-safe.  Since this method executes
in a goroutine, errors must be handled/logged here and
not just returned.

Each cycle we need to:
1. verify service is Up and running.
2. capture process metrics
3. run rules
4. trigger any necessary actions
*/
func (i *Inspeqtor) collectService(svc *Service) {
	if svc.Manager == nil {
		// Couldn't resolve it when we started up so we can't collect it.
		return
	}
	if svc.Status == services.Starting {
		pid, status, err := svc.Manager.LookupService(svc.Name)
		if err != nil {
			util.Warn(err.Error())
		}
		svc.PID = pid
		svc.Status = status
	}
	if svc.Status == services.Down {
		return
	}
	if svc.Status == services.Up {
		err := i.captureProcess(svc)
		if err != nil {
			util.Warn("Error capturing process " + string(svc.PID) + ", marking as Down: " + err.Error())
			svc.Status = services.Down
		}
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
