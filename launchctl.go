package main

import (
  "os/exec"
  "fmt"
  "errors"
  "strings"
  "strconv"
)

type Launchctl struct {}


func (l *Launchctl) FindService(serviceName string) (string, int, error) {
  cmd := exec.Command("launchctl", "list")
  sout, err := cmd.CombinedOutput()
  if err != nil { return "", 0, err }

  lines, err := readLines(sout)
  if err != nil { return "", 0, err }

  for _, line := range(lines) {
    if strings.Contains(line, serviceName) {
      fmt.Println("Found " + serviceName)
      parts := strings.SplitN(line, "\t", 3)
      pid, err := strconv.Atoi(parts[0])
      if err != nil { return "", 0, err }

      pname := parts[len(parts)-1]
      return pname, pid, nil
    }
  }

  return "", 0, errors.New("Couldn't find " + serviceName)
}


