package services

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectUpstart(t *testing.T) {
	t.Parallel()
	init, err := detectUpstart("etc/init")
	assert.Nil(t, err)

	var output string
	output = "mysql start/running, process 14190"
	upstart := init.(*Upstart)
	upstart.dummyOutput = &output
	st, err := init.LookupService("mysql")
	assert.Nil(t, err)
	assert.Equal(t, 14190, st.Pid)
	assert.Equal(t, Up, st.Status)
	assert.Equal(t, st.String(), "Up/14190")

	// conf exists, but job is invalid
	output = "initctl: Unknown job: foo"
	upstart.dummyOutput = &output
	st, err = init.LookupService("foo")
	assert.Nil(t, st)
	assert.NotNil(t, err)

	// bad service name
	st, err = init.LookupService("nonexistent")
	assert.Nil(t, st)
	assert.NotNil(t, err)

	// running as non-root
	output = "initctl: Unable to connect to system bus: Failed to connect to socket /var/run/dbus/system_bus_socket: No such file or directory"
	upstart.dummyOutput = &output
	st, err = init.LookupService("foo")
	assert.NotNil(t, err)
	assert.Nil(t, st)

	// garbage
	output = "what the deuce?"
	upstart.dummyOutput = &output
	st, err = init.LookupService("mysql")
	assert.Nil(t, st)
	assert.NotNil(t, err)
	if !strings.Contains(err.Error(), "Unknown upstart output") {
		t.Error("Unexpected error: " + err.Error())
	}

	output = "rsyslog start/running, process 28192"
	upstart.dummyOutput = &output
	err = init.Restart("rsyslog")
	assert.Nil(t, err)

	output = ""
	upstart.dummyOutput = &output
	err = init.Reload("rsyslog")
	assert.Nil(t, err)
}
