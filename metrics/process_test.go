package metrics

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestProcessCapture(t *testing.T) {
	store := NewStore()
	err := CaptureProcess(store, "proc", 100)
	if err == nil {
		t.Error("Expected process 100 to not exist")
	}

	store = NewStore()
	err = CaptureProcess(store, "proc", 9051)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1024*1024, store.Get("memory", "rss"))
	assert.Equal(t, 316964*1024, store.Get("memory", "vsz"))
	assert.Equal(t, 1, store.Get("cpu", "user"))
	assert.Equal(t, 0, store.Get("cpu", "system"))
	assert.Equal(t, 0, store.Get("cpu", "total_user"))
	assert.Equal(t, 0, store.Get("cpu", "total_system"))

	store = NewStore()
	err = CaptureProcess(store, "proc", 14190)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 324072*1024, store.Get("memory", "rss"))
	assert.Equal(t, 1481648*1024, store.Get("memory", "vsz"))
	assert.Equal(t, 524283, store.Get("cpu", "user"))
	assert.Equal(t, 270503, store.Get("cpu", "system"))
	assert.Equal(t, 0, store.Get("cpu", "total_user"))
	assert.Equal(t, 0, store.Get("cpu", "total_system"))

	store = NewStore()
	err = CaptureProcess(store, "proc", 3589)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 19728*1024, store.Get("memory", "rss"))
	assert.Equal(t, 287976*1024, store.Get("memory", "vsz"))
	assert.Equal(t, 258, store.Get("cpu", "user"))
	assert.Equal(t, 28954, store.Get("cpu", "system"))
	assert.Equal(t, 2135754, store.Get("cpu", "total_user"))
	assert.Equal(t, 259400, store.Get("cpu", "total_system"))

	for i := 0; i < 100000000; i++ {
		// eat up some CPU time so we get a non-zero value for user CPU
		// TODO mine bitcoins here, send them to mike AT contribsys.com
	}

	store = NewStore()
	err = CaptureProcess(store, "/etc/proc", os.Getpid())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, true, store.Get("memory", "rss") > 0)
	assert.Equal(t, true, store.Get("memory", "vsz") > 0)
	assert.Equal(t, true, store.Get("cpu", "user") > 0)

	store = NewStore()
	err = CaptureProcess(store, "/etc/proc", -1)
	assert.Error(t, err)
}
