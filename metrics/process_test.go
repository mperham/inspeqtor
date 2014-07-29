package metrics

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestProcessCapture(t *testing.T) {
	m, err := CaptureProcess("proc", 100)
	if err == nil {
		t.Error("Expected process 100 to not exist")
	}

	m, err = CaptureProcess("proc", 9051)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1024*1024, m.VmRSS)
	assert.Equal(t, 316964*1024, m.VmSize)
	assert.Equal(t, 1, m.UserCpu)
	assert.Equal(t, 0, m.SystemCpu)
	assert.Equal(t, 0, m.UserChildCpu)
	assert.Equal(t, 0, m.SystemChildCpu)

	m, err = CaptureProcess("proc", 14190)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 324072*1024, m.VmRSS)
	assert.Equal(t, 1481648*1024, m.VmSize)
	assert.Equal(t, 524283, m.UserCpu)
	assert.Equal(t, 270503, m.SystemCpu)
	assert.Equal(t, 0, m.UserChildCpu)
	assert.Equal(t, 0, m.SystemChildCpu)

	m, err = CaptureProcess("proc", 3589)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 19728*1024, m.VmRSS)
	assert.Equal(t, 287976*1024, m.VmSize)
	assert.Equal(t, 258, m.UserCpu)
	assert.Equal(t, 28954, m.SystemCpu)
	assert.Equal(t, 2135754, m.UserChildCpu)
	assert.Equal(t, 259400, m.SystemChildCpu)

	for i := 0; i < 100000000; i++ {
		// eat up some CPU time so we get a non-zero value for user CPU
		// TODO mine bitcoins here, send them to mike AT contribsys.com
	}

	m, err = CaptureProcess("/etc/proc", os.Getpid())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, true, m.VmRSS > 0)
	assert.Equal(t, true, m.VmSize > 0)
	assert.Equal(t, true, m.UserCpu > 0)

	m, err = CaptureProcess("/etc/proc", -1)
	assert.Nil(t, m)
	assert.Error(t, err)
}
