// Usually found in /etc/service, /etc/sv or /service with subdirectories:
//    <service_name>/
//        run
//        log/
//           run
//        supervise/
//           pid  # => 4994
//           stat # => run / down
//
package main

import (
  "path/filepath"
  "fmt"
  "errors"
  "io/ioutil"
  "strconv"
)

var (
  expectedPaths = []string { "/etc/sv", "/etc/service", "/service" }
)

type Runit struct{
  path string
}

func DetectRunit() (*Runit, error) {
  for _, path := range(expectedPaths) {
    matches, err := filepath.Glob(path + "/*/run")
    if err != nil { return nil, err }

    if len(matches) > 0 {
      fmt.Println("Detected runit in " + path)
      return &Runit{ path }, nil
    }
  }
  return nil, nil
}


func (r *Runit) FindService(serviceName string) (string, int, error) {
  matches, err := filepath.Glob(r.path + "/" + serviceName + "/run")
  if err != nil { return "", 0, err }

  if len(matches) == 0 {
    return "", 0, errors.New("No service matching " + serviceName + " was found in " + r.path)
  }

  content, err := ioutil.ReadFile(r.path + "/" + serviceName + "/supervise/pid")
  if len(content) == 0 {
    // service exists but is not running
    return serviceName, 0, nil
  }
  pid, err := strconv.Atoi(string(content))
  if err != nil { return "", 0, err }

  return serviceName, pid, nil
}

