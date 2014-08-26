package inspeqtor

import (
	"bufio"
	"fmt"
	"inspeqtor/util"
	"io"
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

func (i *Inspeqtor) acceptCommand() {
	c, err := i.Socket.Accept()
	if err != nil {
		util.Warn("Unix socket shutdown: %s", err.Error())
		return
	}
	defer c.Close()

	reader := bufio.NewReader(c)
	line, err := reader.ReadString('\n')
	if len(line) == 0 && err != nil {
		util.Info(err.Error())
	}

	firstChar := []rune(line)[0]
	funk := CommandHandlers[firstChar]
	if funk == nil {
		util.Warn("Unknown command: " + line)
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
		"%s %s, uptime: %s\n", Name, VERSION, time.Now().Sub(i.StartedAt).String()))
	io.WriteString(resp, "\n")

	io.WriteString(resp, fmt.Sprintf("Host: %s\n", i.Host.Hostname))
	for _, rule := range i.Host.Rules {
		io.WriteString(resp, fmt.Sprintf("  %-1s %-20s %-15s %s\n", rule.DisplayState(), rule.MetricName(), rule.DisplayCurrentValue(), rule.DisplayThreshold()))
	}

	for _, svc := range i.Services {
		io.WriteString(resp, "\n")
		io.WriteString(resp, fmt.Sprintf("Service: %s [%s]\n", svc.Name(), svc.Process))

		for _, rule := range svc.Rules {
			io.WriteString(resp, fmt.Sprintf("  %-1s %-20s %-15s %s\n", rule.DisplayState(), rule.MetricName(), rule.DisplayCurrentValue(), rule.DisplayThreshold()))
		}
	}
}

func heart(i *Inspeqtor, resp io.Writer) {
	io.WriteString(resp, "Awwww, I love you too.\n")
}
