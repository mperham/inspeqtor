package channels

import (
	"bytes"
	"strings"
	"testing"

	"github.com/mperham/inspeqtor"
	"github.com/stretchr/testify/assert"
)

// Verify the email templates are packaged correctly in Pro
func TestEmail(t *testing.T) {
	svc := inspeqtor.NewService("sidekiq")
	action, err := inspeqtor.Notifiers["email"](svc, map[string]string{
		"username":    "mike",
		"password":    "fuzzbucket",
		"smtp_server": "smtp.example.com",
		"to_email":    "mike@example.org",
	})
	assert.Nil(t, err)
	emailer := action.(*inspeqtor.EmailNotifier)
	alert := &inspeqtor.Event{Type: inspeqtor.ProcessDoesNotExist, Eventable: svc, Rule: nil}
	err = emailer.TriggerEmail(alert, func(e *inspeqtor.EmailNotifier, doc bytes.Buffer) error {
		content := string(doc.Bytes())
		assert.True(t, strings.Index(content, "can't locate") > 0, "email does not contain expected content")
		assert.True(t, strings.Index(content, "the sidekiq service") > 0, "email does not contain expected content")
		return nil
	})
}
