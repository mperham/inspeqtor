package redacted

import (
	"github.com/mperham/redacted/util"
	"reflect"
)

type Operator uint8

const (
	LT Operator = iota
	GT
)

func (o Operator) String() string {
	switch o {
	case GT:
		return "greater than"
	case LT:
		return "less than"
	default:
		return "???"
	}
}

type RuleState string

const (
	Ok        RuleState = "Ok"
	Triggered RuleState = "Triggered"
	Recovered RuleState = "Recovered"
)

func (o RuleState) String() string {
	return string(o)
}

type Rule struct {
	Entity           Checkable
	MetricFamily     string
	MetricName       string
	Op               Operator
	DisplayThreshold string
	Threshold        float64
	CurrentValue     float64
	CycleCount       int
	TrippedCount     int
	State            RuleState
	Actions          []Action
}

func (r *Rule) Consequence() string {
	for _, a := range r.Actions {
		// So clean!
		if reflect.ValueOf(a).Elem().Type().Name() == "Restarter" {
			return ", restarting"
		}
	}
	return ""
}

func (r *Rule) Metric() string {
	s := r.MetricFamily
	if r.MetricName != "" {
		s += ":" + r.MetricName
	}
	return s
}

func (r *Rule) EntityName() string {
	return r.Entity.Name()
}

func (r *Rule) DisplayState() string {
	if r.State == Triggered {
		return "!"
	} else {
		return ""
	}
}

func (r *Rule) FetchLatestMetricValue() float64 {
	return r.Entity.Metrics().Get(r.MetricFamily, r.MetricName)
}

func (r *Rule) FetchDisplayCurrentValue() string {
	return r.Entity.Metrics().Display(r.MetricFamily, r.MetricName)
}

func (r *Rule) Reset() {
	r.TrippedCount = 0
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
	rule.CurrentValue = rule.FetchLatestMetricValue()
	if rule.CurrentValue == -1 {
		return nil
	}

	tripped := false

	switch rule.Op {
	case LT:
		tripped = rule.CurrentValue < rule.Threshold
	case GT:
		tripped = rule.CurrentValue > rule.Threshold
	}

	if tripped {
		rule.TrippedCount = rule.TrippedCount + 1
	} else {
		rule.TrippedCount = 0
	}

	return stateMachine[rule.State](rule, tripped)
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
	if tripped && rule.TrippedCount == rule.CycleCount {
		util.Warn("%s[%s] triggered.  Current value = %.1f", rule.EntityName(), rule.Metric(), rule.CurrentValue)
		rule.State = Triggered
		return &Event{RuleFailed, rule.Entity, rule}
	} else if tripped {
		util.Debug("%s[%s] tripped. Current: %.1f, Threshold: %.1f", rule.EntityName(), rule.Metric(), rule.CurrentValue, rule.Threshold)
	}
	return nil
}

func recoveredHandler(rule *Rule, tripped bool) *Event {
	if tripped && rule.TrippedCount == rule.CycleCount {
		util.Warn("%s[%s] flapped.  Current value = %.1f", rule.EntityName(), rule.Metric(), rule.CurrentValue)
		rule.State = Triggered
	} else {
		rule.State = Ok
		return &Event{RuleRecovered, rule.Entity, rule}
	}
	return nil
}

func triggeredHandler(rule *Rule, tripped bool) *Event {
	if !tripped {
		util.Info("%s[%s] recovered.", rule.EntityName(), rule.Metric())
		rule.State = Recovered
		return nil
	} else {
		util.Debug("%s[%s] still triggered. Current: %.1f, Threshold: %.1f", rule.EntityName(), rule.Metric(), rule.CurrentValue, rule.Threshold)
	}
	return nil
}
