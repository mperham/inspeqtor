package inspeqtor

import (
	"expvar"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/mperham/inspeqtor/metrics"
	// Pull in daemon-specific sources
	_ "github.com/mperham/inspeqtor/metrics/daemon"
	"github.com/mperham/inspeqtor/services"
	"github.com/mperham/inspeqtor/util"
)

const (
	VERSION = "1.0.0"
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
	Expose          net.Listener
	SilenceUntil    time.Time
	Stopping        chan struct{}
	Handlers        map[string][]func(*Inspeqtor) error
}

func New(dir string, socketpath string) (*Inspeqtor, error) {
	i := &Inspeqtor{RootDir: dir,
		SocketPath:   socketpath,
		StartedAt:    time.Now(),
		SilenceUntil: time.Now(),
		Host:         &Host{&Entity{name: "localhost", metrics: metrics.NewMockStore()}},
		GlobalConfig: &ConfigFile{Defaults, map[string]*AlertRoute{}},
		Stopping:     make(chan struct{}),
		Handlers:     map[string][]func(*Inspeqtor) error{},
	}
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
	Name      = "Inspeqtor"
	Licensing = "Licensed under the GNU Public License 3.0"
	Singleton *Inspeqtor
	Reloaders = []func(*Inspeqtor, *Inspeqtor) error{basicReloader}

	counters = expvar.NewMap("inspeqtor")
)

func init() {
	counters.Add("deploy", 0)
	counters.Add("events", 0)
}

func (i *Inspeqtor) Start() {
	util.Debug("Starting command socket")
	err := i.openSocket(i.SocketPath)
	if err != nil {
		util.Warn("Could not create Unix socket: %s", err.Error())
		exit(i)
	}

	go func() {
		for {
			if !i.safelyAccept() {
				util.Debug("Shutting down command socket")
				return
			}
		}
	}()

	// if expose_port is 0, disable the feature altogether
	if i.GlobalConfig.ExposePort != 0 {
		sock, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", i.GlobalConfig.ExposePort))
		if err != nil {
			util.Warn("Could not listen on port %d: %s", i.GlobalConfig.ExposePort, err.Error())
			exit(i)
		}
		i.Expose = sock
		go func() {
			// TODO How do we error handling here?
			util.Info("Expose now available at port %d", i.GlobalConfig.ExposePort)
			err := http.Serve(i.Expose, nil)
			// Don't log an "error" when we shut down normally and close the socket
			if err != nil && !strings.Contains(err.Error(), "use of closed network") {
				util.Warn("HTTP server error: %s", err.Error())
			}
		}()
	}

	util.Debug("Starting main run loop")
	go i.runLoop()

	Singleton = i
}

func (i *Inspeqtor) safelyAccept() bool {
	defer func() {
		if err := recover(); err != nil {
			// TODO Is there a way to print out the backtrace of the goroutine where it crashed?
			util.Warn("Command crashed:\n%s", err)
		}
	}()

	return i.acceptCommand()
}

func (i *Inspeqtor) Parse() error {
	i.ServiceManagers = services.Detect()

	config, err := ParseGlobal(i.RootDir)
	if err != nil {
		return err
	}
	util.DebugDebug("Global config: %+v", config)
	i.GlobalConfig = config

	host, err := ParseHost(i.GlobalConfig, i.RootDir+"/host.inq")
	if err != nil {
		return err
	}
	i.Host = host

	services, err := ParseServices(i.GlobalConfig, i.RootDir+"/services.d")
	if err != nil {
		return err
	}
	i.Services = services

	util.DebugDebug("Config: %+v", config)
	util.DebugDebug("Host: %+v", host)
	for _, val := range services {
		util.DebugDebug("Service: %+v", val)
	}

	return nil
}

func HandleSignal(sig os.Signal, handler func(*Inspeqtor)) {
	SignalHandlers[sig] = handler
}

func HandleSignals() {
	signals := make(chan os.Signal)
	for k := range SignalHandlers {
		signal.Notify(signals, k)
	}

	for {
		sig := <-signals
		util.Debug("Received signal %d", sig)
		funk := SignalHandlers[sig]
		funk(Singleton)
	}
}

// private

func basicReloader(oldcopy *Inspeqtor, newcopy *Inspeqtor) error {
	newcopy.SilenceUntil = oldcopy.SilenceUntil
	return nil
}

func reload(i *Inspeqtor) {
	util.Info(Name + " reloading")
	newi, err := New(i.RootDir, i.SocketPath)
	if err != nil {
		util.Warn("Unable to reload: %s", err.Error())
		return
	}

	err = newi.Parse()
	if err != nil {
		util.Warn("Unable to reload: %s", err.Error())
		return
	}

	// we're reloading and newcopy will become the new
	// singleton.  Pro hooks into this to reload its features too.
	for _, callback := range Reloaders {
		err := callback(i, newi)
		if err != nil {
			util.Warn("Unable to reload: %s", err.Error())
			return
		}
	}

	// TODO proper reloading would not throw away the existing metric data
	// in i but defining new metrics can change the storage tree.  Implement
	// deep metric tree ring buffer sync if possible in basicReloader?
	i.Shutdown()
	newi.Start()
}

func exit(i *Inspeqtor) {
	util.Info(Name + " exiting")

	i.Shutdown()
	os.Exit(0)
}

func (i *Inspeqtor) Shutdown() {
	close(i.Stopping)

	if i.Socket != nil {
		err := i.Socket.Close()
		if err != nil {
			util.Warn(err.Error())
		}
	}
	if i.Expose != nil {
		err := i.Expose.Close()
		if err != nil {
			util.Warn(err.Error())
		}
	}
	i.Fire("shutdown")

	// let other goroutines log their exit
	time.Sleep(time.Millisecond)
}

// this method never returns.
//
// since we can't test this method in an automated fashion, it should
// contain as little logic as possible.
func (i *Inspeqtor) runLoop() {
	util.DebugDebug("Resolving services")
	for _, svc := range i.Services {
		err := svc.Resolve(i.ServiceManagers)
		if err != nil {
			util.Warn(err.Error())
		}
	}

	i.scanSystem()

	for {
		select {
		case <-time.After(time.Duration(i.GlobalConfig.CycleTime) * time.Second):
			i.scanSystem()
		case <-i.Stopping:
			util.Debug("Shutting down main run loop")
			return
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
	i.Fire("cycleComplete")
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
		e := i.Host.Verify()
		if len(e) > 0 {
			counters.Add("events", int64(len(e)))
		}
		for _, svc := range i.Services {
			e := svc.Verify()
			if len(e) > 0 {
				counters.Add("events", int64(len(e)))
			}
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
			bad++
			util.Warn("Error creating %s/%s route: %s", route.Channel, nm, err.Error())
			continue
		}
		util.Debug("Triggering notification for %s/%s", route.Channel, nm)
		err = notifier.Trigger(&Event{RuleFailed, i.Host, i.Host.Rules()[0]})
		if err != nil {
			bad++
			util.Warn("Error firing %s/%s route: %s", route.Channel, nm, err.Error())
		}
	}
	return bad
}

func (i *Inspeqtor) Listen(eventName string, handler func(*Inspeqtor) error) {
	x := i.Handlers[eventName]
	if x == nil {
		x = []func(*Inspeqtor) error{}
	}
	i.Handlers[eventName] = append(x, handler)
}

func (i *Inspeqtor) Fire(eventName string) {
	x := i.Handlers[eventName]
	if len(x) > 0 {
		for _, handler := range x {
			handler(i)
		}
	}
}
