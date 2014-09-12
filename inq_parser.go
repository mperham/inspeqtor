package inspeqtor

import (
	"errors"
	"inspeqtor/conf/inq/ast"
	"inspeqtor/conf/inq/lexer"
	"inspeqtor/conf/inq/parser"
	"inspeqtor/metrics"
	"inspeqtor/services"
	"inspeqtor/util"
	"io/ioutil"
	"os"
	"path/filepath"
)

/*
Parses the host- and service-specific rules in /etc/inspeqtor/conf.d/*.inq
*/
func ParseInq(global *ConfigFile, confDir string) (*Host, []Checkable, error) {
	util.Debug("Parsing config in " + confDir)
	files, err := filepath.Glob(confDir + "/*.inq")
	if err != nil {
		return nil, nil, err
	}

	var host *Host
	var checks []Checkable

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
			host, err = convertHost(global, obj.(*ast.HostCheck))
			if err != nil {
				return nil, nil, err
			}
			util.DebugDebug("Host: %+v", *host)
		case *ast.ProcessCheck:
			svc, err := convertService(global, obj.(*ast.ProcessCheck))
			if err != nil {
				return nil, nil, err
			}
			util.DebugDebug("Service: %+v", *svc)
			checks = append(checks, svc)
		}
	}

	if host == nil {
		return nil, nil, errors.New("No " + confDir + "/host.inq file found for host monitoring")
	}

	return host, checks, nil
}

// GACK, so ugly
func convertHost(global *ConfigFile, inqhost *ast.HostCheck) (*Host, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	storage := metrics.NewHostStore("/proc", global.Top.CycleTime)
	h := &Host{&Entity{hostname, nil, storage, inqhost.Parameters}}
	rules := make([]*Rule, len(inqhost.Rules))
	for idx, rule := range inqhost.Rules {
		rule, err := convertRule(global, h, rule)
		util.DebugDebug("Rule: %+v", rule)
		if err != nil {
			return nil, err
		}
		rules[idx] = rule
	}
	h.rules = rules
	return h, nil
}

func convertRule(global *ConfigFile, check Checkable, inqrule ast.Rule) (*Rule, error) {
	op := GT
	switch inqrule.Operator {
	case ">":
		op = GT
	case "<":
		op = LT
	default:
		return nil, errors.New("Unknown operator: " + inqrule.Operator)
	}

	threshold, err := check.Metrics().PrepareRule(inqrule.Metric.Family, inqrule.Metric.Name, inqrule.Threshold.Parsed)
	if err != nil {
		return nil, err
	}

	actions := make([]Action, 0)
	for _, action := range inqrule.Actions {
		act, err := convertAction(global, check, action.Name, action.Team)
		if err != nil {
			return nil, err
		}
		actions = append(actions, act)
	}

	return &Rule{check, inqrule.Metric.Family, inqrule.Metric.Name,
		op, inqrule.Threshold.Raw, threshold, 0, inqrule.CycleCount, 0, Ok, actions}, nil
}

func convertAction(global *ConfigFile, check Checkable, name string, team string) (Action, error) {
	switch name {
	case "alert":
		owner := team
		if owner == "" {
			owner = check.Owner()
		}

		route := global.AlertRoutes[owner]
		if owner == "" && route == nil {
			return nil, errors.New("No default alert route configured!")
		}
		if route == nil {
			return nil, errors.New("No such alert route: " + owner)
		}
		return Actions["alert"](check, route)
	case "restart":
		return Actions["restart"](check, nil)
	}
	return nil, nil
}

func convertService(global *ConfigFile, inqsvc *ast.ProcessCheck) (*Service, error) {
	rules := make([]*Rule, len(inqsvc.Rules))
	storage := metrics.NewProcessStore("/proc", global.Top.CycleTime)

	svc := &Service{&Entity{inqsvc.Name, nil, storage, inqsvc.Parameters}, nil, services.NewStatus(), nil}

	action, err := convertAction(global, svc, "alert", inqsvc.Parameters["owner"])
	if err != nil {
		return nil, err
	}
	svc.EventHandler = action

	for idx, rule := range inqsvc.Rules {
		rule, err := convertRule(global, svc, rule)
		if err != nil {
			return nil, err
		}
		util.DebugDebug("Rule: %+v", *rule)
		rules[idx] = rule
	}
	svc.rules = rules
	return svc, nil
}
