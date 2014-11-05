package services

import (
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/mperham/inspeqtor/util"
)

type Initd struct {
	ctlPath    string
	varrunPath string
	pidParser  func([]byte) (int, error)
}

func (i *Initd) LookupService(serviceName string) (*ProcessStatus, error) {
	path := i.ctlPath + serviceName
	result, _ := util.FileExists(path)
	if !result {
		// service script does not exist in etc/init.d, not under
		// init.d control
		return nil, &ServiceError{i.Name(), serviceName, ErrServiceNotFound}
	}

	// First try to find the PID file with same name in /var/run.
	paths := []string{
		i.varrunPath + serviceName + ".pid",
		i.varrunPath + serviceName + "/" + serviceName + ".pid",
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
	path := i.ctlPath + serviceName

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
	ctlpath := root + "etc/init.d/"
	result, err := util.FileExists(ctlpath)
	if err != nil {
		return nil, err
	}

	if !result {
		util.Debug("init.d not detected in " + ctlpath)
		return nil, nil
	}

	matches, err := filepath.Glob(ctlpath + "*")
	if err != nil {
		return nil, err
	}

	if !result {
		util.Debug("init.d not detected in " + ctlpath)
		return nil, nil
	}

	if len(matches) > 0 {
		util.Info("Detected init.d in " + ctlpath)
		return &Initd{ctlpath, root + "var/run/", pidForString}, nil
	} else {
		util.Info(ctlpath + " exists but appears to be empty")
		return nil, nil
	}
}
