package inspeqtor

import (
	"errors"
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
		return Singleton.GlobalConfig.Top.CycleTime
	} else {
		return 15
	}
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

func (s *Service) SetMetrics(newStore metrics.Store) {
	s.metrics = newStore
}

/*
 Host is the local machine.
*/
type Host struct {
	*Entity
}

func (h *Host) Resolve(_ []services.InitSystem) error {
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
	Resolve([]services.InitSystem) error
	Rules() []*Rule
	Verify() []*Event
	Collect(bool, func(Checkable))
}

// A Service is Restartable, Host is not.
type Restartable interface {
	Restart() error
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
			svc.Transition(status, func(et EventType) {
				if !silenced {
					svc.EventHandler.Trigger(&Event{et, svc, nil})
				}
			})
		}
	}

	if svc.Process.Status == services.Up {
		merr := svc.Metrics().Collect(svc.Process.Pid)
		if merr != nil {
			err := syscall.Kill(svc.Process.Pid, syscall.Signal(0))
			if err != nil {
				util.Info("Service %s with process %d does not exist: %s", svc.Name(), svc.Process.Pid, err)
				svc.Transition(&services.ProcessStatus{0, services.Down}, func(et EventType) {
					if !silenced {
						svc.EventHandler.Trigger(&Event{et, svc, nil})
					}
				})
			} else {
				util.Warn("Error capturing metrics for process %d: %s", svc.Process.Pid, merr)
			}
		}
	}
}

func (s *Host) Verify() []*Event {
	events := []*Event{}
	for _, r := range s.Rules() {
		evt := r.Check(s.CycleTime())
		if evt != nil {
			events = append(events, evt)
			for _, a := range r.Actions {
				a.Trigger(evt)
			}
		}
	}
	return events
}

func (s *Service) Verify() []*Event {
	events := []*Event{}

	if s.Process.Status != services.Up {
		// we probably shouldn't verify anything that isn't actually Up
		util.Debug("%s is %s, skipping...", s.Name(), s.Process.Status)
		return events
	}

	for _, r := range s.Rules() {
		evt := r.Check(s.CycleTime())
		if evt != nil {
			events = append(events, evt)
			for _, a := range r.Actions {
				a.Trigger(evt)
			}
		}
	}
	return events
}

func (s *Service) Restart() error {
	s.Process.Pid = 0
	s.Process.Status = services.Starting
	go func() {
		util.Debug("Restarting %s", s.Name())
		err := s.Manager.Restart(s.Name())
		if err != nil {
			util.Warn(err.Error())
		} else {
			util.DebugDebug("Restarted %s", s.Name())
		}
	}()
	return nil
}

/*
  Resolve each defined service to its managing init system.  Called only
  at startup, this is what maps services to init and fires ProcessDoesNotExist events.
*/
func (svc *Service) Resolve(mgrs []services.InitSystem) error {
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
		svc.Transition(ps, func(et EventType) {
			svc.EventHandler.Trigger(&Event{et, svc, nil})
		})
		break
	}
	if svc.Manager == nil {
		return errors.New(fmt.Sprintf("Could not find service %s, did you misspell it?", svc.Name()))
	}
	return nil
}

func (s *Service) Transition(ps *services.ProcessStatus, emitter func(EventType)) {
	oldst := s.Process.Status
	s.Process = ps

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

func (s *Service) String() string {
	return fmt.Sprintf("%s [%s]", s.Name(), s.Process)
}
