package channels

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"text/template"

	"github.com/mperham/inspeqtor"
	"github.com/mperham/inspeqtor/util"
)

var (
	flowdockTemplates map[inspeqtor.EventType]*template.Template
)

func init() {
	inspeqtor.Notifiers["flowdock"] = buildFlowdockNotifier
}

func buildFlowdockTemplates() {
	flowdockTemplates = map[inspeqtor.EventType]*template.Template{}
	for _, event := range inspeqtor.Events {
		str := event.String()
		asset, err := Asset("templates/chat/" + str + ".txt")
		if err != nil {
			panic(err)
		}
		flowdockTemplates[event] = template.Must(template.New(str).Parse(string(asset)))
	}
}

func buildFlowdockNotifier(check inspeqtor.Eventable, config map[string]string) (inspeqtor.Action, error) {
	if _, ok := config["token"]; !ok {
		return nil, errors.New("flowdock integration missing 'token' value")
	}
	return &FlowdockNotifier{config["token"], sendFlowdockAlert}, nil
}

type FlowdockNotifier struct {
	Token  string
	Sender func(string, string, url.Values) error
}

func (e *FlowdockNotifier) url() string {
	return fmt.Sprintf("https://api.flowdock.com/v1/messages/team_inbox/%s", e.Token)
}

func (e *FlowdockNotifier) Trigger(alert *inspeqtor.Event) error {
	return e.trigger(alert, e.Sender)
}

func (e *FlowdockNotifier) trigger(event *inspeqtor.Event, sender func(string, string, url.Values) error) error {
	if flowdockTemplates == nil {
		buildFlowdockTemplates()
	}

	var doc bytes.Buffer
	template := flowdockTemplates[event.Type]
	err := template.Execute(&doc, event)
	if err != nil {
		return err
	}

	txt := string(doc.Bytes())

	values := url.Values{}
	values.Set("from_address", "inspeqtor@contribsys.com")
	values.Set("source", inspeqtor.Name)
	values.Set("subject", "Alert")
	values.Set("content", txt)

	return sender(e.url(), e.Token, values)
}

func sendFlowdockAlert(url, token string, msg url.Values) error {
	util.Debug("Sending flowdock alert to %s", url)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(msg.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if resp != nil {
		if resp.StatusCode != 200 {
			util.Warn("Unable to send flowdock alert: %d", resp.StatusCode)
		}
		resp.Body.Close()
	}
	return err
}
