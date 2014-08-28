package inspeqtor

import (
	"errors"
	"inspeqtor/metrics"
	"inspeqtor/services"
	"inspeqtor/util"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	VERSION = "1.0.0"
)

type Inspeqtor struct {
	RootDir    string
	SocketPath string
	StartedAt  time.Time

	ServiceManagers []services.InitSystem
	Host            *Host
	Services        []*Service
	GlobalConfig    *ConfigFile
	Socket          net.Listener
	SilenceUntil    time.Time
}

func New(dir string, socketpath string) (*Inspeqtor, error) {
	i := &Inspeqtor{RootDir: dir,
		SocketPath:   socketpath,
		StartedAt:    time.Now(),
		SilenceUntil: time.Now(),
		Host:         &Host{&Entity{name: "localhost"}},
		GlobalConfig: &ConfigFile{Defaults, map[string]*AlertRoute{}}}
	return i, nil
}

var (
	Term os.Signal = syscall.SIGTERM

	SignalHandlers = map[os.Signal]func(*Inspeqtor){
		Term:         exit,
		os.Interrupt: exit,
	}
	Name string = "Inspeqtor"
)

func (i *Inspeqtor) Start() {
	err := i.openSocket(i.SocketPath)
	if err != nil {
		util.Warn("Could not create Unix socket: %s", err.Error())
		exit(i)
	}

	go func() {
		for {
			i.acceptCommand()
		}
	}()

	go i.runLoop()

	// This method never returns
	handleSignals(i)
}

func (i *Inspeqtor) Parse() error {
	i.ServiceManagers = services.Detect()

	config, err := ParseGlobal(i.RootDir)
	if err != nil {
		return err
	}
	util.DebugDebug("Global config: %+v", config)
	i.GlobalConfig = config

	host, services, err := ParseInq(i.GlobalConfig, i.RootDir+"/conf.d")
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

func (i *Inspeqtor) openSocket(path string) error {
	if i.Socket != nil {
		return errors.New("Socket is already open!")
	}

	socket, err := net.Listen("unix", path)
	if err != nil {
		return err
	}
	i.Socket = socket
	return nil
}

func HandleSignal(sig os.Signal, handler func(*Inspeqtor)) {
	SignalHandlers[sig] = handler
}

func handleSignals(i *Inspeqtor) {
	signals := make(chan os.Signal)
	for k, _ := range SignalHandlers {
		signal.Notify(signals, k)
	}

	for {
		sig := <-signals
		util.Debug("Received signal %d", sig)
		funk := SignalHandlers[sig]
		funk(i)
	}
}

func exit(i *Inspeqtor) {
	util.Info(Name + " exiting")
	if i.Socket != nil {
		err := i.Socket.Close()
		if err != nil {
			util.Warn(err.Error())
		}
	}
	os.Exit(0)
}

// this method never returns.
//
// since we can't test this method in an automated fashion, it should
// contain as little logic as possible.
func (i *Inspeqtor) runLoop() {
	util.DebugDebug("Resolving services")
	for _, svc := range i.Services {
		i.resolveService(svc)
	}

	i.scanSystem()

	for {
		select {
		case <-time.After(time.Duration(i.GlobalConfig.Top.CycleTime) * time.Second):
			i.scanSystem()
		}
	}
}

func (i *Inspeqtor) silenced() bool {
	return time.Now().Before(i.SilenceUntil)
}

func (i *Inspeqtor) scanSystem() {
	i.collect()
	i.verify(i.Host, i.Services)
}

func (i *Inspeqtor) collect() {
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
}

func (i *Inspeqtor) verify(host *Host, services []*Service) []*Event {
	eventsToTrigger := []*Event{}
	checker := func(rule *Rule) {
		if i.silenced() {
			// We are silenced until some point in the future.
			// We don't want to check rules (as a deploy might use
			// enough resources to trip a threshold) or trigger
			// any actions
			rule.Reset()
		} else {
			event := rule.Check()
			if event != nil {
				eventsToTrigger = append(eventsToTrigger, event)
			}
		}
	}

	if host != nil {
		for _, rule := range host.Rules() {
			checker(rule)
		}
	}

	for _, svc := range services {
		for _, rule := range svc.Rules() {
			checker(rule)
		}
	}

	/*
	   We now have a set of rules which have failed.  We need to trigger
	   the actions for each.
	*/
	i.triggerActions(eventsToTrigger)
	return eventsToTrigger
}

func (i *Inspeqtor) triggerActions(alerts []*Event) error {
	for _, alert := range alerts {
		for _, action := range alert.Rule.actions {
			err := action.Trigger(alert)
			if err != nil {
				util.Warn("Error triggering action: %s", err.Error())
			}
		}
	}
	return nil
}

func (i *Inspeqtor) collectHost(completeCallback func()) {
	defer completeCallback()
	err := metrics.CollectHostMetrics(i.Host.Metrics(), "/proc")
	if err != nil {
		util.Warn("Error collecting host metrics: %s", err.Error())
	}
}

/*
Resolve each defined service to its managing init system.  Called only
at startup, this is what maps services to init and fires ProcessDoesNotExist events.
*/
func (i *Inspeqtor) resolveService(svc *Service) {
	for _, sm := range i.ServiceManagers {
		// TODO There's a bizarre race condition here. Figure out
		// why this is necessary.  We shouldn't be multi-threaded yet.
		if sm == nil {
			continue
		}

		ps, err := sm.LookupService(svc.Name())
		if err != nil {
			serr := err.(*services.ServiceError)
			if serr.Err == services.ErrServiceNotFound {
				util.Debug(sm.Name() + " doesn't have " + svc.Name())
				continue
			}
			util.Warn(err.Error())
			return
		}
		util.Info("Found %s/%s with status %s", sm.Name(), svc.Name(), ps)
		svc.Manager = sm
		svc.Transition(ps, func(et EventType) {
			i.handleProcessEvent(et, svc)
		})
		break
	}
	if svc.Manager == nil {
		util.Warn("Could not find service %s, did you misspell it?", svc.Name())
	}
}

func (i *Inspeqtor) handleProcessEvent(etype EventType, svc *Service) {
	if i.silenced() {
		util.Debug("SILENCED %s %s", etype, svc.Name())
		return
	}

	util.Warn("%s %s", etype, svc.Name())

	evt := Event{etype, svc, nil}
	err := svc.EventHandler.Trigger(&evt)
	if err != nil {
		util.Warn("%s", err)
	}
}

func (i *Inspeqtor) handleRuleEvent(etype EventType, check Checkable, rule *Rule) {
	if i.silenced() {
		util.Debug("SILENCED %s %s", etype, check.Name())
		return
	}

	util.Warn("%s %s", etype, check.Name())

	evt := Event{etype, check, rule}
	for _, action := range rule.actions {
		err := action.Trigger(&evt)
		if err != nil {
			util.Warn("%s", err)
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
	if svc.Process.Status != services.Up {
		status, err := svc.Manager.LookupService(svc.Name())
		if err != nil {
			util.Warn("%s", err)
		} else {
			svc.Transition(status, func(et EventType) {
				i.handleProcessEvent(et, svc)
			})
		}
	}

	if svc.Process.Status == services.Up {
		merr := metrics.CaptureProcess(svc.Metrics(), "/proc", svc.Process.Pid)
		if merr != nil {
			err := syscall.Kill(svc.Process.Pid, syscall.Signal(0))
			if err != nil {
				util.Info("Service %s with process %d does not exist: %s", svc.Name(), svc.Process.Pid, err)
				svc.Transition(&services.ProcessStatus{0, services.Down}, func(et EventType) {
					i.handleProcessEvent(et, svc)
				})
			} else {
				util.Warn("Error capturing metrics for process %d: %s", svc.Process.Pid, merr)
			}
		}
	}
	return
}

func (i *Inspeqtor) TestNotifications() {
	for _, route := range i.GlobalConfig.AlertRoutes {
		nm := route.Name
		if nm == "" {
			nm = "default"
		}
		util.Info("Creating notification for %s/%s", route.Channel, nm)
		notifier, err := Actions["alert"](i.Host, route)
		if err != nil {
			util.Warn("Error creating %s/%s route: %s", route.Channel, nm, err.Error())
			continue
		}
		util.Info("Triggering notification for %s/%s", route.Channel, nm)
		err = notifier.Trigger(&Event{RuleFailed, i.Host, i.Host.Rules()[0]})
		if err != nil {
			util.Warn("Error firing %s/%s route: %s", route.Channel, nm, err.Error())
		}
	}
}
