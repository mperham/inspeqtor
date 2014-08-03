package metrics

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNonexistentProcessCapture(t *testing.T) {
	store := NewProcessStore()
	err := CaptureProcess(store, "proc", 100)
	if err == nil {
		t.Error("Expected process 100 to not exist")
	}
}

// doesn't have real CPU numbers
func TestBasicProcess(t *testing.T) {
	store := NewProcessStore()
	err := CaptureProcess(store, "proc", 9051)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1024*1024, store.Get("memory", "rss"))
	assert.Equal(t, 316964*1024, store.Get("memory", "vsz"))
	assert.Equal(t, 0, store.Get("cpu", "user"))
	assert.Equal(t, 0, store.Get("cpu", "system"))
	assert.Equal(t, 0, store.Get("cpu", "total_user"))
	assert.Equal(t, 0, store.Get("cpu", "total_system"))

	err = CaptureProcess(store, "proc2", 9051)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1024*1024, store.Get("memory", "rss"))
	assert.Equal(t, 316964*1024, store.Get("memory", "vsz"))
	assert.Equal(t, 4, store.Get("cpu", "user"))
	assert.Equal(t, 0, store.Get("cpu", "system"))
	assert.Equal(t, 0, store.Get("cpu", "total_user"))
	assert.Equal(t, 0, store.Get("cpu", "total_system"))
}

// has real stats, no children
func TestMysqlProcess(t *testing.T) {
	store := NewProcessStore()
	err := CaptureProcess(store, "proc", 14190)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 324072*1024, store.Get("memory", "rss"))
	assert.Equal(t, 1481648*1024, store.Get("memory", "vsz"))
	assert.Equal(t, 0, store.Get("cpu", "user"))
	assert.Equal(t, 0, store.Get("cpu", "system"))
	assert.Equal(t, 0, store.Get("cpu", "total_user"))
	assert.Equal(t, 0, store.Get("cpu", "total_system"))

	err = CaptureProcess(store, "proc2", 14190)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 324072*1024, store.Get("memory", "rss"))
	assert.Equal(t, 1481648*1024, store.Get("memory", "vsz"))
	assert.Equal(t, 17, store.Get("cpu", "user"))
	assert.Equal(t, 97, store.Get("cpu", "system"))
	assert.Equal(t, 0, store.Get("cpu", "total_user"))
	assert.Equal(t, 0, store.Get("cpu", "total_system"))
}

// has real stats, child processes
func TestApacheProcess(t *testing.T) {
	store := NewProcessStore()
	err := CaptureProcess(store, "proc", 3589)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 19728*1024, store.Get("memory", "rss"))
	assert.Equal(t, 287976*1024, store.Get("memory", "vsz"))
	assert.Equal(t, 0, store.Get("cpu", "user"))
	assert.Equal(t, 0, store.Get("cpu", "system"))
	assert.Equal(t, 0, store.Get("cpu", "total_user"))
	assert.Equal(t, 0, store.Get("cpu", "total_system"))

	err = CaptureProcess(store, "proc2", 3589)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 19728*1024, store.Get("memory", "rss"))
	assert.Equal(t, 287976*1024, store.Get("memory", "vsz"))
	assert.Equal(t, 2, store.Get("cpu", "user"))
	assert.Equal(t, 6, store.Get("cpu", "system"))
	assert.Equal(t, 46, store.Get("cpu", "total_user"))
	assert.Equal(t, 100, store.Get("cpu", "total_system"))
}

// verify our own process stats
func TestRealProcess(t *testing.T) {
	store := NewProcessStore()
	err := CaptureProcess(store, "/etc/proc", os.Getpid())
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 100000000; i++ {
		// eat up some CPU time so we get a non-zero value for user CPU
		// TODO mine bitcoins here, send them to mike AT contribsys.com
	}

	err = CaptureProcess(store, "/etc/proc", os.Getpid())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, true, store.Get("memory", "rss") > 0)
	assert.Equal(t, true, store.Get("memory", "vsz") > 0)
	assert.Equal(t, true, store.Get("cpu", "user") > 0)
}

// verify we can't capture a non-existent process for real
func TestNonexistentProcess(t *testing.T) {
	store := NewProcessStore()
	err := CaptureProcess(store, "/etc/proc", -1)
	assert.Error(t, err)
}
