package inspeqtor

import (
	"bytes"
	"errors"
	"fmt"
	"net/smtp"
)

type Notifier interface {
	Notify(data interface{}) error
}

type EmailNotifier struct {
	username string
	password string
	host     string
	from     string
	to       string
}

func SetupNotification(name string, vars map[string]string) (Notifier, error) {
	switch name {
	case "email":
		en, err := setupEmail(vars)
		// Go's nil semantics around interfaces are a little jacked up
		// You'd think this could be a simple passthru.  You'd be wrong.
		if en == nil {
			return nil, err
		}
		return en, nil
	case "gmail":
		vars["hostname"] = "smtp.gmail.com"
		en, err := setupEmail(vars)
		if en == nil {
			return nil, err
		}
		return en, nil
	default:
		return nil, nil
	}
}

func setupEmail(hash map[string]string) (*EmailNotifier, error) {
	usr, ok := hash["username"]
	if !ok {
		return nil, errors.New("You must have a username configured to send email")
	}
	pwd, ok := hash["password"]
	if !ok {
		return nil, errors.New("You must have a password configured to send email")
	}
	host, ok := hash["hostname"]
	if !ok {
		return nil, errors.New("You must have a hostname configured to send email")
	}
	to, ok := hash["to"]
	if !ok {
		return nil, errors.New("You must have a to address configured to send email")
	}
	from, ok := hash["from"]
	if !ok {
		from = "Inspeqtor <noreply@example.com>"
	}

	return &EmailNotifier{usr, pwd, host, from, to}, nil
}

func (e *EmailNotifier) Notify(data interface{}) error {
	auth := smtp.PlainAuth("", e.username, "", e.host)
	err := smtp.SendMail(e.host+":587", auth, e.from,
		[]string{e.to},
		bytes.NewBufferString(fmt.Sprint(data)).Bytes())
	if err != nil {
		return err
	}
	return nil
}
