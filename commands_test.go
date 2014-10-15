package inspeqtor

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"regexp"
	"testing"
	"time"

	"github.com/mperham/inspeqtor/metrics"
	"github.com/mperham/inspeqtor/services"
	"github.com/mperham/inspeqtor/util"
	"github.com/stretchr/testify/assert"
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
	proc := CommandHandlers["start"]
	proc(i, []string{}, resp)

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
	proc := CommandHandlers["finish"]
	proc(i, []string{}, resp)

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
	proc := CommandHandlers["♡"]
	proc(i, []string{}, resp)

	output, err := resp.ReadString('\n')
	assert.Nil(t, err)
	assert.Equal(t, "Awwww, I love you too.\n", output)
}

func TestStatus(t *testing.T) {
	t.Parallel()
	i, err := New("_", "")
	i.Services = []Checkable{
		&Service{&Entity{"foo", nil, metrics.NewProcessStore("/proc", 15), nil}, nil, &services.ProcessStatus{99, services.Up}, nil},
	}

	var resp bytes.Buffer

	assert.Nil(t, err)
	proc := CommandHandlers["status"]
	proc(i, []string{}, &resp)

	line, err := resp.ReadString('\n')
	assert.Nil(t, err)

	idxs := regexp.MustCompile(fmt.Sprintf("\\AInspeqtor %s, uptime: ", VERSION)).FindStringIndex(line)
	assert.NotNil(t, idxs)
	assert.Equal(t, 0, idxs[0])
}

type mockDisplayable struct {
	*util.RingBuffer
}

func (*mockDisplayable) Displayable(val float64) string {
	return fmt.Sprintf("%.2fm", val)
}

func TestMetricParse(t *testing.T) {
	t.Parallel()

	var f, n string
	f, n = parseMetric("memory(rss)")
	assert.Equal(t, f, "memory")
	assert.Equal(t, n, "rss")
	f, n = parseMetric("memory:rss")
	assert.Equal(t, f, "memory")
	assert.Equal(t, n, "rss")
	f, n = parseMetric("cpu")
	assert.Equal(t, f, "cpu")
	assert.Equal(t, n, "")
}

func TestSparkline(t *testing.T) {
	t.Parallel()

	i, err := New("_", "")
	assert.Nil(t, err)

	buf := util.NewRingBuffer(120)
	for i := 1; i <= 100; i++ {
		buf.Add(float64(i))
	}
	src := &mockDisplayable{buf}

	output := buildSparkline(i.Host, "memory:rss", func(family, name string) displayable {
		return src
	})

	expected := "localhost memory:rss min 1.00m max 100.00m avg 50.50m\n▁▁▁▁▁▁▁▁▁▁▁▁▁▂▂▂▂▂▂▂▂▂▂▂▂▃▃▃▃▃▃▃▃▃▃▃▃▃▄▄▄▄▄▄▄▄▄▄▄▄▅▅▅▅▅▅▅▅▅▅▅▅▆▆▆▆▆▆▆▆▆▆▆▆▆▇▇▇▇▇▇▇▇▇▇▇▇█████████████\n"
	assert.Equal(t, output, expected)

	// alternate, CLI friendy metric naming format
	expected = "localhost memory:rss min 1.00m max 100.00m avg 50.50m\n▁▁▁▁▁▁▁▁▁▁▁▁▁▂▂▂▂▂▂▂▂▂▂▂▂▃▃▃▃▃▃▃▃▃▃▃▃▃▄▄▄▄▄▄▄▄▄▄▄▄▅▅▅▅▅▅▅▅▅▅▅▅▆▆▆▆▆▆▆▆▆▆▆▆▆▇▇▇▇▇▇▇▇▇▇▇▇█████████████\n"
	output = buildSparkline(i.Host, "memory:rss", func(family, name string) displayable {
		return src
	})
	assert.Equal(t, output, expected)

	// invalid metric
	output = buildSparkline(i.Host, "memory:xxx", func(family, name string) displayable {
		return nil
	})
	assert.Equal(t, output, "Unknown metric: memory:xxx\n")
}
