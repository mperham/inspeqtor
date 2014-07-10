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
package linux

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

var (
	expectedPaths = []string{"/etc/service", "/service"}
)

type Runit struct {
	path string
}

func DetectRunit() (*Runit, error) {
	result, err := FileExists("/etc/sv")
	if err != nil {
		return nil, err
	}

	if !result {
		log.Println("runit not detected, no /etc/sv")
		return nil, nil
	}

	for _, path := range expectedPaths {
		matches, err := filepath.Glob(path + "/*/run")
		if err != nil {
			return nil, err
		}

		if len(matches) > 0 {
			log.Println("Detected runit in " + path)
			return &Runit{path}, nil
		}
	}
	return nil, nil
}

func (r *Runit) FindService(serviceName string) (string, int, error) {
	matches, err := filepath.Glob(r.path + "/" + serviceName + "/run")
	if err != nil {
		return "", 0, err
	}

	if len(matches) == 0 {
		return "", 0, errors.New("No service matching " + serviceName + " was found in " + r.path)
	}

	content, err := ioutil.ReadFile(r.path + "/" + serviceName + "/supervise/pid")
	if len(content) == 0 {
		// service exists but is not running
		return serviceName, 0, nil
	}
	pid, err := strconv.Atoi(string(content))
	if err != nil {
		return "", 0, err
	}

	return serviceName, pid, nil
}
func FileExists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
