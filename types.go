package inspeqtor

import (
	"inspeqtor/metrics"
	"inspeqtor/services"
	"inspeqtor/util"
)

/*
 Core Inspeqtor types, interfaces, etc.
*/

/*
  A service is a logical named entity we wish to monitor, "mysql".
  A logical service maps onto a physical process with a PID.
  PID 0 means the process did not exist during that cycle.
*/
type Service struct {
	ServiceName string
	Process     *services.ProcessStatus
	Rules       []*Rule
	Parameters  map[string]string
	Metrics     *metrics.Storage
	Manager     services.InitSystem
}

type Host struct {
	Hostname   string
	Rules      []*Rule
	Metrics    *metrics.Storage
	Parameters map[string]string
}

func (h *Service) MetricData() *metrics.Storage {
	return h.Metrics
}

func (h *Host) MetricData() *metrics.Storage {
	return h.Metrics
}

func (h *Service) Name() string {
	return h.ServiceName
}

func (h *Host) Name() string {
	return h.Hostname
}

func (h *Service) Owner() string {
	return h.Parameters["owner"]
}

func (h *Host) Owner() string {
	return h.Parameters["owner"]
}

type Checkable interface {
	Name() string
	Owner() string
	MetricData() *metrics.Storage
}

type Restartable interface {
	Restart() error
}

func (s *Service) Restart() error {
	s.Process.Pid = 0
	s.Process.Status = services.Starting
	go func() {
		util.Debug("Restarting %s", s.ServiceName)
		err := s.Manager.Restart(s.ServiceName)
		if err != nil {
			util.Warn(err.Error())
		} else {
			util.DebugDebug("Restarted %s", s.ServiceName)
		}
	}()
	return nil
}

type Operator uint8

const (
	LT Operator = iota
	GT
)

type RuleState uint8

const (
	Ok RuleState = iota
	Triggered
	Recovered
)

type Rule struct {
	Entity           Checkable
	metricFamily     string
	metricName       string
	op               Operator
	displayThreshold string
	threshold        int64
	currentValue     int64
	cycleCount       int
	trippedCount     int
	state            RuleState
	actions          []Action
}

type Action interface {
	Trigger(event *Event) error
}

/*
 There are several different types of Events:

 * Process disappeared (or did not exist when we started up)
 * Process appeared
 * Rule triggered based on metric check
 * Metric has recovered
 * Process is restarting due to rule trigger
 * Process has restarted
*/
type EventType uint8

const (
	ServiceDoesNotExist EventType = iota
	ServiceExists
	HealthFailure
	HealthRecovered
	ServiceRestarting
	ServiceRestarted
)

type Event struct {
	Check Checkable
	Rule  *Rule
	Type  EventType
}

type EventHandler interface {
	Fire(event *Event)
}
