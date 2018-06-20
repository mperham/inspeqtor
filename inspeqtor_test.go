package inspeqtor

import (
	"strings"
	"testing"

	"github.com/mperham/inspeqtor/util"
	"github.com/stretchr/testify/assert"
)

func init() {
	util.LogVerbose = true
	util.LogDebug = true
	util.LogInfo = true
}

func TestInspeqtorParse(t *testing.T) {
	t.Parallel()
	i, err := New("test", "")
	assert.Nil(t, err)
	err = i.Parse()
	assert.Nil(t, err)
	assert.Equal(t, i.GlobalConfig.CycleTime, uint(15))
	assert.Equal(t, i.GlobalConfig.Variables["foo"], "bar")
	assert.Equal(t, 2, len(i.Services))

	i, err = New("testx", "")
	assert.NotNil(t, i)
	assert.Nil(t, err)
	err = i.Parse()
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "Missing required file"))
}

func TestCreateSocket(t *testing.T) {
	// not parallelizable since it uses a Unix sock
	i, err := New("test", "")
	assert.Nil(t, err)

	err = i.Parse()
	assert.Nil(t, err)

	err = i.openSocket("/tmp/tmp.sock")
	assert.Nil(t, err)
	assert.NotNil(t, i.Socket)
	defer i.Socket.Close()
}

func TestTestAlertRoutes(t *testing.T) {
	i, err := New("test", "")
	assert.Nil(t, err)

	err = i.Parse()
	assert.Nil(t, err)

	badCount := i.TestAlertRoutes()
	assert.Equal(t, badCount, 1)
}
