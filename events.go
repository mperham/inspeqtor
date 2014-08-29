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
	Checkable
	*Rule
}

func (e *Event) Service() *Service {
	return e.Checkable.(*Service)
}

func (e *Event) Hostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return err.Error()
	}
	return hostname
}

func (e *Event) Target() string {
	switch x := e.Checkable.(type) {
	case *Service:
		return fmt.Sprintf("%s[%s]", e.Hostname(), x.Name())
	case *Host:
		return fmt.Sprintf("%s", x.Name())
	default:
		return fmt.Sprintf("Unknown: %s", e.Checkable)
	}
}
