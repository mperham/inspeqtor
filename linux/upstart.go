package linux

import (
	"bufio"
	"bytes"
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type Upstart struct {
	Path string
}

var (
	pidScanner *regexp.Regexp = regexp.MustCompile(" (?:start|stop)\\/(?:running|waiting)(?:, process (\\d+))?")
)

func DetectUpstart(path string) (*Upstart, error) {
	result, err := fileExists(path)
	if err != nil {
		return nil, err
	}

	if !result {
		log.Println("upstart not detected, no " + path)
		return nil, nil
	}

	matches, err := filepath.Glob(path + "/*.conf")
	if err != nil {
		return nil, err
	}

	if len(matches) > 0 {
		log.Println("Detected upstart in " + path)
		return &Upstart{path}, nil
	}

	log.Println("upstart not detected, empty " + path)
	return nil, nil
}

func (u *Upstart) FindService(serviceName string) (string, int, error) {
	matches, err := filepath.Glob(u.Path + "/" + serviceName + ".conf")
	if err != nil {
		return "", 0, err
	}

	if len(matches) == 0 {
		return "", 0, errors.New("No service matching " + serviceName + " was found in /etc/init")
	}

	cmd := exec.Command("status", matches[0])
	sout, err := cmd.CombinedOutput()
	if err != nil {
		return "", 0, err
	}

	lines, err := readLines(sout)
	if len(lines) != 1 {
		return "", 0, errors.New("Unexpected output: " + strings.Join(lines, "\n"))
	}

	// mysql start/running, process 14190
	// sshdgenkeys stop/waiting
	line := lines[0]
	results := pidScanner.FindStringSubmatch(line)
	log.Println(results)

	if len(results) > 1 && len(results[1]) > 0 {
		pid, err := strconv.Atoi(results[1])
		if err != nil {
			return "", 0, err
		}
		return matches[0], pid, nil
	}
	if len(results) > 1 {
		return matches[0], 0, nil
	}

	return "", 0, errors.New("Unknown upstart output: " + line)
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
