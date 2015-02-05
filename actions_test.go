package inspeqtor

import (
	"bytes"
	"strings"
	"testing"

	"github.com/mperham/inspeqtor/metrics"
	"github.com/mperham/inspeqtor/services"
	"github.com/stretchr/testify/assert"
)

func makeAction(actionName, notifType string, config map[string]string) (Action, error) {
	return Actions[actionName](nil, &AlertRoute{"", notifType, config})
}

func mockService(name string) *Service {
	return &Service{&Entity{name, nil, nil, nil}, nil, services.WithStatus(999, services.Up), services.MockInit()}
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

func TestReload(t *testing.T) {
	t.Parallel()
	service := mockService("foobar")
	res, err := Actions["reload"](service, nil)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}

func TestGmailNotifier(t *testing.T) {
	t.Parallel()
	action, err := makeAction("alert", "gmail", map[string]string{
		"username": "mike",
		"password": "fuzzbucket",
		"to_email": "mike@example.org",
	})
	assert.Nil(t, err)
	assert.NotNil(t, action)
}

func TestEmailNotifier(t *testing.T) {
	t.Parallel()
	action, err := makeAction("alert", "email", map[string]string{
		"username":    "mike",
		"password":    "fuzzbucket",
		"smtp_server": "smtp.example.com",
		"to_email":    "mike@example.org",
	})
	assert.Nil(t, err)
	assert.NotNil(t, action)
}

func TestUnauthenticatedEmailNotifier(t *testing.T) {
	t.Parallel()
	action, err := makeAction("alert", "email", map[string]string{
		"smtp_server": "smtp.example.com",
		"to_email":    "mike@example.org",
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
		"to_email": "mike@example.org",
	})
	assert.NotNil(t, err)
	assert.Nil(t, action)
}

func TestEmailEventRuleFailed(t *testing.T) {
	t.Parallel()
	alert := validRuleEvent(RuleFailed)

	err := validEmailSetup().TriggerEmail(alert, func(e *EmailNotifier, doc bytes.Buffer) error {
		content := string(doc.Bytes())
		assert.True(t, strings.Index(content, "[mysql]") > 0, "email does not contain expected content: "+content)
		assert.True(t, strings.Index(content, "memory:rss") > 0, "email does not contain expected content: "+content)
		return nil
	})
	assert.Nil(t, err)
}

func TestEmailEventProcessExists(t *testing.T) {
	t.Parallel()
	alert := validProcessEvent(ProcessExists)

	err := validEmailSetup().TriggerEmail(alert, func(e *EmailNotifier, doc bytes.Buffer) error {
		content := string(doc.Bytes())
		assert.True(t, strings.Index(content, "PID 100") > 0, "email does not contain expected content")
		assert.True(t, strings.Index(content, "The mysql service") > 0, "email does not contain expected content")
		return nil
	})
	assert.Nil(t, err)
}

func TestEmailEventProcessDoesNotExist(t *testing.T) {
	t.Parallel()
	alert := validProcessEvent(ProcessDoesNotExist)

	err := validEmailSetup().TriggerEmail(alert, func(e *EmailNotifier, doc bytes.Buffer) error {
		content := string(doc.Bytes())
		assert.True(t, strings.Index(content, "can't locate") > 0, "email does not contain expected content")
		assert.True(t, strings.Index(content, "the mysql service") > 0, "email does not contain expected content")
		return nil
	})
	assert.Nil(t, err)
}

func TestEmailEventRuleRecovered(t *testing.T) {
	t.Parallel()
	alert := validRuleEvent(RuleRecovered)

	err := validEmailSetup().TriggerEmail(alert, func(e *EmailNotifier, doc bytes.Buffer) error {
		content := string(doc.Bytes())
		assert.True(t, strings.Index(content, "has recovered") > 0, "email does not contain expected content")
		assert.True(t, strings.Index(content, "[mysql]") > 0, "email does not contain expected content")
		return nil
	})
	assert.Nil(t, err)
}

func validRuleEvent(etype EventType) *Event {
	svc := &Service{&Entity{"mysql", nil, metrics.NewProcessStore("/proc", 15), nil}, nil, services.WithStatus(100, services.Up), nil}
	return &Event{
		etype, svc, &Rule{svc, "memory", "rss", GT, "64m", 64 * 1024 * 1024, 0, false, 1, 0, Ok, []Action{mockAction()}},
	}
}

func validProcessEvent(etype EventType) *Event {
	svc := &Service{&Entity{"mysql", nil, metrics.NewProcessStore("/proc", 15), nil}, nil, services.WithStatus(100, services.Up), nil}
	return &Event{etype, svc, nil}
}

func validEmailSetup() *EmailNotifier {
	return &EmailNotifier{
		"mike", "fuzzbucket", "smtp.gmail.com", "mike@example.org", "mperham@gmail.com", ""}
}

type TestAction struct {
	events []*Event
}

func (t *TestAction) Latest() *Event {
	if len(t.events) == 0 {
		return nil
	}
	return t.events[len(t.events)-1]
}

func (t *TestAction) Size() int {
	return len(t.events)
}

func (t *TestAction) Trigger(e *Event) error {
	t.events = append(t.events, e)
	return nil
}

func mockAction() *TestAction {
	return &TestAction{}
}
