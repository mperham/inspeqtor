package inspeqtor

import (
	"errors"
	"inspeqtor/conf/inq/ast"
	"inspeqtor/conf/inq/lexer"
	"inspeqtor/conf/inq/parser"
	"inspeqtor/metrics"
	"inspeqtor/util"
	"io/ioutil"
	"os"
	"path/filepath"
)

/*
Parses the host- and service-specific rules in /etc/inspeqtor/conf.d/*.inq
*/
func ParseInq(confDir string) (*Host, []*Service, error) {
	util.Debug("Parsing config in " + confDir)
	files, err := filepath.Glob(confDir + "/*.inq")
	if err != nil {
		return nil, nil, err
	}

	var host *Host
	var checks []*Service

	for _, filename := range files {
		util.DebugDebug("Parsing " + filename)
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, nil, err
		}

		s := lexer.NewLexer([]byte(data))
		p := parser.NewParser()
		obj, err := p.Parse(s)
		if err != nil {
			util.Warn("Unable to parse " + filename + ": " + err.Error())
			continue
		}

		switch obj.(type) {
		case *ast.HostCheck:
			if host != nil {
				panic("Found more than one \"check host\" configuration in " + confDir)
			}
			host, err = convertHost(obj.(*ast.HostCheck))
			if err != nil {
				return nil, nil, err
			}
			util.DebugDebug("Host: %+v", *host)
		case *ast.ProcessCheck:
			svc, err := convertService(obj.(*ast.ProcessCheck))
			if err != nil {
				return nil, nil, err
			}
			util.DebugDebug("Service: %+v", *svc)
			checks = append(checks, svc)
		}
	}

	return host, checks, nil
}

// GACK, so ugly
func convertHost(inqhost *ast.HostCheck) (*Host, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	rules := make([]*Rule, len(inqhost.Rules))
	for i, rule := range inqhost.Rules {
		rule, err := convertRule(rule, nil)
		util.DebugDebug("Rule: %+v", *rule)
		if err != nil {
			return nil, err
		}
		rules[i] = rule
	}
	return &Host{hostname, rules, metrics.NewStore()}, nil
}

func convertRule(inqrule *ast.Rule, actionList []*Action) (*Rule, error) {
	op := GT
	switch inqrule.Operator {
	case ">":
		op = GT
	case "<":
		op = LT
	default:
		return nil, errors.New("Unknown operator: " + inqrule.Operator)
	}

	return &Rule{inqrule.Metric.Family, inqrule.Metric.Name, op,
		inqrule.Value, inqrule.CycleCount, 0, nil}, nil
}

func convertService(inqsvc *ast.ProcessCheck) (*Service, error) {
	rules := make([]*Rule, len(inqsvc.Rules))
	for i, rule := range inqsvc.Rules {
		rule, err := convertRule(rule, nil)
		if err != nil {
			return nil, err
		}
		util.DebugDebug("Rule: %+v", *rule)
		rules[i] = rule
	}
	svc := &Service{inqsvc.Name, 0, 0, rules, metrics.NewStore(), nil}
	return svc, nil
}
