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

type DiskUsage map[string]int8

type SystemMetrics struct {
	When             time.Time
	CpuUsage         CpuMetrics
	Load1            float32
	Load5            float32
	Load15           float32
	PercentSwapInUse int8
	Disk             *DiskUsage
}

type CpuMetrics struct {
	Total     uint64
	User      uint64
	Nice      uint64
	System    uint64
	Idle      uint64
	IOWait    uint64
	Irq       uint64
	SoftIrq   uint64
	Steal     uint64
	Guest     uint64
	GuestNice uint64
}

var (
	meminfoParser *regexp.Regexp = regexp.MustCompile("([^:]+):\\s+(\\d+)")
)

func CollectHostMetrics(path string) (*SystemMetrics, error) {
	var err error
	var metrics *SystemMetrics = &SystemMetrics{}
	metrics.When = time.Now()

	err = collectLoadAverage(path, metrics)
	if err != nil {
		return nil, err
	}
	err = collectMemory(path, metrics)
	if err != nil {
		return nil, err
	}
	err = collectCpu(path, metrics)
	if err != nil {
		return nil, err
	}

	err = collectDisk("", metrics)
	if err != nil {
		return nil, err
	}

	return metrics, nil
}

func collectMemory(path string, metrics *SystemMetrics) error {
	ok, err := util.FileExists(path + "/meminfo")
	if err != nil {
		return err
	}

	if ok {
		contentBytes, err := ioutil.ReadFile(path + "/meminfo")
		if err != nil {
			return err
		}
		lines := strings.Split(string(contentBytes), "\n")

		memMetrics := make(map[string]uint64)
		for _, line := range lines {
			if line == "" {
				continue
			}

			results := meminfoParser.FindStringSubmatch(line)
			if results == nil {
				util.Warn("Unknown input: " + line)
				continue
			}
			val, err := strconv.ParseUint(results[2], 10, 64)
			if err != nil {
				util.Warn("Unexpected input: " + results[2] + " in " + line)
				return err
			}
			memMetrics[results[1]] = uint64(val)
		}

		free := memMetrics["SwapFree"]
		total := memMetrics["SwapTotal"]
		if free == 0 {
			metrics.PercentSwapInUse = 100
		} else if free == total {
			metrics.PercentSwapInUse = 0
		} else {
			metrics.PercentSwapInUse = 100 - int8(100*(float64(free)/float64(total)))
		}
	} else {
		cmd := exec.Command("sysctl", "-n", "vm.swapusage")
		sout, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		lines, err := util.ReadLines(sout)
		if err != nil {
			return err
		}

		rest := lines[0]
		matches := swapRegexp.FindStringSubmatch(rest)
		total := matches[1]
		rest = matches[2]

		matches = swapRegexp.FindStringSubmatch(rest)
		used := matches[1]

		tot, err := strconv.ParseFloat(total[0:len(total)-1], 64)
		if err != nil {
			return err
		}
		usd, err := strconv.ParseFloat(used[0:len(used)-1], 64)
		if err != nil {
			return err
		}

		t := normalizeSwap(tot, rune(total[len(total)-1]))
		u := normalizeSwap(usd, rune(used[len(used)-1]))
		metrics.PercentSwapInUse = int8(100 * (u / t))
	}

	return nil
}

func normalizeSwap(val float64, size rune) float64 {
	switch size {
	case 'M', 'm':
		return val * 1024
	case 'K', 'k':
		return val
	case 'G', 'g':
		return val * 1024 * 1024
	case 'T', 't':
		return val * 1024 * 1024 * 1024
	default:
		// ¯\_( ツ )_/¯
		return val
	}
}

var (
	swapRegexp = regexp.MustCompile("= (\\d+\\.\\d{2}[A-Z])(.*)")
)

func collectLoadAverage(path string, metrics *SystemMetrics) error {
	// TODO make this a one-time check so we don't incur the overhead
	// on every cycle.
	ok, err := util.FileExists(path + "/loadavg")
	if err != nil {
		return err
	}

	var loadavgString string
	if ok {
		contentBytes, err := ioutil.ReadFile(path + "/loadavg")
		if err != nil {
			return err
		}
		loadavgString = string(contentBytes)
	} else {
		cmd := exec.Command("sysctl", "-n", "vm.loadavg")
		sout, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		lines, err := util.ReadLines(sout)
		if err != nil {
			return err
		}
		loadavgString = lines[0][2 : len(lines[0])-2] // trim braces
	}

	slices := strings.Split(loadavgString, " ")
	load1, err := strconv.ParseFloat(slices[0], 32)
	if err != nil {
		return err
	}
	load5, err := strconv.ParseFloat(slices[1], 32)
	if err != nil {
		return err
	}
	load15, err := strconv.ParseFloat(slices[2], 32)
	if err != nil {
		return err
	}

	metrics.Load1 = float32(load1)
	metrics.Load5 = float32(load5)
	metrics.Load15 = float32(load15)
	return nil
}

func collectCpu(path string, metrics *SystemMetrics) error {
	ok, err := util.FileExists(path + "/stat")
	if err != nil {
		return err
	}

	if ok {
		contents, err := ioutil.ReadFile(path + "/stat")
		if err != nil {
			return err
		}

		lines := strings.Split(string(contents), "\n")
		line := lines[0]
		fields := strings.Fields(line)
		metrics.CpuUsage = createCpuMetrics(fields)
	}
	return nil
}

func createCpuMetrics(fields []string) CpuMetrics {
	s := CpuMetrics{}
	s.User, _ = strconv.ParseUint(fields[1], 10, 64)
	s.Nice, _ = strconv.ParseUint(fields[2], 10, 64)
	s.System, _ = strconv.ParseUint(fields[3], 10, 64)
	s.Idle, _ = strconv.ParseUint(fields[4], 10, 64)
	s.IOWait, _ = strconv.ParseUint(fields[5], 10, 64)
	s.Irq, _ = strconv.ParseUint(fields[6], 10, 64)
	s.SoftIrq, _ = strconv.ParseUint(fields[7], 10, 64)
	s.Steal, _ = strconv.ParseUint(fields[8], 10, 64)
	s.Guest, _ = strconv.ParseUint(fields[9], 10, 64)
	s.GuestNice, _ = strconv.ParseUint(fields[10], 10, 64)
	s.Total = s.User + s.Nice + s.System + s.Idle + s.IOWait +
		s.Irq + s.SoftIrq + s.Steal + s.Guest + s.GuestNice
	return s
}

func collectDisk(path string, metrics *SystemMetrics) error {
	var lines []string

	if path == "" {
		cmd := exec.Command("df")
		sout, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		lines, err = util.ReadLines(sout)
		if err != nil {
			return err
		}
	} else {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		lines, err = util.ReadLines(data)
		if err != nil {
			return err
		}
	}

	usage := DiskUsage{}

	for _, line := range lines {
		if line[0] == '/' {
			items := strings.Fields(line)
			if len(items) < 5 {
				util.Debug("Cannot parse df output: %v", items)
				continue
			}
			pct := items[4]
			if pct[len(pct)-1] == '%' {
				val, err := strconv.ParseInt(pct[0:len(pct)-1], 10, 32)
				if err != nil {
					util.Debug("Cannot parse df output: " + line)
				}
				usage[items[len(items)-1]] = int8(val)
			}

		}
	}

	metrics.Disk = &usage
	return nil
}
