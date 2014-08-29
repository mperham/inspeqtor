package inspeqtor

import (
	"inspeqtor/metrics"
	"inspeqtor/services"
	"inspeqtor/util"
)

// A named thing which can checked by Inspeqtor
type Entity struct {
	name       string
	rules      []*Rule
	metrics    *metrics.Storage
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

func (e *Entity) Owner() string {
	return e.Parameter("owner")
}

func (e *Entity) Metrics() *metrics.Storage {
	return e.metrics
}

func NewHost() *Host {
	return &Host{&Entity{"localhost", nil, metrics.NewHostStore(15), nil}}
}

func NewService(name string) *Service {
	return &Service{&Entity{name, nil, metrics.NewProcessStore(), nil}, nil, services.NewStatus(), nil}
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

func (s *Service) Capture(path string) error {
	return metrics.CollectProcess(s.Metrics(), path, s.Process.Pid)
}

/*
 Host is the local machine.
*/
type Host struct {
	*Entity
}

func (h *Host) Capture(path string) error {
	return metrics.CollectHost(h.Metrics(), path)
}

type Checkable interface {
	Name() string
	Owner() string
	Metrics() *metrics.Storage
	Capture(string) error
}

// A Service is Restartable, Host is not.
type Restartable interface {
	Restart() error
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

func (s *Service) Transition(ps *services.ProcessStatus, emitter func(EventType)) {
	oldst := s.Process.Status
	s.Process = ps

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
