package inspeqtor

import (
	"errors"
	"inspeqtor/conf/inq/ast"
	"inspeqtor/conf/inq/lexer"
	"inspeqtor/conf/inq/parser"
	"inspeqtor/util"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func ParseInq(rootDir string) (*Host, []*Service, error) {
	util.Debug("Parsing config in " + rootDir)
	files, err := filepath.Glob(rootDir + "/*.inq")
	if err != nil {
		return nil, nil, err
	}

	var host *Host
	var checks []*Service
	checks = make([]*Service, 0)

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
			util.DebugDebug("%+v", *host)
		case *ast.ProcessCheck:
			svc, err := convertService(obj.(*ast.ProcessCheck))
			if err != nil {
				return nil, nil, err
			}
			util.DebugDebug("%+v", *svc)
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
		util.DebugDebug("%+v", *rule)
		if err != nil {
			return nil, err
		}
		rules[i] = rule
	}
	return &Host{hostname, rules}, nil
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

	return &Rule{inqrule.Metric, op, inqrule.Value, inqrule.CycleCount, Undetermined, nil}, nil
}

func convertService(inqsvc *ast.ProcessCheck) (*Service, error) {
	rules := make([]*Rule, len(inqsvc.Rules))
	for i, rule := range inqsvc.Rules {
		rule, err := convertRule(rule, nil)
		if err != nil {
			return nil, err
		}
		util.DebugDebug("%+v", *rule)
		rules[i] = rule
	}
	svc := &Service{inqsvc.Name, 0, 0, rules, nil}
	return svc, nil
}
