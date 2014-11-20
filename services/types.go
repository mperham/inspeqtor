package services

import (
	"errors"
	"fmt"

	"github.com/mperham/inspeqtor/util"
)

type ProcessStatus struct {
	Pid int
	Status
}

func WithStatus(pid int, status Status) *ProcessStatus {
	return &ProcessStatus{Pid: pid, Status: status}
}

func NewStatus() *ProcessStatus {
	return &ProcessStatus{0, Unknown}
}

func (s *ProcessStatus) String() string {
	return fmt.Sprintf("%s/%d", s.Status, s.Pid)
}

type Status string

func (s Status) String() string {
	return string(s)
}

const (
	Unknown  Status = "Unknown"
	Down            = "Down"
	Starting        = "Starting"
	Up              = "Up"
)

type ServiceError struct {
	Init string
	Name string
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

	// Find the process info for a given service name. Method MUST either return
	// a non-nil ProcessStatus or a ServiceError.
	LookupService(name string) (*ProcessStatus, error)

	// Restart the process associated with the given service name. All errors
	// returned must be of type ServiceError.
	Restart(name string) error

	// Reload the process associated with the service name. Typically this will
	// send a HUP signal to the service. Not all system managers will handle
	// "reload", in which case they'll throw a ServiceError.
	Reload(name string) error
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
		func() (InitSystem, error) {
			return detectSystemd("/etc/systemd")
		},
		func() (InitSystem, error) {
			return detectInitd("/")
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

func (m *MockInitSystem) Reload(name string) error {
	m.Actions = append(m.Actions, "reload "+name)
	return nil
}

func (m *MockInitSystem) LookupService(name string) (*ProcessStatus, error) {
	if m.CurrentStatus != nil {
		return m.CurrentStatus, nil
	} else {
		return &ProcessStatus{123, Up}, nil
	}
}
