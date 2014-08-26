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
	Process *services.ProcessStatus
	Manager services.InitSystem
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
