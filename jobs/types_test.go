package jobs

import (
	"bytes"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/mperham/inspeqtor"
	"github.com/mperham/inspeqtor/channels"
	"github.com/stretchr/testify/assert"
)

func TestJobs(t *testing.T) {
	t.Parallel()

	global := &inspeqtor.ConfigFile{GlobalConfig: inspeqtor.Defaults, AlertRoutes: map[string]*inspeqtor.AlertRoute{
		"":     &inspeqtor.AlertRoute{Name: "", Channel: "null", Config: map[string]string{}},
		"mike": &inspeqtor.AlertRoute{Name: "mike", Channel: "null", Config: map[string]string{}},
	}}

	jobs, err := Parse(global, "fixtures")
	assert.Nil(t, err)
	assert.NotNil(t, jobs)

	for _, j := range jobs {
		j.alerter = mockAction()
	}

	assert.Equal(t, len(jobs), 4)
	assert.Equal(t, jobs["bank_check"].Interval, time.Minute)
	assert.Equal(t, jobs["bank_check"].Parameters["owner"], "mike")
	assert.Equal(t, jobs["something_else"].Interval, 48*time.Hour)
	assert.Equal(t, jobs["something_else"].Parameters["owner"], "")

	var buf bytes.Buffer
	jobDone(nil, []string{}, &buf)

	line, err := buf.ReadString('\n')
	assert.Nil(t, err)
	assert.Equal(t, line, "Usage: job_done <name>\n")

	buf = bytes.Buffer{}
	jobDone(nil, []string{"bunk"}, &buf)

	line, err = buf.ReadString('\n')
	assert.Nil(t, err)
	assert.Equal(t, line, "Error: no such job \"bunk\", verify spelling or reload Inspeqtor\n")

	before := jobs["bank_check"].LastRun
	buf = bytes.Buffer{}
	jobDone(nil, []string{"bank_check"}, &buf)

	line, err = buf.ReadString('\n')
	assert.Nil(t, err)
	assert.Equal(t, line, "OK\n")
	after := jobs["bank_check"].LastRun
	assert.NotEqual(t, before, after)

	sleepTime := check(jobs)
	assert.True(t, sleepTime < time.Minute)
	assert.True(t, sleepTime > time.Duration(59)*time.Second)
}

func TestWatch(t *testing.T) {
	t.Parallel()

	act := mockAction()
	sleepTime := check(map[string]*Job{
		"one": &Job{"one", time.Minute, nil, time.Now(), act, inspeqtor.Ok},
		"two": &Job{"two", time.Hour, nil, time.Now(), act, inspeqtor.Ok},
	})
	assert.True(t, sleepTime < time.Minute)
	assert.True(t, sleepTime > time.Duration(59)*time.Second)
	assert.Equal(t, act.Size(), 0)

	sleepTime = check(map[string]*Job{
		"one": &Job{"one", time.Minute, nil, time.Now(), act, inspeqtor.Ok},
		"two": &Job{"two", time.Hour, nil, time.Unix(time.Now().Unix()-3599, 0), act, inspeqtor.Ok},
	})
	assert.True(t, sleepTime < time.Minute)
	assert.True(t, sleepTime < time.Second)
	assert.Equal(t, act.Size(), 0)

	sleepTime = check(map[string]*Job{
		"one": &Job{"one", time.Minute, nil, time.Now(), act, inspeqtor.Ok},
		"two": &Job{"two", time.Hour, nil, time.Unix(time.Now().Unix()-3601, 0), act, inspeqtor.Ok},
	})
	assert.True(t, sleepTime < time.Minute)
	assert.True(t, sleepTime > time.Duration(59)*time.Second)
	assert.Equal(t, act.Size(), 1)
	assert.Equal(t, act.Latest().Type, JobOverdue)
	assert.Equal(t, act.Latest().Eventable.Name(), "two")
}

func TestNotifications(t *testing.T) {
	job := New("test", time.Hour, map[string]string{})
	expected := ""
	job.alerter = &channels.FlowdockNotifier{Token: "abcdef", Sender: func(url string, token string, msg url.Values) error {
		expected = msg["content"][0]
		return nil
	}}
	job.LastRun = time.Unix(time.Now().Unix()-3601, 0)

	sleepTime := check(map[string]*Job{"test": job})
	assert.Equal(t, sleepTime, time.Hour)
	assert.True(t, strings.Index(expected, "Recurring job \"test\" is overdue") > 0)
}

type TestAction struct {
	events []*inspeqtor.Event
}

func (t *TestAction) Latest() *inspeqtor.Event {
	if len(t.events) == 0 {
		return nil
	}
	return t.events[len(t.events)-1]
}

func (t *TestAction) Size() int {
	return len(t.events)
}

func (t *TestAction) Trigger(e *inspeqtor.Event) error {
	t.events = append(t.events, e)
	return nil
}

func mockAction() *TestAction {
	return &TestAction{}
}
