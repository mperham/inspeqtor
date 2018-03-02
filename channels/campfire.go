package channels

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"text/template"

	"github.com/mperham/inspeqtor"
	"github.com/mperham/inspeqtor/util"
)

var (
	campfireTemplates map[inspeqtor.EventType]*template.Template
)

func init() {
	inspeqtor.Notifiers["campfire"] = buildCampfireNotifier
}

func buildCampfireTemplates() {
	campfireTemplates = map[inspeqtor.EventType]*template.Template{}
	for _, event := range inspeqtor.Events {
		str := event.String()
		asset, err := Asset("templates/chat/" + str + ".txt")
		if err != nil {
			panic(err)
		}
		campfireTemplates[event] = template.Must(template.New(str).Parse(string(asset)))
	}
}

func buildCampfireNotifier(check inspeqtor.Eventable, config map[string]string) (inspeqtor.Action, error) {
	if campfireTemplates == nil {
		buildCampfireTemplates()
	}

	if _, ok := config["team"]; !ok {
		return nil, errors.New("Campfire integration missing 'team' value")
	}
	if _, ok := config["token"]; !ok {
		return nil, errors.New("Campfire integration missing 'token' value")
	}
	if _, ok := config["room"]; !ok {
		return nil, errors.New("Campfire integration missing 'room' value")
	}
	return &campfireNotifier{config["team"], config["token"], config["room"]}, nil
}

type campfireNotifier struct {
	team  string
	token string
	room  string
}

func (e *campfireNotifier) url() string {
	return fmt.Sprintf("https://%s.campfirenow.com/room/%s/speak.json", e.team, e.room)
}

func (e *campfireNotifier) Trigger(alert *inspeqtor.Event) error {
	return e.trigger(alert, sendCampfireAlert)
}

func (e *campfireNotifier) trigger(event *inspeqtor.Event, sender func(string, string, map[string]map[string]string) error) error {
	var doc bytes.Buffer
	template := campfireTemplates[event.Type]
	err := template.Execute(&doc, event)
	if err != nil {
		return err
	}

	txt := string(doc.Bytes())

	msg := map[string]map[string]string{
		"message": map[string]string{
			"body": txt,
			"type": "TextMessage",
		},
	}

	return sender(e.url(), e.token, msg)
}

func sendCampfireAlert(url, token string, msg map[string]map[string]string) error {
	util.Debug("Sending campfire alert to %s", url)

	client := &http.Client{}
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(jsonMsg))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Inspeqtor Pro")
	req.SetBasicAuth(token, "X")
	resp, err := client.Do(req)
	if resp != nil {
		if resp.StatusCode != 201 {
			util.Warn("Unable to send campfire alert: %d", resp.StatusCode)
		}
		resp.Body.Close()
	}
	return err
}
