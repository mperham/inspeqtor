package inspeqtor

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"inspeqtor/metrics"
	"inspeqtor/services"
	"strings"
	"testing"
)

func makeAction(actionName, notifType string, config map[string]string) (Action, error) {
	return Actions[actionName](nil, &AlertRoute{"", notifType, config})
}

func mockService(name string) *Service {
	return &Service{name, &services.ProcessStatus{999, services.Up}, nil, nil, nil, services.MockInit()}
}

func TestRestartAlert(t *testing.T) {
	t.Parallel()
	s := mockService("foobar")
	res, err := Actions["restart"](s, nil)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestRestart(t *testing.T) {
	t.Parallel()
	s := mockService("foobar")
	res, err := Actions["restart"](s, nil)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	assert.Equal(t, 999, s.Process.Pid)
	assert.Equal(t, services.Up, s.Process.Status)
	res.Trigger(nil)
	assert.Equal(t, 0, s.Process.Pid)
	assert.Equal(t, services.Starting, s.Process.Status)
}

func TestGmailNotifier(t *testing.T) {
	t.Parallel()
	action, err := makeAction("alert", "gmail", map[string]string{
		"username": "mike",
		"password": "fuzzbucket",
		"email":    "mike@example.org",
	})
	assert.Nil(t, err)
	assert.NotNil(t, action)
}

func TestEmailNotifier(t *testing.T) {
	t.Parallel()
	action, err := makeAction("alert", "email", map[string]string{
		"username": "mike",
		"password": "fuzzbucket",
		"hostname": "smtp.example.com",
		"email":    "mike@example.org",
	})
	assert.Nil(t, err)
	assert.NotNil(t, action)
}

func TestInvalidNotifier(t *testing.T) {
	t.Parallel()
	action, err := makeAction("alert", "emaul", map[string]string{})
	assert.NotNil(t, err)
	assert.Nil(t, action)
}

func TestMissingEmailNotifier(t *testing.T) {
	t.Parallel()
	action, err := makeAction("alert", "email", map[string]string{
		"username": "mike",
		"password": "fuzzbucket",
		"email":    "mike@example.org",
	})
	assert.NotNil(t, err)
	assert.Nil(t, action)
}

func TestEmailTrigger(t *testing.T) {
	t.Parallel()
	svc := Service{"mysql", nil, nil, nil, metrics.NewProcessStore(), nil}
	alert := &Event{
		&svc, &Rule{&svc, "memory", "rss", GT, "64m", 64 * 1024 * 1024, 0, 1, 0, Ok, nil}, HealthFailure,
	}

	err := validEmailSetup().TriggerEmail(alert, func(e *EmailNotifier, doc bytes.Buffer) error {
		content := string(doc.Bytes())
		assert.True(t, strings.Index(content, "[mysql]") > 0, "email does not contain expected content")
		assert.True(t, strings.Index(content, "memory(rss)") > 0, "email does not contain expected content")
		return nil
	})
	assert.Nil(t, err)
}

func validEmailSetup() *EmailNotifier {
	return &EmailNotifier{
		"mike", "fuzzbucket", "smtp.gmail.com", "mike@example.org", "mperham@gmail.com"}
}
