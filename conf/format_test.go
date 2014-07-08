package conf

import (
	"fmt"
	"inspeqtor/conf/ast"
	"inspeqtor/conf/lexer"
	"inspeqtor/conf/parser"
	"io/ioutil"
	"log"
	"testing"
)

func TestBasicServiceParsing(t *testing.T) {
	data, err := ioutil.ReadFile("conf.d/memcache.inq")
	if err != nil {
		t.Fatal(err)
	}

	s := lexer.NewLexer([]byte(data))
	p := parser.NewParser()
	obj, err := p.Parse(s)
	log.Println("%v", obj)
	if err != nil {
		t.Error(err)
	}

	check := obj.(*ast.Check)
	fmt.Printf("output: %s\n", check)
}
