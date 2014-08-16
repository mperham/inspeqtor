package inspeqtor

import (
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestInspeqtorParse(t *testing.T) {
	t.Parallel()
	i, err := New("test")
	assert.Nil(t, err)
	err = i.Parse()
	assert.Nil(t, err)
	assert.Equal(t, uint16(15), i.GlobalConfig.Top.CycleTime)
}

func TestCreateSocket(t *testing.T) {
	i, err := New("test")
	assert.Nil(t, err)

	err = i.Parse()
	assert.Nil(t, err)

	sock, err := i.openSocket("tmp.sock")
	assert.Nil(t, err)
	defer sock.Close()

	go func() {
		conn, err := net.Dial("unix", "tmp.sock")
		assert.Nil(t, err)
		conn.Write([]byte("start deploy"))
		conn.Close()

		conn, err = net.Dial("unix", "tmp.sock")
		assert.Nil(t, err)
		conn.Write([]byte("finish deploy"))
		conn.Close()
	}()

	assert.False(t, i.silenced())
	i.acceptCommand()
	assert.True(t, i.silenced())
	i.acceptCommand()
	assert.False(t, i.silenced())
}
