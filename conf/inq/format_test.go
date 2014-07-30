package inq

import (
	"fmt"
	"inspeqtor/conf/inq/ast"
	"inspeqtor/conf/inq/lexer"
	"inspeqtor/conf/inq/parser"
	"io/ioutil"
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
	if err != nil {
		t.Fatal(err)
	}

	check := obj.(*ast.ProcessCheck)
	fmt.Printf("output: %s\n", check)
}

func TestBasicHostParsing(t *testing.T) {
	data, err := ioutil.ReadFile("conf.d/system.inq")
	if err != nil {
		t.Fatal(err)
	}

	s := lexer.NewLexer([]byte(data))
	p := parser.NewParser()
	obj, err := p.Parse(s)
	if err != nil {
		t.Fatal(err)
	}

	check := obj.(*ast.HostCheck)
	fmt.Printf("output: %s\n", check)
}
