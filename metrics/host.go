package metrics

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"

	"github.com/mperham/inspeqtor/util"
)

type hostStorage struct {
	*storage
	cycleTicks float64
	path       string
}

func NewHostStore(path string, cycleSeconds uint) Store {

	store := &hostStorage{
		&storage{map[string]*family{}},
		float64(cycleSeconds * ClkTck),
		path,
	}

	tickPercentage := func(cur, prev float64) float64 {
		return float64((float64(cur-prev) / float64(cycleSeconds*ClkTck)) * 100)
	}

	store.DeclareGauge("swap", "", DisplayPercent)
	store.DeclareGauge("load", "1", displayLoad)
	store.DeclareGauge("load", "5", displayLoad)
	store.DeclareGauge("load", "15", displayLoad)
	store.DeclareCounter("cpu", "", tickPercentage, DisplayPercent)
	store.DeclareCounter("cpu", "user", tickPercentage, DisplayPercent)
	store.DeclareCounter("cpu", "system", tickPercentage, DisplayPercent)
	store.DeclareCounter("cpu", "iowait", tickPercentage, DisplayPercent)
	store.DeclareCounter("cpu", "steal", tickPercentage, DisplayPercent)
	store.declareDynamicFamily("disk")
	store.DeclareGauge("disk", "/", DisplayPercent)
	return store
}

func (hs *hostStorage) Prepare() error {
	return nil
}

func (hs *hostStorage) AddSource(name string, config map[string]string) (Source, error) {
	return nil, fmt.Errorf("Host storage does not support dynamic metric sources")
}

func (hs *hostStorage) Watch(metricFamily, metricName string) error {
	return nil
}

func (hs *hostStorage) Collect(_ int) error {
	var err error

	err = hs.collectLoadAverage()
	if err != nil {
		return err
	}
	err = hs.collectMemory()
	if err != nil {
		return err
	}
	err = hs.collectCPU()
	if err != nil {
		return err
	}

	err = hs.collectDisk("")
	if err != nil {
		return err
	}

	return nil
}

func (hs *hostStorage) collectMemory() error {
	ok, err := util.FileExists(hs.path + "/meminfo")
	if err != nil {
		return err
	}

	if ok {
		contentBytes, err := ioutil.ReadFile(hs.path + "/meminfo")
		if err != nil {
			return err
		}
		lines := strings.Split(string(contentBytes), "\n")

		memMetrics := make(map[string]float64)
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
			memMetrics[results[1]] = float64(val)
		}

		free := memMetrics["SwapFree"]
		total := memMetrics["SwapTotal"]
		if free == 0 {
			hs.Save("swap", "", 100)
		} else if free == total {
			hs.Save("swap", "", 0)
		} else {
			hs.Save("swap", "", float64(100-int8(100*(float64(free)/float64(total)))))
		}
	} else {
		cmd := exec.Command("sysctl", "-n", "vm.swapusage")
		cmd.Env = []string{"LANG=C"}
		sout, err := util.SafeRun(cmd)
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
			hs.Save("swap", "", 100)
		} else {
			hs.Save("swap", "", float64(100*(u/t)))
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

func (hs *hostStorage) collectLoadAverage() error {
	// TODO make this a one-time check so we don't incur the overhead
	// on every cycle.
	ok, err := util.FileExists(hs.path + "/loadavg")
	if err != nil {
		return err
	}

	var loadavgString string
	if ok {
		contentBytes, err := ioutil.ReadFile(hs.path + "/loadavg")
		if err != nil {
			return err
		}
		loadavgString = string(contentBytes)
	} else {
		cmd := exec.Command("sysctl", "-n", "vm.loadavg")
		cmd.Env = []string{"LANG=C"}
		sout, err := util.SafeRun(cmd)
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

	hs.Save("load", "1", load1)
	hs.Save("load", "5", load5)
	hs.Save("load", "15", load15)
	return nil
}

func (hs *hostStorage) collectCPU() error {
	ok, err := util.FileExists(hs.path + "/stat")
	if err != nil {
		return err
	}

	if ok {
		contents, err := ioutil.ReadFile(hs.path + "/stat")
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
		hs.Save("cpu", "", float64(total))
		hs.Save("cpu", "user", float64(user))
		hs.Save("cpu", "system", float64(system))
		hs.Save("cpu", "iowait", float64(iowait))
		hs.Save("cpu", "steal", float64(steal))
	} else {
		// TODO
		util.Info("Cannot collect host CPU metrics, not implemented on this platform")
	}
	return nil
}

func (hs *hostStorage) collectDisk(path string) error {
	var lines []string

	if path == "" {
		cmd := exec.Command("df", "-P")
		sout, err := util.SafeRun(cmd)
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

	usage := map[string]float64{}

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
				usage[items[len(items)-1]] = float64(val)
			}

		}
	}

	for name, used := range usage {
		hs.saveType("disk", name, used, Gauge)
	}
	return nil
}
