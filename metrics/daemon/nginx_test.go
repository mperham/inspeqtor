package daemon

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestBadNginxConfig(t *testing.T) {
	t.Parallel()
	src, err := sources["nginx"](map[string]string{"port": "885u"})
	assert.Nil(t, src)
	assert.NotNil(t, err)

	src, err = sources["nginx"](map[string]string{"port": "8080"})
	assert.Nil(t, err)
	assert.NotNil(t, src)
}

func TestNginxCollection(t *testing.T) {
	t.Parallel()
	rs := testNginxSource(nil)
	rs.client = testNginxClient("fixtures/nginx.status.txt")
	assert.NotNil(t, rs)
	hash, err := rs.runCli()
	assert.Nil(t, err)
	assert.NotNil(t, hash)

	assert.Equal(t, metricMap{"Active_connections": 2, "requests": 3, "Waiting": 1}, hash)

	rs.Watch("bad_metric")
	hash, err = rs.runCli()
	assert.Nil(t, err)
	assert.NotNil(t, hash)

	assert.Equal(t, metricMap{"Active_connections": 2, "requests": 3, "Waiting": 1}, hash)
}

func TestRealNginxConnection(t *testing.T) {
	t.Parallel()
	rs := testNginxSource(nil)
	rs.Port = "8080"
	assert.NotNil(t, rs)
	hash, err := rs.Capture()
	assert.Nil(t, err)
	assert.NotNil(t, hash)

	assert.True(t, hash["requests"] > 0, "This test will fail if you don't have nginx installed")
}

func testNginxSource(metrics []string) *nginxSource {
	src, err := sources["nginx"](map[string]string{})
	if err != nil {
		panic(err)
	}
	if metrics == nil {
		metrics = []string{"Active_connections", "requests", "Waiting"}
	}
	for _, x := range metrics {
		src.Watch(x)
	}
	return src.(*nginxSource)
}

func testNginxClient(path string) func(string, string, string) ([]byte, error) {
	return func(host string, port string, ep string) ([]byte, error) {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}
		return data, nil
	}
}
