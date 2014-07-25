package metrics

import (
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
		243376,
		0.02,
		0.03,
		0.05,
		2,
		nil,
	}
	if *metrics != expected {
		t.Errorf("Expected %+v, got %+v", expected, metrics)
	}
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
