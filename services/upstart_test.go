package services

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDetectUpstart(t *testing.T) {
	t.Parallel()
	init, err := detectUpstart("etc/init")
	assert.Nil(t, err)

	upstart := init.(Upstart)
	upstart.dummyOutput = "mysql start/running, process 14190"
	pid, st, err := upstart.LookupService("mysql")
	assert.Nil(t, err)
	assert.Equal(t, 14190, pid)
	assert.Equal(t, Up, st)

	// conf exists, but job is invalid
	upstart.dummyOutput = "initctl: Unknown job: foo"
	pid, st, err = upstart.LookupService("foo")
	assert.Nil(t, err)
	assert.Equal(t, -1, pid)
	assert.Equal(t, Unknown, st)

	// bad service name
	pid, st, err = upstart.LookupService("nonexistent")
	assert.Nil(t, err)
	assert.Equal(t, -1, pid)
	assert.Equal(t, Unknown, st)

	// running as non-root
	upstart.dummyOutput = "initctl: Unable to connect to system bus: Failed to connect to socket /var/run/dbus/system_bus_socket: No such file or directory"
	pid, st, err = upstart.LookupService("foo")
	assert.NotNil(t, err)
	assert.Equal(t, 0, pid)
	assert.Equal(t, Unknown, st)

	// garbage
	upstart.dummyOutput = "what the deuce?"
	pid, st, err = upstart.LookupService("mysql")
	assert.NotNil(t, err)
	assert.Equal(t, 0, pid)
	assert.Equal(t, Unknown, st)
	if !strings.Contains(err.Error(), "Unknown upstart output") {
		t.Error("Unexpected error: " + err.Error())
	}
}
