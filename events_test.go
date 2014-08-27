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
}

func TestEventProcessAppears(t *testing.T) {
	t.Parallel()

	i, err := New("", "")
	assert.Nil(t, err)

	init := services.MockInit()
	init.CurrentStatus = &services.ProcessStatus{os.Getpid(), services.Up}
	act := mockAction()

	assert.Equal(t, 0, act.Size())
	svc := &Service{&Entity{"foo", nil, nil, metrics.NewProcessStore()}, act, &services.ProcessStatus{0, services.Down}, init}
	i.collectService(svc, func(_ *Service) {})
	assert.Equal(t, 1, act.Size())
	assert.Equal(t, services.Up, svc.Process.Status)
	assert.Equal(t, os.Getpid(), svc.Process.Pid)
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
	svc := &Service{&Entity{"dne", nil, nil, metrics.NewProcessStore()}, act, &services.ProcessStatus{0, services.Unknown}, nil}
	i.resolveService(svc)
	assert.Equal(t, 1, act.Size())
	assert.Equal(t, services.Down, svc.Process.Status)
	assert.Equal(t, 0, svc.Process.Pid)
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
	svc := &Service{&Entity{"exists", nil, nil, metrics.NewProcessStore()}, act, &services.ProcessStatus{0, services.Unknown}, init}
	i.resolveService(svc)
	assert.Equal(t, 0, act.Size())
	assert.Equal(t, services.Up, svc.Process.Status)
	assert.Equal(t, 100, svc.Process.Pid)
}

type TestAction struct {
	events []Event
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
