package inspeqtor

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"inspeqtor/metrics"
	"inspeqtor/services"
	"strings"
	"testing"
)

func TestGmailNotifier(t *testing.T) {
	i, err := SetupNotification("gmail", map[string]string{
		"username": "mike",
		"password": "fuzzbucket",
		"to":       "mike@example.org",
	})
	assert.Nil(t, err)
	assert.NotNil(t, i)
}

func TestEmailNotifier(t *testing.T) {
	i, err := SetupNotification("email", map[string]string{
		"username": "mike",
		"password": "fuzzbucket",
		"hostname": "smtp.example.com",
		"to":       "mike@example.org",
	})
	assert.Nil(t, err)
	assert.NotNil(t, i)
}

func TestMissingEmailNotifier(t *testing.T) {
	i, err := SetupNotification("email", map[string]string{
		"username": "mike",
		"password": "fuzzbucket",
		"to":       "mike@example.org",
	})
	assert.NotNil(t, err)
	assert.Nil(t, i)
}

func TestEmailTrigger(t *testing.T) {
	alert := &Alert{
		&Service{"mysql", 0, services.Down, nil, metrics.NewStore(), nil},
		&Rule{"memory", "rss", GT, 64 * 1024 * 1024, 1, 0, nil},
	}

	err := validEmailSetup().TriggerEmail(alert, func(e *EmailConfig, doc bytes.Buffer) error {
		content := string(doc.Bytes())
		assert.True(t, strings.Index(content, "[mysql]") > 0, "email does not contain expected content")
		assert.True(t, strings.Index(content, "memory(rss)") > 0, "email does not contain expected content")
		return nil
	})
	assert.Nil(t, err)
}

func validEmailSetup() *EmailConfig {
	return &EmailConfig{
		"mike", "fuzzbucket", "smtp.gmail.com", "mike@example.org", "mperham@gmail.com"}
}
