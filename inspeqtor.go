package inspeqtor

import (
	"inspeqtor/conf/global"
	"inspeqtor/conf/inq"
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
	ServiceManagers []Init
	Checks          *inq.Checks
	GlobalConfig    *global.ConfigFile
}

type Init interface {
	// Name of the init system: "upstart", "runit", etc.
	Name() string

	// Look up PID for the given service name, returns
	// positive integer if successful, -1 if the service
	// name was not found or error if there was an
	// unexpected failure.
	FindServicePID(name string) (int32, error)
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
	checks, err := inq.ParseChecks(i.RootDir + "/conf.d")
	if err != nil {
		return err
	}
	i.Checks = checks

	config, err := global.Parse(i.RootDir)
	if err != nil {
		return err
	}
	i.GlobalConfig = config

	util.DebugDebug("Config: %+v", config)
	util.DebugDebug("Checks: %+v", checks)
	return nil
}

func (i *Inspeqtor) runLoop() {
	scanSystem(true)
	select {
	case <-time.After(time.Duration(i.GlobalConfig.Top.CycleTime) * time.Second):
		scanSystem(false)
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
