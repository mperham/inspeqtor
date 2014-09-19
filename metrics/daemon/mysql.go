package daemon

import (
	"errors"
	"github.com/mperham/inspeqtor/util"
	"strconv"
	"strings"
)

type MysqlSource struct {
	Hostname    string
	Port        string
	Socket      string
	Username    string
	Password    string
	metrics     map[string]bool
	captureRepl bool
}

func (rs *MysqlSource) Name() string {
	return "mysql"
}

func (rs *MysqlSource) Watch(metricName string) {
	rs.metrics[metricName] = true
}

func (rs *MysqlSource) Capture() (metricMap, error) {
	values, err := rs.runStatus(execCmd)
	if err != nil {
		return nil, err
	}
	if rs.captureRepl {
		values, err = rs.runRepl(values, execCmd)
	}
	return values, nil
}

func (rs *MysqlSource) Prepare(funk executor) error {
	if !rs.metrics["Seconds_Behind_Master"] {
		return nil
	}
	args := rs.buildArgs()
	args = append(args, "-e")
	args = append(args, "show status like 'Slave_running'")
	sout, err := funk("mysql", args, nil)
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

func (rs *MysqlSource) ValidMetrics() []metric {
	return mysqlMetrics
}

func (rs *MysqlSource) runRepl(values metricMap, funk executor) (metricMap, error) {
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
				values["Seconds_Behind_Master"] = val
			}
		}
	}
	return values, nil
}

func (rs *MysqlSource) runStatus(funk executor) (metricMap, error) {
	args := rs.buildArgs()
	args = append(args, "-e")
	args = append(args, "show status")
	sout, err := funk("mysql", args, nil)
	lines, err := util.ReadLines(sout)
	if err != nil {
		return nil, err
	}

	values := map[string]int64{}

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
			values[parts[0]] = val
		}
	}

	if len(rs.metrics) > len(values) {
		for k, _ := range rs.metrics {
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
	rs := &MysqlSource{"localhost", "3306", "/tmp/mysql.sock", "root", "", map[string]bool{}, false}
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
	mysqlMetrics []metric = []metric{
		metric{"Aborted_clients", c, nil},
		metric{"Aborted_connects", c, nil},
		metric{"Bytes_received", c, &funcWrapper{nil, inMB, nil}},
		metric{"Bytes_sent", c, &funcWrapper{nil, inMB, nil}},
		metric{"Com_delete", c, nil},
		metric{"Com_delete_multi", c, nil},
		metric{"Com_insert", c, nil},
		metric{"Com_insert_select", c, nil},
		metric{"Com_select", c, nil},
		metric{"Com_update", c, nil},
		metric{"Com_update_multi", c, nil},
		metric{"Connections", c, nil},
		metric{"Innodb_buffer_pool_pages_data", g, nil},
		metric{"Innodb_buffer_pool_bytes_data", g, &funcWrapper{nil, inMB, nil}},
		metric{"Innodb_buffer_pool_pages_free", g, nil},
		metric{"Innodb_buffer_pool_pages_total", g, nil},
		metric{"Innodb_buffer_pool_reads", c, nil},
		metric{"Innodb_data_read", c, nil},
		metric{"Innodb_data_reads", c, nil},
		metric{"Innodb_data_writes", c, nil},
		metric{"Innodb_data_written", c, nil},
		metric{"Innodb_deadlocks", c, nil},
		metric{"Innodb_pages_created", c, nil},
		metric{"Innodb_pages_read", c, nil},
		metric{"Innodb_pages_written", c, nil},
		metric{"Innodb_rows_deleted", c, nil},
		metric{"Innodb_rows_inserted", c, nil},
		metric{"Innodb_rows_read", c, nil},
		metric{"Innodb_rows_updated", c, nil},
		metric{"Innodb_num_open_files", g, nil},
		metric{"Key_reads", c, nil},
		metric{"Key_writes", c, nil},
		metric{"Max_used_connections", g, nil},
		metric{"Open_files", g, nil},
		metric{"Open_streams", g, nil},
		metric{"Open_tables", g, nil},
		metric{"Opened_files", c, nil},
		metric{"Opened_tables", c, nil},
		metric{"Prepared_stmt_count", g, nil},
		metric{"Qcache_free_blocks", g, nil},
		metric{"Qcache_free_memory", g, &funcWrapper{nil, inMB, nil}},
		metric{"Qcache_hits", c, nil},
		metric{"Qcache_inserts", c, nil},
		metric{"Qcache_lowmem_prunes", c, nil},
		metric{"Qcache_not_cached", c, nil},
		metric{"Qcache_queries_in_cache", g, nil},
		metric{"Queries", c, nil},
		metric{"Questions", c, nil},
		metric{"Seconds_Behind_Master", g, nil},
		metric{"Slow_queries", c, nil},
		metric{"Table_locks_immediate", c, nil},
		metric{"Table_locks_waited", c, nil},
		metric{"Threadpool_idle_threads", g, nil},
		metric{"Threadpool_threads", g, nil},
		metric{"Threads_cached", g, nil},
		metric{"Threads_connected", g, nil},
		metric{"Threads_created", c, nil},
		metric{"Threads_running", g, nil},
	}
)
