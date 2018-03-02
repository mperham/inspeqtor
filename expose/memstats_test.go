package expose

import (
	"io/ioutil"
	"testing"

	"github.com/mperham/inspeqtor/metrics"
	"github.com/stretchr/testify/assert"
)

func TestBadMemstatsConfig(t *testing.T) {
	t.Parallel()
	src, err := buildMemstatsSource(map[string]string{"port": "885u"})
	assert.Nil(t, src)
	assert.NotNil(t, err)

	src, err = buildMemstatsSource(map[string]string{"port": "8123"})
	assert.Nil(t, err)
	assert.NotNil(t, src)
	assert.Equal(t, "http://localhost:8123/debug/vars", src.(*memstatsSource).Location())

	x, ok := src.(metrics.MandatorySource)
	assert.True(t, ok)
	assert.True(t, x.Mandatory())
}

func TestMemstatsCollection(t *testing.T) {
	t.Parallel()
	rs := testMemstatsSource(nil)
	rs.client = testMemstatsClient("fixtures/expvar.output.json")
	assert.NotNil(t, rs)
	hash, err := rs.Capture()
	assert.Nil(t, err)
	assert.NotNil(t, hash)

	assert.Equal(t, metrics.Map{"PauseTotalNs": 2.567713e+06, "TotalAlloc": 2.050072e+06,
		"HeapAlloc": 1.254184e+06, "NumGC": 7, "Alloc": 1.254184e+06}, hash)
}

func testMemstatsSource(metrics []string) *memstatsSource {
	src, err := buildMemstatsSource(map[string]string{})
	if err != nil {
		panic(err)
	}
	return src.(*memstatsSource)
}

func testMemstatsClient(path string) func(string) ([]byte, error) {
	return func(url string) ([]byte, error) {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
}
