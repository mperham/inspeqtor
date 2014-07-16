package inq

import (
	"inspeqtor/conf/global/ast"
	"inspeqtor/conf/global/lexer"
	"inspeqtor/conf/global/parser"
	"io/ioutil"
	"log"
	"testing"
)

func TestBasicParsing(t *testing.T) {
	data, err := ioutil.ReadFile("fixtures/inspeqtor.conf")
	ok(t, err)

	s := lexer.NewLexer([]byte(data))
	p := parser.NewParser()
	obj, err := p.Parse(s)
	ok(t, err)
	log.Printf("%v", obj)

	config := obj.(ast.ConfigFile)
	equals(t, "b\"=!%^&!#@ar", config.TopConfig["foo"])
	equals(t, "mikep", config.Contextual["email"]["username"])

}

func assert(tb testing.TB, condition bool, msg string) {
	if !condition {
		tb.Error(msg)
	}
}
func ok(tb testing.TB, err error) {
	if err != nil {
		tb.Fatal(err)
	}
}
func equals(tb testing.TB, exp, act interface{}) {
	if exp != act {
		tb.Error("Expected", exp, ", received ", act)
	}
}
