/*
 * Integration tests which verify various events fire under various conditions
 */

package inspeqtor

import (
	"fmt"
	"os"
	"testing"

	"github.com/mperham/inspeqtor/metrics"
	"github.com/mperham/inspeqtor/services"
	"github.com/mperham/inspeqtor/util"
	"github.com/stretchr/testify/assert"
)

func findDownPid() int {
	if util.Darwin() {
		return 2
	}

	pid := 2
	for {
		_, err := os.Stat(fmt.Sprintf("/proc/%d", pid))
		if err != nil && os.IsNotExist(err) {
			return pid
		}
		pid++
		if pid == 1000 {
			panic("Unable to find available PID slot")
		}
	}
}

func TestEventProcessDisappears(t *testing.T) {
	t.Parallel()

	init := services.MockInit()
	init.CurrentStatus = services.WithStatus(0, services.Down)
	act := mockAction()

	assert.Equal(t, 0, act.Size())
	svc := &Service{&Entity{"foo", nil, metrics.NewProcessStore("/proc", 15), nil}, act, services.WithStatus(findDownPid(), services.Up), init}
	svc.Collect(false, func(_ Checkable) {})
	assert.Equal(t, services.Down, svc.Process.Status)
	assert.Equal(t, 0, svc.Process.Pid)
	assert.Equal(t, 1, act.Size())
	assert.Equal(t, ProcessDoesNotExist, act.Latest().Type)
}

func TestEventProcessDisappearsDuringDeploy(t *testing.T) {
	t.Parallel()

	var inits []services.InitSystem

	init := services.MockInit()
	act := mockAction()
	inits = append(inits, init)

	init.CurrentStatus = services.WithStatus(0, services.Down)

	assert.Equal(t, 0, act.Size())
	svc := &Service{&Entity{"foo", nil, metrics.NewProcessStore("/proc", 15), nil}, act, services.WithStatus(findDownPid(), services.Up), init}
	assert.Equal(t, services.Down, svc.Process.Status)
	assert.Equal(t, 15, svc.Process.Pid)
	assert.Equal(t, 0, act.Size())
	assert.Nil(t, act.Latest())
}

func TestEventProcessAppears(t *testing.T) {
	t.Parallel()

	init := services.MockInit()
	init.CurrentStatus = services.WithStatus(os.Getpid(), services.Up)
	act := mockAction()

	assert.Equal(t, 0, act.Size())
	svc := &Service{&Entity{"foo", nil, metrics.NewProcessStore("/proc", 15), nil}, act, services.WithStatus(0, services.Down), init}
	svc.Collect(false, func(_ Checkable) {})
	assert.Equal(t, services.Up, svc.Process.Status)
	assert.Equal(t, os.Getpid(), svc.Process.Pid)
	assert.Equal(t, 1, act.Size())
	assert.Equal(t, ProcessExists, act.Latest().Type)
}

func TestEventProcessAppearsDuringDeploy(t *testing.T) {
	t.Parallel()

	init := services.MockInit()
	init.CurrentStatus = services.WithStatus(os.Getpid(), services.Up)
	act := mockAction()

	assert.Equal(t, 0, act.Size())
	svc := &Service{&Entity{"foo", nil, metrics.NewProcessStore("/proc", 15), nil}, act, services.WithStatus(0, services.Down), init}
	svc.Collect(true, func(_ Checkable) {})
	assert.Equal(t, services.Up, svc.Process.Status)
	assert.Equal(t, os.Getpid(), svc.Process.Pid)
	assert.Equal(t, 0, act.Size())
	assert.Nil(t, act.Latest())
}

func TestEventProcessDoesNotChangeDuringDeploy(t *testing.T) {
	t.Parallel()

	init := services.MockInit()
	init.CurrentStatus = services.WithStatus(0, services.Up)

	act := mockAction()
	assert.Equal(t, 0, act.Size())

	silent := true // Begin a deployment
	svc := &Service{&Entity{"dne", nil, metrics.NewProcessStore("/proc", 15), nil}, act, services.WithStatus(0, services.Down), nil}
	svc.Resolve(silent, []services.InitSystem{init})
	assert.Equal(t, services.Down, svc.Process.Status)

	silent = false // Complete the deployment
	svc = &Service{&Entity{"dne", nil, metrics.NewProcessStore("/proc", 15), nil}, act, services.WithStatus(0, services.Up), nil}
	svc.Resolve(silent, []services.InitSystem{init})
	assert.Equal(t, services.Up, svc.Process.Status)

	assert.Equal(t, 0, act.Size())
}

func TestEventProcessDoneAtStartup(t *testing.T) {
	t.Parallel()

	init := services.MockInit()
	init.CurrentStatus = services.WithStatus(0, services.Down)

	act := mockAction()

	assert.Equal(t, 0, act.Size())
	svc := &Service{&Entity{"dne", nil, metrics.NewProcessStore("/proc", 15), nil}, act, services.WithStatus(0, services.Unknown), nil}
	svc.Resolve(false, []services.InitSystem{init})
	assert.Equal(t, services.Down, svc.Process.Status)
	assert.Equal(t, 0, svc.Process.Pid)
	assert.Equal(t, 1, act.Size())
	assert.Equal(t, ProcessDoesNotExist, act.Latest().Type)
}

func TestEventProcessExistsAtStartup(t *testing.T) {
	t.Parallel()

	init := services.MockInit()
	init.CurrentStatus = services.WithStatus(100, services.Up)

	act := mockAction()

	assert.Equal(t, 0, act.Size())
	svc := &Service{&Entity{"exists", nil, metrics.NewProcessStore("/proc", 15), nil}, act, services.WithStatus(0, services.Unknown), init}
	svc.Resolve(false, []services.InitSystem{init})
	assert.Equal(t, services.Up, svc.Process.Status)
	assert.Equal(t, 100, svc.Process.Pid)
	assert.Equal(t, 0, act.Size())
}

func TestEventRuleFails(t *testing.T) {
	t.Parallel()

	act := mockAction()

	svc := &Service{&Entity{"me", nil, metrics.NewProcessStore("/proc", 15), nil}, act, services.WithStatus(os.Getpid(), services.Up), services.MockInit()}
	rule := &Rule{svc, "memory", "rss", LT, "100m", 100 * 1024 * 1024, 0, false, 2, 0, Ok, []Action{act}}
	svc.rules = []*Rule{rule}

	// first collection should trip but not trigger since rule requires 2 cycles
	svc.Collect(false, func(_ Checkable) {})
	events := svc.Verify()
	assert.Equal(t, 0, len(events))
	assert.Equal(t, 0, act.Size())

	svc.Collect(false, func(_ Checkable) {})
	events = svc.Verify()
	assert.Equal(t, 1, len(events))
	assert.Equal(t, 1, act.Size())
	assert.Equal(t, RuleFailed, act.Latest().Type)
}

func TestEventRuleRecovers(t *testing.T) {
	t.Parallel()

	act := mockAction()

	svc := &Service{&Entity{"me", nil, metrics.NewProcessStore("/proc", 15), nil}, act, services.WithStatus(os.Getpid(), services.Up), services.MockInit()}
	rule := &Rule{svc, "memory", "rss", LT, "100m", 100 * 1024 * 1024, 0, false, 1, 0, Ok, []Action{act}}
	svc.rules = []*Rule{rule}

	svc.Collect(false, func(_ Checkable) {})
	events := svc.Verify()
	assert.Equal(t, 1, len(events))
	assert.Equal(t, 1, act.Size())
	assert.Equal(t, RuleFailed, act.Latest().Type)

	// recovery takes 2 cycles so we don't flap unnecessarily
	rule.Threshold = 1
	svc.Collect(false, func(_ Checkable) {})
	events = svc.Verify()
	assert.Equal(t, 0, len(events))

	svc.Collect(false, func(_ Checkable) {})
	events = svc.Verify()
	assert.Equal(t, 1, len(events))
	assert.Equal(t, 2, act.Size())
	assert.Equal(t, RuleRecovered, act.Latest().Type)
}

func TestDontVerifyDownServices(t *testing.T) {
	svc := NewService("foo")
	assert.NotEqual(t, svc.Process.Status, services.Up)
	events := svc.Verify()
	assert.Equal(t, events, []*Event{})
}
