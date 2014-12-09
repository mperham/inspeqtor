package services

import (
	"errors"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/mperham/inspeqtor/util"
)

type Upstart struct {
	path        string
	dummyOutput *string
}

var (
	pidScanner = regexp.MustCompile(" (start|stop)\\/([a-z\\-]+)(?:, process (\\d+))?")
)

func detectUpstart(path string) (InitSystem, error) {
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
		return &Upstart{path, nil}, nil
	}

	util.Debug("upstart not detected, empty " + path)
	return nil, nil
}

func (u *Upstart) serviceCommand(serviceName string, command string, timeout time.Duration, expectedLines int) error {
	var err error
	var sout []byte

	if u.dummyOutput != nil {
		sout = []byte(*u.dummyOutput)
	} else {
		cmd := exec.Command("initctl", command, serviceName)
		sout, err = util.SafeRun(cmd, timeout)
		if err != nil {
			return &ServiceError{u.Name(), serviceName, err}
		}
	}

	lines, err := util.ReadLines(sout)
	if len(lines) != expectedLines {
		return &ServiceError{u.Name(), serviceName, errors.New("Unexpected output: " + strings.Join(lines, "\n"))}
	}
	return nil
}

func (u *Upstart) Name() string {
	return "upstart"
}

func (u *Upstart) LookupService(serviceName string) (*ProcessStatus, error) {
	matches, err := filepath.Glob(u.path + "/" + serviceName + ".conf")
	if err != nil {
		return nil, &ServiceError{u.Name(), serviceName, err}
	}

	if len(matches) == 0 {
		return nil, &ServiceError{u.Name(), serviceName, ErrServiceNotFound}
	}

	var sout []byte
	if u.dummyOutput != nil {
		sout = []byte(*u.dummyOutput)
	} else {
		cmd := exec.Command("initctl", "status", serviceName)
		sout, err = util.SafeRun(cmd)
		if err != nil {
			return nil, &ServiceError{u.Name(), serviceName, err}
		}
	}

	lines, err := util.ReadLines(sout)
	if len(lines) != 1 {
		return nil, &ServiceError{u.Name(), serviceName, errors.New("Unexpected output: " + strings.Join(lines, "\n"))}
	}

	// mysql start/running, process 14190
	// sshdgenkeys stop/waiting
	line := lines[0]
	if strings.Contains(line, "Unknown job") {
		return nil, &ServiceError{u.Name(), serviceName, ErrServiceNotFound}
	}

	results := pidScanner.FindStringSubmatch(line)

	if len(results) == 4 && len(results[3]) > 0 {
		pid, err := strconv.ParseInt(results[3], 10, 32)
		if err != nil {
			return nil, &ServiceError{u.Name(), serviceName, err}
		}
		return &ProcessStatus{int(pid), Up}, nil
	}

	if len(results) == 3 {
		switch {
		case results[1] == "start":
			return &ProcessStatus{0, Starting}, nil
		case results[1] == "stop":
			return &ProcessStatus{0, Down}, nil
		}
	}

	return nil, &ServiceError{u.Name(), serviceName, errors.New("Unknown upstart output: " + line)}
}

func (u *Upstart) Restart(serviceName string) error {
	return u.serviceCommand(serviceName, "restart", util.RestartTimeout, 1)
}

func (u *Upstart) Reload(serviceName string) error {
	return u.serviceCommand(serviceName, "reload", util.CmdTimeout, 0)
}
