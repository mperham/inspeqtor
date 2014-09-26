package inspeqtor

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/mperham/inspeqtor/util"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

/*
 * Commands are ways for the external world to communicate with Inspeqtor
 * via its command socket.
 */

type commandFunc func(*Inspeqtor, []string, io.Writer)

var (
	CommandHandlers = map[string]commandFunc{
		"start":  startDeploy,
		"finish": finishDeploy,
		"status": currentStatus,
		"show":   sparkline,
		"♡":      heart,
	}
)

func (i *Inspeqtor) openSocket(path string) error {
	if i.Socket != nil {
		return errors.New("Socket is already open!")
	}

	socket, err := net.Listen("unix", path)
	if err != nil {
		return err
	}
	i.Socket = socket
	return nil
}

func (i *Inspeqtor) acceptCommand() {
	c, err := i.Socket.Accept()
	if err != nil {
		if i.Valid {
			util.Warn("Unix socket shutdown: %s", err.Error())
		}
		return
	}
	defer c.Close()
	c.SetDeadline(time.Now().Add(2 * time.Second))

	reader := bufio.NewReader(c)
	line, err := reader.ReadString('\n')
	if err != nil {
		util.Info("Did not receive a command line in time: %s", err.Error())
	}

	fields := strings.Fields(line)
	funk := CommandHandlers[fields[0]]
	if funk == nil {
		util.Warn("Unknown command: %s", strings.TrimSpace(line))
		io.WriteString(c, "Unknown command: "+line)
		return
	}

	funk(i, fields[1:], c)
}

func startDeploy(i *Inspeqtor, args []string, resp io.Writer) {
	length := time.Duration(i.GlobalConfig.Top.DeployLength) * time.Second
	i.SilenceUntil = time.Now().Add(length)

	util.Info("Starting deploy")
	io.WriteString(resp, "Starting deploy, now silenced\n")
}

func finishDeploy(i *Inspeqtor, args []string, resp io.Writer) {
	i.SilenceUntil = time.Now()
	util.Info("Finished deploy")
	io.WriteString(resp, "Finished deploy, volume turned to 11\n")
}

func sparkline(i *Inspeqtor, args []string, resp io.Writer) {
	targetName := args[0]
	var target Checkable

	if targetName == "host" {
		target = i.Host
	} else {
		for _, s := range i.Services {
			if s.Name() == targetName {
				target = s
			}
		}
		if target == nil {
			io.WriteString(resp, "Invalid target: "+targetName)
			return
		}
	}

	output := buildSparkline(target, args[1], func(family, name string) *util.RingBuffer {
		return target.Metrics().Buffer(family, name)
	})
	io.WriteString(resp, output)
}

func currentStatus(i *Inspeqtor, args []string, resp io.Writer) {
	io.WriteString(resp, fmt.Sprintf(
		"%s %s, uptime: %s, pid: %d\n", Name, VERSION, time.Now().Sub(i.StartedAt).String(), os.Getpid()))
	io.WriteString(resp, "\n")

	io.WriteString(resp, fmt.Sprintf("Host: %s\n", i.Host.Name()))
	store := i.Host.Metrics()
	for _, fam := range store.Families() {
		for _, met := range store.Metrics(fam) {
			name := fam
			if met != "" {
				name = name + "(" + met + ")"
			}
			var rule *Rule
			for _, r := range i.Host.Rules() {
				if r.Metric() == name {
					rule = r
				}
			}
			if rule != nil {
				io.WriteString(resp, fmt.Sprintf("  %-1s %-30s %-15s %s\n", rule.DisplayState(), name, store.Display(fam, met), rule.DisplayThreshold))
			} else {
				io.WriteString(resp, fmt.Sprintf("    %-30s %-15s\n", name, store.Display(fam, met)))
			}
		}
	}

	for _, svc := range i.Services {
		io.WriteString(resp, "\n")
		io.WriteString(resp, fmt.Sprintf("Service: %s\n", svc))

		store := svc.Metrics()
		for _, fam := range store.Families() {
			for _, met := range store.Metrics(fam) {
				name := fam
				if met != "" {
					name = name + "(" + met + ")"
				}
				var rule *Rule
				for _, r := range svc.Rules() {
					if r.Metric() == name {
						rule = r
					}
				}
				if rule != nil {
					io.WriteString(resp, fmt.Sprintf("  %-1s %-30s %-15s %s\n", rule.DisplayState(), name, store.Display(fam, met), rule.DisplayThreshold))
				} else {
					io.WriteString(resp, fmt.Sprintf("    %-30s %-15s\n", name, store.Display(fam, met)))
				}
			}
		}
	}
}

func heart(i *Inspeqtor, args []string, resp io.Writer) {
	io.WriteString(resp, "Awwww, I love you too.\n")
}

func buildSparkline(target Checkable, metric string, buf func(string, string) *util.RingBuffer) string {
	fields := strings.Split(metric, "(")
	family := fields[0]
	name := ""
	if len(fields) > 1 {
		name = fields[1]
		name = name[0 : len(name)-1]
	}

	buff := buf(family, name)
	values := buff.Export()

	var min, max, sum, avg float64
	for _, val := range values {
		if min > val {
			min = val
		}
		if max < val {
			max = val
		}
		sum += val
	}
	if len(values) > 0 {
		avg = sum / float64(len(values))
	}

	var resp bytes.Buffer

	resp.WriteString(fmt.Sprintf("%s %s min: %.2f max: %.2f avg: %.2f\n", target.Name(), metric, min, max, avg))
	runes := []string{"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█"}
	tick := (max - min) / 8

	for _, x := range values {
		diff := int((x - min) / tick)
		if diff > 7 {
			diff = 7
		}

		resp.WriteString(runes[diff])
	}

	resp.WriteString("\n")
	return string(resp.Bytes())
}
