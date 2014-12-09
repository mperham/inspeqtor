package services

import (
	"errors"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/mperham/inspeqtor/util"
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

func (s *Systemd) serviceCommand(serviceName string, command string, timeout time.Duration) error {
	if len(s.dummyOutput) == 0 {
		cmd := exec.Command("systemctl", command, serviceName)
		_, err := util.SafeRun(cmd, timeout)
		if err != nil {
			return &ServiceError{s.Name(), serviceName, err}
		}
	}

	return nil
}

func (s *Systemd) Name() string {
	return "systemd"
}

func (s *Systemd) LookupService(serviceName string) (*ProcessStatus, error) {
	var sout []byte
	var err error

	if len(s.dummyOutput) != 0 {
		sout = []byte(s.dummyOutput)
	} else {
		cmd := exec.Command("systemctl", "show", "-p", "MainPID", serviceName)
		sout, err = util.SafeRun(cmd)
	}

	if err != nil {
		return nil, &ServiceError{s.Name(), serviceName, ErrServiceNotFound}
	}
	lines, err := util.ReadLines(sout)
	if len(lines) != 1 {
		return nil, &ServiceError{s.Name(), serviceName, errors.New("Unexpected output: " + strings.Join(lines, "\n"))}
	}

	// Output will be "MainPID=1234" or
	// "MainPID=0" if service does not exist.
	line := lines[0]
	fields := strings.Split(line, "=")
	if fields[1] != "0" {
		pid, err := strconv.ParseInt(fields[1], 10, 32)
		if err != nil {
			return nil, &ServiceError{s.Name(), serviceName, err}
		}
		return &ProcessStatus{int(pid), Up}, nil
	}

	if len(s.dummyOutput2) != 0 {
		sout = []byte(s.dummyOutput2)
	} else {
		cmd := exec.Command("systemctl", "is-enabled", serviceName)
		sout, err = util.SafeRun(cmd)
	}

	if err != nil || string(sout) != "enabled\n" {
		return nil, &ServiceError{s.Name(), serviceName, ErrServiceNotFound}
	}
	return &ProcessStatus{0, Down}, nil
}

func (s *Systemd) Restart(serviceName string) error {
	return s.serviceCommand(serviceName, "restart", util.RestartTimeout)
}

func (s *Systemd) Reload(serviceName string) error {
	return s.serviceCommand(serviceName, "reload", util.CmdTimeout)
}
