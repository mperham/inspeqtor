package inspeqtor

import (
	"fmt"
	"syscall"

	"github.com/mperham/inspeqtor/metrics"
	"github.com/mperham/inspeqtor/services"
	"github.com/mperham/inspeqtor/util"
)

// A named thing which can checked by Inspeqtor
type Entity struct {
	name       string
	rules      []*Rule
	metrics    metrics.Store
	parameters map[string]string
}

func (e *Entity) Name() string {
	return e.name
}

func (e *Entity) Rules() []*Rule {
	return e.rules
}

func (e *Entity) Parameter(key string) string {
	return e.parameters[key]
}

func (e *Entity) Parameters() map[string]string {
	return e.parameters
}

func (e *Entity) Metrics() metrics.Store {
	return e.metrics
}

func (e *Entity) CycleTime() uint {
	if Singleton != nil {
		return Singleton.GlobalConfig.CycleTime
	}
	return 15
}

func NewHost() *Host {
	return &Host{&Entity{"localhost", nil, metrics.NewHostStore("/proc", 15), nil}}
}

func NewService(name string) *Service {
	return &Service{&Entity{name, nil, metrics.NewProcessStore("/proc", 15), nil}, nil, services.NewStatus(), nil}
}

/*
  A service is an Entity which resolves to a Process
  we can monitor.
*/
type Service struct {
	*Entity
	// Handles process events: exists, doesn't exist
	EventHandler Action
	Process      *services.ProcessStatus
	Manager      services.InitSystem
}

func (svc *Service) SetMetrics(newStore metrics.Store) {
	svc.metrics = newStore
}

/*
 Host is the local machine.
*/
type Host struct {
	*Entity
}

func (h *Host) Resolve(_ bool, _ []services.InitSystem) error {
	return nil
}

func (h *Host) Collect(silenced bool, completeCallback func(Checkable)) {
	defer completeCallback(h)
	err := h.Metrics().Collect(0)
	if err != nil {
		util.Warn("Error collecting host metrics: %s", err.Error())
	}
}

type Eventable interface {
	Name() string
	Parameter(string) string
}

type Checkable interface {
	Name() string
	Parameter(string) string
	Metrics() metrics.Store
	Resolve(bool, []services.InitSystem) error
	Rules() []*Rule
	Verify() []*Event
	Collect(bool, func(Checkable))
}

// A Service is Controllable, a Host is not.
type Controllable interface {
	Restart() error
	Reload() error
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
func (svc *Service) Collect(silenced bool, completeCallback func(Checkable)) {
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
			svc.transitionWithEventTrigger(status, silenced)
		}
	}

	if svc.Process.Status == services.Up {
		merr := svc.Metrics().Collect(svc.Process.Pid)
		if merr != nil {
			err := syscall.Kill(svc.Process.Pid, syscall.Signal(0))
			if err != nil {
				util.Info("Service %s with process %d does not exist: %s", svc.Name(), svc.Process.Pid, err)
				svc.transitionWithEventTrigger(services.WithStatus(0, services.Down), silenced)
			} else {
				util.Warn("Error capturing metrics for process %d: %s", svc.Process.Pid, merr)
			}
		}
	}
}

func (h *Host) Verify() []*Event {
	events := []*Event{}
	for _, r := range h.Rules() {
		// When running "make real", the race detector will complain
		// of a race condition here.  I believe it's harmless.
		evt := r.Check(h.CycleTime())
		if evt != nil {
			events = append(events, evt)
			for _, a := range r.Actions {
				err := a.Trigger(evt)
				if err != nil {
					util.Warn("Error firing event: %s", err.Error())
				}
			}
		}
	}
	return events
}

func (svc *Service) Verify() []*Event {
	events := []*Event{}

	if svc.Process.Status != services.Up {
		// we probably shouldn't verify anything that isn't actually Up
		util.Debug("%s is %s, skipping...", svc.Name(), svc.Process.Status)
		return events
	}

	for _, r := range svc.Rules() {
		evt := r.Check(svc.CycleTime())
		if evt != nil {
			events = append(events, evt)
			for _, a := range r.Actions {
				err := a.Trigger(evt)
				if err != nil {
					util.Warn("Error firing event: %s", err.Error())
				}
			}
		}
	}
	return events
}

func (svc *Service) Restart() error {
	svc.Process.Pid = 0
	svc.Process.Status = services.Starting
	go func() {
		util.Debug("Restarting %s", svc.Name())
		err := svc.Manager.Restart(svc.Name())
		if err != nil {
			util.Warn(err.Error())
		} else {
			util.DebugDebug("Restarted %s", svc.Name())
		}
	}()
	return nil
}

func (svc *Service) Reload() error {
	go func() {
		util.Debug("Reloading %s", svc.Name())
		err := svc.Manager.Reload(svc.Name())
		if err != nil {
			util.Warn(err.Error())
		} else {
			util.DebugDebug("Reloaded %s", svc.Name())
		}
	}()
	return nil
}

/*
	Resolve each defined service to its managing init system. Called only at
	startup, this is what maps services to init and fires ProcessDoesNotExist
	events.
*/
func (svc *Service) Resolve(silenced bool, mgrs []services.InitSystem) error {
	for _, sm := range mgrs {
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
			return err
		}

		util.Info("Found %s/%s with status %s", sm.Name(), svc.Name(), ps)
		svc.Manager = sm
		svc.transitionWithEventTrigger(ps, silenced)

		break
	}
	if svc.Manager == nil {
		return fmt.Errorf("Could not find service %s, did you misspell it?", svc.Name())
	}
	return nil
}

func (svc *Service) Transition(ps *services.ProcessStatus, emitter func(EventType)) {
	oldst := svc.Process.Status
	svc.Process = ps

	if oldst == ps.Status {
		// don't fire PDNE events every cycle
		return
	}

	switch ps.Status {
	case services.Up:
		if oldst != services.Unknown {
			// Don't need to fire the event when first starting up and
			// transitioning from Unknown to Up.
			emitter(ProcessExists)
		}
	case services.Down:
		emitter(ProcessDoesNotExist)
	default:
		// do nothing
	}
}

func (svc *Service) String() string {
	return fmt.Sprintf("%s [%s]", svc.Name(), svc.Process)
}

func (svc *Service) transitionWithEventTrigger(status *services.ProcessStatus, silenced bool) {
	if !silenced {
		svc.Transition(status, func(et EventType) {
			counters.Add("events", 1)
			err := svc.EventHandler.Trigger(&Event{et, svc, nil})
			if err != nil {
				util.Warn("Error firing event: %s", err.Error())
			}
		})
	}
}
