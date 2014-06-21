package main

import (
	"os"
	"flag"
	"fmt"
	"bufio"
	"bytes"
	"strings"
	"strconv"
	"io/ioutil"
	"net/smtp"
	"os/exec"
  "gopkg.in/yaml.v1"
)

var VERSION = "1.0.0"

type Configuration struct {
  Services []string
}

type cliOptions struct {
  testConfig bool
  configDirectory string
}

// - parse cmd line flags
// - open logger
// - parse configuration
// - locate processes
func main() {
  var config Configuration

  parseArguments()

  b, err := ioutil.ReadFile("inspeqtor.yml")
  if err != nil { panic(err) }

  err = yaml.Unmarshal(b, &config)
  if err != nil { panic(err) }

  result, err := FileExists("/mach_kernel")
  if err != nil { panic(err) }

  if result {
    data := LaunchctlResolve(config.Services)
    auth := smtp.PlainAuth("", "mperham", "", "smtp.gmail.com")
    err := smtp.SendMail("smtp.gmail.com:587", auth,
            "mperham@gmail.com",
            []string{"mperham@gmail.com"},
            bytes.NewBufferString(fmt.Sprint(data)).Bytes())
    if err != nil { panic(err) }
  }
}

func parseArguments() cliOptions {
  defaults := cliOptions{}
  flag.BoolVar(&defaults.testConfig, "t", false, "Verify configuration and exit")
  flag.StringVar(&defaults.configDirectory, "c", "/etc/inspeqtor", "Configuration directory")
  verPtr := flag.Bool("v", false, "Print version and exit")
  helpPtr := flag.Bool("help", false, "You're looking at it")
  help2Ptr := flag.Bool("h", false, "You're looking at it")
  flag.Parse()

  if *verPtr {
    fmt.Println("inspeqtor", VERSION)
    os.Exit(0)
  }

  if *helpPtr || *help2Ptr {
    fmt.Println("inspeqtor", VERSION)
    fmt.Println("Copyright (c) 2014 Q Systems Corp")
    fmt.Println("")
    fmt.Println("-c [dir]\tConfiguration directory")
    fmt.Println("-t\t\tVerify configuration and exit")
    fmt.Println("-v\t\tPrint version and exit")
    fmt.Println("")
    fmt.Println("-h\t\tYou're looking at it")
    os.Exit(0)
  }

  return defaults
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

func LaunchctlResolve(services []string) map[string]int {
  cmd := exec.Command("launchctl", "list")
  sout, err := cmd.CombinedOutput()
  if err != nil { panic(err) }

  lines, err := readLines(sout)

  data := make(map[string]int)
  for _, sname := range(services) {
    var found = false
    for _, line := range(lines) {
      if strings.Contains(line, sname) {
        fmt.Println("Found " + sname)
        parts := strings.SplitN(line, "\t", 3)
        pid, err := strconv.Atoi(parts[0])
        if err != nil {
          continue
        }
        pname := parts[len(parts)-1]
        data[pname] = pid
        found = true
        break
      }
    }

    if !found { panic("Couldn't find " + sname) }
  }
  return data
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
