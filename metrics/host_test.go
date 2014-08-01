package metrics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollectHostMetrics(t *testing.T) {
	store := NewStore()
	err := CollectHostMetrics(store, "proc")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, store.Get("cpu", ""), 5662844)
	assert.Equal(t, store.Get("cpu", "user"), 4670673)
	assert.Equal(t, store.Get("cpu", "system"), 768153)
	assert.Equal(t, store.Get("cpu", "iowait"), 143718)
	assert.Equal(t, store.Get("cpu", "steal"), 68601)
	assert.Equal(t, store.Get("load", "1"), 2)
	assert.Equal(t, store.Get("load", "5"), 3)
	assert.Equal(t, store.Get("load", "15"), 5)
	assert.Equal(t, store.Get("swap", ""), 2)
	//expected := SystemMetrics{
	//when,
	//CpuMetrics{
	//1304544815, 4670673, 0, 768153, 1298881971, 143718, 844, 10855, 68601, 0, 0,
	//},
	//2,
	//3,
	//5,
	//2,
	//nil,
	//}
	//if *metrics != expected {
	//t.Errorf("Expected %+v, got %+v", expected, metrics)
	//}
}

func TestCollectRealHostMetrics(t *testing.T) {
	store := NewStore()
	err := CollectHostMetrics(store, "/proc")
	if err != nil {
		t.Fatal(err)
	}
	// Can't really know what we'll collect so we'll check for non-zero.
	assert.True(t, store.Get("load", "1") > 0)
	assert.True(t, store.Get("load", "5") > 0)
	assert.True(t, store.Get("load", "15") > 0)
	assert.True(t, store.Get("swap", "") > 0)
}

func TestCollectDiskMetrics(t *testing.T) {
	store := NewStore()
	err := collectDisk("fixtures/df.linux.txt", store)
	if err != nil {
		t.Error(err)
	}
	if store.Get("disk", "/") != 17 {
		t.Error("Unexpected results: %v", store.Get("disk", "/"))
	}
	if store.Get("disk", "/old") != 30 {
		t.Error("Unexpected results: %v", store.Get("disk", "/old"))
	}

	store = NewStore()
	err = collectDisk("fixtures/df.darwin.txt", store)
	if err != nil {
		t.Error(err)
	}
	if store.Get("disk", "/") != 7 {
		t.Error("Unexpected results: %v", store.Get("disk", "/"))
	}

	store = NewStore()
	err = collectDisk("", store)
	if err != nil {
		t.Error(err)
	}
	if store.Get("disk", "/") <= 0 {
		t.Error("Expected root disk to have more than 0% usage")
	}
}
