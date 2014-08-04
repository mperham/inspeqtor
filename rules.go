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
	op           Operator
	threshold    int64
	cycleCount   int
	trippedCount int
	actions      []*Action
}

func (r Rule) MetricName() string {
	s := r.metricFamily
	if r.metricName != "" {
		s += "(" + r.metricName + ")"
	}
	return s
}

func (r Rule) Threshold() int64 {
	return r.threshold
}

func (r Rule) Op() Operator {
	return r.op
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

	switch rule.op {
	case LT:
		tripped = curval < rule.threshold
	case GT:
		tripped = curval > rule.threshold
	default:
		util.Warn("Unknown operator: %d", rule.op)
	}

	if tripped {
		util.Debug("%s[%s] tripped.  Current value = %d", svcName, rule.MetricName(), curval)
		rule.trippedCount++
	}

	if rule.trippedCount == rule.cycleCount && tripped {
		util.Warn("%s[%s] triggered.  Current value = %d", svcName, rule.MetricName(), curval)
		return Triggered
	}
	if rule.trippedCount != 0 && !tripped {
		util.Info("%s[%s] recovered.", svcName, rule.MetricName())
		rule.trippedCount = 0
		return Recovered
	}

	return Unchanged
}
