package redacted

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/mperham/redacted/util"
	"net/smtp"
	"strings"
	"text/template"
)

type Action interface {
	Trigger(event *Event) error
}

var (
	emailTemplates = map[EventType]*template.Template{}
)

func init() {
	for _, event := range Events {
		str := event.String()
		asset, err := Asset("templates/email/" + str + ".txt")
		if err != nil {
			panic(err)
		}
		emailTemplates[event] = template.Must(template.New(str).Parse(string(asset)))
	}
}

/*
 Global parser parses "send alert" statements and creates alert route objects.
 Inq parser parses rules and creates action objects based on the name and param.
*/

/*
 An Action is something which is triggered when a rule is broken.  This is typically
 either a Notification or to Restart the service.
*/
type ActionBuilder func(Checkable, *AlertRoute) (Action, error)

/*
 A Notifier is a route to send an alert somewhere else.  The global
 conf sets up the necessary params for the notification to work.
*/
type NotifierBuilder func(Checkable, map[string]string) (Action, error)

var (
	Actions = map[string]ActionBuilder{
		"alert":   buildAlerter,
		"restart": buildRestarter,
	}
	Notifier = map[string]NotifierBuilder{
		"email": buildEmailNotifier,
		"gmail": buildGmailNotifier,
	}
)

func buildAlerter(check Checkable, route *AlertRoute) (Action, error) {
	funk := Notifier[route.Channel]
	if funk == nil {
		// TODO Include valid channels
		return nil, errors.New(fmt.Sprintf("No such notification scheme: %s", route.Channel))
	}
	return funk(check, route.Config)
}

func buildRestarter(check Checkable, _ *AlertRoute) (Action, error) {
	switch check.(type) {
	case *Service:
		return &Restarter{check.(*Service)}, nil
	default:
		return nil, errors.New("Cannot restart " + check.Name())
	}
}

type Restarter struct {
	*Service
}

func (r Restarter) Trigger(event *Event) error {
	return r.Service.Restart()
}

func buildEmailNotifier(check Checkable, config map[string]string) (Action, error) {
	en := &EmailNotifier{}
	err := en.setup(config)
	if err != nil {
		return nil, err
	}
	return en, nil
}

func buildGmailNotifier(check Checkable, params map[string]string) (Action, error) {
	params["smtp_server"] = "smtp.gmail.com"
	return buildEmailNotifier(check, params)
}

type EmailSender func(e *EmailNotifier, doc bytes.Buffer) error

type EmailNotifier struct {
	Username string
	Password string
	Host     string
	From     string
	To       string
}

type EmailEvent struct {
	*Event
	Config *EmailNotifier
}

func ValidateChannel(name string, channel string, config map[string]string) (*AlertRoute, error) {
	_, ok := Notifier[channel]
	if !ok {
		return nil, errors.New("No such notification type: " + channel)
	}
	return &AlertRoute{name, channel, config}, nil
}

func (e EmailNotifier) Trigger(event *Event) error {
	return e.TriggerEmail(event, sendEmail)
}

func (e *EmailNotifier) TriggerEmail(event *Event, sender EmailSender) error {
	var doc bytes.Buffer
	template := emailTemplates[event.Type]
	err := template.Execute(&doc, &EmailEvent{event, e})
	if err != nil {
		return err
	}
	return sender(e, doc)
}

func sendEmail(e *EmailNotifier, doc bytes.Buffer) error {
	if strings.Index(e.To, "@example.com") > 0 {
		util.Warn("Invalid email configured: %s", e.To)
		util.Warn(string(doc.Bytes()))
	} else {
		util.Debug("Sending email to %s", e.To)
		util.Debug("Sending email:\n%s", string(doc.Bytes()))
		auth := smtp.PlainAuth("", e.Username, e.Password, e.Host)
		err := smtp.SendMail(e.Host+":587", auth, e.From,
			[]string{e.To}, doc.Bytes())
		if err != nil {
			return err
		}
	}
	return nil
}

func (e *EmailNotifier) setup(hash map[string]string) error {
	usr, ok := hash["username"]
	if !ok {
		return errors.New("You must have a username configured to send email")
	}
	pwd, ok := hash["password"]
	if !ok {
		return errors.New("You must have a password configured to send email")
	}
	host, ok := hash["smtp_server"]
	if !ok {
		return errors.New("You must have a 'smtp_server' parameter pointing to your SMTP server")
	}
	to, ok := hash["to_email"]
	if !ok {
		return errors.New("You are missing the 'to_email' parameter, needed to specify a To address for your alert emails")
	}
	from, ok := hash["from_email"]
	if !ok {
		from = "redacted@example.com"
	} else {
		from = from
	}

	e.Username = usr
	e.Password = pwd
	e.Host = host
	e.From = from
	e.To = to

	return nil
}
