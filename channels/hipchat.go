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
	hipchatTemplates map[inspeqtor.EventType]*template.Template
)

func init() {
	inspeqtor.Notifiers["hipchat"] = buildHipchatNotifier
}

func buildHipchatTemplates() {
	hipchatTemplates = map[inspeqtor.EventType]*template.Template{}
	for _, event := range inspeqtor.Events {
		str := event.String()
		asset, err := Asset("templates/chat/" + str + ".txt")
		if err != nil {
			panic(err)
		}
		hipchatTemplates[event] = template.Must(template.New(str).Parse(string(asset)))
	}
}

func buildHipchatNotifier(check inspeqtor.Eventable, config map[string]string) (inspeqtor.Action, error) {
	if hipchatTemplates == nil {
		buildHipchatTemplates()
	}

	if _, ok := config["token"]; !ok {
		return nil, errors.New("Hipchat integration missing 'token' value")
	}
	if _, ok := config["room"]; !ok {
		return nil, errors.New("Hipchat integration missing 'room' value")
	}
	return &hipchatNotifier{config["token"], config["room"]}, nil
}

type hipchatNotifier struct {
	token string
	room  string
}

func (e *hipchatNotifier) url() string {
	return fmt.Sprintf("https://api.hipchat.com/v1/rooms/message?auth_token=%s", e.token)
}

func (e *hipchatNotifier) Trigger(alert *inspeqtor.Event) error {
	return e.trigger(alert, sendHipchatAlert)
}

func (e *hipchatNotifier) trigger(event *inspeqtor.Event, sender func(string, string, url.Values) error) error {
	var doc bytes.Buffer
	template := hipchatTemplates[event.Type]
	err := template.Execute(&doc, event)
	if err != nil {
		return err
	}

	txt := string(doc.Bytes())

	values := url.Values{}
	values.Set("room_id", e.room)
	values.Set("from", inspeqtor.Name)
	values.Set("message", txt)

	return sender(e.url(), e.token, values)
}

func sendHipchatAlert(url, token string, msg url.Values) error {
	util.Debug("Sending hipchat alert to %s", url)

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(msg.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if resp != nil {
		if resp.StatusCode != 200 {
			util.Warn("Unable to send hipchat alert: %d", resp.StatusCode)
		}
		resp.Body.Close()
	}
	return err
}
