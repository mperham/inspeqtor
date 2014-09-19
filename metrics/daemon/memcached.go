package daemon

import (
	"errors"
	"github.com/mperham/inspeqtor/util"
	"strconv"
	"strings"
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

func (rs *MemcachedSource) Capture() (metricMap, error) {
	return rs.runCli(execCmd)
}

func (rs *MemcachedSource) Prepare(funk executor) error {
	return nil
}

func (rs *MemcachedSource) ValidMetrics() []metric {
	return memcachedMetrics
}

func (rs *MemcachedSource) runCli(funk executor) (metricMap, error) {
	sout, err := funk("nc", []string{rs.Hostname, rs.Port}, []byte("stats\n"))
	lines, err := util.ReadLines(sout)
	if err != nil {
		return nil, err
	}

	values := map[string]int64{}

	for _, line := range lines {
		if line == "" || line[0] != 'S' {
			continue
		}
		parts := strings.Fields(line)
		if rs.metrics[parts[1]] {
			val, err := strconv.ParseInt(parts[2], 10, 64)
			if err != nil {
				return nil, errors.New("Invalid metric input for '" + line + "': " + err.Error())
			}
			values[parts[1]] = val
		}
	}

	if len(rs.metrics) > len(values) {
		for k, _ := range rs.metrics {
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
	memcachedMetrics []metric = []metric{
		metric{"curr_connections", g, nil},
		metric{"total_connections", c, nil},
		metric{"cmd_get", c, nil},
		metric{"cmd_set", c, nil},
		metric{"cmd_flush", c, nil},
		metric{"cmd_touch", c, nil},
		metric{"get_hits", c, nil},
		metric{"get_misses", c, nil},
		metric{"delete_hits", c, nil},
		metric{"delete_misses", c, nil},
		metric{"incr_hits", c, nil},
		metric{"incr_misses", c, nil},
		metric{"decr_hits", c, nil},
		metric{"decr_misses", c, nil},
		metric{"cas_hits", c, nil},
		metric{"cas_misses", c, nil},
		metric{"cas_badval", c, nil},
		metric{"touch_hits", c, nil},
		metric{"touch_misses", c, nil},
		metric{"auth_cmds", c, nil},
		metric{"auth_errors", c, nil},
		metric{"bytes_read", c, &funcWrapper{nil, inMB, nil}},
		metric{"bytes_written", c, &funcWrapper{nil, inMB, nil}},
		metric{"threads", g, nil},
		metric{"malloc_fails", c, nil},
		metric{"bytes", g, nil},
		metric{"curr_items", g, nil},
		metric{"total_items", c, nil},
		metric{"expired_unfetched", c, nil},
		metric{"evicted_unfetched", c, nil},
		metric{"evictions", c, nil},
		metric{"reclaimed", c, nil},
		metric{"crawler_reclaimed", c, nil},
	}
)
