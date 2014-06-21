package main

import (
	"os"
	"crypto/sha256"
	"flag"
	"fmt"
	"bufio"
	"errors"
	"bytes"
	"strings"
	"strconv"
	"io/ioutil"
	"encoding/hex"
	"net/smtp"
	"os/exec"
  "gopkg.in/yaml.v1"
  "github.com/blackjack/syslog"
)

var VERSION = "1.0.0"

type cliOptions struct {
  testConfig bool
  configDirectory string
}

// - parse cmd line flags
// - open logger
// - parse configuration
// - locate processes
func main() {

  options := parseArguments()

  syslog.Openlog("inspeqtor", syslog.LOG_PID, syslog.LOG_USER)
  syslog.Noticef("Parsed options: %s", options)

  defer syslog.Closelog()

  result, err := FileExists("license.yml")
  if err != nil { panic(err) }

  if result {
    license, err := verifyLicense()
    if err != nil {
      fmt.Println(err)
      os.Exit(120)
    }
    syslog.Noticef("Licensed to %s <%s>, maximum of %d hosts.",
                   license.Name, license.Email, license.Hosts)
  }

  result, err = FileExists("/mach_kernel")
  if err != nil { panic(err) }

  if result {
    data := LaunchctlResolve([]string{"percona", "redis"})
    os.Exit(0)
    auth := smtp.PlainAuth("", "mperham", "", "smtp.gmail.com")
    err := smtp.SendMail("smtp.gmail.com:587", auth,
            "mperham@gmail.com",
            []string{"mperham@gmail.com"},
            bytes.NewBufferString(fmt.Sprint(data)).Bytes())
    if err != nil { panic(err) }
  }
}

type License struct {
  Name string
  Email string
  Hosts int
  Key string
  Nonce int
  V int
}

func (lic *License) verify() (*License, error) {
  if len(lic.Key) != 64 {
    return nil, errors.New("Invalid license")
  }
  if lic.V == 1 {
    cat := []byte("TastySalt" + lic.Name + lic.Email +
          strconv.Itoa(lic.Hosts) +
          strconv.Itoa(lic.Nonce))
    hash := sha256.Sum256(cat)
    should_be := hex.EncodeToString(hash[:])
//    fmt.Println(should_be)
    if lic.Key == should_be {
      return lic, nil
    } else {
      return nil, errors.New("Invalid license")
    }
  } else {
    panic("Unknown license format")
  }
}

func verifyLicense() (*License, error) {
  var license License
  b, err := ioutil.ReadFile("license.yml")
  if err != nil { return nil, err }

  err = yaml.Unmarshal(b, &license)
  if err != nil { return nil, err }

  fmt.Println(license)
  return license.verify()
}

func parseArguments() cliOptions {
  defaults := cliOptions{}
  flag.BoolVar(&defaults.testConfig, "t", false, "Verify configuration and exit")
  flag.StringVar(&defaults.configDirectory, "c", ".", "Configuration directory")
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
