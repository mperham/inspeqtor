package inq

import (
	"inspeqtor/conf/inq/ast"
	"inspeqtor/conf/inq/lexer"
	"inspeqtor/conf/inq/parser"
	"inspeqtor/util"
	"io/ioutil"
	"log"
	"path/filepath"
)

type Checks struct {
	Host      *ast.HostCheck
	Processes []*ast.ProcessCheck
}

func ParseChecks(rootDir string) (*Checks, error) {
	util.Debug("Parsing config in " + rootDir)
	files, err := filepath.Glob(rootDir + "/*.inq")
	if err != nil {
		return nil, err
	}

	var checks Checks
	checks.Processes = make([]*ast.ProcessCheck, len(files))

	for _, filename := range files {
		log.Println("Parsing " + filename)
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}

		s := lexer.NewLexer([]byte(data))
		p := parser.NewParser()
		obj, err := p.Parse(s)
		switch obj.(type) {
		case *ast.HostCheck:
			checks.Host = obj.(*ast.HostCheck)
		case *ast.ProcessCheck:
			checks.Processes = append(checks.Processes, obj.(*ast.ProcessCheck))
		}
	}

	return &checks, nil
}
