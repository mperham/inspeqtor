package inq

import (
	"errors"
	"inspeqtor/conf/inq/ast"
	"inspeqtor/conf/inq/lexer"
	"inspeqtor/conf/inq/parser"
	"inspeqtor/core"
	"inspeqtor/util"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func ParseChecks(rootDir string) (*core.Host, []*core.Service, error) {
	util.Debug("Parsing config in " + rootDir)
	files, err := filepath.Glob(rootDir + "/*.inq")
	if err != nil {
		return nil, nil, err
	}

	var host *core.Host
	var checks []*core.Service
	checks = make([]*core.Service, len(files))

	for _, filename := range files {
		log.Println("Parsing " + filename)
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, nil, err
		}

		s := lexer.NewLexer([]byte(data))
		p := parser.NewParser()
		obj, err := p.Parse(s)
		switch obj.(type) {
		case *ast.HostCheck:
			host, err = convertHost(obj.(*ast.HostCheck))
			if err != nil {
				return nil, nil, err
			}
		case *ast.ProcessCheck:
			svc, err := convertService(obj.(*ast.ProcessCheck))
			if err != nil {
				return nil, nil, err
			}
			checks = append(checks, svc)
		}
	}

	return host, checks, nil
}

// GACK, so ugly
func convertHost(inqhost *ast.HostCheck) (*core.Host, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	rules := make([]*core.Rule, len(inqhost.Rules))
	for i, rule := range inqhost.Rules {
		rule, err := convertRule(rule, nil)
		if err != nil {
			return nil, err
		}
		rules[i] = rule
	}
	return &core.Host{hostname, rules}, nil
}

func convertRule(inqrule *ast.Rule, actionList []*core.Action) (*core.Rule, error) {
	op := core.GT
	switch inqrule.Operator {
	case ">":
		op = core.GT
	case "<":
		op = core.LT
	default:
		return nil, errors.New("Unknown operator: " + inqrule.Operator)
	}

	return &core.Rule{inqrule.Metric, op, inqrule.Value, inqrule.CycleCount, core.Unknown, nil}, nil
}

func convertService(inqsvc *ast.ProcessCheck) (*core.Service, error) {
	rules := make([]*core.Rule, len(inqsvc.Rules))
	for i, rule := range inqsvc.Rules {
		rule, err := convertRule(rule, nil)
		if err != nil {
			return nil, err
		}
		rules[i] = rule
	}
	svc := &core.Service{inqsvc.Name, nil, rules}
	return svc, nil
}
