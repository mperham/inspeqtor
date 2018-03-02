package channels

import (
	"net/url"
	"strings"
	"testing"

	"github.com/mperham/inspeqtor"
	"github.com/mperham/inspeqtor/util"
	"github.com/stretchr/testify/assert"
)

func TestSlack(t *testing.T) {
	util.LogInfo = true
	t.Parallel()

	check := inspeqtor.NewHost()
	action, err := buildSlackNotifier(check, map[string]string{})
	assert.Nil(t, action)
	assert.NotNil(t, err)

	action, err = buildSlackNotifier(check, map[string]string{"team": "contribsys", "url": "https://acmecorp.slack.com/services/hooks/incoming-webhook?token=xxx/xxx/xxx", "icon_emoji": "beer"})
	assert.Nil(t, err)
	assert.NotNil(t, action)
	sn := action.(*slackNotifier)

	var params url.Values
	var theurl string

	sendHere := func(url string, values url.Values) error {
		theurl = url
		params = values
		return nil
	}

	alert := swapEvent(check, inspeqtor.RuleFailed)
	err = sn.triggerSlack(alert, sendHere)
	assert.Nil(t, err)
	assert.Equal(t, "https://acmecorp.slack.com/services/hooks/incoming-webhook?token=xxx/xxx/xxx", theurl)
	assert.True(t, strings.Index(params["payload"][0], "localhost: swap is greater than than 20%") > 0)

	alert = swapEvent(check, inspeqtor.RuleRecovered)
	err = sn.triggerSlack(alert, sendHere)
	assert.Nil(t, err)
	assert.True(t, strings.Index(params["payload"][0], "localhost: swap has recovered.") > 0)

	svc := inspeqtor.NewService("sidekiq")
	alert = &inspeqtor.Event{Type: inspeqtor.ProcessDoesNotExist, Eventable: svc, Rule: nil}
	err = sn.triggerSlack(alert, sendHere)
	assert.Nil(t, err)
	assert.True(t, strings.Index(params["payload"][0], "[sidekiq] does not exist.") > 0)

	alert = &inspeqtor.Event{Type: inspeqtor.ProcessExists, Eventable: svc, Rule: nil}
	err = sn.triggerSlack(alert, sendHere)
	assert.Nil(t, err)
	assert.True(t, strings.Index(params["payload"][0], "[sidekiq] now running with PID 0") > 0)
}
