package inspeqtor

import (
	"inspeqtor/util"
)

type Operator uint8

const (
	LT Operator = iota
	GT
)

type RuleState uint8

const (
	Ok RuleState = iota
	Triggered
	Recovered
)

type Rule struct {
	Entity           Checkable
	metricFamily     string
	metricName       string
	op               Operator
	displayThreshold string
	threshold        int64
	currentValue     int64
	cycleCount       int
	trippedCount     int
	state            RuleState
	actions          []Action
}

func (r *Rule) MetricName() string {
	s := r.metricFamily
	if r.metricName != "" {
		s += "(" + r.metricName + ")"
	}
	return s
}

func (r *Rule) EntityName() string {
	return r.Entity.Name()
}

func (r *Rule) DisplayState() string {
	if r.state == Triggered {
		return "!"
	} else {
		return ""
	}
}

func (r *Rule) DisplayThreshold() string {
	return r.displayThreshold
}

func (r *Rule) Threshold() int64 {
	return r.threshold
}

func (r *Rule) CurrentValue() int64 {
	return r.Entity.Metrics().Get(r.metricFamily, r.metricName)
}

func (r *Rule) DisplayCurrentValue() string {
	return r.Entity.Metrics().Display(r.metricFamily, r.metricName)
}

func (r *Rule) Op() string {
	switch r.op {
	case GT:
		return "greater than"
	case LT:
		return "less than"
	default:
		return "?"
	}
}

func (r *Rule) Reset() {
	r.trippedCount = 0
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
func (rule *Rule) Check() *Event {
	rule.currentValue = rule.Entity.Metrics().Get(rule.metricFamily, rule.metricName)
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

type stateHandler func(*Rule, bool) *Event

var (
	stateMachine = map[RuleState]stateHandler{
		Ok:        okHandler,
		Recovered: recoveredHandler,
		Triggered: triggeredHandler,
	}
)

func okHandler(rule *Rule, tripped bool) *Event {
	if tripped && rule.trippedCount == rule.cycleCount {
		util.Warn("%s[%s] triggered.  Current value = %d", rule.EntityName(), rule.MetricName(), rule.currentValue)
		rule.state = Triggered
		return &Event{RuleFailed, rule.Entity, rule}
	} else if tripped {
		util.Debug("%s[%s] tripped. Current: %d, Threshold: %d", rule.EntityName(), rule.MetricName(), rule.currentValue, rule.threshold)
	}
	return nil
}

func recoveredHandler(rule *Rule, tripped bool) *Event {
	if tripped && rule.trippedCount == rule.cycleCount {
		util.Warn("%s[%s] flapped.  Current value = %d", rule.EntityName(), rule.MetricName(), rule.currentValue)
		rule.state = Triggered
	} else {
		rule.state = Ok
		return &Event{RuleRecovered, rule.Entity, rule}
	}
	return nil
}

func triggeredHandler(rule *Rule, tripped bool) *Event {
	if !tripped {
		util.Info("%s[%s] recovered.", rule.EntityName(), rule.MetricName())
		rule.state = Recovered
		return nil
	} else {
		util.Debug("%s[%s] still triggered. Current: %d, Threshold: %d", rule.EntityName(), rule.MetricName(), rule.currentValue, rule.threshold)
	}
	return nil
}
