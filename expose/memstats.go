package expose

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/mperham/inspeqtor/metrics"
	"github.com/mperham/inspeqtor/util"
)

func init() {
	metrics.Sources["memstats"] = buildMemstatsSource
}

var (
	c    = metrics.Counter
	g    = metrics.Gauge
	inMB = metrics.DisplayInMB
)

// Gathems runtime.MemStats data from a running Go process.
// This can be exposed by the `expvar` stdlib package.
type memstatsSource struct {
	Hostname string
	Port     string
	Path     string
	client   func(string) ([]byte, error)
	url      string
}

func (ms *memstatsSource) Location() string {
	if ms.url == "" {
		ms.url = fmt.Sprintf("http://%s:%s%s", ms.Hostname, ms.Port, ms.Path)
	}
	return ms.url
}

func (ms *memstatsSource) Name() string {
	return "memstats"
}

func (ms *memstatsSource) Prepare() error {
	return nil
}

func (ms *memstatsSource) Mandatory() bool {
	return true
}

func (ms *memstatsSource) Watch(name string) {
}

func (ms *memstatsSource) ValidMetrics() []metrics.Descriptor {
	return memstatsMetrics
}

func (ms *memstatsSource) Capture() (metrics.Map, error) {
	sout, err := ms.client(ms.Location())
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}

	err = json.Unmarshal(sout, &data)
	if err != nil {
		return nil, err
	}

	values := map[string]float64{}
	memstats := data["memstats"].(map[string]interface{})

	for _, metric := range memstatsMetrics {
		val, ok := memstats[metric.Name]
		if !ok {
			return nil, fmt.Errorf("No such memstats metric: %s", metric.Name)
		}
		values[metric.Name] = val.(float64)
	}

	return values, nil
}

func defaultMemstatsClient(url string) ([]byte, error) {
	util.Debug("Fetching memstats from %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func buildMemstatsSource(params map[string]string) (metrics.Source, error) {
	ms := &memstatsSource{"localhost", "8080", "/debug/vars", defaultMemstatsClient, ""}
	for k, v := range params {
		switch k {
		case "port":
			_, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				return nil, err
			}
			ms.Port = v
		case "path":
			ms.Path = v
		}
	}
	return ms, nil
}

var (
	memstatsMetrics = []metrics.Descriptor{
		metrics.D("Alloc", g, inMB, nil),
		metrics.D("TotalAlloc", c, inMB, nil),
		metrics.D("HeapAlloc", g, inMB, nil),
		metrics.D("NumGC", c, nil, nil),
		metrics.D("PauseTotalNs", c, nil, nil),
	}
)
