package main

import (
  "path/filepath"
  "fmt"
  "errors"
  "regexp"
  "strings"
  "strconv"
  "os"
  "os/exec"
)

type Upstart struct{}

var (
  pidScanner *regexp.Regexp = regexp.MustCompile(" (?:start|stop)\\/(?:running|waiting)(?:, process (\\d+))?")
)

func serviceList(serviceName string) ([]string, error) {
  var matches []string
  var done bool = false

  err := filepath.Walk("/etc/init", func (path string, info os.FileInfo, err error) error {
    if done { return nil }
    if !info.IsDir() {
      if strings.HasPrefix(info.Name(), ".") {
        return nil
      }

      if info.Name() == (serviceName + ".conf") {
        matches = append(matches, serviceName)
        done = true
        return nil
      }
      if strings.Contains(info.Name(), serviceName) {
        name := info.Name()
        fmt.Println("Found " + name)
        matches = append(matches, name[:len(name)-5])
      }
    }
    return nil
  })
  if err != nil { return nil, err }

  return matches, nil
}

func (u *Upstart) FindService(serviceName string) (string, int, error) {
  matches, err := serviceList(serviceName)
  if err != nil { return "", 0, err }

  if len(matches) == 0 {
    return "", 0, errors.New("No service matching " + serviceName + " was found in /etc/init")
  }
  if len(matches) > 1 {
    return "", 0, errors.New("Found multiple services matching " + serviceName + " in /etc/init")
  }

  cmd := exec.Command("status", matches[0])
  sout, err := cmd.CombinedOutput()
  if err != nil { return "", 0, err }

  lines, err := readLines(sout)
  if len(lines) != 1 {
    return "", 0, errors.New("Unexpected output: " + strings.Join(lines, "\n"))
  }

  // mysql start/running, process 14190
  // sshdgenkeys stop/waiting
  line := lines[0]
  results := pidScanner.FindStringSubmatch(line)
  fmt.Println(results)

  if len(results) > 1 && len(results[1]) > 0 {
    pid, err := strconv.Atoi(results[1])
    if err != nil { return "", 0, err }
    return matches[0], pid, nil
  }
  if len(results) > 1 {
    return matches[0], 0, nil
  }

  return "", 0, errors.New("Unknown upstart output: " + line)
}

//func readLines(data []byte) ([]string, error) {
  //var lines []string
  //scan := bufio.NewScanner(bytes.NewReader(data))
  //for scan.Scan() {
    //lines = append(lines, scan.Text())
  //}
  //return lines, scan.Err()
//}
