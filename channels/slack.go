package channels

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"text/template"

	"github.com/mperham/inspeqtor"
	"github.com/mperham/inspeqtor/util"
)

var (
	slackTemplates map[inspeqtor.EventType]*template.Template
)

func init() {
	inspeqtor.Notifiers["slack"] = buildSlackNotifier
}

func buildSlackTemplates() {
	slackTemplates = map[inspeqtor.EventType]*template.Template{}
	for _, event := range inspeqtor.Events {
		str := event.String()
		asset, err := Asset("templates/chat/" + str + ".txt")
		if err != nil {
			panic(err)
		}
		slackTemplates[event] = template.Must(template.New(str).Parse(string(asset)))
	}
}

func buildSlackNotifier(check inspeqtor.Eventable, config map[string]string) (inspeqtor.Action, error) {
	if slackTemplates == nil {
		buildSlackTemplates()
	}

	if _, ok := config["url"]; !ok {
		return nil, errors.New("Slack notifier missing 'url' value")
	}
	icon := "ghost"
	if _, ok := config["icon_emoji"]; ok {
		icon = config["icon_emoji"]
	}
	return &slackNotifier{config["url"], "Inspeqtor Pro", icon}, nil
}

type slackNotifier struct {
	url       string
	username  string
	iconEmoji string
}

type slackMessage struct {
	Text     string `json:"text"`
	Username string `json:"username"`
	Emoji    string `json:"icon_emoji"`
}

func (e *slackNotifier) Trigger(alert *inspeqtor.Event) error {
	return e.triggerSlack(alert, sendSlackAlert)
}

func (e *slackNotifier) triggerSlack(event *inspeqtor.Event, sender func(string, url.Values) error) error {
	var doc bytes.Buffer
	template := slackTemplates[event.Type]
	err := template.Execute(&doc, event)
	if err != nil {
		return err
	}

	msg := slackMessage{string(doc.Bytes()), e.username, ":" + e.iconEmoji + ":"}
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	values := url.Values{}
	values.Set("payload", string(data))
	return sender(e.url, values)
}

func sendSlackAlert(url string, params url.Values) error {
	util.Debug("Sending slack alert to %s", url)
	resp, err := http.PostForm(url, params)
	if resp != nil {
		resp.Body.Close()
	}
	return err
}
