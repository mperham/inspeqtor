package inspeqtor

import (
	"inspeqtor/metrics"
	"inspeqtor/services"
	"inspeqtor/util"
)

// A named thing which can checked by Inspeqtor
type Entity struct {
	EntityName string
	Rules      []*Rule
	Parameters map[string]string
	Metrics    *metrics.Storage
}

func (e *Entity) Name() string {
	return e.EntityName
}

func (e *Entity) Owner() string {
	return e.Parameters["owner"]
}

func (e *Entity) MetricData() *metrics.Storage {
	return e.Metrics
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

/*
 Host is the local machine.
*/
type Host struct {
	*Entity
}

type Checkable interface {
	Name() string
	Owner() string
	MetricData() *metrics.Storage
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
