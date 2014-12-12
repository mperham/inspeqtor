package metrics

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/mperham/inspeqtor/util"
)

var (
	timeRegexp = regexp.MustCompile("\\A(\\d+):(\\d\\d).(\\d\\d)\\z")
)

type dynamicCollector func(int, *processStorage) error

type processStorage struct {
	*storage
	path string
	// TODO refactor this to be a Source
	dyncol         []dynamicCollector
	daemonSpecific []Source
}

func NewProcessStore(path string, cycleSeconds uint) Store {
	store := &processStorage{
		storage: &storage{map[string]*family{}},
		path:    path,
	}

	tickPercentage := func(cur, prev float64) float64 {
		return float64(((cur - prev) / float64(cycleSeconds*ClkTck)) * 100)
	}

	store.DeclareGauge("memory", "rss", DisplayInMB)
	store.DeclareCounter("cpu", "user", tickPercentage, DisplayPercent)
	store.DeclareCounter("cpu", "system", tickPercentage, DisplayPercent)
	store.DeclareCounter("cpu", "total_user", tickPercentage, DisplayPercent)
	store.DeclareCounter("cpu", "total_system", tickPercentage, DisplayPercent)
	return store
}

func (ps *processStorage) Prepare() error {
	for _, x := range ps.daemonSpecific {
		err := x.Prepare()
		if err != nil {
			return err
		}
	}
	return nil
}

func (ps *processStorage) AddSource(name string, config map[string]string) (Source, error) {
	for _, x := range ps.daemonSpecific {
		if x.Name() == name {
			return x, nil
		}
	}
	builder := Sources[name]
	if builder == nil {
		return nil, nil
	}
	util.Info("Activating metrics for %s", name)
	src, err := builder(config)
	if err != nil {
		return nil, err
	}
	ps.daemonSpecific = append(ps.daemonSpecific, src)
	return src, nil
}

func (ps *processStorage) Watch(family, name string) error {
	if family == "memory" && name == "total_rss" {
		ps.DeclareGauge("memory", "total_rss", DisplayInMB)
		ps.dyncol = append(ps.dyncol, totalRssCollector)
		return nil
	}

	for _, x := range ps.daemonSpecific {
		if x.Name() == family {
			descs := x.ValidMetrics()
			for _, d := range descs {
				if d.Name == name {
					if d.MetricType == Counter {
						ps.DeclareCounter(family, name, nil, d.Display)
					} else {
						ps.DeclareGauge(family, name, d.Display)
					}
					x.Watch(name)
					return nil
				}
			}
			return fmt.Errorf("No such metric: %s:%s", family, name)
		}
	}

	return nil
}

func (ps *processStorage) Load(values ...interface{}) {
	if len(values) > 0 {
		ps.fill(values...)
	}
}

type processEntry struct {
	pid  int
	ppid int
	rss  int64
}

/*
  Collecting total RSS for a process is actually rather involved on Linux.
	Because of this, we only collect the value if the user defines a rule
	for it.  This is a "dynamicCollector".
*/
func totalRssCollector(mypid int, ps *processStorage) error {
	matches, err := filepath.Glob(fmt.Sprintf("%s/[1-9][0-9]*/status", ps.path))
	if err != nil {
		return err
	}
	var live []processEntry

	for _, file := range matches {
		pe := processEntry{}

		data, err := ioutil.ReadFile(file)
		if err != nil {
			// race condition between globbing and reading, process
			// can disappear at any moment
			continue
		}

		lines, err := util.ReadLines(data)
		if err != nil {
			return err
		}
		for _, line := range lines {
			if line[0] == 'V' || line[0] == 'P' {
				items := strings.Split(line, ":")
				switch items[0] {
				case "Pid":
					pid, err := strconv.Atoi(strings.TrimSpace(items[1]))
					if err != nil {
						return err
					}
					pe.pid = pid
				case "PPid":
					ppid, err := strconv.Atoi(strings.TrimSpace(items[1]))
					if err != nil {
						return err
					}
					pe.ppid = ppid
				case "VmRSS":
					vals := strings.Fields(items[1])
					val, err := strconv.ParseInt(vals[0], 10, 64)
					if err != nil {
						return err
					}
					pe.rss = val * 1024
				}
			}
		}

		if pe.rss != 0 {
			live = append(live, pe)
		}
	}
	util.DebugDebug("Calculating %d processes", len(live))

	rss := memoryFor(live, mypid)
	util.DebugDebug("Total RSS for %d: %d", mypid, rss)

	ps.Save("memory", "total_rss", float64(rss))
	return nil
}

// recursively walk the array, looking for children, grandchildren, etc
func memoryFor(procs []processEntry, curpid int) int64 {
	var rss int64
	for _, entry := range procs {
		if entry.pid == curpid {
			rss += entry.rss
		} else if entry.ppid == curpid {
			rss += memoryFor(procs, entry.pid)
		}
	}
	return rss
}

func (ps *processStorage) Collect(pid int) error {
	var err error

	ok, err := util.FileExists(ps.path)
	if err != nil {
		return err
	}

	if !ok {
		// we don't have the /proc filesystem, e.g. darwin or freebsd
		// use `ps` output instead.
		err = ps.capturePs(pid)
		if err != nil {
			return err
		}
	} else {
		err = ps.captureVM(pid)
		if err != nil {
			return err
		}

		err = ps.captureCPU(pid)
		if err != nil {
			return err
		}
	}

	for _, fn := range ps.dyncol {
		err = fn(pid, ps)
		if err != nil {
			return err
		}
	}

	for _, x := range ps.daemonSpecific {
		data, err := x.Capture()
		if err != nil {
			return err
		}
		for k, v := range data {
			ps.Save(x.Name(), k, v)
		}
	}
	return nil
}

/*
 * So many hacks in this.  OSX support can be seen as "bad" at best.
 */
func (ps *processStorage) capturePs(pid int) error {
	cmd := exec.Command("ps", "So", "rss,time,utime", "-p", strconv.Itoa(pid))
	sout, err := util.SafeRun(cmd)
	if err != nil {
		return err
	}

	lines, err := util.ReadLines(sout)
	if err != nil {
		return err
	}

	if len(lines) < 2 {
		return errors.New("Insufficient output from ps")
	}

	fields := strings.Fields(lines[1])
	val, err := strconv.ParseInt(fields[0], 10, 64)
	if err != nil {
		return err
	}

	ps.Save("memory", "rss", float64(1024*val))

	times := timeRegexp.FindStringSubmatch(fields[1])
	if times == nil {
		util.Debug("Unable to parse CPU time in " + lines[1])
		return nil
	}
	min, _ := strconv.ParseUint(times[1], 10, 32)
	sec, _ := strconv.ParseUint(times[2], 10, 32)
	cs, _ := strconv.ParseUint(times[3], 10, 32)

	ticks := min*60*100 + sec*100 + cs

	times = timeRegexp.FindStringSubmatch(fields[2])
	if times == nil {
		util.Debug("Unable to parse User time in " + lines[1])
		return nil
	}
	min, _ = strconv.ParseUint(times[1], 10, 32)
	sec, _ = strconv.ParseUint(times[2], 10, 32)
	cs, _ = strconv.ParseUint(times[3], 10, 32)

	uticks := min*60*100 + sec*100 + cs

	ps.Save("cpu", "user", float64(uticks))
	ps.Save("cpu", "system", float64(ticks-uticks))

	return nil
}

func (ps *processStorage) captureCPU(pid int) error {
	dir := ps.path + "/" + strconv.Itoa(int(pid))
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
		ps.Save("cpu", "user", float64(utime))
		ps.Save("cpu", "system", float64(stime))
		ps.Save("cpu", "total_user", float64(cutime))
		ps.Save("cpu", "total_system", float64(cstime))
	}

	return nil
}

func (ps *processStorage) captureVM(pid int) error {
	dir := ps.path + "/" + strconv.Itoa(int(pid))
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
				ps.Save("memory", "rss", float64(1024*val))
			}
		}

	}

	return nil
}
