package daemon

import (
	"errors"
	"strconv"
	"strings"

	"github.com/mperham/inspeqtor/util"
)

type MysqlSource struct {
	Hostname    string
	Port        string
	Socket      string
	Username    string
	Password    string
	metrics     map[string]bool
	captureRepl bool
	exec        executor
}

func (rs *MysqlSource) Name() string {
	return "mysql"
}

func (rs *MysqlSource) Watch(metricName string) {
	rs.metrics[metricName] = true
}

func (rs *MysqlSource) Capture() (MetricMap, error) {
	values, err := rs.runStatus(execCmd)
	if err != nil {
		return nil, err
	}
	if rs.captureRepl {
		values, err = rs.runRepl(values, execCmd)
	}
	return values, nil
}

func (rs *MysqlSource) Prepare() error {
	if !rs.metrics["Seconds_Behind_Master"] {
		return nil
	}
	args := rs.buildArgs()
	args = append(args, "-e")
	args = append(args, "show status like 'Slave_running'")
	sout, err := rs.exec("mysql", args, nil)
	if err != nil {
		return err
	}
	lines, err := util.ReadLines(sout)
	if err != nil {
		return err
	}

	parts := strings.Fields(lines[1])
	if parts[1] != "ON" {
		return errors.New("Cannot monitor mysql replication, slave not running")
	}
	delete(rs.metrics, "Seconds_Behind_Master")
	rs.captureRepl = true
	return nil
}

func (rs *MysqlSource) ValidMetrics() []Metric {
	return mysqlMetrics
}

func (rs *MysqlSource) runRepl(values MetricMap, funk executor) (MetricMap, error) {
	args := rs.buildArgs()
	args = append(args, "-e")
	args = append(args, "show slave status\\G")
	sout, err := funk("mysql", args, nil)
	lines, err := util.ReadLines(sout)
	if err != nil {
		return nil, err
	}

	for _, line := range lines {
		if line == "" || line[0] == '*' {
			continue
		}
		parts := strings.Fields(line)
		if parts[0] == "Seconds_Behind_Master:" {
			if parts[0] == "NULL" {
				values["Seconds_Behind_Master"] = 999999
			} else {
				val, err := strconv.ParseInt(parts[1], 10, 64)
				if err != nil {
					return nil, errors.New("Invalid metric input for '" + line + "': " + err.Error())
				}
				values["Seconds_Behind_Master"] = float64(val)
			}
		}
	}
	return values, nil
}

func (rs *MysqlSource) runStatus(funk executor) (MetricMap, error) {
	args := rs.buildArgs()
	args = append(args, "-e")
	args = append(args, "show global status")
	sout, err := funk("mysql", args, nil)
	lines, err := util.ReadLines(sout)
	if err != nil {
		return nil, err
	}

	values := map[string]float64{}

	for _, line := range lines {
		if line == "" || line[0] == '#' {
			continue
		}
		parts := strings.Fields(line)
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
				util.Warn("Could not find metric mysql(%s), did you spell it right?", k)
			}
		}
	}

	return values, nil
}

func (rs *MysqlSource) buildArgs() []string {
	args := []string{"-B"}
	socket := false

	if rs.Socket != "" {
		result, err := util.FileExists(rs.Socket)
		if result && err == nil {
			socket = true
		}
	}

	if socket {
		args = append(args, "-S")
		args = append(args, rs.Socket)
	} else {
		if rs.Hostname != "" {
			args = append(args, "-h")
			args = append(args, rs.Hostname)
		}
		if rs.Port != "" {
			args = append(args, "-P")
			args = append(args, rs.Port)
		}
	}

	if rs.Username != "" {
		args = append(args, "-u")
		args = append(args, rs.Username)
	}
	if rs.Password != "" {
		args = append(args, "-p")
		args = append(args, rs.Password)
	}

	return args
}

func buildMysqlSource(params map[string]string) (Collector, error) {
	rs := &MysqlSource{"localhost", "3306", "/tmp/mysql.sock", "root", "", map[string]bool{}, false, execCmd}
	for k, v := range params {
		switch k {
		case "username":
			rs.Username = v
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
	mysqlMetrics = []Metric{
		Metric{"Aborted_clients", c, nil, nil},
		Metric{"Aborted_connects", c, nil, nil},
		Metric{"Bytes_received", c, inMB, nil},
		Metric{"Bytes_sent", c, inMB, nil},
		Metric{"Com_delete", c, nil, nil},
		Metric{"Com_delete_multi", c, nil, nil},
		Metric{"Com_insert", c, nil, nil},
		Metric{"Com_insert_select", c, nil, nil},
		Metric{"Com_select", c, nil, nil},
		Metric{"Com_update", c, nil, nil},
		Metric{"Com_update_multi", c, nil, nil},
		Metric{"Connections", c, nil, nil},
		Metric{"Created_tmp_disk_tables", c, nil, nil},
		Metric{"Created_tmp_files", c, nil, nil},
		Metric{"Created_tmp_tables", c, nil, nil},
		Metric{"Innodb_buffer_pool_pages_data", g, nil, nil},
		Metric{"Innodb_buffer_pool_bytes_data", g, inMB, nil},
		Metric{"Innodb_buffer_pool_pages_free", g, nil, nil},
		Metric{"Innodb_buffer_pool_pages_total", g, nil, nil},
		Metric{"Innodb_buffer_pool_reads", c, nil, nil},
		Metric{"Innodb_data_read", c, nil, nil},
		Metric{"Innodb_data_reads", c, nil, nil},
		Metric{"Innodb_data_writes", c, nil, nil},
		Metric{"Innodb_data_written", c, nil, nil},
		Metric{"Innodb_deadlocks", c, nil, nil},
		Metric{"Innodb_pages_created", c, nil, nil},
		Metric{"Innodb_pages_read", c, nil, nil},
		Metric{"Innodb_pages_written", c, nil, nil},
		Metric{"Innodb_rows_deleted", c, nil, nil},
		Metric{"Innodb_rows_inserted", c, nil, nil},
		Metric{"Innodb_rows_read", c, nil, nil},
		Metric{"Innodb_rows_updated", c, nil, nil},
		Metric{"Innodb_num_open_files", g, nil, nil},
		Metric{"Key_reads", c, nil, nil},
		Metric{"Key_writes", c, nil, nil},
		Metric{"Max_used_connections", g, nil, nil},
		Metric{"Open_files", g, nil, nil},
		Metric{"Open_tables", g, nil, nil},
		Metric{"Prepared_stmt_count", g, nil, nil},
		Metric{"Qcache_free_blocks", g, nil, nil},
		Metric{"Qcache_free_memory", g, inMB, nil},
		Metric{"Qcache_hits", c, nil, nil},
		Metric{"Qcache_inserts", c, nil, nil},
		Metric{"Qcache_lowmem_prunes", c, nil, nil},
		Metric{"Qcache_not_cached", c, nil, nil},
		Metric{"Qcache_queries_in_cache", g, nil, nil},
		Metric{"Queries", c, nil, nil},
		Metric{"Questions", c, nil, nil},
		Metric{"Seconds_Behind_Master", g, nil, nil},
		Metric{"Slow_queries", c, nil, nil},
		Metric{"Table_locks_immediate", c, nil, nil},
		Metric{"Table_locks_waited", c, nil, nil},
		Metric{"Threads_cached", g, nil, nil},
		Metric{"Threads_connected", g, nil, nil},
		Metric{"Threads_created", c, nil, nil},
		Metric{"Threads_running", g, nil, nil},
	}
)
