package inspeqtor

import (
	"fmt"
)

/*
 There are 3 pairs of events we need to handle:
 1. Process does not exist, either on startup or disappearing during a cycle.
 2. Process exists, it appears after previously not existing.

 3. Rule triggered after metric crossed threshold for N cycles (simple alert)
 4. Metric recovered

 5. Service needs restarting due to rule triggering.
 6. Service started correctly (shared with 2?)
*/

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
	ProcessDoesNotExist EventType = iota
	ProcessExists
	MetricFailed
	MetricRecovered
	ServiceRestarting
	ServiceRestarted
)

// Go question: is there a way to automate / DRY up
// this boilerplate?
func (s EventType) String() string {
	switch s {
	case ProcessDoesNotExist:
		return "ProcessDoesNotExist"
	case ProcessExists:
		return "ProcessExists"
	case MetricFailed:
		return "MetricFailed"
	case MetricRecovered:
		return "MetricRecovered"
	case ServiceRestarting:
		return "ServiceRestarting"
	case ServiceRestarted:
		return "ServiceRestarted"
	default:
		return fmt.Sprintf("Oops: %d", s)
	}
}

type Event struct {
	Type  EventType
	Check Checkable
	Rule  *Rule
}
