package inspeqtor

import (
	"inspeqtor/util"
)

func (r Rule) MetricName() string {
	s := r.metricFamily
	if r.metricName != "" {
		s += "(" + r.metricName + ")"
	}
	return s
}

func (r Rule) EntityName() string {
	return r.entity.Name()
}

func (r Rule) Threshold() int64 {
	return r.threshold
}

func (r Rule) Op() string {
	switch r.op {
	case GT:
		return ">"
	case LT:
		return "<"
	default:
		return "?"
	}
}

/*
 Run through all Rules and check if we need to trigger actions.

 "tripped" means the Rule threshold was breached **this cycle**.
 "triggered" means the Rule threshold was breached enough cycles in
 a row to fire the alerts associated with the Rule.

 There are three possible states for Rules:
 Ok - rule was fine last time and passed this time too.
 Triggered - threshold breached enough times, action should be taken
 Recovered - rule is currently Triggered but threshold was not breached this time
*/
func (rule *Rule) Check() *Rule {
	rule.currentValue = rule.entity.MetricData().Get(rule.metricFamily, rule.metricName)
	if rule.currentValue == -1 {
		return nil
	}

	tripped := false

	switch rule.op {
	case LT:
		tripped = rule.currentValue < rule.threshold
	case GT:
		tripped = rule.currentValue > rule.threshold
	}

	if tripped {
		rule.trippedCount = rule.trippedCount + 1
	} else {
		rule.trippedCount = 0
	}

	return stateMachine[rule.state](rule, tripped)
}

type stateHandler func(*Rule, bool) *Rule

var (
	stateMachine = map[RuleState]stateHandler{
		Ok:        okHandler,
		Recovered: recoveredHandler,
		Triggered: triggeredHandler,
	}
)

func okHandler(rule *Rule, tripped bool) *Rule {
	if tripped && rule.trippedCount == rule.cycleCount {
		util.Warn("%s[%s] triggered.  Current value = %d", rule.entity.Name(), rule.MetricName(), rule.currentValue)
		rule.state = Triggered
		return rule
	} else if tripped {
		util.Debug("%s[%s] tripped. Current: %d, Threshold: %d", rule.entity.Name(), rule.MetricName(), rule.currentValue, rule.threshold)
	}
	return nil
}

func recoveredHandler(rule *Rule, tripped bool) *Rule {
	if tripped && rule.trippedCount == rule.cycleCount {
		util.Warn("%s[%s] flapped.  Current value = %d", rule.entity.Name(), rule.MetricName(), rule.currentValue)
		rule.state = Triggered
		return rule
	} else {
		rule.state = Ok
	}
	return nil
}

func triggeredHandler(rule *Rule, tripped bool) *Rule {
	if !tripped {
		util.Info("%s[%s] recovered.", rule.entity.Name(), rule.MetricName())
		rule.state = Recovered
		return rule
	} else {
		util.Debug("%s[%s] still triggered. Current: %d, Threshold: %d", rule.entity.Name(), rule.MetricName(), rule.currentValue, rule.threshold)
	}
	return nil
}
