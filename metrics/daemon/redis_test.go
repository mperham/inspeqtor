package daemon

import (
	"testing"

	"github.com/mperham/inspeqtor/metrics"
	"github.com/stretchr/testify/assert"
)

func TestBadRedisConfig(t *testing.T) {
	t.Parallel()
	src, err := metrics.Sources["redis"](map[string]string{"port": "885u"})
	assert.Nil(t, src)
	assert.NotNil(t, err)

	src, err = metrics.Sources["redis"](map[string]string{"socket": "/foo/bar.sock", "password": "fuzzy"})
	assert.Nil(t, err)
	assert.NotNil(t, src)
}

func TestExec(t *testing.T) {
	t.Parallel()
	sout, err := execCmd("/bin/echo", []string{"mike", "perham"}, nil)
	assert.Nil(t, err)
	assert.Equal(t, "mike perham\n", string(sout))
}

func TestRedisCollection(t *testing.T) {
	t.Parallel()
	rs := testRedisSource(nil)
	assert.NotNil(t, rs)
	hash, err := rs.runCli(testExec("fixtures/redis.6379.txt"))
	assert.Nil(t, err)
	assert.NotNil(t, hash)

	assert.Equal(t, metrics.Map{"connected_clients": 3, "latest_fork_usec": 758, "master_repl_offset": 0}, hash)

	rs.Watch("bad_metric")
	hash, err = rs.runCli(testExec("fixtures/redis.6379.txt"))
	assert.Nil(t, err)
	assert.NotNil(t, hash)

	assert.Equal(t, metrics.Map{"connected_clients": 3, "latest_fork_usec": 758, "master_repl_offset": 0}, hash)
}

func TestRealRedisConnection(t *testing.T) {
	t.Parallel()
	rs := testRedisSource(nil)
	assert.NotNil(t, rs)
	hash, err := rs.Capture()
	assert.Nil(t, err)
	assert.NotNil(t, hash)

	assert.True(t, hash["connected_clients"] > 0, "This test will fail if you don't have redis-cli installed")
}

func testRedisSource(mets []string) *redisSource {
	src, err := metrics.Sources["redis"](map[string]string{})
	if err != nil {
		panic(err)
	}
	if mets == nil {
		mets = []string{"latest_fork_usec", "connected_clients", "master_repl_offset"}
	}
	for _, x := range mets {
		src.Watch(x)
	}
	return src.(*redisSource)
}
