package inspeqtor

import (
	"inspeqtor/metrics"
	"inspeqtor/services"
	"inspeqtor/util"
	"os"
	"os/signal"
	"strconv"
	"sync"
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
	i.ServiceManagers = services.Detect()
	return i, nil
}

var (
	Quit           os.Signal = syscall.SIGQUIT
	SignalHandlers           = map[os.Signal]func(){
		Quit:         exit,
		os.Interrupt: exit,
	}
	Name string = "Inspeqtor"
)

func (i *Inspeqtor) Start() {
	go i.runLoop()

	// This method never returns
	handleSignals()
}

func HandleSignal(sig os.Signal, handler func()) {
	SignalHandlers[sig] = handler
}

func (i *Inspeqtor) Parse() error {
	config, err := ParseGlobal(i.RootDir)
	if err != nil {
		return err
	}
	util.DebugDebug("Global config: %+v", config)
	i.GlobalConfig = config

	host, services, err := ParseInq(i.RootDir + "/conf.d")
	if err != nil {
		return err
	}
	i.Host = host
	i.Services = services

	util.DebugDebug("Config: %+v", config)
	util.DebugDebug("Host: %+v", host)
	for _, val := range services {
		util.DebugDebug("Service: %+v", *val)
	}
	return nil
}

// private

func handleSignals() {
	signals := make(chan os.Signal)
	for k, _ := range SignalHandlers {
		signal.Notify(signals, k)
	}

	for {
		sig := <-signals
		util.Debug("Received signal %d", sig)
		funk := SignalHandlers[sig]
		funk()
	}
}

func exit() {
	util.Info(Name + " exiting")
	os.Exit(0)
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
	if firstTime {
		util.DebugDebug("Resolving services")
		i.resolveServices()
	}

	start := time.Now()
	var barrier sync.WaitGroup
	barrier.Add(1)
	barrier.Add(len(i.Services))

	go i.collectHost(func() {
		barrier.Done()
	})
	for _, svc := range i.Services {
		go i.collectService(svc, func(svc *Service) {
			barrier.Done()
		})
	}
	barrier.Wait()
	util.Debug("Collection complete in " + time.Now().Sub(start).String())

	rulesToAlert := []*Alert{}
	for _, rule := range i.Host.Rules {
		rule := rule.Check()
		if rule != nil {
			rulesToAlert = append(rulesToAlert, &Alert{rule})
		}
	}
	for _, svc := range i.Services {
		for _, rule := range svc.Rules {
			rule := rule.Check()
			if rule != nil {
				rulesToAlert = append(rulesToAlert, &Alert{rule})
			}
		}
	}

	/*
	   We now have a set of rules which have triggered.  We need to run
	   the alerts for each.
	*/
	i.fireAlerts(rulesToAlert)
}

func (i *Inspeqtor) fireAlerts(alerts []*Alert) error {
	//for _, alert := range alerts {
	//for _, action := range alert.Rule.actions {
	//action.Trigger(alert)
	//}
	//}
	return nil
}

func (i *Inspeqtor) collectHost(completeCallback func()) {
	defer completeCallback()
	err := metrics.CollectHostMetrics(i.Host.Metrics, "/proc")
	if err != nil {
		util.Warn("%v", err)
	}
}

/*
Resolve each defined service to its managing init system.
*/
func (i *Inspeqtor) resolveServices() {
	for _, svc := range i.Services {
		for _, sm := range i.ServiceManagers {
			pid, status, err := sm.LookupService(svc.Name())
			if err != nil {
				util.Warn(err.Error())
				return
			}
			if pid == -1 {
				util.Debug(sm.Name() + " doesn't have " + svc.Name())
				return
			}
			util.Info("Found " + sm.Name() + "/" + svc.Name() + " with PID " + strconv.Itoa(int(pid)))
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
func (i *Inspeqtor) collectService(svc *Service, completeCallback func(*Service)) {
	defer completeCallback(svc)

	if svc.Manager == nil {
		// Couldn't resolve it when we started up so we can't collect it.
		return
	}
	if svc.Status == services.Starting {
		pid, status, err := svc.Manager.LookupService(svc.Name())
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
		err := metrics.CaptureProcess(svc.Metrics, "/proc", int(svc.PID))
		if err != nil {
			util.Warn("Error capturing process " + strconv.Itoa(int(svc.PID)) + ", marking as Down: " + err.Error())
			svc.Status = services.Down
		}
	}
}
