package metrics

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCollectHostMetrics(t *testing.T) {
	metrics, err := CollectHostMetrics("proc")
	if err != nil {
		t.Fatal(err)
	}

	when := time.Now()
	metrics.When = when
	metrics.Disk = nil

	expected := SystemMetrics{
		when,
		CpuMetrics{
			1304544815, 4670673, 0, 768153, 1298881971, 143718, 844, 10855, 68601, 0, 0,
		},
		2,
		3,
		5,
		2,
		nil,
	}
	if *metrics != expected {
		t.Errorf("Expected %+v, got %+v", expected, metrics)
	}
}

func TestCollectRealHostMetrics(t *testing.T) {
	m, err := CollectHostMetrics("/proc")
	if err != nil {
		t.Fatal(err)
	}
	// Can't really know what we'll collect so we'll check for non-zero.
	assert.True(t, m.Load1 > 0)
	assert.True(t, m.Load5 > 0)
	assert.True(t, m.Load15 > 0)
	assert.True(t, m.PercentSwapInUse > 0)
}

func TestCollectDiskMetrics(t *testing.T) {
	sm := SystemMetrics{}
	err := collectDisk("fixtures/df.linux.txt", &sm)
	if err != nil {
		t.Error(err)
	}
	if (*sm.Disk)["/"] != 17 {
		t.Error("Unexpected results: %v", *sm.Disk)
	}
	if (*sm.Disk)["/old"] != 30 {
		t.Error("Unexpected results: %v", *sm.Disk)
	}

	err = collectDisk("fixtures/df.darwin.txt", &sm)
	if err != nil {
		t.Error(err)
	}
	if (*sm.Disk)["/"] != 7 {
		t.Error("Unexpected results: %v", *sm.Disk)
	}

	err = collectDisk("", &sm)
	if err != nil {
		t.Error(err)
	}
}
