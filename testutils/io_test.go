package testutils

import (
	"github.com/phayes/freeport"
	"testing"
)

func TestPortIsAvailable(t *testing.T) {
	freeTest, err := freeport.GetFreePort()
	if err != nil {
		t.Error(err)
	}
	if !PortIsAvailable(freeTest) {
		t.Errorf("port %d is available according to freeport, but not iohelper", freeTest)
	}
	port, err := GetUnFreePort()
	if err != nil {
		t.Error(err)
	}
	if PortIsAvailable(port) {
		t.Errorf("port %d is available, should not be", port)
	}
}
