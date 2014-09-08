package metrics

import (
	"inspeqtor/util"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var (
	timeRegexp = regexp.MustCompile("\\A(\\d+):(\\d\\d).(\\d\\d)\\z")
)

type processStorage struct {
	*storage
	path string
}

func NewProcessStore(path string, cycleSeconds uint) Store {
	store := &processStorage{
		&storage{map[string]*family{}},
		path,
	}

	tickPercentage := func(cur, prev int64) int64 {
		return int64((float64(cur-prev) / float64(cycleSeconds*CLK_TCK)) * 100)
	}

	store.DeclareGauge("memory", "rss", nil, displayInMB)
	store.DeclareCounter("cpu", "user", tickPercentage, displayPercent)
	store.DeclareCounter("cpu", "system", tickPercentage, displayPercent)
	store.DeclareCounter("cpu", "total_user", tickPercentage, displayPercent)
	store.DeclareCounter("cpu", "total_system", tickPercentage, displayPercent)
	return store
}

func (ps *processStorage) Load(values ...interface{}) {
	if len(values) > 0 {
		ps.fill(values...)
	}
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
		err = ps.captureVm(pid)
		if err != nil {
			return err
		}

		err = ps.captureCpu(pid)
		if err != nil {
			return err
		}
	}

	return nil
}

/*
 * So many hacks in this.  OSX support can be seen as "bad" at best.
 */
func (ps *processStorage) capturePs(pid int) error {
	cmd := exec.Command("ps", "So", "rss,time,utime", "-p", strconv.Itoa(pid))
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

	ps.Save("memory", "rss", 1024*val)

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

	ps.Save("cpu", "user", int64(uticks))
	ps.Save("cpu", "system", int64(ticks-uticks))

	return nil
}

func (ps *processStorage) captureCpu(pid int) error {
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
		ps.Save("cpu", "user", utime)
		ps.Save("cpu", "system", stime)
		ps.Save("cpu", "total_user", cutime)
		ps.Save("cpu", "total_system", cstime)
	}

	return nil
}

func (ps *processStorage) captureVm(pid int) error {
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
				ps.Save("memory", "rss", 1024*val)
			}
		}

	}

	return nil
}
