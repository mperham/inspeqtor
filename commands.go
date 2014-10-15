package inspeqtor

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"os"
	"strings"
	"time"

	"github.com/mperham/inspeqtor/util"
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

func (i *Inspeqtor) acceptCommand() bool {
	c, err := i.Socket.Accept()
	if err != nil {
		select {
		case <-i.Stopping:
			// we're stopping or reloading, no big deal...
		default:
			util.Warn("%v", err)
		}
		return false
	}
	defer c.Close()
	c.SetDeadline(time.Now().Add(2 * time.Second))

	reader := bufio.NewReader(c)
	line, err := reader.ReadString('\n')
	if err != nil {
		util.Info("Did not receive command line in time: %s", err.Error())
		return true
	}

	fields := strings.Fields(line)
	funk := CommandHandlers[fields[0]]
	if funk == nil {
		util.Warn("Unknown command: %s", strings.TrimSpace(line))
		io.WriteString(c, "Unknown command: "+line)
		return true
	}

	funk(i, fields[1:], c)
	return true
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

func currentStatus(i *Inspeqtor, args []string, resp io.Writer) {
	io.WriteString(resp, fmt.Sprintf(
		"%s %s, uptime: %s, pid: %d\n", Name, VERSION, time.Now().Sub(i.StartedAt).String(), os.Getpid()))
	io.WriteString(resp, "\n")

	io.WriteString(resp, fmt.Sprintf("Host: %s\n", i.Host.Name()))
	store := i.Host.Metrics()
	for _, fam := range store.Families() {
		for _, met := range store.MetricNames(fam) {
			name := fam
			if met != "" {
				name = name + ":" + met
			}
			var rule *Rule
			for _, r := range i.Host.Rules() {
				if r.Metric() == name {
					rule = r
				}
			}
			if rule != nil {
				io.WriteString(resp, fmt.Sprintf("%-1s %-35s %-15s %s\n", rule.DisplayState(), name, store.Display(fam, met), rule.DisplayThreshold))
			} else {
				io.WriteString(resp, fmt.Sprintf("  %-35s %-15s\n", name, store.Display(fam, met)))
			}
		}
	}

	for _, svc := range i.Services {
		io.WriteString(resp, "\n")
		io.WriteString(resp, fmt.Sprintf("Service: %s\n", svc))

		store := svc.Metrics()
		for _, fam := range store.Families() {
			for _, met := range store.MetricNames(fam) {
				name := fam
				if met != "" {
					name = name + ":" + met
				}
				var rule *Rule
				for _, r := range svc.Rules() {
					if r.Metric() == name {
						rule = r
					}
				}
				if rule != nil {
					io.WriteString(resp, fmt.Sprintf("%-1s %-35s %-15s %s\n", rule.DisplayState(), name, store.Display(fam, met), rule.DisplayThreshold))
				} else {
					io.WriteString(resp, fmt.Sprintf("  %-35s %-15s\n", name, store.Display(fam, met)))
				}
			}
		}
	}
}

func heart(i *Inspeqtor, args []string, resp io.Writer) {
	io.WriteString(resp, "Awwww, I love you too.\n")
}

// Beautiful, love this Go feature where you
// can slice out only the methods you need for simplicity
// of testing...
type displayable interface {
	At(int) *float64
	Displayable(float64) string
	Size() int
}

func sparkline(i *Inspeqtor, args []string, resp io.Writer) {
	if len(args) < 2 {
		io.WriteString(resp, "show [target] [metric]\n")
		return
	}

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
	}

	if target == nil {
		io.WriteString(resp, fmt.Sprintf("Invalid target: %s\n", targetName))
		return
	}

	output := buildSparkline(target, args[1], func(family, name string) displayable {
		return target.Metrics().Metric(family, name)
	})
	io.WriteString(resp, output)
}

func buildSparkline(target Checkable, metric string, buf func(string, string) displayable) string {
	family, name := parseMetric(metric)

	buff := buf(family, name)
	if buff == nil {
		return fmt.Sprintf("Unknown metric: %s\n", metric)
	}

	sz := buff.Size()
	values := make([]float64, sz)

	for i := 0; i > -sz; i-- {
		v := buff.At(i)
		if v == nil {
			util.Warn("BUG: Nil data in ring buffer: %d %d", sz, i)
			return "Inspeqtor bug, error building graph\n"
		}
		values[-i] = *v
	}

	// does not work for some reason, SO to the rescue!
	//sort.Reverse(sort.Float64Slice(values))
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}

	var min, max, sum, avg float64
	min = math.MaxFloat64

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

	resp.WriteString(fmt.Sprintf("%s %s min %s max %s avg %s\n",
		target.Name(),
		metric,
		buff.Displayable(min),
		buff.Displayable(max),
		buff.Displayable(avg)))

	runes := []string{"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█"}
	tick := (max - min) / 8

	for _, x := range values {
		diff := int((x - min) / tick)
		if diff > 7 {
			diff = 7
		}
		if diff < 0 {
			diff = 0
		}

		resp.WriteString(runes[diff])
	}

	resp.WriteString("\n")
	return string(resp.Bytes())
}

func parseMetric(metric string) (string, string) {
	if strings.Index(metric, ":") > 0 {
		fields := strings.Split(metric, ":")
		family := fields[0]
		name := ""
		if len(fields) > 1 {
			name = fields[1]
		}
		return family, name
	} else if strings.Index(metric, "(") > 0 {
		fields := strings.Split(metric, "(")
		family := fields[0]
		name := ""
		if len(fields) > 1 {
			name = fields[1]
			name = name[0 : len(name)-1]
		}
		return family, name
	} else {
		return metric, ""
	}
}
