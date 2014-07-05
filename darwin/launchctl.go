package darwin

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Launchctl struct {
	Path string
}

func DetectLaunchctl() (*Launchctl, error) {
	file, err := fileExists("/mach_kernel")
	if err != nil {
		return nil, err
	}
	if !file {
		return nil, nil
	}
	return &Launchctl{}, nil
}

func (l *Launchctl) FindService(serviceName string) (string, int, error) {
	cmd := exec.Command("launchctl", "list")
	sout, err := cmd.CombinedOutput()
	if err != nil {
		return "", 0, err
	}

	lines, err := readLines(sout)
	if err != nil {
		return "", 0, err
	}

	for _, line := range lines {
		if strings.Contains(line, serviceName) {
			fmt.Println("Found " + serviceName)
			parts := strings.SplitN(line, "\t", 3)
			pid, err := strconv.Atoi(parts[0])
			if err != nil {
				return "", 0, err
			}

			pname := parts[len(parts)-1]
			return pname, pid, nil
		}
	}

	return "", 0, errors.New("Couldn't find " + serviceName)
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(data []byte) ([]string, error) {
	var lines []string
	scan := bufio.NewScanner(bytes.NewReader(data))
	for scan.Scan() {
		lines = append(lines, scan.Text())
	}
	return lines, scan.Err()
}

func fileExists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
