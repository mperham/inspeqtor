package inspeqtor

import (
	"inspeqtor/metrics"
	"inspeqtor/services"
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
	PID         services.ProcessId
	Status      services.Status
	Rules       []*Rule
	Parameters  map[string]string
	Metrics     metrics.Storage

	// Upon bootup, we scan each init system looking for the service
	// and cache which init system manages it for our lifetime.
	Manager services.InitSystem
}

func (s Service) Owner() string {
	return s.Parameters["owner"]
}

type Host struct {
	Hostname string
	Rules    []*Rule
	Metrics  metrics.Storage
}

func (h Service) MetricData() metrics.Storage {
	return h.Metrics
}

func (h Host) MetricData() metrics.Storage {
	return h.Metrics
}

func (h Service) Name() string {
	return h.ServiceName
}

func (h Host) Name() string {
	return h.Hostname
}

type Checkable interface {
	Name() string
	MetricData() metrics.Storage
}

type Operator uint8

const (
	LT Operator = iota
	GT
)

type Alert struct {
	Rule *Rule
}

type Action interface {
	Name() string
	Setup(map[string]string) error
	Trigger(alert *Alert) error
}
