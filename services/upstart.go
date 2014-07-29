package services

import (
	"errors"
	"inspeqtor/util"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type Upstart struct {
	name        string
	path        string
	dummyOutput string
}

var (
	pidScanner *regexp.Regexp = regexp.MustCompile(" (start|stop)\\/([a-z\\-]+)(?:, process (\\d+))?")
)

func DetectUpstart(path string) (InitSystem, error) {
	result, err := util.FileExists(path)
	if err != nil {
		return nil, err
	}

	if !result {
		util.Debug("upstart not detected, no " + path)
		return nil, nil
	}

	matches, err := filepath.Glob(path + "/*.conf")
	if err != nil {
		return nil, err
	}

	if len(matches) > 0 {
		util.Info("Detected upstart in " + path)
		return Upstart{"upstart", path, ""}, nil
	}

	util.Debug("upstart not detected, empty " + path)
	return nil, nil
}

func (u Upstart) Name() string {
	return u.name
}

func (u Upstart) LookupService(serviceName string) (ProcessId, Status, error) {
	matches, err := filepath.Glob(u.path + "/" + serviceName + ".conf")
	if err != nil {
		return 0, Unknown, err
	}

	if len(matches) == 0 {
		return -1, Unknown, nil
	}

	var sout []byte
	if len(u.dummyOutput) != 0 {
		sout = []byte(u.dummyOutput)
	} else {
		cmd := exec.Command("status", serviceName)
		sout, err = cmd.CombinedOutput()
		if err != nil {
			return 0, Unknown, err
		}
	}

	lines, err := util.ReadLines(sout)
	if len(lines) != 1 {
		return 0, Unknown, errors.New("Unexpected output: " + strings.Join(lines, "\n"))
	}

	// mysql start/running, process 14190
	// sshdgenkeys stop/waiting
	line := lines[0]
	if strings.Contains(line, "Unknown job") {
		return -1, Unknown, nil
	}

	results := pidScanner.FindStringSubmatch(line)

	if len(results) == 4 && len(results[3]) > 0 {
		pid, err := strconv.ParseInt(results[3], 10, 32)
		if err != nil {
			return 0, Unknown, err
		}
		return ProcessId(pid), Up, nil
	}

	if len(results) == 3 {
		switch {
		case results[1] == "start":
			return 0, Starting, nil
		case results[1] == "stop" && results[2] != "waiting":
			return 0, Stopping, nil
		case results[1] == "stop":
			return 0, Down, nil
		}
	}

	return 0, Unknown, errors.New("Unknown upstart output: " + line)
}
