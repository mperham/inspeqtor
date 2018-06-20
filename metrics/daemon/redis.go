package daemon

import (
	"errors"
	"strconv"
	"strings"

	"github.com/mperham/inspeqtor/metrics"
	"github.com/mperham/inspeqtor/util"
)

func init() {
	metrics.Sources["redis"] = buildRedisSource
}

type redisSource struct {
	Hostname string
	Port     string
	Socket   string
	Password string
	metrics  map[string]bool
	args     []string
}

func (rs *redisSource) Name() string {
	return "redis"
}

func (rs *redisSource) Watch(metricName string) {
	rs.metrics[metricName] = true
}

func (rs *redisSource) Capture() (metrics.Map, error) {
	return rs.runCli(execCmd)
}

func (rs *redisSource) ValidMetrics() []metrics.Descriptor {
	return redisMetrics
}

func (rs *redisSource) Prepare() error {
	return nil
}

func (rs *redisSource) runCli(funk executor) (metrics.Map, error) {
	sout, err := funk("redis-cli", rs.buildArgs(), nil)
	if err != nil {
		return nil, err
	}
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

func (rs *redisSource) buildArgs() []string {
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

func buildRedisSource(params map[string]string) (metrics.Source, error) {
	rs := &redisSource{"localhost", "6379", "", "", map[string]bool{}, nil}
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
	redisMetrics = []metrics.Descriptor{
		metrics.D("connected_clients", g, nil, nil),
		metrics.D("client_longest_output_list", g, nil, nil),
		metrics.D("client_biggest_input_buf", g, nil, nil),
		metrics.D("blocked_clients", g, nil, nil),
		metrics.D("used_memory", g, inMB, nil),
		metrics.D("used_memory_rss", g, inMB, nil),
		metrics.D("used_memory_peak", g, inMB, nil),
		metrics.D("used_memory_lua", g, inMB, nil),
		metrics.D("rdb_changes_since_last_save", g, nil, nil),
		metrics.D("rdb_last_bgsave_time_sec", g, nil, nil),
		metrics.D("rdb_current_bgsave_time_sec", g, nil, nil),
		metrics.D("aof_last_rewrite_time_sec", g, nil, nil),
		metrics.D("aof_current_rewrite_time_sec", g, nil, nil),
		metrics.D("total_connections_received", c, nil, nil),
		metrics.D("total_commands_processed", c, nil, nil),
		metrics.D("instantaneous_ops_per_sec", g, nil, nil),
		metrics.D("rejected_connections", c, nil, nil),
		metrics.D("sync_full", c, nil, nil),
		metrics.D("sync_partial_ok", c, nil, nil),
		metrics.D("sync_partial_err", c, nil, nil),
		metrics.D("expired_keys", c, nil, nil),
		metrics.D("evicted_keys", c, nil, nil),
		metrics.D("keyspace_hits", c, nil, nil),
		metrics.D("keyspace_misses", c, nil, nil),
		metrics.D("pubsub_channels", c, nil, nil),
		metrics.D("pubsub_patterns", c, nil, nil),
		metrics.D("latest_fork_usec", g, nil, nil),
		metrics.D("master_last_io_seconds_ago", g, nil, nil),
		metrics.D("connected_slaves", g, nil, nil),
		metrics.D("repl_backlog_size", g, inMB, nil),
	}
)
