package cli

import (
	"flag"
	"inspeqtor/util"
	"log"
	"os"
)

type CmdOptions struct {
	Verbose         bool
	TestConfig      bool
	ConfigDirectory string
}

func SetupLogging() {
	// Modern POSIX processes should log to STDOUT only
	// and use a smart init system to manage logging.  That
	// logging system should add things like PID, timestamp, etc
	// to the logging output so we don't add any context at all.
	log.SetOutput(os.Stdout)
	log.SetFlags(0)
}

func ParseArguments(name string, version string) CmdOptions {
	defaults := CmdOptions{false, false, "/etc/inspeqtor"}

	flag.BoolVar(&defaults.Verbose, "v", false, "Enable verbose logging")
	flag.BoolVar(&defaults.TestConfig, "t", false, "Verify configuration and exit")
	flag.StringVar(&defaults.ConfigDirectory, "c", "/etc/inspeqtor", "Configuration directory")
	helpPtr := flag.Bool("help", false, "You're looking at it")
	help2Ptr := flag.Bool("h", false, "You're looking at it")
	flag.Parse()

	if *helpPtr || *help2Ptr {
		log.Println(name, version)
		log.Println("Copyright (c) 2014 Contributed Systems LLC")
		log.Println("")
		log.Println("Upgrading to Inspeqtor Pro gives you more features and dedicated support.")
		log.Println("See http://contribsys.com/inspeqtor for details.")
		log.Println("")
		log.Println("-c [dir]\tConfiguration directory")
		log.Println("-t\t\tVerify configuration and exit")
		log.Println("-v\t\tEnable verbose logging")
		log.Println("")
		log.Println("-h\t\tYou're looking at it")
		os.Exit(0)
	}

	if defaults.Verbose {
		util.Verbose = true
	}

	return defaults
}
