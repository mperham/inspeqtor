package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLaunchd(t *testing.T) {
	t.Parallel()
	l, err := detectLaunchd("darwin/")
	assert.Nil(t, err)
	assert.NotNil(t, l)

	// Verify we can find a known good service.
	// Should be running on all OSX machines, right?
	st, err := l.LookupService("com.apple.Finder")
	assert.NotNil(t, st)
	assert.Nil(t, err)
	assert.True(t, st.Pid > 0, "Expected positive value for PID")
	assert.Equal(t, Up, st.Status)

	st, err = l.LookupService("some.fake.service")
	assert.Nil(t, st)
	assert.NotNil(t, err)

	err = l.Restart("some.jacked.up.name")
	assert.NotNil(t, err)

	st1, err := l.LookupService("homebrew.mxcl.memcached")
	if st1.Pid > 0 && st1.Status == Up {
		err = l.Restart("homebrew.mxcl.memcached")
		assert.Nil(t, err)
		st2, err := l.LookupService("homebrew.mxcl.memcached")
		assert.Nil(t, err)
		assert.Equal(t, Up, st2.Status)
		assert.NotEqual(t, st1.Pid, st2.Pid)
	}
}
