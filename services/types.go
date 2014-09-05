package services

import (
	"errors"
	"fmt"
	"inspeqtor/util"
)

type ProcessStatus struct {
	Pid int
	Status
}

func NewStatus() *ProcessStatus {
	return &ProcessStatus{}
}

func (s *ProcessStatus) String() string {
	return fmt.Sprintf("%s/%d", s.Status, s.Pid)
}

type Status uint8

func (s Status) String() string {
	switch s {
	case Unknown:
		return "Unknown"
	case Down:
		return "Down"
	case Starting:
		return "Starting"
	case Up:
		return "Up"
	case Stopping:
		return "Stopping"
	default:
		return fmt.Sprintf("Oops: %d", s)
	}
}

const (
	Unknown Status = iota
	Down
	Starting
	Up
	Stopping
)

type ServiceError struct {
	Name string
	Init string
	Err  error
}

func (e *ServiceError) Error() string { return e.Init + "/" + e.Name + ": " + e.Err.Error() }

var ErrServiceNotFound = errors.New("No such service")

// Your init system(s) manages services.  We use
// the init system to:
// 1. find the associated process
// 2. restart the service
type InitSystem interface {
	// Name of the init system: "upstart", "runit", etc.
	Name() string

	// Find the process info for a given service name.  All errors
	// returned must be of type ServiceError.
	LookupService(name string) (*ProcessStatus, error)

	// Restart the process associated with the given service name.  All errors
	// returned must be of type ServiceError.
	Restart(name string) error
}

var (
	SupportedInits = []func() (InitSystem, error){
		func() (InitSystem, error) {
			return detectLaunchd("/")
		},
		func() (InitSystem, error) {
			return detectUpstart("/etc/init")
		},
		func() (InitSystem, error) {
			return detectRunit("/")
		},
	}
)

func Detect() []InitSystem {
	inits := make([]InitSystem, 0)

	for name, funk := range SupportedInits {
		sm, err := funk()
		if err != nil {
			util.Warn("Couldn't detect %s: %s", name, err.Error())
			continue
		}

		if sm != nil {
			inits = append(inits, sm)
		}
	}

	if len(inits) == 0 {
		util.Warn("No init system detected.  Inspeqtor cannot control any services!")
	}

	return inits
}

func MockInit() *MockInitSystem {
	return &MockInitSystem{}
}

type MockInitSystem struct {
	Actions       []string
	CurrentStatus *ProcessStatus
}

func (m *MockInitSystem) Name() string { return "mock" }

func (m *MockInitSystem) Restart(name string) error {
	m.Actions = append(m.Actions, "restart "+name)
	return nil
}

func (m *MockInitSystem) LookupService(name string) (*ProcessStatus, error) {
	if m.CurrentStatus != nil {
		return m.CurrentStatus, nil
	} else {
		return &ProcessStatus{123, Up}, nil
	}
}
