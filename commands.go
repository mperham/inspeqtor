package inspeqtor

import (
	"bufio"
	"inspeqtor/util"
	"time"
)

type commandFunc func(*Inspeqtor, *bufio.Writer)

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

	response := bufio.NewWriter(c)
	defer response.Flush()

	funk(i, response)
}

func startDeploy(i *Inspeqtor, resp *bufio.Writer) {
	length := time.Duration(i.GlobalConfig.Top.DeployLength) * time.Second
	i.SilenceUntil = time.Now().Add(length)

	util.Info("Starting deploy")
	resp.WriteString("Starting deploy, now silenced\n")
}

func finishDeploy(i *Inspeqtor, resp *bufio.Writer) {
	i.SilenceUntil = time.Now()
	util.Info("Finished deploy")
	resp.WriteString("Finished deploy, volume turned to 11\n")
}

func currentInfo(i *Inspeqtor, resp *bufio.Writer) {

}

func heart(i *Inspeqtor, resp *bufio.Writer) {
	resp.WriteString("Awwww, I love you too.\n")
}
