package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectSystemd(t *testing.T) {
	t.Parallel()
	init, err := detectSystemd("etc/systemd")
	assert.Nil(t, err)

	upstart := init.(*Systemd)
	upstart.dummyOutput = "MainPID=12345"
	st, err := init.LookupService("mysql")
	assert.Nil(t, err)
	assert.Equal(t, 12345, st.Pid)
	assert.Equal(t, Up, st.Status.String())

	// bad name
	upstart.dummyOutput = "MainPID=0"
	upstart.dummyOutput2 = "Failed to issue method call: No such file or directory"
	st, err = init.LookupService("foo")
	assert.Nil(t, st)
	assert.NotNil(t, err)

	// process is down
	upstart.dummyOutput = "MainPID=0"
	upstart.dummyOutput2 = "enabled\n"
	st, err = init.LookupService("memcached")
	assert.NotNil(t, st)
	assert.Nil(t, err)
	assert.Equal(t, 0, st.Pid)
	assert.Equal(t, Down, st.Status.String())

	err = init.Restart("rsyslog")
	assert.Nil(t, err)

	err = init.Reload("rsyslog")
	assert.Nil(t, err)
}
