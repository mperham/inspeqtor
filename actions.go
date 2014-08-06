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
Subject: [{{.Alert.Rule.EntityName}}] {{.Alert.Rule.MetricName}} is {{.Alert.Rule.Op}} than {{.Alert.Rule.Threshold}}

[{{.Alert.Rule.EntityName}}] {{.Alert.Rule.MetricName}} is {{.Alert.Rule.Op}} than {{.Alert.Rule.Threshold}}
`
)

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

func SetupNotification(name string, vars map[string]string) (Action, error) {
	switch name {
	case "email":
		en := &EmailNotifier{}
		err := en.Setup(vars)
		if err != nil {
			return nil, err
		}
		return en, nil
	case "gmail":
		vars["hostname"] = "smtp.gmail.com"
		en := &EmailNotifier{}
		err := en.Setup(vars)
		if err != nil {
			return nil, err
		}
		return en, nil
	default:
		return nil, nil
	}
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
