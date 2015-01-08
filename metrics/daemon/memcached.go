package daemon

import (
	"errors"
	"strconv"
	"strings"

	"github.com/mperham/inspeqtor/metrics"
	"github.com/mperham/inspeqtor/util"
)

func init() {
	metrics.Sources["memcached"] = buildMemcachedSource
}

type memcachedSource struct {
	Hostname string
	Port     string
	metrics  map[string]bool
	args     []string
}

func (rs *memcachedSource) Name() string {
	return "memcached"
}

func (rs *memcachedSource) Watch(metricName string) {
	rs.metrics[metricName] = true
}

func (rs *memcachedSource) Capture() (metrics.Map, error) {
	return rs.runCli(execCmd)
}

func (rs *memcachedSource) Prepare() error {
	return nil
}

func (rs *memcachedSource) ValidMetrics() []metrics.Descriptor {
	return memcachedMetrics
}

func (rs *memcachedSource) runCli(funk executor) (metrics.Map, error) {
	sout, err := funk("nc", []string{rs.Hostname, rs.Port}, []byte("stats\n"))
	if err != nil {
		return nil, err
	}
	lines, err := util.ReadLines(sout)
	if err != nil {
		return nil, err
	}

	values := map[string]float64{}

	for _, line := range lines {
		if line == "" || line[0] != 'S' {
			continue
		}
		parts := strings.Fields(line)
		if rs.metrics[parts[1]] {
			val, err := strconv.ParseFloat(parts[2], 64)
			if err != nil {
				return nil, errors.New("Invalid metric input for '" + line + "': " + err.Error())
			}
			values[parts[1]] = val
		}
	}

	if len(rs.metrics) > len(values) {
		for k := range rs.metrics {
			if _, ok := values[k]; !ok {
				util.Info("Could not find metric %s(%s), did you spell it right?", rs.Name(), k)
			}
		}
	}

	return values, nil
}

func buildMemcachedSource(params map[string]string) (metrics.Source, error) {
	rs := &memcachedSource{"localhost", "11211", map[string]bool{}, nil}
	for k, v := range params {
		switch k {
		case "hostname":
			rs.Hostname = v
		case "port":
			_, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				return nil, err
			}
			rs.Port = v
		}
	}
	return rs, nil
}

var (
	memcachedMetrics = []metrics.Descriptor{
		metrics.D("curr_connections", g, nil, nil),
		metrics.D("total_connections", c, nil, nil),
		metrics.D("cmd_get", c, nil, nil),
		metrics.D("cmd_set", c, nil, nil),
		metrics.D("cmd_flush", c, nil, nil),
		metrics.D("cmd_touch", c, nil, nil),
		metrics.D("get_hits", c, nil, nil),
		metrics.D("get_misses", c, nil, nil),
		metrics.D("delete_hits", c, nil, nil),
		metrics.D("delete_misses", c, nil, nil),
		metrics.D("incr_hits", c, nil, nil),
		metrics.D("incr_misses", c, nil, nil),
		metrics.D("decr_hits", c, nil, nil),
		metrics.D("decr_misses", c, nil, nil),
		metrics.D("cas_hits", c, nil, nil),
		metrics.D("cas_misses", c, nil, nil),
		metrics.D("cas_badval", c, nil, nil),
		metrics.D("touch_hits", c, nil, nil),
		metrics.D("touch_misses", c, nil, nil),
		metrics.D("auth_cmds", c, nil, nil),
		metrics.D("auth_errors", c, nil, nil),
		metrics.D("bytes_read", c, inMB, nil),
		metrics.D("bytes_written", c, inMB, nil),
		metrics.D("threads", g, nil, nil),
		metrics.D("malloc_fails", c, nil, nil),
		metrics.D("bytes", g, nil, nil),
		metrics.D("curr_items", g, nil, nil),
		metrics.D("total_items", c, nil, nil),
		metrics.D("expired_unfetched", c, nil, nil),
		metrics.D("evicted_unfetched", c, nil, nil),
		metrics.D("evictions", c, nil, nil),
		metrics.D("reclaimed", c, nil, nil),
		metrics.D("crawler_reclaimed", c, nil, nil),
	}
)
