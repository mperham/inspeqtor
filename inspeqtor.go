package inspeqtor

import (
	"github.com/mperham/inspeqtor/metrics"
	"github.com/mperham/inspeqtor/metrics/daemon"
	"github.com/mperham/inspeqtor/services"
	"github.com/mperham/inspeqtor/util"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const (
	VERSION = "0.5.0"
)

type Inspeqtor struct {
	RootDir    string
	SocketPath string
	StartedAt  time.Time

	ServiceManagers []services.InitSystem
	Host            *Host
	Services        []Checkable
	GlobalConfig    *ConfigFile
	Socket          net.Listener
	SilenceUntil    time.Time
	Valid           bool
}

func New(dir string, socketpath string) (*Inspeqtor, error) {
	i := &Inspeqtor{RootDir: dir,
		SocketPath:   socketpath,
		StartedAt:    time.Now(),
		SilenceUntil: time.Now(),
		Host:         &Host{&Entity{name: "localhost", metrics: metrics.NewMockStore()}},
		GlobalConfig: &ConfigFile{Defaults, map[string]*AlertRoute{}},
		Valid:        true}
	return i, nil
}

var (
	Term os.Signal = syscall.SIGTERM
	Hup  os.Signal = syscall.SIGHUP

	SignalHandlers = map[os.Signal]func(*Inspeqtor){
		Term:         exit,
		os.Interrupt: exit,
		Hup:          reload,
	}
	Name      string     = "Inspeqtor"
	Licensing string     = "Licensed under the GNU Public License 3.0"
	singleton *Inspeqtor = nil
)

func (i *Inspeqtor) Start() {
	err := i.openSocket(i.SocketPath)
	if err != nil {
		util.Warn("Could not create Unix socket: %s", err.Error())
		exit(i)
	}

	go func() {
		for {
			i.acceptCommand()
		}
	}()

	go i.runLoop()

	singleton = i
}

func (i *Inspeqtor) Parse() error {
	i.ServiceManagers = services.Detect()

	config, err := ParseGlobal(i.RootDir)
	if err != nil {
		return err
	}
	util.DebugDebug("Global config: %+v", config)
	i.GlobalConfig = config

	host, services, err := ParseInq(i.GlobalConfig, i.RootDir+"/conf.d")
	if err != nil {
		return err
	}
	i.Host = host
	i.Services = services

	util.DebugDebug("Config: %+v", config)
	util.DebugDebug("Host: %+v", host)
	for _, val := range services {
		util.DebugDebug("Service: %+v", val)
	}

	for _, s := range i.Services {
		err := wrapService(s.(*Service))
		if err != nil {
			return err
		}
	}
	return nil
}

func HandleSignal(sig os.Signal, handler func(*Inspeqtor)) {
	SignalHandlers[sig] = handler
}

func HandleSignals() {
	signals := make(chan os.Signal)
	for k, _ := range SignalHandlers {
		signal.Notify(signals, k)
	}

	for {
		sig := <-signals
		util.Debug("Received signal %d", sig)
		funk := SignalHandlers[sig]
		funk(singleton)
	}
}

// private

func wrapService(s *Service) error {
	var store *daemon.Store

	for _, r := range s.Rules() {
		funk := daemon.Sources[r.MetricFamily]
		if funk != nil {
			if store == nil {
				source, err := funk(s.Parameters())
				if err != nil {
					return err
				}
				util.Info("Activating %s-specific metrics", r.MetricFamily)
				store = daemon.NewStore(s.Metrics(), source)
				s.SetMetrics(store)
			}
			util.Debug("Watching %s(%s)", r.MetricFamily, r.MetricName)
			store.Watch(r.MetricName)
		}
	}
	if store != nil {
		err := daemon.Prepare(store)
		if err != nil {
			return err
		}
	}
	return nil
}

func reload(i *Inspeqtor) {
	util.Info(Name + " reloading")
	newi, err := New(i.RootDir, i.SocketPath)
	if err != nil {
		util.Warn("Unable to reload: %s", err.Error())
		return
	}
	newi.SilenceUntil = i.SilenceUntil

	err = newi.Parse()
	if err != nil {
		util.Warn("Unable to reload: %s", err.Error())
		return
	}

	// TODO proper reloading would not throw away the existing metric data
	// in i but defining new metrics can change the storage tree.  Implement
	// deep metric tree ring buffer sync if possible?
	i.Shutdown()
	newi.Start()
}

func exit(i *Inspeqtor) {
	util.Info(Name + " exiting")

	i.Shutdown()
	os.Exit(0)
}

func (i *Inspeqtor) Shutdown() {
	i.Valid = false
	if i.Socket != nil {
		err := i.Socket.Close()
		if err != nil {
			util.Warn(err.Error())
		}
	}
}

// this method never returns.
//
// since we can't test this method in an automated fashion, it should
// contain as little logic as possible.
func (i *Inspeqtor) runLoop() {
	util.DebugDebug("Resolving services")
	for _, svc := range i.Services {
		svc.Resolve(i.ServiceManagers)
	}

	i.scanSystem()

	for {
		select {
		case <-time.After(time.Duration(i.GlobalConfig.Top.CycleTime) * time.Second):
			if !i.Valid {
				return
			}
			i.scanSystem()
		}
	}
}

func (i *Inspeqtor) silenced() bool {
	return time.Now().Before(i.SilenceUntil)
}

func (i *Inspeqtor) scanSystem() {
	// "Trust, but verify"
	// https://en.wikipedia.org/wiki/Trust%2C_but_verify
	i.scan()
	i.verify()
}

func (i *Inspeqtor) scan() {
	start := time.Now()
	var barrier sync.WaitGroup
	barrier.Add(1)
	barrier.Add(len(i.Services))

	go i.Host.Collect(i.silenced(), func(_ Checkable) {
		barrier.Done()
	})
	for _, svc := range i.Services {
		go svc.Collect(i.silenced(), func(_ Checkable) {
			barrier.Done()
		})
	}
	barrier.Wait()
	util.Debug("Collection complete in " + time.Now().Sub(start).String())
}

func (i *Inspeqtor) verify() {
	if i.silenced() {
		// We are silenced until some point in the future.
		// We don't want to check rules (as a deploy might use
		// enough resources to trip a threshold) or trigger
		// any actions
		for _, rule := range i.Host.Rules() {
			rule.Reset()
		}
		for _, svc := range i.Services {
			for _, rule := range svc.Rules() {
				rule.Reset()
			}
		}
	} else {
		i.Host.Verify()
		for _, svc := range i.Services {
			svc.Verify()
		}
	}
}

func (i *Inspeqtor) TestAlertRoutes() int {
	bad := 0
	util.Info("Testing alert routes")
	for _, route := range i.GlobalConfig.AlertRoutes {
		nm := route.Name
		if nm == "" {
			nm = "default"
		}
		util.Debug("Creating notification for %s/%s", route.Channel, nm)
		notifier, err := Actions["alert"](i.Host, route)
		if err != nil {
			bad += 1
			util.Warn("Error creating %s/%s route: %s", route.Channel, nm, err.Error())
			continue
		}
		util.Debug("Triggering notification for %s/%s", route.Channel, nm)
		err = notifier.Trigger(&Event{RuleFailed, i.Host, i.Host.Rules()[0]})
		if err != nil {
			bad += 1
			util.Warn("Error firing %s/%s route: %s", route.Channel, nm, err.Error())
		}
	}
	return bad
}
