package services

import (
	"github.com/mperham/inspeqtor/util"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

type Initd struct {
	root      string
	path      string
	pidParser func([]byte) (int, error)
}

func (i *Initd) LookupService(serviceName string) (*ProcessStatus, error) {
	ctlpath := i.path + "/" + serviceName
	result, _ := util.FileExists(ctlpath)
	if !result {
		// service script does not exist in etc/init.d, not under
		// init.d control
		return nil, nil
	}

	// First try to find the PID file with same name in /var/run.
	paths := []string{
		i.root + "var/run/" + serviceName + ".pid",
		i.root + "var/run/" + serviceName + "/" + serviceName + ".pid",
	}

	for _, path := range paths {
		st, err := i.readPidFile(path)
		if err != nil {
			util.Info("Error processing PID file %s: %s", path, err.Error())
			continue
		} else if st != nil {
			return st, nil
		}
	}

	return &ProcessStatus{0, Down}, nil
}

func (i *Initd) Name() string {
	return "init.d"
}

func (i *Initd) Restart(serviceName string) error {
	path := "/etc/init.d/" + serviceName

	cmd := exec.Command(path, "restart")
	_, err := util.SafeRun(cmd, util.RestartTimeout)
	if err != nil {
		return err
	}
	return nil
}

func pidForString(data []byte) (int, error) {
	pid, err := strconv.ParseInt(strings.TrimSpace(string(data)), 10, 32)
	if err != nil {
		return 0, err
	}
	return int(pid), nil
}

func (i *Initd) readPidFile(path string) (*ProcessStatus, error) {
	result, err := util.FileExists(path)
	if err != nil {
		return nil, err
	}
	if !result {
		return nil, nil
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	pid, err := i.pidParser(data)
	if err != nil {
		return nil, err
	}

	err = syscall.Kill(pid, syscall.Signal(0))
	if err != nil {
		return nil, err
	}
	return &ProcessStatus{pid, Up}, nil
}

func detectInitd(root string) (InitSystem, error) {
	path := root + "etc/init.d"
	result, err := util.FileExists(path)
	if err != nil {
		return nil, err
	}

	if !result {
		util.Debug("init.d not detected in /etc/init.d")
		return nil, nil
	}

	matches, err := filepath.Glob(path + "/*")
	if err != nil {
		return nil, err
	}

	if len(matches) > 0 {
		util.Info("Detected init.d in " + path)
		return &Initd{root, path, pidForString}, nil
	} else {
		util.Debug(path + " exists but appears to be empty")
		return nil, nil
	}
}
