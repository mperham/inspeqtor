package daemon

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestBadMysqlConfig(t *testing.T) {
	t.Parallel()
	src, err := Sources["mysql"](map[string]string{"port": "885u"})
	assert.Nil(t, src)
	assert.NotNil(t, err)

	src, err = Sources["mysql"](map[string]string{"socket": "/foo/bar.sock", "password": "fuzzy"})
	assert.Nil(t, err)
	assert.NotNil(t, src)
}

func TestMysqlCollection(t *testing.T) {
	t.Parallel()
	rs := mysqlSource()
	assert.NotNil(t, rs)
	rs.Watch("Seconds_Behind_Master")
	assert.True(t, rs.metrics["Seconds_Behind_Master"])

	err := rs.Prepare(testExec("fixtures/mysql.slave.running.txt"))
	assert.Nil(t, err)
	assert.True(t, rs.captureRepl)
	_, ok := rs.metrics["Seconds_Behind_Master"]
	assert.False(t, ok)

	hash, err := rs.runStatus(testExec("fixtures/mysql.output.txt"))
	assert.Nil(t, err)
	assert.NotNil(t, hash)

	hash, err = rs.runRepl(hash, testExec("fixtures/mysql.slave.status.txt"))
	assert.Nil(t, err)
	assert.NotNil(t, hash)

	assert.Equal(t, metricMap{"Connections": 20, "Queries": 62, "Table_locks_waited": 0, "Seconds_Behind_Master": 12}, hash)

	rs.Watch("bad_metric")
	hash, err = rs.runStatus(testExec("fixtures/mysql.output.txt"))
	assert.Nil(t, err)
	assert.NotNil(t, hash)

	assert.Equal(t, metricMap{"Connections": 20, "Queries": 62, "Table_locks_waited": 0}, hash)
}

func TestRealMysqlConnection(t *testing.T) {
	t.Parallel()
	rs := mysqlSource("Connections", "Seconds_Behind_Master")
	err := rs.Prepare(execCmd)
	assert.NotNil(t, err, "This test will fail if you don't have mysql installed")
	assert.True(t, strings.Contains(err.Error(), "slave not running"))

	rs = mysqlSource()
	assert.NotNil(t, rs)
	err = rs.Prepare(execCmd)
	assert.Nil(t, err)
	hash, err := rs.Capture()
	assert.Nil(t, err)
	assert.NotNil(t, hash)

	assert.True(t, hash["Connections"] > 0, "This test will fail if you don't have mysql installed")
}

func mysqlSource(metrics ...string) *MysqlSource {
	src, err := Sources["mysql"](map[string]string{})
	if err != nil {
		panic(err)
	}
	if len(metrics) == 0 {
		metrics = []string{"Connections", "Queries", "Table_locks_waited"}
	}
	for _, x := range metrics {
		src.Watch(x)
	}
	return src.(*MysqlSource)
}
