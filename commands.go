package inspeqtor

import (
	"bufio"
	"inspeqtor/util"
	"io"
	"time"
)

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
		util.Info("Unix socket shutdown: %s", err.Error())
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

}

func heart(i *Inspeqtor, resp io.Writer) {
	io.WriteString(resp, "Awwww, I love you too.\n")
}
