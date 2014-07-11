package services

import (
	"testing"
)

func TestDetectRunit(t *testing.T) {
	runit, err := DetectRunit("./")
	if err != nil {
		t.Fatal(err)
	}
	if runit == nil {
		t.Fatal("Runit not detected")
	}

	expected := Runit{"./etc/service"}
	if *runit != expected {
		t.Errorf("Expected %+v, got %+v", expected, runit)
	}

	pid, err := runit.FindServicePID("memcached")
	if err != nil {
		t.Error(err)
	}
	if pid != 1234 {
		t.Errorf("Expected positive PID, got %d\n", pid)
	}

	// bad service name
	pid, err = runit.FindServicePID("nonexistent")
	if err != nil {
		t.Error(err)
	}
	if pid != -1 {
		t.Errorf("Expected not found PID, got %d\n", pid)
	}
}
