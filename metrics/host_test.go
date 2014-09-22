package metrics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollectHost(t *testing.T) {
	t.Parallel()
	store := NewHostStore("proc", 15)
	err := store.Collect(0)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, store.Get("cpu", ""), 0)
	assert.Equal(t, store.Get("cpu", "user"), 0)
	assert.Equal(t, store.Get("cpu", "system"), 0)
	assert.Equal(t, store.Get("cpu", "iowait"), 0)
	assert.Equal(t, store.Get("cpu", "steal"), 0)
	assert.Equal(t, store.Get("load", "1"), 0.02)
	assert.Equal(t, store.Get("load", "5"), 0.03)
	assert.Equal(t, store.Get("load", "15"), 0.05)
	assert.Equal(t, store.Get("swap", ""), 2)

	assert.Equal(t, []string{"cpu", "disk", "load", "swap"}, store.Families())
	assert.Equal(t, []string{"", "iowait", "steal", "system", "user"}, store.Metrics(store.Families()[0]))

	store.(*hostStorage).path = "proc2"
	err = store.Collect(0)
	if err != nil {
		t.Fatal(err)
	}
	assert.InDelta(t, store.Get("cpu", ""), 147, 1)
	assert.InDelta(t, store.Get("cpu", "user"), 66, 1)
	assert.InDelta(t, store.Get("cpu", "system"), 40, 1)
	assert.InDelta(t, store.Get("cpu", "iowait"), 20, 1)
	assert.InDelta(t, store.Get("cpu", "steal"), 20, 1)
	assert.Equal(t, store.Display("cpu", "user"), "66.7%")
	assert.Equal(t, store.Display("cpu", "system"), "40.1%")
	assert.Equal(t, store.Display("cpu", "iowait"), "20.2%")
	assert.Equal(t, store.Display("cpu", "steal"), "20.3%")
	assert.InDelta(t, store.Get("load", "1"), 0.02, 0.01)
	assert.InDelta(t, store.Get("load", "5"), 0.03, 0.01)
	assert.InDelta(t, store.Get("load", "15"), 0.05, 0.01)
	assert.Equal(t, store.Display("load", "1"), "0.02")
	assert.Equal(t, store.Display("load", "5"), "0.03")
	assert.Equal(t, store.Display("load", "15"), "0.05")

	assert.Equal(t, store.Get("swap", ""), 2)
	assert.Equal(t, store.Display("swap", ""), "2.0%")
}

func TestCollectRealHostMetrics(t *testing.T) {
	t.Parallel()
	store := NewHostStore("/proc", 15)
	err := store.Collect(0)
	if err != nil {
		t.Fatal(err)
	}
	// Can't really know what we'll collect so we'll check for non-zero.
	assert.True(t, store.Get("load", "1") > 0)
	assert.True(t, store.Get("load", "5") > 0)
	assert.True(t, store.Get("load", "15") > 0)
	assert.True(t, store.Get("swap", "") > 0, "Will fail if your swap usage is 0%")
}

func TestCollectDiskMetrics(t *testing.T) {
	t.Parallel()
	store := NewHostStore("", 15).(*hostStorage)
	err := store.collectDisk("fixtures/df.linux.txt")
	if err != nil {
		t.Error(err)
	}
	if store.Get("disk", "/") != 17 {
		t.Error("Unexpected results: %v", store.Get("disk", "/"))
	}
	if store.Get("disk", "/old") != 30 {
		t.Error("Unexpected results: %v", store.Get("disk", "/old"))
	}

	store = NewHostStore("", 15).(*hostStorage)
	err = store.collectDisk("fixtures/df.darwin.txt")
	if err != nil {
		t.Error(err)
	}
	if store.Get("disk", "/") != 7 {
		t.Error("Unexpected results: %v", store.Get("disk", "/"))
	}

	store = NewHostStore("", 15).(*hostStorage)
	err = store.collectDisk("")
	if err != nil {
		t.Error(err)
	}
	if store.Get("disk", "/") <= 0 {
		t.Error("Expected root disk to have more than 0% usage")
	}
}
