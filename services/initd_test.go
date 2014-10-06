package services

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestInitd(t *testing.T) {
	t.Parallel()

	l, err := detectInitd("./")
	assert.Nil(t, err)
	assert.NotNil(t, l)

	assert.Equal(t, "init.d", l.Name())

	// service does not exist
	st, err := l.LookupService("apache2")
	assert.Nil(t, err)
	assert.Nil(t, st)

	// service exists but pidfile doesn't
	st, err = l.LookupService("mysql")
	assert.Nil(t, err)
	assert.NotNil(t, st)
	assert.Equal(t, 0, st.Pid)
	assert.Equal(t, Down, st.Status)

	// Need to be able to kill -0 the PID and our own process
	// is the only one we can be sure of.
	i := l.(*Initd)
	i.pidParser = func(_ []byte) (int, error) { return os.Getpid(), nil }

	// service exists and pidfile exists
	st, err = l.LookupService("redis")
	assert.Nil(t, err)
	assert.Equal(t, os.Getpid(), st.Pid)
	assert.Equal(t, Up, st.Status)
}
