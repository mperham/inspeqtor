package channels

import (
	"strings"
	"testing"

	"github.com/mperham/inspeqtor"
	"github.com/mperham/inspeqtor/util"
	"github.com/stretchr/testify/assert"
)

func swapEvent(check inspeqtor.Checkable, status inspeqtor.EventType) *inspeqtor.Event {
	return &inspeqtor.Event{Type: status, Eventable: check, Rule: &inspeqtor.Rule{
		Entity: check, MetricFamily: "swap", MetricName: "", Op: inspeqtor.GT,
		DisplayThreshold: "20%", Threshold: 20, CurrentValue: 0, PerSec: false,
		CycleCount: 1, TrippedCount: 0, State: inspeqtor.Ok, Actions: nil}}
}

func TestCampfire(t *testing.T) {
	util.LogInfo = true
	t.Parallel()

	check := inspeqtor.NewHost()
	action, err := buildCampfireNotifier(check, map[string]string{})
	assert.Nil(t, action)
	assert.NotNil(t, err)

	action, err = buildCampfireNotifier(check, map[string]string{"team": "contribsys", "token": "abcdef123456", "room": "12345"})
	assert.Nil(t, err)
	assert.NotNil(t, action)
	sn := action.(*campfireNotifier)

	var params map[string]string
	var theurl string

	sendHere := func(url, token string, body map[string]map[string]string) error {
		theurl = url
		params = body["message"]
		return nil
	}

	alert := swapEvent(check, inspeqtor.RuleFailed)
	err = sn.trigger(alert, sendHere)
	assert.Nil(t, err)
	assert.Equal(t, theurl, "https://contribsys.campfirenow.com/room/12345/speak.json")
	assert.True(t, strings.Index(params["body"], "localhost: swap is greater than than 20%") > -1)

	alert = swapEvent(check, inspeqtor.RuleRecovered)
	err = sn.trigger(alert, sendHere)
	assert.Nil(t, err)
	assert.True(t, strings.Index(params["body"], "localhost: swap has recovered.") > -1)

	svc := inspeqtor.NewService("sidekiq")
	alert = &inspeqtor.Event{Type: inspeqtor.ProcessDoesNotExist, Eventable: svc, Rule: nil}
	err = sn.trigger(alert, sendHere)
	assert.Nil(t, err)
	assert.True(t, strings.Index(params["body"], "[sidekiq] does not exist.") > -1)

	alert = &inspeqtor.Event{Type: inspeqtor.ProcessExists, Eventable: svc, Rule: nil}
	err = sn.trigger(alert, sendHere)
	assert.Nil(t, err)
	assert.True(t, strings.Index(params["body"], "[sidekiq] now running with PID 0") > -1)
}
