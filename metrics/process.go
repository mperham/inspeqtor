package metrics

import (
	"inspeqtor/util"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	/*
	  CPU time is stored in system clock ticks.  Most
	  modern linux systems use 100 clock ticks per second,
	  or 100 Hz.  Use "getconf CLK_TCK" to verify this.
	  Since our cycle time is 15 seconds, there's 1500
	  ticks to spend each cycle.  If a process uses 750
	  ticks on the CPU, that means it used 50% of the CPU
	  during that cycle.  Multithreaded processes running on
	  systems with multiple CPUs/cores can use more than 100% CPU.
	*/
	CLK_TCK = 100
)

type metricFamily string

var (
	SupportedProcessMetrics = map[metricFamily]func(*ProcessMetrics, string) int64{
		"memory": func(p *ProcessMetrics, param string) int64 {
			switch param {
			case "rss":
				return p.VmRSS
			case "vsz":
				return p.VmSize
			default:
				panic("Unknown process metric \"memory(" + param + ")\"")
			}
		},
		"cpu": func(p *ProcessMetrics, param string) int64 {
			switch param {
			case "user":
				return p.UserCpu
			case "system":
				return p.SystemCpu
			default:
				panic("Unknown process metric \"cpu(" + param + ")\"")
			}
		},
	}
)

func Lookup(family string, name string, buffer *util.RingBuffer, idx int) int64 {
	getter := SupportedProcessMetrics[metricFamily(family)]
	if getter == nil {
		return 0
	}

	m := buffer.At(idx).(*ProcessMetrics)
	return getter(m, name)
}

type ProcessMetrics struct {
	When           time.Time
	PID            int
	UserCpu        int64
	SystemCpu      int64
	UserChildCpu   int64
	SystemChildCpu int64
	VmRSS          int64
	VmSize         int64
}

func CaptureProcess(rootPath string, pid int) (*ProcessMetrics, error) {
	m := &ProcessMetrics{time.Now(), pid, 0, 0, 0, 0, 0, 0}

	var err error

	ok, err := util.FileExists(rootPath)
	if err != nil {
		return nil, err
	}

	if !ok {
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

/*
So many hacks in this.  OSX support can be seen as "bad" at best.
*/
func capturePs(m *ProcessMetrics, pid int) error {
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
	val, err := strconv.ParseInt(fields[0], 10, 64)
	if err != nil {
		return err
	}
	m.VmRSS = 1024 * val
	val, err = strconv.ParseInt(fields[1], 10, 64)
	if err != nil {
		return err
	}
	m.VmSize = 1024 * val

	times := timeRegexp.FindStringSubmatch(fields[2])
	if times == nil {
		util.Debug("Unable to parse CPU time in " + lines[1])
		return nil
	}
	min, _ := strconv.ParseUint(times[1], 10, 32)
	sec, _ := strconv.ParseUint(times[2], 10, 32)
	cs, _ := strconv.ParseUint(times[3], 10, 32)

	ticks := min*60*100 + sec*100 + cs

	times = timeRegexp.FindStringSubmatch(fields[3])
	if times == nil {
		util.Debug("Unable to parse User time in " + lines[1])
		return nil
	}
	min, _ = strconv.ParseUint(times[1], 10, 32)
	sec, _ = strconv.ParseUint(times[2], 10, 32)
	cs, _ = strconv.ParseUint(times[3], 10, 32)

	uticks := min*60*100 + sec*100 + cs

	m.UserCpu = int64(uticks)
	m.SystemCpu = int64(ticks - uticks)

	return nil
}

var (
	timeRegexp = regexp.MustCompile("\\A(\\d+):(\\d\\d).(\\d\\d)\\z")
)

func captureCpu(m *ProcessMetrics, rootPath string, pid int) error {
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
		utime, err := strconv.ParseInt(fields[13], 10, 64)
		if err != nil {
			return err
		}
		stime, err := strconv.ParseInt(fields[14], 10, 64)
		if err != nil {
			return err
		}
		cutime, err := strconv.ParseInt(fields[15], 10, 64)
		if err != nil {
			return err
		}
		cstime, err := strconv.ParseInt(fields[16], 10, 64)
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

func captureVm(m *ProcessMetrics, rootPath string, pid int) error {
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
				val, err := strconv.ParseInt(items[1], 10, 64)
				if err != nil {
					return err
				}
				m.VmRSS = 1024 * val
			case "VmSize:":
				val, err := strconv.ParseInt(items[1], 10, 64)
				if err != nil {
					return err
				}
				m.VmSize = 1024 * val
			}
		}

	}

	return nil
}
