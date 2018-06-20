package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectRunit(t *testing.T) {
	t.Parallel()
	runit, err := detectRunit("./")
	assert.Nil(t, err)
	assert.NotNil(t, runit)

	st, err := runit.LookupService("memcached")
	assert.NotNil(t, st)
	assert.Nil(t, err)
	assert.Equal(t, 1234, st.Pid)
	assert.Equal(t, Up, st.Status.String())

	// bad service name
	st, err = runit.LookupService("nonexistent")
	assert.Nil(t, st)
	assert.NotNil(t, err)

	output := "ok: run: memcached: (pid 28125) 1s"
	runit.(*Runit).dummyOutput = &output
	err = runit.Restart("memcached")
	assert.Nil(t, err)

	output = "ok: run: memcached: (pid 28125) 1s"
	runit.(*Runit).dummyOutput = &output
	err = runit.Reload("memcached")
	assert.Nil(t, err)
}
