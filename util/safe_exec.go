package util

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"time"
)

var (
	// default command timeout
	CmdTimeout = 3 * time.Second
	// restarting services can take a few seconds so
	// allow a little more leeway here
	RestartTimeout = 10 * time.Second
)

/*
 Here's a lot of complexity to ensure that any child process
 we execute has a limited time to return.  Otherwise goroutines
 can leak and the main run loop stall due to shell commands
 which lock up for some reason.

 My kingdom for default parameter values!  timeout is a hack
 to make it optional.
*/
func SafeRun(cmd *exec.Cmd, timeout ...time.Duration) ([]byte, error) {
	if cmd.Stdout != nil {
		return nil, errors.New("exec: Stdout already set")
	}
	if cmd.Stderr != nil {
		return nil, errors.New("exec: Stderr already set")
	}
	Debug("Executing %v %v", cmd.Path, cmd.Args)
	var b bytes.Buffer
	cmd.Stdout = &b
	cmd.Stderr = &b
	err := cmd.Start()
	if err != nil {
		return b.Bytes(), err
	}

	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	timelimit := CmdTimeout
	if len(timeout) == 1 {
		timelimit = timeout[0]
	}

	select {
	case <-time.After(timelimit):
		if err := cmd.Process.Kill(); err != nil {
			Warn("failed to kill command: %s", err)
		}
		<-done // allow goroutine to exit
		return b.Bytes(), fmt.Errorf("Command timed out: %s", cmd.Args)
	case err := <-done:
		if err != nil {
			return b.Bytes(), err
		}
	}

	return b.Bytes(), nil
}
