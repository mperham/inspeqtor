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

	svc := &Service{&Entity{"foo", nil, nil, nil}, &services.ProcessStatus{99, services.Up}, init}
	i.collectService(svc, func(_ *Service) {})
	assert.Equal(t, services.Down, svc.Process.Status)
	assert.Equal(t, 0, svc.Process.Pid)
}

func TestEventProcessAppears(t *testing.T) {
	t.Parallel()

	i, err := New("", "")
	assert.Nil(t, err)

	init := services.MockInit()
	init.CurrentStatus = &services.ProcessStatus{os.Getpid(), services.Up}

	svc := &Service{&Entity{"foo", nil, nil, metrics.NewProcessStore()}, &services.ProcessStatus{0, services.Down}, init}
	i.collectService(svc, func(_ *Service) {})
	assert.Equal(t, services.Up, svc.Process.Status)
	assert.Equal(t, os.Getpid(), svc.Process.Pid)
}
