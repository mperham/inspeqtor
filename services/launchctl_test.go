package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLaunchctl(t *testing.T) {
	t.Parallel()
	l, err := detectLaunchctl("darwin/")
	assert.Nil(t, err)
	assert.NotNil(t, l)

	// Verify we can find a known good service.
	// Should be running on all OSX machines, right?
	pid, status, err := l.LookupService("com.apple.Finder")
	assert.Nil(t, err)
	assert.True(t, pid > 0, "Expected positive value for PID")
	assert.Equal(t, Up, status)

	pid, status, err = l.LookupService("some.fake.service")
	assert.Nil(t, err)
	assert.Equal(t, Unknown, status)
	assert.Equal(t, -1, pid)

	err = l.Restart("some.jacked.up.name")
	assert.NotNil(t, err)

	pid1, status, err := l.LookupService("homebrew.mxcl.memcached")
	if pid > 0 && status == Up {
		err = l.Restart("homebrew.mxcl.memcached")
		assert.Nil(t, err)
		pid2, status, err := l.LookupService("homebrew.mxcl.memcached")
		assert.Nil(t, err)
		assert.Equal(t, Up, status)
		assert.NotEqual(t, pid1, pid2)
	}
}
