package services

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"

	"github.com/mperham/inspeqtor/util"
)

type Launchd struct {
	dirs []string
}

func detectLaunchd(rootDir string) (InitSystem, error) {
	if !util.Darwin() {
		return nil, nil
	}
	util.Info("Detected OSX, using launchd")

	usr, err := user.Current()
	if err != nil {
		return nil, err
	}

	dir := usr.HomeDir
	paths := []string{
		dir + "/Library/LaunchAgents",
		"/Library/LaunchAgents",
		"/Library/LaunchDaemons",
		"/System/Library/LaunchDaemons",
	}
	return &Launchd{paths}, nil
}

func (l *Launchd) resolvePlist(serviceName string) string {
	for _, path := range l.dirs {
		candidate := fmt.Sprintf("%s/%s.plist", path, serviceName)
		_, err := os.Lstat(candidate)
		if err == nil {
			return candidate
		}
	}
	return ""
}

func (l *Launchd) Name() string {
	return "launchd"
}

func (l *Launchd) LookupService(serviceName string) (*ProcessStatus, error) {
	cmd := exec.Command("launchctl", "list")
	sout, err := util.SafeRun(cmd)
	if err != nil {
		return nil, &ServiceError{l.Name(), serviceName, err}
	}

	lines, err := util.ReadLines(sout)
	if err != nil {
		return nil, &ServiceError{l.Name(), serviceName, err}
	}

	for _, line := range lines {
		if strings.Contains(line, serviceName) {
			util.Debug("launchctl found " + serviceName)
			parts := strings.SplitN(line, "\t", 3)
			pid, err := strconv.ParseInt(parts[0], 10, 32)
			if err != nil {
				return nil, &ServiceError{l.Name(), serviceName, err}
			}

			return &ProcessStatus{int(pid), Up}, nil
		}
	}

	path := l.resolvePlist(serviceName)
	if path != "" {
		return &ProcessStatus{0, Down}, nil
	}

	return nil, &ServiceError{l.Name(), serviceName, ErrServiceNotFound}
}

func (l *Launchd) Restart(serviceName string) error {
	path := l.resolvePlist(serviceName)
	if path == "" {
		return &ServiceError{l.Name(), serviceName, ErrServiceNotFound}
	}

	cmd := exec.Command("launchctl", "unload", path)
	sout, err := util.SafeRun(cmd)
	if err != nil {
		return &ServiceError{l.Name(), serviceName, err}
	}

	lines, err := util.ReadLines(sout)
	if len(lines) != 0 {
		return &ServiceError{l.Name(), serviceName, errors.New("Unexpected output: " + strings.Join(lines, "\n"))}
	}

	cmd = exec.Command("launchctl", "load", path)
	sout, err = util.SafeRun(cmd)
	if err != nil {
		return &ServiceError{l.Name(), serviceName, err}
	}

	lines, err = util.ReadLines(sout)
	if len(lines) != 0 {
		return &ServiceError{l.Name(), serviceName, errors.New("Unexpected output: " + strings.Join(lines, "\n"))}
	}
	return nil
}

func (l *Launchd) Reload(serviceName string) error {
	return &ServiceError{l.Name(), serviceName, errors.New("Reload isn't supported")}
}
