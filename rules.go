package inspeqtor

import (
	"inspeqtor/metrics"
	"inspeqtor/util"
)

/*
 Run through all Rules and check if we need to trigger actions
*/
func checkRule(svcName string, svcData interface{}, rule *Rule) RuleStatus {
	curval := metrics.Lookup(svcData, rule.MetricFamily, rule.MetricName)
	tripped := false

	switch rule.Op {
	case LT:
		tripped = curval < rule.Threshold
	case GT:
		tripped = curval > rule.Threshold
	default:
		util.Warn("Unknown operator: %d", rule.Op)
	}

	if rule.Status == Ok && tripped {
		util.Warn(svcName + "[" + rule.Metric() + "] just tripped")
		rule.Status = Tripped
		return Tripped
	}
	if rule.Status == Tripped && !tripped {
		util.Warn(svcName + "[" + rule.Metric() + "] just recovered")
		rule.Status = Ok
		return Recovered
	}

	return Unchanged
}
