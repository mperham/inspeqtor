package daemon

import (
	"errors"
	"strconv"
	"strings"

	"github.com/mperham/inspeqtor/util"
)

type RedisSource struct {
	Hostname string
	Port     string
	Socket   string
	Password string
	metrics  map[string]bool
	args     []string
}

func (rs *RedisSource) Name() string {
	return "redis"
}

func (rs *RedisSource) Watch(metricName string) {
	rs.metrics[metricName] = true
}

func (rs *RedisSource) Capture() (MetricMap, error) {
	return rs.runCli(execCmd)
}

func (rs *RedisSource) ValidMetrics() []Metric {
	return redisMetrics
}

func (rs *RedisSource) Prepare() error {
	return nil
}

func (rs *RedisSource) runCli(funk executor) (MetricMap, error) {
	sout, err := funk("redis-cli", rs.buildArgs(), nil)
	lines, err := util.ReadLines(sout)
	if err != nil {
		return nil, err
	}

	values := map[string]float64{}

	for _, line := range lines {
		if line == "" || line[0] == '#' {
			continue
		}
		parts := strings.Split(line, ":")
		if rs.metrics[parts[0]] {
			val, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				return nil, errors.New("Invalid metric input for '" + line + "': " + err.Error())
			}
			values[parts[0]] = float64(val)
		}
	}

	if len(rs.metrics) > len(values) {
		for k := range rs.metrics {
			if _, ok := values[k]; !ok {
				util.Warn("Could not find metric redis(%s), did you spell it right?", k)
			}
		}
	}

	return values, nil
}

func (rs *RedisSource) buildArgs() []string {
	if rs.args == nil {
		args := []string{}
		if rs.Socket != "" {
			args = append(args, "-s")
			args = append(args, rs.Socket)
		} else {
			args = append(args, "-h")
			args = append(args, rs.Hostname)
			args = append(args, "-p")
			args = append(args, rs.Port)
		}
		if rs.Password != "" {
			args = append(args, "-a")
			args = append(args, rs.Password)
		}
		args = append(args, "info")
		rs.args = args
	}
	return rs.args
}

func buildRedisSource(params map[string]string) (Collector, error) {
	rs := &RedisSource{"localhost", "6379", "", "", map[string]bool{}, nil}
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
		case "password":
			rs.Password = v
		case "socket":
			rs.Socket = v
		}
	}
	return rs, nil
}

var (
	redisMetrics = []Metric{
		Metric{"connected_clients", g, nil, nil},
		Metric{"client_longest_output_list", g, nil, nil},
		Metric{"client_biggest_input_buf", g, nil, nil},
		Metric{"blocked_clients", g, nil, nil},
		Metric{"used_memory", g, inMB, nil},
		Metric{"used_memory_rss", g, inMB, nil},
		Metric{"used_memory_peak", g, inMB, nil},
		Metric{"used_memory_lua", g, inMB, nil},
		Metric{"rdb_changes_since_last_save", g, nil, nil},
		Metric{"rdb_last_bgsave_time_sec", g, nil, nil},
		Metric{"rdb_current_bgsave_time_sec", g, nil, nil},
		Metric{"aof_last_rewrite_time_sec", g, nil, nil},
		Metric{"aof_current_rewrite_time_sec", g, nil, nil},
		Metric{"total_connections_received", c, nil, nil},
		Metric{"total_commands_processed", c, nil, nil},
		Metric{"instantaneous_ops_per_sec", g, nil, nil},
		Metric{"rejected_connections", c, nil, nil},
		Metric{"sync_full", c, nil, nil},
		Metric{"sync_partial_ok", c, nil, nil},
		Metric{"sync_partial_err", c, nil, nil},
		Metric{"expired_keys", c, nil, nil},
		Metric{"evicted_keys", c, nil, nil},
		Metric{"keyspace_hits", c, nil, nil},
		Metric{"keyspace_misses", c, nil, nil},
		Metric{"pubsub_channels", c, nil, nil},
		Metric{"pubsub_patterns", c, nil, nil},
		Metric{"latest_fork_usec", g, nil, nil},
		Metric{"master_last_io_seconds_ago", g, nil, nil},
		Metric{"connected_slaves", g, nil, nil},
		Metric{"repl_backlog_size", g, inMB, nil},
	}
)
