package linux

import "testing"

func TestSystemMetricsCollection(t *testing.T) {
	metrics, err := CollectSystemMetrics("proc")
	if err != nil {
		t.Fatal("stat read fail")
	}
	t.Logf("%+v", metrics)
}
