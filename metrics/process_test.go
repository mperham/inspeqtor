package metrics

import (
	"fmt"
	"testing"
)

func TestProcessCapture(t *testing.T) {
	m, err := CaptureProcess("proc", 9051)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%+v", m)
}
