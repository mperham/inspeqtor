package ast

import (
	"inspeqtor/conf/token"
)

type Check struct {
	InitType string
	Name     string
	Rules    []*Rule
}

type RuleList []*Rule

type Rule struct {
	Metric   string
	Operator string
	Value    string
	Action   string
}

func NewCheck(initType interface{}, checkType interface{}, name interface{}, rules interface{}) *Check {
	return &Check{
		string(initType.(*token.Token).Lit),
		string(name.(*token.Token).Lit),
		rules.([]*Rule),
	}
}

func NewRuleList(rule interface{}) []*Rule {
	return RuleList{rule.(*Rule)}
}

func AppendRule(list interface{}, rule interface{}) []*Rule {
	return append(list.(RuleList), rule.(*Rule))
}

func NewRule(metric interface{}, operator interface{}, value interface{}, action interface{}) *Rule {
	return &Rule{
		string(metric.(*token.Token).Lit),
		string(operator.(*token.Token).Lit),
		string(value.(*token.Token).Lit),
		string(action.(*token.Token).Lit),
	}
}
