package inq

import (
	"io/ioutil"
	"sort"
	"strings"
	"testing"

	"github.com/mperham/inspeqtor/conf/inq/ast"
	"github.com/mperham/inspeqtor/conf/inq/lexer"
	"github.com/mperham/inspeqtor/conf/inq/parser"
	"github.com/stretchr/testify/assert"
)

func TestMysqlParsing(t *testing.T) {
	data, err := ioutil.ReadFile("fixtures/mysql.inq")
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
	assert.Equal(t, len(check.Rules), 4)
	assert.Equal(t, check.Rules[1].Threshold.Parsed, 50)
	assert.Equal(t, check.Rules[1].Threshold.Raw, "50")
	assert.Equal(t, check.Rules[1].Threshold.PerSec, false)

	assert.Equal(t, check.Rules[2].Threshold.Parsed, 1024)
	assert.Equal(t, check.Rules[2].Threshold.Raw, "1k/sec")
	assert.Equal(t, check.Rules[2].Threshold.PerSec, true)

	assert.Equal(t, check.Rules[3].Threshold.Parsed, 2)
	assert.Equal(t, check.Rules[3].Threshold.Raw, "2/sec")
	assert.Equal(t, check.Rules[3].Threshold.PerSec, true)
}

func TestBadAmount(t *testing.T) {
	data, err := ioutil.ReadFile("fixtures/bad_amount.inq")
	if err != nil {
		t.Fatal(err)
	}

	s := lexer.NewLexer([]byte(data))
	p := parser.NewParser()
	obj, err := p.Parse(s)
	assert.Nil(t, obj)
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "Invalid amount"))
}

func TestExpose(t *testing.T) {
	data, err := ioutil.ReadFile("fixtures/inspeqtor.inq")
	if err != nil {
		t.Fatal(err)
	}

	s := lexer.NewLexer([]byte(data))
	p := parser.NewParser()
	obj, err := p.Parse(s)
	assert.NotNil(t, obj)
	assert.Nil(t, err)

	check := obj.(*ast.ProcessCheck)
	assert.Equal(t, check.Name, "inspeqtor")
	assert.Equal(t, len(check.Parameters), 0)
	assert.Equal(t, len(check.Rules), 0)
	assert.Equal(t, len(check.Exposed), 1)
	assert.Equal(t, "memstats", check.Exposed[0])
}

func TestBasicServiceParsing(t *testing.T) {
	data, err := ioutil.ReadFile("fixtures/memcached.inq")
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
	assert.Equal(t, "memcached", check.Name)
	assert.Equal(t, 2, len(check.Rules))
	assert.Equal(t, 4, len(check.Parameters))
	assert.Equal(t, "dev", check.Parameters["owner"])
	assert.Equal(t, "bar", check.Parameters["foo"])
	assert.Equal(t, "/foo", check.Parameters["endpoint"])
	assert.Equal(t, "whoa sp\"aces", check.Parameters["quoted"])
	assert.Equal(t, "cpu", check.Rules[1].Metric.Family)
	assert.Equal(t, "user", check.Rules[1].Metric.Name)

	var names []string
	for _, act := range check.Rules[0].Actions {
		names = append(names, act.Name())
	}
	sort.Strings(names)

	assert.Equal(t, "alert", names[0])
	assert.Equal(t, "reload", names[1])
	assert.Equal(t, "restart", names[2])
}

func TestBasicHostParsing(t *testing.T) {
	data, err := ioutil.ReadFile("fixtures/host.inq")
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
