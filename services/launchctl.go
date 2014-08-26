package services

import (
	"errors"
	"fmt"
	"inspeqtor/util"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
)

type Launchctl struct {
	dirs []string
}

func detectLaunchctl(rootDir string) (InitSystem, error) {
	file, err := util.FileExists(rootDir + "mach_kernel")
	if err != nil {
		return nil, err
	}
	if !file {
		return nil, nil
	}
	util.Info("Detected OSX, using launchctl")

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
	return &Launchctl{paths}, nil
}

func (l *Launchctl) Name() string {
	return "launchctl"
}

func (l *Launchctl) resolvePlist(serviceName string) (string, error) {
	for _, path := range l.dirs {
		candidate := fmt.Sprintf("%s/%s.plist", path, serviceName)
		_, err := os.Lstat(candidate)
		if err == nil {
			return candidate, nil
		}
	}
	return "", errors.New("Could not find a plist for " + serviceName)
}

func (l *Launchctl) Restart(serviceName string) error {
	path, err := l.resolvePlist(serviceName)
	if err != nil {
		return &ServiceError{l.Name(), serviceName, err}
	}

	cmd := exec.Command("launchctl", "unload", path)
	sout, err := cmd.CombinedOutput()
	if err != nil {
		return &ServiceError{l.Name(), serviceName, err}
	}

	lines, err := util.ReadLines(sout)
	if len(lines) != 0 {
		return &ServiceError{l.Name(), serviceName, errors.New("Unexpected output: " + strings.Join(lines, "\n"))}
	}

	cmd = exec.Command("launchctl", "load", path)
	sout, err = cmd.CombinedOutput()
	if err != nil {
		return &ServiceError{l.Name(), serviceName, err}
	}

	lines, err = util.ReadLines(sout)
	if len(lines) != 0 {
		return &ServiceError{l.Name(), serviceName, errors.New("Unexpected output: " + strings.Join(lines, "\n"))}
	}
	return nil
}

func (l *Launchctl) LookupService(serviceName string) (*ProcessStatus, error) {
	cmd := exec.Command("launchctl", "list")
	sout, err := cmd.CombinedOutput()
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

	return nil, &ServiceError{l.Name(), serviceName, ErrServiceNotFound}
}
