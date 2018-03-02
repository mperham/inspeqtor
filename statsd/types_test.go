package statsd

import (
	"bytes"
	"testing"

	"github.com/mperham/inspeqtor"
	"github.com/stretchr/testify/assert"
)

func TestStatsd(t *testing.T) {

	conn, err := Dial("localhost:9909")
	assert.NotNil(t, conn)
	assert.Nil(t, err)

	conn.Close()

	i, err := inspeqtor.New("../test", "")
	assert.NotNil(t, i)
	assert.Nil(t, err)

	err = i.Parse()
	assert.Nil(t, err)

	var buff bytes.Buffer
	Export(&buff, i)

	actual := string(buff.Bytes())

	assert.Contains(t, actual, "local.host.cpu:0.00|c")
	assert.Contains(t, actual, "local.host.cpu.iowait:0.00|c")
	assert.Contains(t, actual, "local.host.cpu.steal:0.00|c")
	assert.Contains(t, actual, "local.host.cpu.system:0.00|c")
	assert.Contains(t, actual, "local.host.cpu.user:0.00|c")
	assert.Contains(t, actual, "local.host.disk./:-1.00|g")
	assert.Contains(t, actual, "local.host.load.1:-1.00|g")
	assert.Contains(t, actual, "local.host.load.15:-1.00|g")
	assert.Contains(t, actual, "local.host.load.5:-1.00|g")
	assert.Contains(t, actual, "local.host.swap:-1.00|g")
	assert.Contains(t, actual, "local.homebrew.mxcl.memcached.cpu.system:0.00|c")
	assert.Contains(t, actual, "local.homebrew.mxcl.memcached.cpu.total_system:0.00|c")
	assert.Contains(t, actual, "local.homebrew.mxcl.memcached.cpu.total_user:0.00|c")
	assert.Contains(t, actual, "local.homebrew.mxcl.memcached.cpu.user:0.00|c")
	assert.Contains(t, actual, "local.homebrew.mxcl.memcached.memory.rss:-1.00|g")
}
