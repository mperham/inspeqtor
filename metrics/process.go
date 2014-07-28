package metrics

import (
	"inspeqtor/util"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type ProcessMetrics struct {
	When           time.Time
	PID            int32
	UserCpu        uint64
	SystemCpu      uint64
	UserChildCpu   uint64
	SystemChildCpu uint64
	VmRSS          uint64
	VmSize         uint64
}

func CaptureProcess(rootPath string, pid int32) (*ProcessMetrics, error) {
	m := &ProcessMetrics{time.Now(), pid, 0, 0, 0, 0, 0, 0}

	var err error

	if util.Darwin() {
		err = capturePs(m, pid)
		if err != nil {
			return nil, err
		}
	} else {
		err = captureVm(m, rootPath, pid)
		if err != nil {
			return nil, err
		}

		err = captureCpu(m, rootPath, pid)
		if err != nil {
			return nil, err
		}
	}

	return m, nil
}

func capturePs(m *ProcessMetrics, pid int32) error {
	cmd := exec.Command("ps", "So", "rss,vsz,time,utime", "-p", strconv.Itoa(int(pid)))
	sout, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	lines, err := util.ReadLines(sout)
	if err != nil {
		return err
	}

	fields := strings.Fields(lines[1])
	val, err := strconv.ParseUint(fields[0], 10, 64)
	if err != nil {
		return err
	}
	m.VmRSS = 1024 * val
	val, err = strconv.ParseUint(fields[1], 10, 64)
	if err != nil {
		return err
	}
	m.VmSize = 1024 * val
	return nil
}

func captureCpu(m *ProcessMetrics, rootPath string, pid int32) error {
	dir := rootPath + "/" + strconv.Itoa(int(pid))
	data, err := ioutil.ReadFile(dir + "/stat")
	if err != nil {
		return err
	}

	lines, err := util.ReadLines(data)
	if err != nil {
		return err
	}
	for _, line := range lines {
		fields := strings.Fields(line)
		utime, err := strconv.ParseUint(fields[13], 10, 64)
		if err != nil {
			return err
		}
		stime, err := strconv.ParseUint(fields[14], 10, 64)
		if err != nil {
			return err
		}
		cutime, err := strconv.ParseUint(fields[15], 10, 64)
		if err != nil {
			return err
		}
		cstime, err := strconv.ParseUint(fields[16], 10, 64)
		if err != nil {
			return err
		}
		m.UserCpu = utime
		m.SystemCpu = stime
		m.UserChildCpu = cutime
		m.SystemChildCpu = cstime
	}

	return nil
}

func captureVm(m *ProcessMetrics, rootPath string, pid int32) error {
	dir := rootPath + "/" + strconv.Itoa(int(pid))
	data, err := ioutil.ReadFile(dir + "/status")
	if err != nil {
		return err
	}

	lines, err := util.ReadLines(data)
	if err != nil {
		return err
	}
	for _, line := range lines {
		if line[0] == 'V' {
			items := strings.Fields(line)
			switch items[0] {
			case "VmRSS:":
				val, err := strconv.ParseUint(items[1], 10, 64)
				if err != nil {
					return err
				}
				m.VmRSS = 1024 * val
			case "VmSize:":
				val, err := strconv.ParseUint(items[1], 10, 64)
				if err != nil {
					return err
				}
				m.VmSize = 1024 * val
			}
		}

	}

	return nil
}
