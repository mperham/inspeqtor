package inspeqtor

import (
	"inspeqtor/metrics"
	"inspeqtor/util"
)

type RuleStatus uint8

const (
	Undetermined RuleStatus = iota
	Ok
	Tripped
	Recovered
	Unchanged
)

type Rule struct {
	metricFamily string
	metricName   string
	Op           Operator
	Threshold    int64
	CycleCount   uint8
	Status       RuleStatus
	Actions      []*Action
}

func (r Rule) MetricName() string {
	s := r.metricFamily
	if r.metricName != "" {
		s += "(" + r.metricName + ")"
	}
	return s
}

func (r Rule) Trippable() bool {
	return r.Status == Ok ||
		r.Status == Undetermined
}

/*
 Run through all Rules and check if we need to trigger actions
*/
func (rule *Rule) Check(svcName string, svcData metrics.Storage) RuleStatus {
	curval := svcData.Get(rule.metricFamily, rule.metricName)
	if curval == -1 {
		rule.Status = Undetermined
		return Undetermined
	}

	tripped := false

	switch rule.Op {
	case LT:
		tripped = curval < rule.Threshold
	case GT:
		tripped = curval > rule.Threshold
	default:
		util.Warn("Unknown operator: %d", rule.Op)
	}

	if rule.Trippable() && tripped {
		util.Warn(svcName + "[" + rule.MetricName() + "] just tripped")
		rule.Status = Tripped
		return Tripped
	}
	if rule.Status == Tripped && !tripped {
		util.Warn(svcName + "[" + rule.MetricName() + "] just recovered")
		rule.Status = Ok
		return Recovered
	}

	return Unchanged
}
