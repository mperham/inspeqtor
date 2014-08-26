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

	upstart := init.(*Upstart)
	upstart.dummyOutput = "mysql start/running, process 14190"
	st, err := init.LookupService("mysql")
	assert.Nil(t, err)
	assert.Equal(t, 14190, st.Pid)
	assert.Equal(t, Up, st.Status)

	// conf exists, but job is invalid
	upstart.dummyOutput = "initctl: Unknown job: foo"
	st, err = init.LookupService("foo")
	assert.Nil(t, st)
	assert.NotNil(t, err)

	// bad service name
	st, err = init.LookupService("nonexistent")
	assert.Nil(t, st)
	assert.NotNil(t, err)

	// running as non-root
	upstart.dummyOutput = "initctl: Unable to connect to system bus: Failed to connect to socket /var/run/dbus/system_bus_socket: No such file or directory"
	st, err = init.LookupService("foo")
	assert.NotNil(t, err)
	assert.Nil(t, st)

	// garbage
	upstart.dummyOutput = "what the deuce?"
	st, err = init.LookupService("mysql")
	assert.Nil(t, st)
	assert.NotNil(t, err)
	if !strings.Contains(err.Error(), "Unknown upstart output") {
		t.Error("Unexpected error: " + err.Error())
	}

	upstart.dummyOutput = "rsyslog start/running, process 28192"
	err = init.Restart("rsyslog")
	assert.Nil(t, err)
}
