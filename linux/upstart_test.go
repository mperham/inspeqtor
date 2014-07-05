package linux

import (
	"testing"
)

func TestDetectUpstart(t *testing.T) {
	upstart, err := DetectUpstart("etc/init")
	if err != nil {
		t.Fatal(err)
	}

	expected := Upstart{"etc/init"}
	if *upstart != expected {
		t.Errorf("Expected %+v, got %+v", expected, upstart)
	}
}
