package global

import (
	"github.com/stretchr/testify/assert"
	"inspeqtor/conf/global/ast"
	"inspeqtor/conf/global/lexer"
	"inspeqtor/conf/global/parser"
	"io/ioutil"
	"log"
	"testing"
)

func TestBasicParsing(t *testing.T) {
	data, err := ioutil.ReadFile("fixtures/inspeqtor.conf")
	assert.Nil(t, err)

	s := lexer.NewLexer([]byte(data))
	p := parser.NewParser()
	obj, err := p.Parse(s)
	assert.Nil(t, err)
	assert.NotNil(t, obj)

	config := obj.(ast.Config)
	assert.Equal(t, "15", config.Variables["cycle_time"])
	assert.Equal(t, 3, len(config.Routes))
	assert.Equal(t, "b!l$a%rgh^fazz\"", config.Routes["analytics"].Config["password"])
	assert.Equal(t, "smtp.example.com:587", config.Routes["analytics"].Config["server"])
	log.Printf("%+v", config)
}
