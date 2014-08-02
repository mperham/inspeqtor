package services

import (
	"inspeqtor/util"
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
