package inspeqtor

import (
	"fmt"
	"os"
)

/*
 There are several different types of Events:

 * Process disappeared (or did not exist when we started up)
 * Process appeared
 * Rule failed check
 * Rule has recovered
*/
type EventType string

const (
	ProcessDoesNotExist EventType = "ProcessDoesNotExist"
	ProcessExists       EventType = "ProcessExists"
	RuleFailed          EventType = "RuleFailed"
	RuleRecovered       EventType = "RuleRecovered"
)

var (
	Events = []EventType{ProcessDoesNotExist, ProcessExists, RuleFailed, RuleRecovered}
)

func (s EventType) String() string {
	return string(s)
}

type Event struct {
	Type EventType
	Eventable
	*Rule
}

func (e *Event) Thing() Eventable {
	return e.Eventable
}

func (e *Event) Service() *Service {
	return e.Eventable.(*Service)
}

func (e *Event) Hostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return err.Error()
	}
	return hostname
}

func (e *Event) Target() string {
	switch x := e.Eventable.(type) {
	case *Service:
		return fmt.Sprintf("%s[%s]", e.Hostname(), x.Name())
	case *Host:
		return x.Name()
	default:
		return fmt.Sprintf("Unknown: %s", e.Eventable)
	}
}
