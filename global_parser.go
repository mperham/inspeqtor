package redacted

import (
	"errors"
	"fmt"
	"github.com/mperham/redacted/conf/global/ast"
	"github.com/mperham/redacted/conf/global/lexer"
	"github.com/mperham/redacted/conf/global/parser"
	"github.com/mperham/redacted/util"
	"io/ioutil"
	"strconv"
)

/*
Parses the global redacted configuration in /etc/redacted/redacted.conf.
*/
type GlobalConfig struct {
	CycleTime    uint
	DeployLength uint
}

var Defaults = GlobalConfig{15, 300}

/*
  An alert route is a way to send an alert to a recipient.

  Channel is the notification mechanism: email, campfire, etc.
  Config is an undefined set of kv pairs for configuring the channel.

  The configuration looks like this:

    send alerts
      via CHANNEL with K V, K V, K V

  You'd then write a rule like:

    if foo > 10 then alert

*/
type AlertRoute struct {
	Name    string
	Channel string
	Config  map[string]string
}

type ConfigFile struct {
	Top         GlobalConfig
	AlertRoutes map[string]*AlertRoute
}

func ParseGlobal(rootDir string) (*ConfigFile, error) {
	path := rootDir + "/redacted.conf"
	exists, err := util.FileExists(path)
	if err != nil {
		return nil, err
	}

	if exists {
		util.Debug("Parsing " + path)
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}

		s := lexer.NewLexer([]byte(data))
		p := parser.NewParser()
		obj, err := p.Parse(s)
		if err != nil {
			return nil, err
		}
		ast := obj.(ast.Config)

		var config ConfigFile
		config.Top = Defaults
		if val, has := ast.Variables["log_level"]; has {
			util.SetLogLevel(val)
		}
		parseValue(ast, &config.Top.CycleTime, "cycle_time", 15)
		parseValue(ast, &config.Top.DeployLength, "deploy_length", 300)

		config.AlertRoutes = map[string]*AlertRoute{}
		for _, v := range ast.Routes {
			ar, err := ValidateChannel(v.Name, v.Channel, v.Config)
			if err != nil {
				return nil, err
			}
			if _, ok := config.AlertRoutes[v.Name]; ok {
				return nil, errors.New(fmt.Sprintf("Duplicate alert config for '%s'", v.Name))
			}
			config.AlertRoutes[v.Name] = ar
		}
		return &config, nil
	} else {
		util.Info("No configuration file found at " + rootDir + "/inspector.conf")
		return &ConfigFile{Defaults, nil}, nil
	}
}

func parseValue(ast ast.Config, store *uint, name string, def uint) {
	if val, has := ast.Variables[name]; has {
		ival, err := strconv.ParseUint(val, 10, 32)
		if err != nil {
			util.Warn("Invalid %s: %d", name, val)
			ival = uint64(def)
		}
		*store = uint(ival)
	}
}
