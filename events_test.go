/*
 * Integration tests which verify various events fire under various conditions
 */

package inspeqtor

import (
	"github.com/stretchr/testify/assert"
	"inspeqtor/metrics"
	"inspeqtor/services"
	"os"
	"testing"
)

func TestEventProcessDisappears(t *testing.T) {
	t.Parallel()

	i, err := New("", "")
	assert.Nil(t, err)

	init := services.MockInit()
	init.CurrentStatus = &services.ProcessStatus{0, services.Down}
	act := mockAction()

	assert.Equal(t, 0, act.Size())
	svc := &Service{&Entity{"foo", nil, nil, nil}, act, &services.ProcessStatus{99, services.Up}, init}
	i.collectService(svc, func(_ *Service) {})
	assert.Equal(t, services.Down, svc.Process.Status)
	assert.Equal(t, 0, svc.Process.Pid)
	assert.Equal(t, 1, act.Size())
	assert.Equal(t, ProcessDoesNotExist, act.Latest().Type)
}

func TestEventProcessAppears(t *testing.T) {
	t.Parallel()

	i, err := New("", "")
	assert.Nil(t, err)

	init := services.MockInit()
	init.CurrentStatus = &services.ProcessStatus{os.Getpid(), services.Up}
	act := mockAction()

	assert.Equal(t, 0, act.Size())
	svc := &Service{&Entity{"foo", nil, metrics.NewProcessStore(), nil}, act, &services.ProcessStatus{0, services.Down}, init}
	i.collectService(svc, func(_ *Service) {})
	assert.Equal(t, 1, act.Size())
	assert.Equal(t, services.Up, svc.Process.Status)
	assert.Equal(t, os.Getpid(), svc.Process.Pid)
	assert.Equal(t, ProcessExists, act.Latest().Type)
}

func TestEventProcessDneAtStartup(t *testing.T) {
	t.Parallel()

	i, err := New("", "")
	assert.Nil(t, err)

	init := services.MockInit()
	init.CurrentStatus = &services.ProcessStatus{0, services.Down}
	i.ServiceManagers = append(i.ServiceManagers, init)

	act := mockAction()

	assert.Equal(t, 0, act.Size())
	svc := &Service{&Entity{"dne", nil, metrics.NewProcessStore(), nil}, act, &services.ProcessStatus{0, services.Unknown}, nil}
	i.resolveService(svc)
	assert.Equal(t, services.Down, svc.Process.Status)
	assert.Equal(t, 0, svc.Process.Pid)
	assert.Equal(t, 1, act.Size())
	assert.Equal(t, ProcessDoesNotExist, act.Latest().Type)
}

func TestEventProcessExistsAtStartup(t *testing.T) {
	t.Parallel()

	i, err := New("", "")
	assert.Nil(t, err)

	init := services.MockInit()
	init.CurrentStatus = &services.ProcessStatus{100, services.Up}
	i.ServiceManagers = append(i.ServiceManagers, init)

	act := mockAction()

	assert.Equal(t, 0, act.Size())
	svc := &Service{&Entity{"exists", nil, metrics.NewProcessStore(), nil}, act, &services.ProcessStatus{0, services.Unknown}, init}
	i.resolveService(svc)
	assert.Equal(t, services.Up, svc.Process.Status)
	assert.Equal(t, 100, svc.Process.Pid)
	assert.Equal(t, 0, act.Size())
}

func TestEventRuleFails(t *testing.T) {
	t.Parallel()

	i, err := New("", "")
	assert.Nil(t, err)
	act := mockAction()

	svc := &Service{&Entity{"me", nil, metrics.NewProcessStore(), nil}, act, &services.ProcessStatus{os.Getpid(), services.Up}, services.MockInit()}
	rule := &Rule{svc, "memory", "rss", LT, "100m", 100 * 1024 * 1024, 0, 2, 0, Ok, []Action{act}}
	svc.rules = []*Rule{rule}

	// first collection should trip but not trigger since rule requires 2 cycles
	i.collectService(svc, func(_ *Service) {})
	events := i.verify(nil, []*Service{svc})
	assert.Equal(t, 0, len(events))
	assert.Equal(t, 0, act.Size())

	i.collectService(svc, func(_ *Service) {})
	events = i.verify(nil, []*Service{svc})
	assert.Equal(t, 1, len(events))
	assert.Equal(t, 1, act.Size())
	assert.Equal(t, RuleFailed, act.Latest().Type)
}

func TestEventRuleRecovers(t *testing.T) {
	t.Parallel()

	i, err := New("", "")
	assert.Nil(t, err)
	act := mockAction()

	svc := &Service{&Entity{"me", nil, metrics.NewProcessStore(), nil}, act, &services.ProcessStatus{os.Getpid(), services.Up}, services.MockInit()}
	rule := &Rule{svc, "memory", "rss", LT, "100m", 100 * 1024 * 1024, 0, 1, 0, Ok, []Action{act}}
	svc.rules = []*Rule{rule}

	i.collectService(svc, func(_ *Service) {})
	events := i.verify(nil, []*Service{svc})
	assert.Equal(t, 1, len(events))
	assert.Equal(t, 1, act.Size())
	assert.Equal(t, RuleFailed, act.Latest().Type)

	// recovery takes 2 cycles so we don't flap unnecessarily
	rule.Threshold = 1
	i.collectService(svc, func(_ *Service) {})
	events = i.verify(nil, []*Service{svc})
	assert.Equal(t, 0, len(events))

	i.collectService(svc, func(_ *Service) {})
	events = i.verify(nil, []*Service{svc})
	assert.Equal(t, 1, len(events))
	assert.Equal(t, 2, act.Size())
	assert.Equal(t, RuleRecovered, act.Latest().Type)
}

type TestAction struct {
	events []Event
}

func (t *TestAction) Latest() Event {
	return t.events[len(t.events)-1]
}

func (t *TestAction) Size() int {
	return len(t.events)
}

func (t *TestAction) Trigger(e *Event) error {
	t.events = append(t.events, *e)
	return nil
}

func mockAction() *TestAction {
	return &TestAction{}
}
