package services

import (
	"errors"
	"fmt"
	"inspeqtor/util"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Launchctl struct {
}

func detectLaunchctl(rootDir string) (InitSystem, error) {
	file, err := util.FileExists(rootDir + "mach_kernel")
	if err != nil {
		return nil, err
	}
	if !file {
		return nil, nil
	}
	util.Info("Found launchctl")
	return Launchctl{}, nil
}

func (l Launchctl) Name() string {
	return "launchctl"
}

var (
	OSX_PATHS = []string{
		"/Library/LaunchAgents",
		"~/Library/LaunchAgents",
		"/Library/LaunchDaemons",
		"/System/Library/LaunchDaemons",
	}
)

func resolvePlist(serviceName string) (string, error) {
	for _, path := range OSX_PATHS {
		candidate := fmt.Sprintf("%s/%s.plist", path, serviceName)
		_, err := os.Lstat(candidate)
		if err == nil {
			return candidate, nil
		}
	}
	return "", errors.New("Could not find a plist for " + serviceName)
}

func (l Launchctl) Restart(serviceName string) error {
	path, err := resolvePlist(serviceName)
	if err != nil {
		return err
	}

	cmd := exec.Command("launchctl", "unload", path)
	sout, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	lines, err := util.ReadLines(sout)
	if len(lines) != 0 {
		return errors.New("Unexpected output: " + strings.Join(lines, "\n"))
	}

	cmd = exec.Command("launchctl", "load", path)
	sout, err = cmd.CombinedOutput()
	if err != nil {
		return err
	}

	lines, err = util.ReadLines(sout)
	if len(lines) != 0 {
		return errors.New("Unexpected output: " + strings.Join(lines, "\n"))
	}
	return nil
}

func (l Launchctl) LookupService(serviceName string) (ProcessId, Status, error) {
	cmd := exec.Command("launchctl", "list")
	sout, err := cmd.CombinedOutput()
	if err != nil {
		return 0, 0, err
	}

	lines, err := util.ReadLines(sout)
	if err != nil {
		return 0, 0, err
	}

	for _, line := range lines {
		if strings.Contains(line, serviceName) {
			util.Debug("Found " + serviceName)
			parts := strings.SplitN(line, "\t", 3)
			pid, err := strconv.ParseInt(parts[0], 10, 32)
			if err != nil {
				return 0, 0, err
			}

			return ProcessId(pid), Up, nil
		}
	}

	return -1, Unknown, nil
}
