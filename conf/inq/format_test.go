package inq

import (
	"github.com/mperham/inspeqtor/conf/inq/ast"
	"github.com/mperham/inspeqtor/conf/inq/lexer"
	"github.com/mperham/inspeqtor/conf/inq/parser"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
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
	assert.Equal(t, len(check.Parameters), 4)
	assert.Equal(t, check.Parameters["owner"], "dev")
	assert.Equal(t, check.Parameters["foo"], "bar")
	assert.Equal(t, check.Parameters["endpoint"], "/foo")
	assert.Equal(t, check.Parameters["quoted"], "whoa sp\"aces")
	assert.Equal(t, check.Rules[0].Actions[1].Name(), "alert")
}

func TestBasicHostParsing(t *testing.T) {
	data, err := ioutil.ReadFile("conf.d/host.inq")
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
