package channels

import (
	"net/url"
	"strings"
	"testing"

	"github.com/mperham/inspeqtor"
	"github.com/mperham/inspeqtor/util"
	"github.com/stretchr/testify/assert"
)

func TestFlowdock(t *testing.T) {
	util.LogInfo = true
	t.Parallel()

	check := inspeqtor.NewHost()
	action, err := buildFlowdockNotifier(check, map[string]string{})
	assert.Nil(t, action)
	assert.NotNil(t, err)

	action, err = buildFlowdockNotifier(check, map[string]string{"token": "abcdef123456"})
	assert.Nil(t, err)
	assert.NotNil(t, action)
	sn := action.(*FlowdockNotifier)

	var params url.Values
	var theurl string

	sendHere := func(url, token string, body url.Values) error {
		theurl = url
		params = body
		return nil
	}

	alert := swapEvent(check, inspeqtor.RuleFailed)
	err = sn.trigger(alert, sendHere)
	assert.Nil(t, err)
	assert.Equal(t, theurl, "https://api.flowdock.com/v1/messages/team_inbox/abcdef123456")
	assert.True(t, strings.Index(params["content"][0], "localhost: swap is greater than than 20%") > -1)

	alert = swapEvent(check, inspeqtor.RuleRecovered)
	err = sn.trigger(alert, sendHere)
	assert.Nil(t, err)
	assert.True(t, strings.Index(params["content"][0], "localhost: swap has recovered.") > -1)

	svc := inspeqtor.NewService("sidekiq")
	alert = &inspeqtor.Event{Type: inspeqtor.ProcessDoesNotExist, Eventable: svc, Rule: nil}
	err = sn.trigger(alert, sendHere)
	assert.Nil(t, err)
	assert.True(t, strings.Index(params["content"][0], "[sidekiq] does not exist.") > -1)

	alert = &inspeqtor.Event{Type: inspeqtor.ProcessExists, Eventable: svc, Rule: nil}
	err = sn.trigger(alert, sendHere)
	assert.Nil(t, err)
	assert.True(t, strings.Index(params["content"][0], "[sidekiq] now running with PID 0") > -1)
}
