package inspeqtor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInspeqtorParse(t *testing.T) {
	t.Parallel()
	i, err := New("test", "")
	assert.Nil(t, err)
	err = i.Parse()
	assert.Nil(t, err)
	assert.Equal(t, uint16(15), i.GlobalConfig.Top.CycleTime)
}

func TestCreateSocket(t *testing.T) {
	// not parallelizable since it uses a Unix sock
	i, err := New("test", "")
	assert.Nil(t, err)

	err = i.Parse()
	assert.Nil(t, err)

	err = i.openSocket("tmp.sock")
	assert.Nil(t, err)
	assert.NotNil(t, i.Socket)
	defer i.Socket.Close()
}
