package daemon

import (
	"fmt"
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
	rs := testMysqlSource(map[string]string{})
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
	rs := testMysqlSource(map[string]string{}, "Connections", "Seconds_Behind_Master")
	err := rs.Prepare()
	assert.NotNil(t, err, "This test will fail if you don't have mysql installed")
	assert.True(t, strings.Contains(err.Error(), "slave not running"))

	rs = testMysqlSource(map[string]string{})
	assert.NotNil(t, rs)
	err = rs.Prepare()
	assert.Nil(t, err)
	hash, err := rs.Capture()
	assert.Nil(t, err)
	assert.NotNil(t, hash)

	assert.True(t, hash["Connections"] > 0, "This test will fail if you don't have mysql installed")
}

func TestRealMysqlConnectionWithPassword(t *testing.T) {
	// This test doesn't run in parallel with the other test as it changes the password from the default.
	assert.True(t, changeMysqlPassword("", "password   .test.%$'#_(@*[|~` &"), "This test will fail if your database has a password set")
	rs := testMysqlSource(map[string]string{"password": "password   .test.%$'#_(@*[|~` &"}, "Connections", "Seconds_Behind_Master")
	err := rs.Prepare()
	assert.NotNil(t, err, "This test will fail if you don't have mysql installed")
	assert.True(t, strings.Contains(err.Error(), "slave not running"))
	rs = testMysqlSource(map[string]string{"password": "password   .test.%$'#_(@*[|~` &"})
	assert.NotNil(t, rs)
	err = rs.Prepare()
	assert.Nil(t, err)
	hash, err := rs.Capture()
	assert.Nil(t, err)
	assert.NotNil(t, hash)

	assert.True(t, changeMysqlPassword("password   .test.%$'#_(@*[|~` &", "password.test.%$'#_(@*[|~`&"), "This test will fail if your database has a password set")
	rs = testMysqlSource(map[string]string{"password": "password.test.%$'#_(@*[|~`&"}, "Connections", "Seconds_Behind_Master")
	err = rs.Prepare()
	assert.NotNil(t, err, "This test will fail if you don't have mysql installed")
	assert.True(t, strings.Contains(err.Error(), "slave not running"))
	rs = testMysqlSource(map[string]string{"password": "password.test.%$'#_(@*[|~`&"})
	assert.NotNil(t, rs)
	err = rs.Prepare()
	assert.Nil(t, err)
	hash, err = rs.Capture()
	assert.Nil(t, err)
	assert.NotNil(t, hash)

	assert.True(t, hash["Connections"] > 0, "This test will fail if you don't have mysql installed")
	assert.True(t, changeMysqlPassword("password.test.%$'#_(@*[|~`&", ""), "This test will fail if your database has a password set")

	rs = testMysqlSource(map[string]string{}, "Connections", "Seconds_Behind_Master")
	err = rs.Prepare()
	assert.NotNil(t, err, "This test will fail if you don't have mysql installed")
	assert.True(t, strings.Contains(err.Error(), "slave not running"))

	rs = testMysqlSource(map[string]string{})
	assert.NotNil(t, rs)
	err = rs.Prepare()
	assert.Nil(t, err)
	hash, err = rs.Capture()
	assert.Nil(t, err)
	assert.NotNil(t, hash)

	assert.True(t, hash["Connections"] > 0, "This test will fail if you don't have mysql installed")
}

func testMysqlSource(connectionStrings map[string]string, mets ...string) *mysqlSource {
	src, err := metrics.Sources["mysql"](connectionStrings)
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

func changeMysqlPassword(currentPassword string, newPassword string) bool {
	args := []string{"-B"}
	args = append(args, "-u")
	args = append(args, "root")
	if currentPassword != "" {
		args = append(args, fmt.Sprintf("-p%s", currentPassword))
	}
	args = append(args, "-e")
	args = append(args, fmt.Sprintf("SET PASSWORD FOR 'root'@'localhost' = PASSWORD(\"%s\")", strings.Replace(newPassword, "'", "\\'", -1)))
	_, err := execCmd("mysql", args, nil)
	return (err == nil)
}
