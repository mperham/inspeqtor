package ast

import (
	"github.com/mperham/inspeqtor/conf/inq/token"
	"strconv"
	"strings"
)

type HostCheck struct {
	Rules      []Rule
	Parameters map[string]string
}

type ProcessCheck struct {
	Name       string
	Rules      RuleList
	Parameters map[string]string
}

type RuleList []Rule

type RuleMetric struct {
	Family string
	Name   string
}

type Rule struct {
	Metric     RuleMetric
	Operator   string
	Threshold  Amount
	Actions    []Action
	CycleCount int
}

type Action interface {
	Name() string
}

type SimpleAction struct {
	ActionName string
}

func (s *SimpleAction) Name() string { return s.ActionName }

type Amount struct {
	Raw    string
	Parsed int64
	PerSec bool
}

var (
	CreateAction = func(props ...string) (Action, error) {
		return &SimpleAction{props[0]}, nil
	}
)

func AppendAction(action interface{}, list interface{}) ([]Action, error) {
	lst := list.([]Action)
	lst = append(lst, action.(Action))
	return lst, nil
}

func AddAction(name interface{}, arg interface{}) (Action, error) {
	argStr := ""
	if arg != nil {
		argStr = string(arg.(*token.Token).Lit)
	}
	return CreateAction(string(name.(*token.Token).Lit), argStr)
}

func AddParam(key interface{}, val interface{}, hash interface{}) (map[string]string, error) {
	k := string(key.(*token.Token).Lit)
	v := string(val.(*token.Token).Lit)

	// remove quotes from quoted strings
	if v[0] == '"' {
		val, err := strconv.Unquote(v)
		if err != nil {
			return nil, err
		}
		v = val
	}

	var h map[string]string

	if hash == nil {
		h = map[string]string{}
	} else {
		h = hash.(map[string]string)
	}
	h[k] = v

	return h, nil
}

func NewProcessCheck(checkType interface{}, name interface{}, rules interface{}, params interface{}) *ProcessCheck {
	return &ProcessCheck{
		string(name.(*token.Token).Lit),
		rules.(RuleList),
		params.(map[string]string),
	}
}
func NewHostCheck(rules interface{}, params interface{}) *HostCheck {
	return &HostCheck{
		rules.(RuleList),
		params.(map[string]string),
	}
}

func NewRuleList(rule interface{}) RuleList {
	return RuleList{rule.(Rule)}
}

func AppendRule(list interface{}, rule interface{}) RuleList {
	return append(list.(RuleList), rule.(Rule))
}

func NewRule(metric interface{}, operator interface{}, value interface{}, actions interface{}, cycleCount interface{}) Rule {
	return Rule{
		*metric.(*RuleMetric),
		string(operator.(*token.Token).Lit),
		*(value.(*Amount)),
		actions.([]Action),
		int(cycleCount.(*Amount).Parsed),
	}
}

func Metric(family interface{}, name interface{}) (*RuleMetric, error) {
	m := &RuleMetric{string(family.(*token.Token).Lit), ""}
	if name != nil {
		m.Name = string(name.(*token.Token).Lit)
	}
	return m, nil
}

func HumanAmount(digits interface{}) (*Amount, error) {
	orig := string(digits.(*token.Token).Lit)
	str := orig

	perSec := false
	slen := len(str)
	if slen > 4 && str[slen-4:] == "/sec" {
		perSec = true
		str = str[0 : slen-4]
	}

	amt, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		return &Amount{orig, amt, perSec}, nil
	}

	sizecode := str[len(str)-1:]
	val := str[0 : len(str)-1]
	amt, err = strconv.ParseInt(val, 10, 64)
	if err != nil {
		return nil, err
	}

	sizecode = strings.ToLower(sizecode)
	if sizecode == "k" {
		amt *= 1024
	} else if sizecode == "m" {
		amt *= 1024 * 1024
	} else if sizecode == "g" {
		amt *= 1024 * 1024 * 1024
	} else if sizecode == "t" {
		amt *= 1024 * 1024 * 1024 * 1024
	} else if sizecode == "p" {
		amt *= 1024 * 1024 * 1024 * 1024 * 1024
	} else if sizecode == "%" {
		// nothing to do
	}
	return &Amount{orig, amt, perSec}, nil
}

func ToInt64(v interface{}) (*Amount, error) {
	raw := string(v.(*token.Token).Lit)
	parsed, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return nil, err
	}
	return &Amount{raw, parsed, false}, nil
}
