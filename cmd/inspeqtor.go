package main

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"flag"
	"inspeqtor"
	"inspeqtor/util"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type cliOptions struct {
	verbose         bool
	testConfig      bool
	configDirectory string
}

func main() {
	// Modern POSIX processes should log to STDOUT only
	// and use a smart init system to manage logging.  That
	// logging system should add things like PID, timestamp, etc
	// to the logging output so we don't add any context at all.
	log.SetOutput(os.Stdout)
	log.SetFlags(0)

	options := parseArguments()
	if options.verbose {
		util.Verbose = true
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
		util.Info("Licensed to %s <%s>, maximum of %s hosts.",
			license.Name, license.Email, license.Hosts)
	}

	ins, err := inspeqtor.New(options.configDirectory)
	if err != nil {
		log.Fatalln(err)
	}
	err = ins.Parse()
	if err != nil {
		log.Fatalln(err)
	}

	if options.testConfig {
		util.Info("Configuration parsed ok.")
		os.Exit(0)
	} else {
		ins.Start()
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

	util.DebugDebug("%v", lic)
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
	flag.StringVar(&defaults.configDirectory, "c", "/etc/inspeqtor", "Configuration directory")
	helpPtr := flag.Bool("help", false, "You're looking at it")
	help2Ptr := flag.Bool("h", false, "You're looking at it")
	flag.Parse()

	if *helpPtr || *help2Ptr {
		log.Println("inspeqtor", inspeqtor.VERSION)
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
