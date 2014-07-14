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
	}
	if *metrics != expected {
		t.Errorf("Expected %+v, got %+v", expected, metrics)
	}
}
