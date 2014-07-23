package core

/*
 Core Inspeqtor types, interfaces, etc.
*/
const (
	CYCLE_TIME = 15
	ONE_HOUR   = 3600
	SLOTS      = ONE_HOUR / CYCLE_TIME
)

//
// Your init system(s) manages services.  We use
// the init system to:
// 1. find the associated PID
// 2. start/stop/restart the service
//
type InitSystem interface {
	// Name of the init system: "upstart", "runit", etc.
	Name() string

	// Look up PID for the given service name, returns
	// positive integer if successful, -1 if the service
	// name was not found or error if there was an
	// unexpected failure.
	FindServicePID(name string) (int32, error)

	Start(name string)
	Stop(name string)
	Status(name string)
}

/*
  A service is a logical named entity we wish to monitor, "mysql".
  A logical service maps onto a physical process with a PID.
  PID 0 means the process did not exist during that cycle.
*/
type Service struct {
	Name  string
	PID   []int32
	Rules []*Rule
}

type Host struct {
	Name  string
	Rules []*Rule
}

type Operator uint8

const (
	LT Operator = iota
	GT
)

type RuleStatus uint8

const (
	Unknown RuleStatus = iota
	Ok
	Failed
)

type Rule struct {
	Metric     string
	Op         Operator
	Threshold  uint64
	CycleCount uint8
	Status     RuleStatus
	Actions    []*Action
}

type Alert struct {
	*Service
	*Rule
}

type Action interface {
	Name() string
	Setup(map[string]string) error
	Trigger(alert *Alert) error
}
