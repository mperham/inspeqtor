package daemon

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBadPostgresqlConfig(t *testing.T) {
	t.Parallel()
	src, err := Sources["postgresql"](map[string]string{"port": "885u"})
	assert.Nil(t, src)
	assert.NotNil(t, err)
}

func TestPostgresqlCollection(t *testing.T) {
	t.Parallel()
	rs := psqlSource()
	assert.NotNil(t, rs)

	rs.execFunk = testExec("fixtures/pg.dbStats.output.txt")
	data := metricMap{}
	err := populate(rs, data, "deadlocks")
	assert.Nil(t, err)
	assert.Equal(t, metricMap{"blk_hit_rate": 99.55268047622324}, data)

	rs.execFunk = testExec("fixtures/pg.sizeStats.output.txt")
	err = populate(rs, data, "total_size")
	assert.Nil(t, err)
	assert.Equal(t, metricMap{"blk_hit_rate": 99.55268047622324, "total_size": 122880.0}, data)

	rs.execFunk = testExec("fixtures/pg.userStats.output.txt")
	err = populate(rs, data, "seq_scans")
	assert.Nil(t, err)
	assert.Equal(t, metricMap{"blk_hit_rate": 99.55268047622324, "total_size": 122880.0, "seq_scans": 6.0}, data)

	err = populate(rs, data, "bad_metric")
	assert.NotNil(t, err)
}

func TestRealPostgresqlConnection(t *testing.T) {
	t.Parallel()

	/*
		To get this test running locally, you must have a postgres user:
		  createuser -s postgres
	*/

	rs := psqlSource()
	assert.NotNil(t, rs)
	err := rs.Prepare(execCmd)
	assert.Nil(t, err)
	hash, err := rs.Capture()
	assert.Nil(t, err)
	assert.NotNil(t, hash)

	fmt.Printf("%v", hash)

	assert.True(t, hash["blk_hit_rate"] > 0, "This test will fail if you don't have postgresql installed")
}

func psqlSource(metrics ...string) *pgSource {
	src, err := Sources["postgresql"](map[string]string{})
	if err != nil {
		panic(err)
	}
	if len(metrics) == 0 {
		metrics = []string{"blk_hit_rate", "total_size", "seq_scans"}
	}
	for _, x := range metrics {
		src.Watch(x)
	}
	return src.(*pgSource)
}
