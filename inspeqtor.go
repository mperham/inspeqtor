package inspeqtor

import (
	"inspeqtor/conf/global"
	"inspeqtor/conf/inq"
	"inspeqtor/core"
	"inspeqtor/metrics"
	"inspeqtor/util"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	VERSION = "1.0.0"
)

type Inspeqtor struct {
	RootDir         string
	ServiceManagers []*core.InitSystem
	Host            *core.Host
	Services        []*core.Service
	GlobalConfig    *global.ConfigFile
}

func New(dir string) (*Inspeqtor, error) {
	return &Inspeqtor{RootDir: dir}, nil
}

var (
	Quit os.Signal = syscall.SIGQUIT
)

func (i *Inspeqtor) Start() {
	go i.runLoop()

	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt)
	signal.Notify(signals, Quit)
	<-signals

	util.Debug("Inspeqtor shutting down...")
	os.Exit(0)
}

func (i *Inspeqtor) Parse() error {
	host, services, err := inq.ParseChecks(i.RootDir + "/conf.d")
	if err != nil {
		return err
	}
	i.Host = host
	i.Services = services

	config, err := global.Parse(i.RootDir)
	if err != nil {
		return err
	}
	i.GlobalConfig = config

	util.DebugDebug("Config: %+v", config)
	util.DebugDebug("Host: %+v", host)
	util.DebugDebug("Services: %+v", services)
	return nil
}

func (i *Inspeqtor) runLoop() {
	scanSystem(true)
	for {
		select {
		case <-time.After(time.Duration(i.GlobalConfig.Top.CycleTime) * time.Second):
			scanSystem(false)
		}
	}
}

func scanSystem(firstTime bool) {
	util.DebugDebug("Scanning...")
	metrics, err := metrics.CollectHostMetrics("/proc")
	if err != nil {
		util.Warn("%v", err)
	} else {
		util.DebugDebug("%+v", metrics)
	}
}
