package metrics

import (
	"inspeqtor/util"
	"io/ioutil"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func NewProcessStore(values ...interface{}) Storage {
	store := Storage{
		make(map[string]map[string]metric),
	}

	store.declare("memory", "rss", Gauge)
	store.declare("memory", "vsz", Gauge)
	store.declare("cpu", "user", Counter)
	store.declare("cpu", "system", Counter)
	store.declare("cpu", "total_user", Counter)
	store.declare("cpu", "total_system", Counter)
	store.fill(values...)
	return store
}

func CaptureProcess(store Storage, rootPath string, pid int) error {
	var err error

	ok, err := util.FileExists(rootPath)
	if err != nil {
		return err
	}

	if !ok {
		// we don't have the /proc filesystem, e.g. darwin or freebsd
		// use `ps` output instead.
		err = capturePs(store, pid)
		if err != nil {
			return err
		}
	} else {
		err = captureVm(store, rootPath, pid)
		if err != nil {
			return err
		}

		err = captureCpu(store, rootPath, pid)
		if err != nil {
			return err
		}
	}

	return nil
}

/*
 * So many hacks in this.  OSX support can be seen as "bad" at best.
 */
func capturePs(store Storage, pid int) error {
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

	store.save("memory", "rss", 1024*val)
	val, err = strconv.ParseInt(fields[1], 10, 64)
	if err != nil {
		return err
	}
	store.save("memory", "vsz", 1024*val)

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

	store.save("cpu", "user", int64(uticks))
	store.save("cpu", "system", int64(ticks-uticks))

	return nil
}

var (
	timeRegexp = regexp.MustCompile("\\A(\\d+):(\\d\\d).(\\d\\d)\\z")
)

func captureCpu(store Storage, rootPath string, pid int) error {
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
		store.save("cpu", "user", utime)
		store.save("cpu", "system", stime)
		store.save("cpu", "total_user", cutime)
		store.save("cpu", "total_system", cstime)
	}

	return nil
}

func captureVm(store Storage, rootPath string, pid int) error {
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
				store.save("memory", "rss", 1024*val)
			case "VmSize:":
				val, err := strconv.ParseInt(items[1], 10, 64)
				if err != nil {
					return err
				}
				store.save("memory", "vsz", 1024*val)
			}
		}

	}

	return nil
}
