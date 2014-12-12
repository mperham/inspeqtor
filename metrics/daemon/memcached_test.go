package daemon

import (
	"testing"

	"github.com/mperham/inspeqtor/metrics"
	"github.com/stretchr/testify/assert"
)

func TestBadMemcachedConfig(t *testing.T) {
	t.Parallel()
	src, err := buildMemcachedSource(map[string]string{"port": "885u"})
	assert.Nil(t, src)
	assert.NotNil(t, err)

	src, err = buildMemcachedSource(map[string]string{"port": "22122"})
	assert.Nil(t, err)
	assert.NotNil(t, src)
}

func TestMemcachedCollection(t *testing.T) {
	t.Parallel()
	rs := testMemcachedSource(nil)
	assert.NotNil(t, rs)
	hash, err := rs.runCli(testExec("fixtures/memcached.output.txt"))
	assert.Nil(t, err)
	assert.NotNil(t, hash)

	assert.Equal(t, metrics.Map{"curr_connections": 15, "total_items": 0}, hash)

	rs.Watch("bad_metric")
	hash, err = rs.runCli(testExec("fixtures/memcached.output.txt"))
	assert.Nil(t, err)
	assert.NotNil(t, hash)

	assert.Equal(t, metrics.Map{"curr_connections": 15, "total_items": 0}, hash)
}

func TestRealMemcachedConnection(t *testing.T) {
	t.Parallel()
	rs := testMemcachedSource(nil)
	assert.NotNil(t, rs)
	hash, err := rs.Capture()
	assert.Nil(t, err)
	assert.NotNil(t, hash)

	assert.True(t, hash["curr_connections"] > 0, "This test will fail if you don't have memcached installed")
}

func testMemcachedSource(metrics []string) *memcachedSource {
	src, err := buildMemcachedSource(map[string]string{})
	if err != nil {
		panic(err)
	}
	if metrics == nil {
		metrics = []string{"curr_connections", "total_items"}
	}
	for _, x := range metrics {
		src.Watch(x)
	}
	return src.(*memcachedSource)
}
