package main

import (
	"inspeqtor"
	"inspeqtor/cli"
	"inspeqtor/util"
	"log"
	"os"
)

func main() {
	cli.SetupLogging()
	options := cli.ParseArguments("Inspeqtor", inspeqtor.VERSION)

	ins, err := inspeqtor.New(options.ConfigDirectory)
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
	} else if options.TestNotifications {
		ins.TestNotifications()
	} else {
		ins.Start()
	}
}
