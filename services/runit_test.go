package services

import (
	"inspeqtor/core"
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

	pid, status, err := runit.LookupService("memcached")
	if err != nil {
		t.Error(err)
	}
	if pid != 1234 {
		t.Errorf("Expected positive PID, got %d\n", pid)
	}
	if status != core.Up {
		t.Errorf("Service should be unknown, got %v\n", status)
	}

	// bad service name
	pid, status, err = runit.LookupService("nonexistent")
	if err != nil {
		t.Error(err)
	}
	if pid != -1 {
		t.Errorf("Expected not found PID, got %d\n", pid)
	}
	if status != core.Unknown {
		t.Errorf("Service should be unknown, got %v\n", status)
	}
}
