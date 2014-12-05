package inspeqtor

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mperham/inspeqtor/conf/inq/ast"
	"github.com/mperham/inspeqtor/conf/inq/lexer"
	"github.com/mperham/inspeqtor/conf/inq/parser"
	"github.com/mperham/inspeqtor/metrics"
	"github.com/mperham/inspeqtor/services"
	"github.com/mperham/inspeqtor/util"
)

/*
Parses the host-specific rules in /etc/inspeqtor/host.inq
*/
func ParseHost(global *ConfigFile, hostInq string) (*Host, error) {
	var host *Host

	result, err := util.FileExists(hostInq)
	if err != nil {
		return nil, err
	}
	if !result {
		return nil, fmt.Errorf("Missing required file: %s", hostInq)
	}

	util.DebugDebug("Parsing " + hostInq)
	data, err := ioutil.ReadFile(hostInq)
	if err != nil {
		return nil, err
	}

	s := lexer.NewLexer([]byte(data))
	p := parser.NewParser()
	obj, err := p.Parse(s)
	if err != nil {
		return nil, err
	}

	switch x := obj.(type) {
	case *ast.HostCheck:
		host, err = BuildHost(global, x)
		if err != nil {
			return nil, err
		}
		util.DebugDebug("Host: %+v", *host)
	default:
		return nil, fmt.Errorf("Invalid host.inq configuration file")
	}

	return host, nil
}

/*
Parses the service-specific rules in /etc/inspeqtor/services.d/*.inq
*/
func ParseServices(global *ConfigFile, confDir string) ([]Checkable, error) {
	util.Debug("Parsing config in " + confDir)
	files, err := filepath.Glob(confDir + "/*.inq")
	if err != nil {
		return nil, err
	}

	var checks []Checkable

	for _, filename := range files {
		util.DebugDebug("Parsing " + filename)
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}

		s := lexer.NewLexer([]byte(data))
		p := parser.NewParser()
		obj, err := p.Parse(s)
		if err != nil {
			util.Warn("Unable to parse " + filename + ": " + err.Error())
			continue
		}

		switch x := obj.(type) {
		case *ast.ProcessCheck:
			svc, err := BuildService(global, x)
			if err != nil {
				return nil, err
			}
			util.DebugDebug("Service: %+v", *svc)
			checks = append(checks, svc)
		default:
			return nil, fmt.Errorf("Invalid configuration file: %s", filename)
		}
	}

	return checks, nil
}

var (
	BuildHost    = convertHost
	BuildService = convertService
	BuildRule    = convertRule
	BuildAction  = convertAction
	BuildExpose  = convertExpose
)

// GACK, so ugly
func convertHost(global *ConfigFile, inqhost *ast.HostCheck) (*Host, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	storage := metrics.NewHostStore("/proc", global.CycleTime)
	h := &Host{&Entity{hostname, nil, storage, inqhost.Parameters}}
	rules := make([]*Rule, len(inqhost.Rules))
	for idx, rule := range inqhost.Rules {
		rule, err := BuildRule(global, h, rule)
		util.DebugDebug("Rule: %+v", rule)
		if err != nil {
			return nil, err
		}
		err = storage.Prepare(rule.MetricFamily, rule.MetricName)
		if err != nil {
			return nil, err
		}
		rules[idx] = rule
	}
	h.rules = rules
	return h, nil
}

func convertExpose(global *ConfigFile, check Checkable, elements []string, options map[string]string) error {
	return nil
}

func convertRule(global *ConfigFile, check Checkable, inqrule ast.Rule) (*Rule, error) {
	op := GT
	switch inqrule.Operator {
	case ">":
		op = GT
	case "<":
		op = LT
	default:
		return nil, fmt.Errorf("Unknown operator: %s", inqrule.Operator)
	}

	actions := make([]Action, 0)
	for _, action := range inqrule.Actions {
		act, err := BuildAction(global, check, action)
		if err != nil {
			return nil, err
		}
		actions = append(actions, act)
	}

	return &Rule{check, inqrule.Metric.Family, inqrule.Metric.Name,
		op, inqrule.Threshold.Raw, float64(inqrule.Threshold.Parsed), 0, inqrule.Threshold.PerSec, inqrule.CycleCount, 0, Ok, actions}, nil
}

func convertAction(global *ConfigFile, check Eventable, action ast.Action) (Action, error) {
	switch action.Name() {
	case "alert":
		route := global.AlertRoutes[""]
		if route == nil {
			return nil, fmt.Errorf("Please configure a \"send alerts\" statement in inspeqtor.conf.")
		}
		return Actions["alert"](check, route)
	case "reload":
		return Actions["reload"](check, nil)
	case "restart":
		return Actions["restart"](check, nil)
	default:
		return nil, fmt.Errorf("Unknown action: %", action.Name())
	}
}

func convertService(global *ConfigFile, inqsvc *ast.ProcessCheck) (*Service, error) {
	rules := make([]*Rule, len(inqsvc.Rules))
	storage := metrics.NewProcessStore("/proc", global.CycleTime)

	svc := &Service{&Entity{inqsvc.Name, nil, storage, inqsvc.Parameters}, nil, services.NewStatus(), nil}

	action, err := BuildAction(global, svc, &ast.SimpleAction{ActionName: "alert"})
	if err != nil {
		return nil, err
	}
	svc.EventHandler = action

	for idx, rule := range inqsvc.Rules {
		rule, err := convertRule(global, svc, rule)
		if err != nil {
			return nil, err
		}
		err = storage.Prepare(rule.MetricFamily, rule.MetricName)
		if err != nil {
			return nil, err
		}
		util.DebugDebug("Rule: %+v", *rule)
		rules[idx] = rule
	}
	svc.rules = rules

	if len(inqsvc.Exposed) > 0 {
		err := BuildExpose(global, svc, inqsvc.Exposed, inqsvc.Parameters)
		if err != nil {
			return nil, err
		}
	}
	return svc, nil
}
