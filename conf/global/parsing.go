package inq

import (
	"inspeqtor/conf/global/ast"
	"inspeqtor/conf/global/lexer"
	"inspeqtor/conf/global/parser"
	"inspeqtor/util"
	"io/ioutil"
	"log"
	"path/filepath"
)

func Parse(rootDir string) (*ast.ConfigFile, error) {
	util.Debug("Parsing global config in " + rootDir)
	files, err := filepath.Glob(rootDir + "/inspeqtor.conf")
	if err != nil {
		return nil, err
	}

	var global ast.ConfigFile
	for _, filename := range files {
		log.Println("Parsing " + filename)
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}

		s := lexer.NewLexer([]byte(data))
		p := parser.NewParser()
		obj, err := p.Parse(s)
		global = obj.(ast.ConfigFile)
		util.DebugDebug("%v", global)
	}

	return &global, nil
}
