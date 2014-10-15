package global

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/mperham/inspeqtor/conf/global/ast"
	"github.com/mperham/inspeqtor/conf/global/lexer"
	"github.com/mperham/inspeqtor/conf/global/parser"
	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, "smtp.example.com", config.Routes["analytics"].Config["smtp_server"])
	log.Printf("%+v", config)
}
