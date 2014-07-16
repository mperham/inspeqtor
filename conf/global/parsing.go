package global

import (
	"inspeqtor/conf/global/ast"
	"inspeqtor/conf/global/lexer"
	"inspeqtor/conf/global/parser"
	"inspeqtor/util"
	"io/ioutil"
)

var (
	Defaults = map[string]string{
		"cycle":         "30",
		"default_alert": "email",
	}
)

type ConfigFile struct {
	TopConfig  map[string]string
	Contextual map[string]map[string]string
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
		ast := obj.(*ast.ConfigFile)
		return &ConfigFile{ast.TopConfig, ast.Contextual}, nil
	} else {
		util.Info("No configuration file found at " + rootDir + "/inspector.conf")
		return &ConfigFile{Defaults, nil}, nil
	}
}
