package metrics

import (
	"inspeqtor/util"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func NewHostStore() Storage {
	store := Storage{
		map[string]*family{},
	}

	store.declareGauge("swap", "", nil)
	store.declareGauge("load", "1", multiplyBy100)
	store.declareGauge("load", "5", multiplyBy100)
	store.declareGauge("load", "15", multiplyBy100)
	store.declareCounter("cpu", "", PERCENTAGE)
	store.declareCounter("cpu", "user", PERCENTAGE)
	store.declareCounter("cpu", "system", PERCENTAGE)
	store.declareCounter("cpu", "iowait", PERCENTAGE)
	store.declareCounter("cpu", "steal", PERCENTAGE)
	store.declareDynamicFamily("disk")
	store.declareGauge("disk", "/", nil)
	return store
}

func multiplyBy100(val int64) int64 {
	return val * 100
}

const (
	CYCLE_TICKS float64 = CLK_TCK * 15
)

var (
	PERCENTAGE = func(cur, prev int64) int64 {
		return int64((float64(cur-prev) / CYCLE_TICKS) * 100)
	}
	meminfoParser *regexp.Regexp = regexp.MustCompile("([^:]+):\\s+(\\d+)")
)

func CollectHostMetrics(store Storage, path string) error {
	var err error

	err = collectLoadAverage(path, store)
	if err != nil {
		return err
	}
	err = collectMemory(path, store)
	if err != nil {
		return err
	}
	err = collectCpu(path, store)
	if err != nil {
		return err
	}

	err = collectDisk("", store)
	if err != nil {
		return err
	}

	return nil
}

func collectMemory(path string, store Storage) error {
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

		memMetrics := make(map[string]int64)
		for _, line := range lines {
			if line == "" {
				continue
			}

			results := meminfoParser.FindStringSubmatch(line)
			if results == nil {
				util.Warn("Unknown input: " + line)
				continue
			}
			val, err := strconv.ParseInt(results[2], 10, 64)
			if err != nil {
				util.Warn("Unexpected input: " + results[2] + " in " + line)
				return err
			}
			memMetrics[results[1]] = val
		}

		free := memMetrics["SwapFree"]
		total := memMetrics["SwapTotal"]
		if free == 0 {
			store.save("swap", "", 100)
		} else if free == total {
			store.save("swap", "", 0)
		} else {
			store.save("swap", "", int64(100-int8(100*(float64(free)/float64(total)))))
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
		if t == 0 {
			store.save("swap", "", 100)
		} else {
			store.save("swap", "", int64(100*(u/t)))
		}
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

func collectLoadAverage(path string, store Storage) error {
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
	load1, err := strconv.ParseFloat(slices[0], 64)
	if err != nil {
		return err
	}
	load5, err := strconv.ParseFloat(slices[1], 64)
	if err != nil {
		return err
	}
	load15, err := strconv.ParseFloat(slices[2], 64)
	if err != nil {
		return err
	}

	store.save("load", "1", int64(load1*100))
	store.save("load", "5", int64(load5*100))
	store.save("load", "15", int64(load15*100))
	return nil
}

func collectCpu(path string, store Storage) error {
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

		user, _ := strconv.ParseInt(fields[1], 10, 64)
		nice, _ := strconv.ParseInt(fields[2], 10, 64)
		system, _ := strconv.ParseInt(fields[3], 10, 64)
		iowait, _ := strconv.ParseInt(fields[5], 10, 64)
		irq, _ := strconv.ParseInt(fields[6], 10, 64)
		softIrq, _ := strconv.ParseInt(fields[7], 10, 64)
		steal, _ := strconv.ParseInt(fields[8], 10, 64)
		total := user + nice + system + iowait + irq + softIrq + steal

		// These are the five I can envision writing rules against.
		// Open an issue if you want access to the other values.
		store.save("cpu", "", total)
		store.save("cpu", "user", user)
		store.save("cpu", "system", system)
		store.save("cpu", "iowait", iowait)
		store.save("cpu", "steal", steal)
	}
	return nil
}

func collectDisk(path string, store Storage) error {
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

	usage := map[string]int64{}

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
				usage[items[len(items)-1]] = val
			}

		}
	}

	for name, used := range usage {
		store.saveType("disk", name, used, Gauge)
	}
	return nil
}
