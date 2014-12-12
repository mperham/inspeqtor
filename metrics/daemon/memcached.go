package daemon

import (
	"errors"
	"strconv"
	"strings"

	"github.com/mperham/inspeqtor/util"
)

type MemcachedSource struct {
	Hostname string
	Port     string
	metrics  map[string]bool
	args     []string
}

func (rs *MemcachedSource) Name() string {
	return "memcached"
}

func (rs *MemcachedSource) Watch(metricName string) {
	rs.metrics[metricName] = true
}

func (rs *MemcachedSource) Capture() (MetricMap, error) {
	return rs.runCli(execCmd)
}

func (rs *MemcachedSource) Prepare() error {
	return nil
}

func (rs *MemcachedSource) ValidMetrics() []Metric {
	return memcachedMetrics
}

func (rs *MemcachedSource) runCli(funk executor) (MetricMap, error) {
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

func buildMemcachedSource(params map[string]string) (Collector, error) {
	rs := &MemcachedSource{"localhost", "11211", map[string]bool{}, nil}
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
	memcachedMetrics = []Metric{
		Metric{"curr_connections", g, nil, nil},
		Metric{"total_connections", c, nil, nil},
		Metric{"cmd_get", c, nil, nil},
		Metric{"cmd_set", c, nil, nil},
		Metric{"cmd_flush", c, nil, nil},
		Metric{"cmd_touch", c, nil, nil},
		Metric{"get_hits", c, nil, nil},
		Metric{"get_misses", c, nil, nil},
		Metric{"delete_hits", c, nil, nil},
		Metric{"delete_misses", c, nil, nil},
		Metric{"incr_hits", c, nil, nil},
		Metric{"incr_misses", c, nil, nil},
		Metric{"decr_hits", c, nil, nil},
		Metric{"decr_misses", c, nil, nil},
		Metric{"cas_hits", c, nil, nil},
		Metric{"cas_misses", c, nil, nil},
		Metric{"cas_badval", c, nil, nil},
		Metric{"touch_hits", c, nil, nil},
		Metric{"touch_misses", c, nil, nil},
		Metric{"auth_cmds", c, nil, nil},
		Metric{"auth_errors", c, nil, nil},
		Metric{"bytes_read", c, inMB, nil},
		Metric{"bytes_written", c, inMB, nil},
		Metric{"threads", g, nil, nil},
		Metric{"malloc_fails", c, nil, nil},
		Metric{"bytes", g, nil, nil},
		Metric{"curr_items", g, nil, nil},
		Metric{"total_items", c, nil, nil},
		Metric{"expired_unfetched", c, nil, nil},
		Metric{"evicted_unfetched", c, nil, nil},
		Metric{"evictions", c, nil, nil},
		Metric{"reclaimed", c, nil, nil},
		Metric{"crawler_reclaimed", c, nil, nil},
	}
)
