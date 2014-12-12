package daemon

import (
	"strings"
	"testing"

	"github.com/mperham/inspeqtor/metrics"
	"github.com/stretchr/testify/assert"
)

func TestBadMysqlConfig(t *testing.T) {
	t.Parallel()
	src, err := metrics.Sources["mysql"](map[string]string{"port": "885u"})
	assert.Nil(t, src)
	assert.NotNil(t, err)

	src, err = metrics.Sources["mysql"](map[string]string{"socket": "/foo/bar.sock", "password": "fuzzy"})
	assert.Nil(t, err)
	assert.NotNil(t, src)
}

func TestMysqlCollection(t *testing.T) {
	t.Parallel()
	rs := testMysqlSource()
	assert.NotNil(t, rs)
	rs.Watch("Seconds_Behind_Master")
	assert.True(t, rs.metrics["Seconds_Behind_Master"])

	rs.exec = testExec("fixtures/mysql.slave.running.txt")
	err := rs.Prepare()
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

	assert.Equal(t, metrics.Map{"Connections": 20, "Queries": 62, "Table_locks_waited": 0, "Seconds_Behind_Master": 12}, hash)

	rs.Watch("bad_metric")
	hash, err = rs.runStatus(testExec("fixtures/mysql.output.txt"))
	assert.Nil(t, err)
	assert.NotNil(t, hash)

	assert.Equal(t, metrics.Map{"Connections": 20, "Queries": 62, "Table_locks_waited": 0}, hash)
}

func TestRealMysqlConnection(t *testing.T) {
	t.Parallel()
	rs := testMysqlSource("Connections", "Seconds_Behind_Master")
	err := rs.Prepare()
	assert.NotNil(t, err, "This test will fail if you don't have mysql installed")
	assert.True(t, strings.Contains(err.Error(), "slave not running"))

	rs = testMysqlSource()
	assert.NotNil(t, rs)
	err = rs.Prepare()
	assert.Nil(t, err)
	hash, err := rs.Capture()
	assert.Nil(t, err)
	assert.NotNil(t, hash)

	assert.True(t, hash["Connections"] > 0, "This test will fail if you don't have mysql installed")
}

func testMysqlSource(mets ...string) *mysqlSource {
	src, err := metrics.Sources["mysql"](map[string]string{})
	if err != nil {
		panic(err)
	}
	if len(mets) == 0 {
		mets = []string{"Connections", "Queries", "Table_locks_waited"}
	}
	for _, x := range mets {
		src.Watch(x)
	}
	return src.(*mysqlSource)
}
