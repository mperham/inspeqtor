package services

import (
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
	pid, err := upstart.FindServicePID("foo")
	if err != nil {
		t.Error(err)
	}
	if pid <= 0 {
		t.Errorf("Expected positive PID, got %d\n", pid)
	}

	// bad service name
	pid, err = upstart.FindServicePID("nonexistent")
	if err != nil {
		t.Error(err)
	}
	if pid != -1 {
		t.Errorf("Expected not found PID, got %d\n", pid)
	}

	// running as non-root
	upstart.dummyOutput = "initctl: Unable to connect to system bus: Failed to connect to socket /var/run/dbus/system_bus_socket: No such file or directory"
	pid, err = upstart.FindServicePID("foo")
	if pid != 0 {
		t.Errorf("Expected zero PID, got %d\n", pid)
	}
	if err == nil {
		t.Error(err)
	}
}
