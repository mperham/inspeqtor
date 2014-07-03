package main

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"log/syslog"
	"net/smtp"
	"inspeqtor/linux"
	"os"
	"os/signal"
	"os/user"
	"strconv"
	"time"
)

var VERSION = "1.0.0"

type cliOptions struct {
	verbose         bool
	testConfig      bool
	configDirectory string
}

// - parse cmd line flags
// - open logger
// - parse configuration
// - locate processes
func main() {
	u, err := user.Current()
	if u.Uid != "0" {
		fmt.Println(u)
		fmt.Println("Must be root")
		os.Exit(-1)
	}
	options := parseArguments()
	if options.verbose {

	}

	log, err := syslog.New(syslog.LOG_USER|syslog.LOG_NOTICE, "inspeqtor")
	if err != nil {
		panic("Unable to open log: " + err.Error())
	}
	defer log.Close()

	result, err := FileExists("license.yml")
	if err != nil {
		panic(err)
	}

	if result {
		license, err := verifyLicense()
		if err != nil {
			fmt.Println(err)
			os.Exit(120)
		}
		log.Warning(fmt.Sprintf("Licensed to %s <%s>, maximum of %d hosts.",
			license.Name, license.Email, license.Hosts))
	}

	macosx, err := FileExists("/mach_kernel")
	if err != nil {
		panic(err)
	}

	var serviceMapping map[string]int = make(map[string]int)

	if macosx {
		fmt.Println("Detected OSX, using launchctl")
		osx := Launchctl{}
		services := []string{"percona", "redis", "bob"}
		for _, service := range services {
			name, pid, err := osx.FindService(service)
			if err != nil {
				log.Warning("Couldn't find service " + service + ", skipping...")
			} else {
				serviceMapping[name] = pid
			}
		}
	}

	upstart, err := FileExists("/etc/init")
	if err != nil {
		panic(err)
	}

	if upstart {
		fmt.Println("Found upstart in /etc/init")
		upstart := Upstart{}
		services := []string{"mysql", "pass", "bob"}
		for _, service := range services {
			name, pid, err := upstart.FindService(service)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				serviceMapping[name] = pid
			}
		}
	}

	fmt.Println(serviceMapping)

	shutdown := make(chan int)
	go pollSystem(shutdown)

	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt)
	<-signals

	shutdown <- 1
	fmt.Println("Complete, bye!")
}

func pollSystem(shutdown chan int) {
	scanSystem()
	select {
	case <-shutdown:
		fmt.Println("Exiting poll loop")
		return
	case <-time.After(30 * time.Second):
		scanSystem()
	}
}

func scanSystem() {
	fmt.Println("Scanning...")
	metrics, err := linux.CollectHostMetrics("/proc")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(metrics)
	}
}

func sendEmail(data interface{}) {
	auth := smtp.PlainAuth("", "mperham", "", "smtp.gmail.com")
	err := smtp.SendMail("smtp.gmail.com:587", auth,
		"mperham@gmail.com",
		[]string{"mperham@gmail.com"},
		bytes.NewBufferString(fmt.Sprint(data)).Bytes())
	if err != nil {
		panic(err)
	}
}

type License struct {
	Name  string
	Email string
	Hosts int
	Key   string
	Nonce int
	V     int
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
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(b, &license)
	if err != nil {
		return nil, err
	}

	fmt.Println(license)
	return license.verify()
}

func parseArguments() cliOptions {
	defaults := cliOptions{}
	flag.BoolVar(&defaults.verbose, "v", false, "Enable verbose logging")
	flag.BoolVar(&defaults.testConfig, "t", false, "Verify configuration and exit")
	flag.StringVar(&defaults.configDirectory, "c", ".", "Configuration directory")
	helpPtr := flag.Bool("help", false, "You're looking at it")
	help2Ptr := flag.Bool("h", false, "You're looking at it")
	flag.Parse()

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
