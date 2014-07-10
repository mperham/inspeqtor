package init

import (
	"testing"
)

func TestLaunchctl(t *testing.T) {
	l, err := DetectLaunchctl("darwin/")
	if err != nil {
		t.Fatal(err)
	}

	expected := Launchctl{}
	if *l != expected {
		t.Fatalf("Expected %+v, got %+v", expected, *l)
	}

	// Verify we can find a known good service.
	// Should be running on all OSX machines, right?
	pid, err := l.FindServicePID("com.apple.Finder")
	if err != nil {
		t.Error(err)
	}
	if pid <= 0 {
		t.Errorf("Expected positive value for PID, got %d\n", pid)
	}

	pid, err = l.FindServicePID("some.fake.service")
	if err != nil {
		t.Error(err)
	}
	if pid != -1 {
		t.Errorf("Expected not found result PID, got %d\n", pid)
	}
}
