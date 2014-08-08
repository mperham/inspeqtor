package inspeqtor

import (
	"bytes"
	"errors"
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

type ActionBuilder func(map[string]string) (Action, error)

var (
	Actions = map[string]ActionBuilder{
		"email": buildEmailAction,
		"gmail": buildGmailAction,
	}
)

func buildEmailAction(params map[string]string) (Action, error) {
	en := &EmailNotifier{}
	err := en.Setup(params)
	if err != nil {
		return nil, err
	}
	return en, nil
}

func buildGmailAction(params map[string]string) (Action, error) {
	params["hostname"] = "smtp.gmail.com"
	return buildEmailAction(params)
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

func ValidateChannel(name string, channel string, config map[string]string) (AlertRoute, error) {
	return AlertRoute{name, channel, config}, nil
}

func (e EmailNotifier) Name() string {
	return "email"
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

func (e *EmailNotifier) Setup(hash map[string]string) error {
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
