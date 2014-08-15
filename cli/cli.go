package cli

import (
	"flag"
	"inspeqtor/util"
	"log"
	"os"
)

type CmdOptions struct {
	TestConfig        bool
	TestNotifications bool
	ConfigDirectory   string
	LogLevel          string
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
	defaults := CmdOptions{false, false, "/etc/inspeqtor", "info"}

	flag.BoolVar(&defaults.TestConfig, "t", false, "Verify configuration and exit")
	flag.BoolVar(&defaults.TestNotifications, "tn", false, "Verify notifications and exit")
	flag.StringVar(&defaults.ConfigDirectory, "c", "/etc/inspeqtor", "Configuration directory")
	flag.StringVar(&defaults.LogLevel, "l", "info", "Logging level (warn, info, debug, verbose)")
	helpPtr := flag.Bool("help", false, "You're looking at it")
	help2Ptr := flag.Bool("h", false, "You're looking at it")
	flag.Parse()

	log.Println(name, version)
	log.Println("Copyright © 2014 Contributed Systems LLC")
	log.Println("Licensed under the GNU Public License Version 3")
	log.Println("")
	log.Println("Want more? Upgrade to Inspeqtor Pro for more features and official support.")
	log.Println("See http://contribsys.com/inspeqtor for details.")
	log.Println("")

	if *helpPtr || *help2Ptr {
		log.Println("")
		log.Println("-c [dir]\tConfiguration directory, default: /etc/inspeqtor")
		log.Println("-l [level]\tSet logging level (warn, info, debug, verbose), default: info")
		log.Println("-t\t\tVerify configuration and exit")
		log.Println("-tn\t\tVerify notifications and exit")
		log.Println("")
		log.Println("-h\t\tYou're looking at it")
		os.Exit(0)
	}

	util.SetLogLevel(defaults.LogLevel)

	return defaults
}
