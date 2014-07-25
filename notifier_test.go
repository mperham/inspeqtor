package inspeqtor

import (
	"bytes"
	"inspeqtor/services"
	"log"
	"strings"
	"testing"
)

func TestGmailNotifier(t *testing.T) {
	i, err := SetupNotification("gmail", map[string]string{
		"username": "mike",
		"password": "fuzzbucket",
		"to":       "mike@example.org",
	})
	ok(t, err)
	assert(t, i != nil, "Expecting valid notifier")
}

func TestEmailNotifier(t *testing.T) {
	i, err := SetupNotification("email", map[string]string{
		"username": "mike",
		"password": "fuzzbucket",
		"hostname": "smtp.example.com",
		"to":       "mike@example.org",
	})
	ok(t, err)
	assert(t, i != nil, "Expecting valid notifier")
}

func TestMissingEmailNotifier(t *testing.T) {
	i, err := SetupNotification("email", map[string]string{
		"username": "mike",
		"password": "fuzzbucket",
		"to":       "mike@example.org",
	})
	assert(t, err != nil, "Missing data should cause error")
	log.Printf("%v", i)
	assert(t, i == nil, "Missing data should not return notifier")
}

func TestEmailTrigger(t *testing.T) {
	alert := &Alert{
		&Service{"mysql", 0, services.Down, nil, nil},
		&Rule{"rss", GT, 64 * 1024 * 1024, 1, Ok, nil},
	}

	err := validEmailSetup().TriggerEmail(alert, func(e *EmailConfig, doc bytes.Buffer) error {
		content := string(doc.Bytes())
		assert(t, strings.Index(content, "[mysql]") > 0, "email does not contain expected content")
		return nil
	})
	ok(t, err)
}

func validEmailSetup() *EmailConfig {
	return &EmailConfig{
		"mike", "fuzzbucket", "smtp.gmail.com", "mike@example.org", "mperham@gmail.com"}
}
