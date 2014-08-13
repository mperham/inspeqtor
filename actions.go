package inspeqtor

import (
	"bytes"
	"errors"
	"fmt"
	"inspeqtor/services"
	"net/smtp"
	"text/template"
)

const (
	emailTemplate = `
From: {{.Config.From}}
To: {{.Config.To}}
Subject: [{{.Rule.EntityName}}] {{.Rule.MetricName}} is {{.Rule.Op}} than {{.Rule.Threshold}}

[{{.Rule.EntityName}}] {{.Rule.MetricName}} is {{.Rule.Op}} than {{.Rule.Threshold}}
`
)

/*
 Global parser parses "send alert" statements and creates alert route objects.
 Inq parser parses rules and creates action objects based on the name and param.
*/

/*
 An Action is something which is triggered when a rule is broken.  This is typically
 either a notification or to restart the service.
*/

type ActionBuilder func(Checkable, *AlertRoute) (Action, error)
type NotificationBuilder func(Checkable, map[string]string) (Action, error)

var (
	Actions = map[string]ActionBuilder{
		"alert":   buildNotifier,
		"restart": buildRestarter,
	}
	Notifiers = map[string]NotificationBuilder{
		"email": buildEmailNotifier,
		"gmail": buildGmailNotifier,
	}
)

func buildNotifier(check Checkable, route *AlertRoute) (Action, error) {
	funk := Notifiers[route.Channel]
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

func (r Restarter) Trigger(alert *Alert) error {
	r.Service.PID = 0
	r.Service.Status = services.Unknown

	go r.Service.Restart()
	return nil
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
	params["hostname"] = "smtp.gmail.com"
	return buildEmailNotifier(check, params)
}

var (
	email = template.Must(template.New("emailTemplate").Parse(emailTemplate))
)

type EmailSender func(e *EmailNotifier, doc bytes.Buffer) error

type EmailNotifier struct {
	Username string
	Password string
	Host     string
	From     string
	To       string
}

type EmailAlert struct {
	*Alert
	Config *EmailNotifier
}

func ValidateChannel(name string, channel string, config map[string]string) (*AlertRoute, error) {
	return &AlertRoute{name, channel, config}, nil
}

func (e EmailNotifier) Trigger(alert *Alert) error {
	return e.TriggerEmail(alert, sendEmail)
}

func (e *EmailNotifier) TriggerEmail(alert *Alert, sender EmailSender) error {
	var doc bytes.Buffer
	err := email.Execute(&doc, &EmailAlert{alert, e})
	if err != nil {
		return err
	}
	return sender(e, doc)
}

func sendEmail(e *EmailNotifier, doc bytes.Buffer) error {
	auth := smtp.PlainAuth("", e.Username, "", e.Host)
	err := smtp.SendMail(e.Host+":587", auth, e.From,
		[]string{e.To}, doc.Bytes())
	if err != nil {
		return err
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
	host, ok := hash["hostname"]
	if !ok {
		return errors.New("You must have a 'hostname' parameter pointing to your SMTP server")
	}
	to, ok := hash["email"]
	if !ok {
		return errors.New("You are missing the 'email' parameter, needed to specify a To address for your alert emails")
	}
	from, ok := hash["from"]
	if !ok {
		from = "Inspeqtor <noreply@example.com>"
	}

	e.Username = usr
	e.Password = pwd
	e.Host = host
	e.From = from
	e.To = to

	return nil
}
