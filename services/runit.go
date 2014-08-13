// runit manages services usually found in /etc/service or /service, which
// are soft links to the actual service directories in /etc/sv:
//    <service_name>/
//        run
//        log/
//           run
//        supervise/
//           pid  # => 4994
//           stat # => run / down
//
package services

import (
	"errors"
	"inspeqtor/util"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type Runit struct {
	path        string
	dummyOutput string
}

func detectRunit(root string) (InitSystem, error) {
	path := root + "etc/service"
	result, err := util.FileExists(path)
	if err != nil {
		return nil, err
	}

	if !result {
		path = root + "service"
		result, err = util.FileExists(path)
		if err != nil {
			return nil, err
		}
	}

	if !result {
		util.Debug("runit not detected in /etc/service or /service")
		return nil, nil
	}

	matches, err := filepath.Glob(path + "/*/run")
	if err != nil {
		return nil, err
	}

	if len(matches) > 0 {
		util.Info("Detected runit in " + path)
		return &Runit{path, ""}, nil
	}

	return nil, nil
}

func (r *Runit) Name() string {
	return "runit"
}

func (r *Runit) Restart(serviceName string) error {
	out := []byte{}

	if r.dummyOutput == "" {
		cmd := exec.Command("sv", "restart", serviceName)
		sout, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		out = sout
	} else {
		out = []byte(r.dummyOutput)
		r.dummyOutput = ""
	}

	lines, err := util.ReadLines(out)
	if err != nil {
		return err
	}
	if len(lines) != 1 {
		return errors.New("Unexpected output: " + strings.Join(lines, "\n"))
	}
	return nil
}

func (r *Runit) LookupService(serviceName string) (ProcessId, Status, error) {
	matches, err := filepath.Glob(r.path + "/" + serviceName + "/run")
	if err != nil {
		return 0, Unknown, err
	}

	if len(matches) == 0 {
		return -1, Unknown, nil
	}

	content, err := ioutil.ReadFile(r.path + "/" + serviceName + "/supervise/pid")
	if len(content) == 0 {
		// service exists but is not running
		return 0, Down, nil
	}
	pid, err := strconv.ParseInt(strings.TrimSpace(string(content)), 10, 32)
	if err != nil {
		return 0, Unknown, err
	}

	return ProcessId(pid), Up, nil
}
