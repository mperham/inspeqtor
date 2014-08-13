package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDetectRunit(t *testing.T) {
	t.Parallel()
	runit, err := detectRunit("./")
	assert.Nil(t, err)
	assert.NotNil(t, runit)

	pid, status, err := runit.LookupService("memcached")
	assert.Nil(t, err)
	assert.Equal(t, 1234, pid)
	assert.Equal(t, Up, status)

	// bad service name
	pid, status, err = runit.LookupService("nonexistent")
	assert.Nil(t, err)
	assert.Equal(t, Unknown, status)
	assert.Equal(t, -1, pid)

	runit.(*Runit).dummyOutput = "ok: run: memcached: (pid 28125) 1s"
	err = runit.Restart("memcached")
	assert.Nil(t, err)
}
