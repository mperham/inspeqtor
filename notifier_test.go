package inspeqtor_test

import (
	"inspeqtor"
	"log"
	"testing"
)

func TestGmailNotifier(t *testing.T) {
	i, err := inspeqtor.SetupNotification("gmail", map[string]string{
		"username": "mike",
		"password": "fuzzbucket",
		"to":       "mike@example.org",
	})
	ok(t, err)
	assert(t, i != nil, "Expecting valid notifier")
}

func TestEmailNotifier(t *testing.T) {
	i, err := inspeqtor.SetupNotification("email", map[string]string{
		"username": "mike",
		"password": "fuzzbucket",
		"hostname": "smtp.example.com",
		"to":       "mike@example.org",
	})
	ok(t, err)
	assert(t, i != nil, "Expecting valid notifier")
}

func TestMissingEmailNotifier(t *testing.T) {
	i, err := inspeqtor.SetupNotification("email", map[string]string{
		"username": "mike",
		"password": "fuzzbucket",
		"to":       "mike@example.org",
	})
	assert(t, err != nil, "Missing data should cause error")
	log.Printf("%v", i)
	assert(t, i == nil, "Missing data should not return notifier")
}
