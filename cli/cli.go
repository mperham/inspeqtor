package cli

import (
	"flag"
	"inspeqtor"
	"inspeqtor/util"
	"log"
	"os"
)

type CmdOptions struct {
	TestConfig      bool
	TestAlertRoutes bool
	ConfigDirectory string
	LogLevel        string
	SocketPath      string
}

var (
	SalesPitch = func() {
		log.Println("Want more? Upgrade to Inspeqtor Pro for more features and official support.")
		log.Println("See http://contribsys.com/inspeqtor for details.")
		log.Println("")
	}
)

func SetupLogging() {
	// Modern POSIX processes should log to STDOUT only
	// and use a smart init system to manage logging.  That
	// logging system should add things like PID, timestamp, etc
	// to the logging output so we don't add any context at all.
	log.SetOutput(os.Stdout)
	log.SetFlags(0)
}

func ParseArguments() CmdOptions {
	defaults := CmdOptions{false, false, "/etc/inspeqtor", "info", "/var/run/inspeqtor.sock"}

	log.Println(inspeqtor.Name, inspeqtor.VERSION)
	log.Println("Copyright © 2014 Contributed Systems LLC")

	if SalesPitch != nil {
		log.Println(inspeqtor.Licensing)
		log.Println("")
		SalesPitch()
	}

	flag.Usage = help
	flag.BoolVar(&defaults.TestConfig, "tc", false, "Verify configuration and exit")
	flag.BoolVar(&defaults.TestAlertRoutes, "ta", false, "Verify alert routes and exit")
	flag.StringVar(&defaults.LogLevel, "l", "info", "Logging level (warn, info, debug, verbose)")

	// undocumented on purpose, for testing only
	flag.StringVar(&defaults.SocketPath, "s", "/var/run/inspeqtor.sock", "")
	flag.StringVar(&defaults.ConfigDirectory, "c", "/etc/inspeqtor", "")
	helpPtr := flag.Bool("help", false, "You're looking at it")
	help2Ptr := flag.Bool("h", false, "You're looking at it")
	flag.Parse()

	if *helpPtr || *help2Ptr {
		help()
		os.Exit(0)
	}

	util.SetLogLevel(defaults.LogLevel)

	return defaults
}

func help() {
	log.Println("-l [level]\tSet logging level (warn, info, debug, verbose), default: info")
	log.Println("-tc\t\tVerify configuration and exit")
	log.Println("-ta\t\tVerify alert routes and exit")
}
