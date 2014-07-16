package global

import (
	"inspeqtor/conf/global/ast"
	"inspeqtor/conf/global/lexer"
	"inspeqtor/conf/global/parser"
	"inspeqtor/util"
	"io/ioutil"
	"strconv"
)

type GlobalConfig struct {
	CycleTime uint16
}

var Defaults = GlobalConfig{30}

type Config map[string]string
type ConfigFile struct {
	Top              GlobalConfig
	ContextualConfig map[string]Config
}

func Parse(rootDir string) (*ConfigFile, error) {
	path := rootDir + "/inspeqtor.conf"
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
		ast := obj.(ast.ConfigFile)

		var config ConfigFile
		config.Top = Defaults
		if val, has := ast.TopConfig["cycle_count"]; has {
			time, err := strconv.Atoi(val)
			if err != nil {
				util.Warn("Invalid cycle time: " + val)
				time = 30
			}
			config.Top.CycleTime = uint16(time)
		}
		config.ContextualConfig = map[string]Config{}
		for k, v := range ast.Contextual {
			config.ContextualConfig[k] = v
		}
		return &config, nil
	} else {
		util.Info("No configuration file found at " + rootDir + "/inspector.conf")
		return &ConfigFile{Defaults, nil}, nil
	}
}
