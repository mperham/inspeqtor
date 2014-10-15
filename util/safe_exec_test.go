package util

import (
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSafeRun(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("echo", "hello", "world")
	sout, err := SafeRun(cmd)
	assert.Nil(t, err)
	assert.Equal(t, sout, []byte("hello world\n"))
}

func TestInvalidCommand(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("lasdjfaslkdf")
	sout, err := SafeRun(cmd)
	assert.NotNil(t, err)
	assert.Nil(t, sout)
}

func TestTimeout(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("sleep", "1")
	sout, err := SafeRun(cmd, time.Duration(10)*time.Millisecond)
	assert.NotNil(t, err)
	assert.Nil(t, sout)
}
