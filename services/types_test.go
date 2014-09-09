package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	assert.Equal(t, NewStatus().String(), "Unknown/0")
}

func TestDetect(t *testing.T) {
	inits := Detect()
	assert.True(t, len(inits) > 0)
}

func TestMock(t *testing.T) {
	mock := MockInit()
	assert.Equal(t, mock.Name(), "mock")
	assert.Nil(t, mock.Restart("foo"))
	st, err := mock.LookupService("bar")
	assert.Nil(t, err)
	assert.NotNil(t, st)
	assert.Equal(t, st.Pid, 123)
	assert.Equal(t, st.Status, Up)
}
