package main

import (
	"github.com/mperham/redacted"
	"github.com/mperham/redacted/cli"
	"github.com/mperham/redacted/util"
	"log"
	"os"
)

func main() {
	cli.SetupLogging()
	options := cli.ParseArguments()

	ins, err := redacted.New(options.ConfigDirectory, options.SocketPath)
	if err != nil {
		log.Fatalln(err)
	}
	err = ins.Parse()
	if err != nil {
		log.Fatalln(err)
	}

	if options.TestConfig {
		util.Info("Configuration parsed ok.")
		os.Exit(0)
	} else if options.TestAlertRoutes {
		ins.TestAlertRoutes()
	} else {
		// Fire up the Redacted singleton
		ins.Start()

		// Install the global signal handlers
		// This method never returns.
		redacted.HandleSignals()
	}
}
