package init

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type Launchctl struct {
}

func DetectLaunchctl(rootDir string) (*Launchctl, error) {
	file, err := fileExists(rootDir + "mach_kernel")
	if err != nil {
		return nil, err
	}
	if !file {
		return nil, nil
	}
	return &Launchctl{}, nil
}

func (l *Launchctl) FindServicePID(serviceName string) (int32, error) {
	cmd := exec.Command("launchctl", "list")
	sout, err := cmd.CombinedOutput()
	if err != nil {
		return 0, err
	}

	lines, err := readLines(sout)
	if err != nil {
		return 0, err
	}

	for _, line := range lines {
		if strings.Contains(line, serviceName) {
			log.Println("Found " + serviceName)
			parts := strings.SplitN(line, "\t", 3)
			pid, err := strconv.ParseInt(parts[0], 10, 32)
			if err != nil {
				return 0, err
			}

			return int32(pid), nil
		}
	}

	return -1, nil
}
