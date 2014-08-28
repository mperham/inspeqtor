package inspeqtor

import (
	"fmt"
)

/*
 There are several different types of Events:

 * Process disappeared (or did not exist when we started up)
 * Process appeared
 * Rule failed check
 * Rule has recovered
*/
type EventType uint8

const (
	ProcessDoesNotExist EventType = iota
	ProcessExists
	RuleFailed
	RuleRecovered
)

var (
	Events = []EventType{ProcessDoesNotExist, ProcessExists, RuleFailed, RuleRecovered}
)

// Go question: is there a way to automate / DRY up
// this boilerplate?
func (s EventType) String() string {
	switch s {
	case ProcessDoesNotExist:
		return "ProcessDoesNotExist"
	case ProcessExists:
		return "ProcessExists"
	case RuleFailed:
		return "RuleFailed"
	case RuleRecovered:
		return "RuleRecovered"
	default:
		return fmt.Sprintf("Oops: %d", s)
	}
}

type Event struct {
	Type  EventType
	Check Checkable
	Rule  *Rule
}
