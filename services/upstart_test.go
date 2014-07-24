package services

import (
	"inspeqtor/core"
	"strings"
	"testing"
)

func TestDetectUpstart(t *testing.T) {
	upstart, err := DetectUpstart("etc/init")
	if err != nil {
		t.Fatal(err)
	}

	expected := Upstart{"etc/init", ""}
	if *upstart != expected {
		t.Errorf("Expected %+v, got %+v", expected, upstart)
	}

	upstart.dummyOutput = "mysql start/running, process 14190"
	pid, st, err := upstart.LookupService("mysql")
	if err != nil {
		t.Error(err)
	}
	if pid <= 0 {
		t.Errorf("Expected positive PID, got %d\n", pid)
	}
	if st != core.Up {
		t.Errorf("Expected Up status, got %v\n", st)
	}

	// conf exists, but job is invalid
	upstart.dummyOutput = "initctl: Unknown job: foo"
	pid, st, err = upstart.LookupService("foo")
	if err != nil {
		t.Error(err)
	}
	if pid != -1 {
		t.Errorf("Expected not found PID, got %d\n", pid)
	}
	if st != core.Unknown {
		t.Errorf("Expected Unknown status, got %v\n", st)
	}

	// bad service name
	pid, st, err = upstart.LookupService("nonexistent")
	if err != nil {
		t.Error(err)
	}
	if pid != -1 {
		t.Errorf("Expected not found PID, got %d\n", pid)
	}
	if st != core.Unknown {
		t.Errorf("Expected Unknown status, got %v\n", st)
	}

	// running as non-root
	upstart.dummyOutput = "initctl: Unable to connect to system bus: Failed to connect to socket /var/run/dbus/system_bus_socket: No such file or directory"
	pid, st, err = upstart.LookupService("foo")
	if err == nil {
		t.Error("Expected an error")
	}
	if pid != 0 {
		t.Errorf("Expected zero PID, got %d\n", pid)
	}
	if st != core.Unknown {
		t.Errorf("Expected Unknown status, got %v\n", st)
	}

	// garbage
	upstart.dummyOutput = "what the deuce?"
	pid, st, err = upstart.LookupService("mysql")
	if err == nil {
		t.Error("Expected an error")
	}
	if !strings.Contains(err.Error(), "Unknown upstart output") {
		t.Error("Unexpected error: " + err.Error())
	}
	if pid != 0 {
		t.Errorf("Expected zero PID, got %d\n", pid)
	}
	if st != core.Unknown {
		t.Errorf("Expected Unknown status, got %v\n", st)
	}
}
