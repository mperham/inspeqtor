package main

import (
	"log"
	"os"

	"github.com/mperham/inspeqtor"
	_ "github.com/mperham/inspeqtor/channels"
	"github.com/mperham/inspeqtor/cli"
	"github.com/mperham/inspeqtor/expose"
	"github.com/mperham/inspeqtor/jobs"
	_ "github.com/mperham/inspeqtor/ownership"
	"github.com/mperham/inspeqtor/statsd"
	"github.com/mperham/inspeqtor/util"
)

func main() {
	cli.SetupLogging()
	options := cli.ParseArguments()

	ins, err := inspeqtor.New(options.ConfigDirectory, options.SocketPath)
	if err != nil {
		log.Fatalln(err)
	}
	err = ins.Parse()
	if err != nil {
		log.Fatalln(err)
	}

	err = bootstrapJobs(ins, options.ConfigDirectory)
	if err != nil {
		log.Fatalln(err)
	}

	err = bootstrapStatsd(ins, options.ConfigDirectory)
	if err != nil {
		log.Fatalln(err)
	}

	err = expose.Bootstrap(ins)
	if err != nil {
		log.Fatalln(err)
	}

	if options.TestConfig {
		util.Info("Configuration parsed ok.")
		os.Exit(0)
	} else if options.TestAlertRoutes {
		ins.TestAlertRoutes()
	} else {
		ins.Start()
		inspeqtor.HandleSignals()
	}
}

func bootstrapStatsd(ins *inspeqtor.Inspeqtor, cfgDir string) error {
	err := statsdReload(nil, ins)
	if err != nil {
		return err
	}

	inspeqtor.Reloaders = append(inspeqtor.Reloaders, statsdReload)
	return nil
}

func statsdReload(_ *inspeqtor.Inspeqtor, newi *inspeqtor.Inspeqtor) error {
	val, ok := newi.GlobalConfig.Variables["statsd_location"]
	if !ok {
		util.Debug("No statsd_location configured, skipping...")
		return nil
	}

	util.Info("Pushing metrics to statsd at %s", val)
	conn, err := statsd.Dial(val)
	if err != nil {
		return err
	}
	newi.Listen("cycleComplete", func(ins *inspeqtor.Inspeqtor) error {
		return statsd.Export(conn, ins)
	})
	newi.Listen("shutdown", func(ins *inspeqtor.Inspeqtor) error {
		return conn.Close()
	})

	return nil
}

func bootstrapJobs(ins *inspeqtor.Inspeqtor, cfgDir string) error {
	recurring, err := jobs.Parse(ins.GlobalConfig, cfgDir)
	if err != nil {
		return err
	}

	jobs.Watch(ins, recurring)

	inspeqtor.Reloaders = append(inspeqtor.Reloaders, func(old *inspeqtor.Inspeqtor, newi *inspeqtor.Inspeqtor) error {
		recurring, err := jobs.Parse(newi.GlobalConfig, cfgDir)
		if err != nil {
			return err
		}

		jobs.Watch(newi, recurring)
		return nil
	})

	return nil
}
