package inspeqtor

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"inspeqtor/metrics"
	"inspeqtor/services"
	"io"
	"net"
	"regexp"
	"testing"
	"time"
)

func TestAcceptSocket(t *testing.T) {
	// not parallelizable since it uses a Unix sock
	i, err := New("test", "")
	assert.Nil(t, err)

	err = i.Parse()
	assert.Nil(t, err)

	err = i.openSocket("/tmp/tmp.sock")
	assert.Nil(t, err)
	defer i.Socket.Close()

	go func() {
		conn, err := net.Dial("unix", "/tmp/tmp.sock")
		assert.Nil(t, err)
		conn.Write([]byte("start deploy\n"))
		conn.Close()

		conn, err = net.Dial("unix", "/tmp/tmp.sock")
		assert.Nil(t, err)
		conn.Write([]byte("finish deploy\n"))
		conn.Close()

		conn, err = net.Dial("unix", "/tmp/tmp.sock")
		assert.Nil(t, err)
		conn.Write([]byte("?\n"))
		buf := make([]byte, 19)
		_, err = io.ReadFull(conn, buf)
		assert.Nil(t, err)
		conn.Close()

		assert.Equal(t, "Unknown command: ?\n", string(buf))
	}()

	assert.False(t, i.silenced())
	i.acceptCommand()
	assert.True(t, i.silenced())
	i.acceptCommand()
	assert.False(t, i.silenced())
	i.acceptCommand()
}

func TestStartDeploy(t *testing.T) {
	t.Parallel()
	i, err := New("_", "")
	i.SilenceUntil = time.Now()

	outbuf := make([]byte, 0)
	resp := bytes.NewBuffer(outbuf)

	assert.Nil(t, err)
	proc := CommandHandlers['s']
	proc(i, resp)

	assert.True(t, i.SilenceUntil.After(time.Now()))
	assert.True(t, i.silenced())
	output, err := resp.ReadString('\n')
	assert.Nil(t, err)
	assert.Equal(t, "Starting deploy, now silenced\n", output)
}

func TestFinishDeploy(t *testing.T) {
	t.Parallel()
	i, err := New("_", "")

	outbuf := make([]byte, 0)
	resp := bytes.NewBuffer(outbuf)

	assert.Nil(t, err)
	proc := CommandHandlers['f']
	proc(i, resp)

	assert.True(t, i.SilenceUntil.Before(time.Now()))
	output, err := resp.ReadString('\n')
	assert.Nil(t, err)
	assert.Equal(t, "Finished deploy, volume turned to 11\n", output)
}

func TestTheLove(t *testing.T) {
	t.Parallel()
	i, err := New("_", "")

	outbuf := make([]byte, 0)
	resp := bytes.NewBuffer(outbuf)

	assert.Nil(t, err)
	proc := CommandHandlers['♡']
	proc(i, resp)

	output, err := resp.ReadString('\n')
	assert.Nil(t, err)
	assert.Equal(t, "Awwww, I love you too.\n", output)
}

func TestInfo(t *testing.T) {
	t.Parallel()
	i, err := New("_", "")
	i.Services = []Checkable{
		&Service{&Entity{"foo", nil, metrics.NewProcessStore("/proc", 15), nil}, nil, &services.ProcessStatus{99, services.Up}, nil},
	}

	var resp bytes.Buffer

	assert.Nil(t, err)
	proc := CommandHandlers['i']
	proc(i, &resp)

	line, err := resp.ReadString('\n')
	assert.Nil(t, err)

	idxs := regexp.MustCompile(fmt.Sprintf("\\AInspeqtor %s, uptime: ", VERSION)).FindStringIndex(line)
	assert.NotNil(t, idxs)
	assert.Equal(t, 0, idxs[0])
}
