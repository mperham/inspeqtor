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
Subject: [{{.Alert.Service.Name}}] {{.Alert.Rule.MetricName}} is {{.Alert.Rule.Op}} than {{.Alert.Rule.Threshold}}

[{{.Alert.Service.Name}}] {{.Alert.Rule.MetricName}} is {{.Alert.Rule.Op}} than {{.Alert.Rule.Threshold}}
`
)

var (
	email = template.Must(template.New("emailTemplate").Parse(emailTemplate))
)

type EmailSender func(e *EmailConfig, doc bytes.Buffer) error

type EmailConfig struct {
	Username string
	Password string
	Host     string
	From     string
	To       string
}

type EmailAlert struct {
	Alert  *Alert
	Config *EmailConfig
}

func SetupNotification(name string, vars map[string]string) (Action, error) {
	switch name {
	case "email":
		en := &EmailConfig{}
		err := en.Setup(vars)
		if err != nil {
			return nil, err
		}
		return en, nil
	case "gmail":
		vars["hostname"] = "smtp.gmail.com"
		en := &EmailConfig{}
		err := en.Setup(vars)
		if err != nil {
			return nil, err
		}
		return en, nil
	default:
		return nil, nil
	}
}

func (e *EmailConfig) Name() string {
	return "email"
}

func (e *EmailConfig) Trigger(alert *Alert) error {
	return e.TriggerEmail(alert, sendEmail)
}

func (e *EmailConfig) TriggerEmail(alert *Alert, sender EmailSender) error {
	var doc bytes.Buffer
	err := email.Execute(&doc, &EmailAlert{alert, e})
	if err != nil {
		return err
	}
	return sender(e, doc)
}

func sendEmail(e *EmailConfig, doc bytes.Buffer) error {
	auth := smtp.PlainAuth("", e.Username, "", e.Host)
	err := smtp.SendMail(e.Host+":587", auth, e.From,
		[]string{e.To}, doc.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func (e *EmailConfig) Setup(hash map[string]string) error {
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
		return errors.New("You must have a hostname configured to send email")
	}
	to, ok := hash["to"]
	if !ok {
		return errors.New("You must have a to address configured to send email")
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
