package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mperham/inspeqtor"
	"github.com/mperham/inspeqtor/util"
)

type CmdOptions struct {
	TestConfig      bool
	TestAlertRoutes bool
	ConfigDirectory string
	LogLevel        string
	SocketPath      string
}

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
	log.Println(fmt.Sprintf("Copyright © %d Contributed Systems LLC", time.Now().Year()))
	log.Println(inspeqtor.Licensing)

	flag.Usage = help
	flag.BoolVar(&defaults.TestConfig, "tc", false, "Test configuration and exit")
	flag.BoolVar(&defaults.TestAlertRoutes, "ta", false, "Test alert routes and exit")
	flag.StringVar(&defaults.LogLevel, "l", "info", "Logging level (warn, info, debug, verbose)")

	// undocumented on purpose, for testing only, we don't want people changing these
	// if possible
	flag.StringVar(&defaults.SocketPath, "s", "/var/run/inspeqtor.sock", "")
	flag.StringVar(&defaults.ConfigDirectory, "c", "/etc/inspeqtor", "")
	helpPtr := flag.Bool("help", false, "You're looking at it")
	help2Ptr := flag.Bool("h", false, "You're looking at it")
	versionPtr := flag.Bool("v", false, "Show version")
	flag.Parse()

	if *helpPtr || *help2Ptr {
		help()
		os.Exit(0)
	}

	if *versionPtr {
		os.Exit(0)
	}

	util.SetLogLevel(defaults.LogLevel)

	return defaults
}

func help() {
	log.Println("-l [level]\tSet logging level (warn, info, debug, verbose), default: info")
	log.Println("-tc\t\tTest configuration and exit")
	log.Println("-ta\t\tTest alert routes and exit")
	log.Println("-v\t\tShow version and license information")
	log.Println("-h\t\tThis help screen")
}
