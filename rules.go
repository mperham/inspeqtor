package inspeqtor

import (
	"inspeqtor/metrics"
	"inspeqtor/util"
)

type RuleStatus uint8

const (
	Undetermined RuleStatus = iota
	Ok
	Triggered
	Recovered
	Unchanged
)

type Rule struct {
	metricFamily string
	metricName   string
	Op           Operator
	Threshold    int64
	CycleCount   int
	TrippedCount int
	Actions      []*Action
}

func (r Rule) MetricName() string {
	s := r.metricFamily
	if r.metricName != "" {
		s += "(" + r.metricName + ")"
	}
	return s
}

/*
 Run through all Rules and check if we need to trigger actions.

 "tripped" means the Rule threshold was breached **this cycle**.
 "triggered" means the Rule threshold was breached enough cycles in
 a row to fire the alerts associated with the Rule.
*/
func (rule *Rule) Check(svcName string, svcData metrics.Storage) RuleStatus {
	curval := svcData.Get(rule.metricFamily, rule.metricName)
	if curval == -1 {
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

	if tripped {
		rule.TrippedCount++
	}

	if rule.TrippedCount == rule.CycleCount && tripped {
		util.Warn("%s[%s] triggered.  Current value = %d", svcName, rule.MetricName(), curval)
		return Triggered
	}
	if rule.TrippedCount != 0 && !tripped {
		util.Info("%s[%s] recovered.", svcName, rule.MetricName())
		rule.TrippedCount = 0
		return Recovered
	}

	return Unchanged
}
