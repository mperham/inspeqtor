package inq

import (
	"github.com/stretchr/testify/assert"
	"inspeqtor/conf/inq/ast"
	"inspeqtor/conf/inq/lexer"
	"inspeqtor/conf/inq/parser"
	"io/ioutil"
	"log"
	"testing"
)

func TestMysqlParsing(t *testing.T) {
	data, err := ioutil.ReadFile("conf.d/mysql.inq")
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
	assert.Equal(t, check.Name, "mysql")
	assert.Equal(t, len(check.Parameters), 0)
	assert.Equal(t, len(check.Rules), 2)
}

func TestBasicServiceParsing(t *testing.T) {
	data, err := ioutil.ReadFile("conf.d/memcached.inq")
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
	assert.Equal(t, check.Name, "memcached")
	assert.Equal(t, len(check.Rules), 2)
	assert.Equal(t, len(check.Parameters), 2)
	assert.Equal(t, check.Parameters["key"], "value")
	assert.Equal(t, check.Parameters["foo"], "bar")
	for _, x := range check.Rules {
		log.Printf("%+v", *x)
	}
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
	assert.Equal(t, len(check.Rules), 3)
}
