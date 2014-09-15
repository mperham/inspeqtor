package inspeqtor

import (
	"bufio"
	"errors"
	"fmt"
	"inspeqtor/util"
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

type commandFunc func(*Inspeqtor, io.Writer)

var (
	CommandHandlers = map[rune]commandFunc{
		's': startDeploy,
		'f': finishDeploy,
		'i': currentInfo,
		'♡': heart,
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

	firstChar := []rune(line)[0]
	funk := CommandHandlers[firstChar]
	if funk == nil {
		util.Warn("Unknown command: %s", strings.TrimSpace(line))
		io.WriteString(c, "Unknown command: "+line)
		return
	}

	funk(i, c)
}

func startDeploy(i *Inspeqtor, resp io.Writer) {
	length := time.Duration(i.GlobalConfig.Top.DeployLength) * time.Second
	i.SilenceUntil = time.Now().Add(length)

	util.Info("Starting deploy")
	io.WriteString(resp, "Starting deploy, now silenced\n")
}

func finishDeploy(i *Inspeqtor, resp io.Writer) {
	i.SilenceUntil = time.Now()
	util.Info("Finished deploy")
	io.WriteString(resp, "Finished deploy, volume turned to 11\n")
}

func currentInfo(i *Inspeqtor, resp io.Writer) {
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

func heart(i *Inspeqtor, resp io.Writer) {
	io.WriteString(resp, "Awwww, I love you too.\n")
}
