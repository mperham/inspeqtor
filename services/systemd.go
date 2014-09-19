package services

import (
	"errors"
	"github.com/mperham/inspeqtor/util"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type Systemd struct {
	path         string
	dummyOutput  string
	dummyOutput2 string
}

func detectSystemd(path string) (InitSystem, error) {
	result, err := util.FileExists(path)
	if err != nil {
		return nil, err
	}

	if !result {
		util.Debug("systemd not detected, no " + path)
		return nil, nil
	}

	matches, err := filepath.Glob(path + "/*.conf")
	if err != nil {
		return nil, err
	}

	if len(matches) > 0 {
		util.Info("Detected systemd in " + path)
		return &Systemd{path, "", ""}, nil
	}

	util.Debug("systemd not detected, empty " + path)
	return nil, nil
}

func (u *Systemd) Name() string {
	return "systemd"
}

func (u *Systemd) Restart(serviceName string) error {
	if len(u.dummyOutput) != 0 {
		//sout = []byte(u.dummyOutput)
	} else {
		cmd := exec.Command("systemctl", "restart", serviceName)
		_, err := cmd.CombinedOutput()
		if err != nil {
			return &ServiceError{u.Name(), serviceName, err}
		}
	}

	return nil
}

func (u *Systemd) LookupService(serviceName string) (*ProcessStatus, error) {
	var sout []byte
	var err error

	if len(u.dummyOutput) != 0 {
		sout = []byte(u.dummyOutput)
	} else {
		cmd := exec.Command("systemctl", "show", "-p", "MainPID", serviceName)
		sout, err = cmd.CombinedOutput()
	}

	if err != nil {
		return nil, &ServiceError{u.Name(), serviceName, err}
	}
	lines, err := util.ReadLines(sout)
	if len(lines) != 1 {
		return nil, &ServiceError{u.Name(), serviceName, errors.New("Unexpected output: " + strings.Join(lines, "\n"))}
	}

	// Output will be "MainPID=1234" or
	// "MainPID=0" if service does not exist.
	line := lines[0]
	fields := strings.Split(line, "=")
	if fields[1] != "0" {
		pid, err := strconv.ParseInt(fields[1], 10, 32)
		if err != nil {
			return nil, &ServiceError{u.Name(), serviceName, err}
		}
		return &ProcessStatus{int(pid), Up}, nil
	}

	if len(u.dummyOutput2) != 0 {
		sout = []byte(u.dummyOutput2)
	} else {
		cmd := exec.Command("systemctl", "is-enabled", serviceName)
		sout, err = cmd.CombinedOutput()
	}
	if err != nil || string(sout) != "enabled\n" {
		return nil, &ServiceError{u.Name(), serviceName, ErrServiceNotFound}
	} else {
		return &ProcessStatus{0, Down}, nil
	}
}
