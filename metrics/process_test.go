package metrics

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNonexistentProcessCollect(t *testing.T) {
	t.Parallel()
	store := NewProcessStore("proc", 15)
	err := store.Collect(100)
	if err == nil {
		t.Error("Expected process 100 to not exist")
	}
}

// doesn't have real CPU numbers
func TestBasicProcess(t *testing.T) {
	t.Parallel()
	store := NewProcessStore("proc", 15)
	err := store.Collect(9051)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1024*1024, store.Get("memory", "rss"))
	assert.Equal(t, 0, store.Get("cpu", "user"))
	assert.Equal(t, 0, store.Get("cpu", "system"))
	assert.Equal(t, 0, store.Get("cpu", "total_user"))
	assert.Equal(t, 0, store.Get("cpu", "total_system"))

	store.(*processStorage).path = "proc2"
	err = store.Collect(9051)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1024*1024, store.Get("memory", "rss"))
	// 500 ticks, 1500 cycle ticks = 33% CPU usage
	assert.Equal(t, 33, store.Get("cpu", "user"))
	assert.Equal(t, 0, store.Get("cpu", "system"))
	assert.Equal(t, 0, store.Get("cpu", "total_user"))
	assert.Equal(t, 0, store.Get("cpu", "total_system"))
}

// has real stats, no children
func TestMysqlProcess(t *testing.T) {
	t.Parallel()
	store := NewProcessStore("proc", 15)
	err := store.Collect(14190)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 324072*1024, store.Get("memory", "rss"))
	assert.Equal(t, 0, store.Get("cpu", "user"))
	assert.Equal(t, 0, store.Get("cpu", "system"))
	assert.Equal(t, 0, store.Get("cpu", "total_user"))
	assert.Equal(t, 0, store.Get("cpu", "total_system"))

	store.(*processStorage).path = "proc2"
	err = store.Collect(14190)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 324072*1024, store.Get("memory", "rss"))
	assert.Equal(t, 1, store.Get("cpu", "user"))
	assert.Equal(t, 6, store.Get("cpu", "system"))
	assert.Equal(t, 0, store.Get("cpu", "total_user"))
	assert.Equal(t, 0, store.Get("cpu", "total_system"))
}

// has real stats, child processes
func TestApacheProcess(t *testing.T) {
	t.Parallel()
	store := NewProcessStore("proc", 15)
	err := store.Collect(3589)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 19728*1024, store.Get("memory", "rss"))
	assert.Equal(t, 0, store.Get("cpu", "user"))
	assert.Equal(t, 0, store.Get("cpu", "system"))
	assert.Equal(t, 0, store.Get("cpu", "total_user"))
	assert.Equal(t, 0, store.Get("cpu", "total_system"))

	store.(*processStorage).path = "proc2"
	err = store.Collect(3589)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 19728*1024, store.Get("memory", "rss"))
	assert.Equal(t, 0, store.Get("cpu", "user"))
	assert.Equal(t, 0, store.Get("cpu", "system"))
	assert.Equal(t, 3, store.Get("cpu", "total_user"))
	assert.Equal(t, 6, store.Get("cpu", "total_system"))
}

// verify our own process stats
func TestRealProcess(t *testing.T) {
	t.Parallel()
	store := NewProcessStore("/etc/proc", 15)
	err := store.Collect(os.Getpid())
	if err != nil {
		t.Fatal(err)
	}

	err = store.Collect(os.Getpid())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, true, store.Get("memory", "rss") > 0)
}

// verify we can't capture a non-existent process for real
func TestNonexistentProcess(t *testing.T) {
	t.Parallel()
	store := NewProcessStore("/etc/proc", 15)
	err := store.Collect(-1)
	assert.Error(t, err)
}
