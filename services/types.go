package services

import (
	"inspeqtor/util"
)

type ProcessId int
type Status uint8

const (
	Unknown Status = iota
	Down
	Starting
	Up
	Stopping
)

//
// Your init system(s) manages services.  We use
// the init system to:
// 1. find the associated PID
// 2. restart the service
//
type InitSystem interface {
	// Name of the init system: "upstart", "runit", etc.
	Name() string

	// Look up PID for the given service name, returns
	// positive integer if successful, -1 if the service
	// name was not found or error if there was an
	// unexpected failure.
	LookupService(name string) (ProcessId, Status, error)

	Restart(name string) error
}

var (
	SupportedInits = map[string]func() (InitSystem, error){
		"launchctl": func() (InitSystem, error) {
			return detectLaunchctl("/")
		},
		"upstart": func() (InitSystem, error) {
			return detectUpstart("/etc/init")
		},
		"runit": func() (InitSystem, error) {
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
		} else {
			inits = append(inits, sm)
		}
	}

	if len(inits) == 0 {
		util.Warn("No init system detected.  Inspeqtor cannot control any services!")
	}

	return inits
}

func MockInit() InitSystem {
	return &MockInitSystem{}
}

type MockInitSystem struct {
	Actions []string
}

func (m *MockInitSystem) Name() string { return "mock" }

func (m *MockInitSystem) Restart(name string) error {
	m.Actions = append(m.Actions, "restart "+name)
	return nil
}

func (m *MockInitSystem) LookupService(name string) (ProcessId, Status, error) {
	return ProcessId(123), Up, nil
}
