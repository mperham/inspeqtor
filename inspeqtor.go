package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"flag"
	"fmt"
	"inspeqtor/darwin"
	"inspeqtor/linux"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"os/signal"
	"strings"
	"time"
)

var (
	VERSION = "1.0.0"
)

type cliOptions struct {
	verbose         bool
	testConfig      bool
	configDirectory string
}

func main() {
	// Modern POSIX applications should log to STDOUT only
	// and use a smart init system to manage logging.
	log.SetOutput(os.Stdout)
	log.SetPrefix("inspeqtor ")

	options := parseArguments()
	if options.verbose {

	}

	result, err := FileExists("license.yml")
	if err != nil {
		panic(err)
	}

	if result {
		license, err := verifyLicense()
		if err != nil {
			log.Println(err)
			os.Exit(120)
		}
		log.Println(fmt.Sprintf("Licensed to %s <%s>, maximum of %s hosts.",
			license.Name, license.Email, license.Hosts))
	}

	serviceMapping := make(map[string]int)

	launchctl, err := darwin.DetectLaunchctl()
	if err != nil {
		panic(err)
	}

	if launchctl != nil {
		services := []string{"homebrew.mxcl.memcached", "bob"}
		for _, service := range services {
			name, pid, err := launchctl.FindService(service)
			if err != nil {
				log.Println("Couldn't find service " + service + ", skipping...")
			} else {
				serviceMapping[name] = pid
			}
		}
	}

	upstart, err := linux.DetectUpstart("/etc/init")
	if err != nil {
		panic(err)
	}

	if upstart != nil {
		services := []string{"mysql", "pass", "bob"}
		for _, service := range services {
			name, pid, err := upstart.FindService(service)
			if err != nil {
				log.Fatalln(err.Error())
			} else {
				serviceMapping[name] = pid
			}
		}
	}

	log.Println(serviceMapping)

	shutdown := make(chan int)
	go pollSystem(shutdown)

	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt)
	<-signals

	shutdown <- 1
	log.Println("Complete, bye!")
}

func pollSystem(shutdown chan int) {
	scanSystem()
	select {
	case <-shutdown:
		log.Println("Exiting poll loop")
		return
	case <-time.After(30 * time.Second):
		scanSystem()
	}
}

func scanSystem() {
	log.Println("Scanning...")
	metrics, err := linux.CollectHostMetrics("/proc")
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println(metrics)
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
	Hosts string
}

func verifyLicense() (*License, error) {
	lic := make(map[string]string)
	b, err := ioutil.ReadFile("license.yml")
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		kv := strings.Split(line, ":")
		if len(kv) < 2 {
			continue
		}
		lic[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
	}

	log.Println(lic)
	l := License{
		Name:  lic["name"],
		Email: lic["email"],
		Hosts: lic["hosts"],
	}

	if len(lic["key"]) != 64 {
		return nil, errors.New("Invalid license key")
	}
	if lic["v"] == "1" {
		cat := []byte("TastySalt" + l.Name + l.Email +
			l.Hosts + lic["nonce"])
		hash := sha256.Sum256(cat)
		should_be := hex.EncodeToString(hash[:])
		//    fmt.Println(should_be)
		if lic["key"] == should_be {
			return &l, nil
		} else {
			return nil, errors.New("Invalid license")
		}
	} else {
		return nil, errors.New("Invalid license")
	}
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
		log.Println("inspeqtor", VERSION)
		log.Println("Copyright (c) 2014 Contributed Systems LLC")
		log.Println("")
		log.Println("-c [dir]\tConfiguration directory")
		log.Println("-t\t\tVerify configuration and exit")
		log.Println("-v\t\tPrint version and exit")
		log.Println("")
		log.Println("-h\t\tYou're looking at it")
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
