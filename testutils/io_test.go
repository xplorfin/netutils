package testutils_test

import (
	"testing"

	"github.com/xplorfin/netutils/testutils"

	"github.com/phayes/freeport"
)

func TestPortIsAvailable(t *testing.T) {
	freeTest, err := freeport.GetFreePort()
	if err != nil {
		t.Error(err)
	}
	if !testutils.PortIsAvailable(freeTest) {
		t.Errorf("port %d is available according to freeport, but not iohelper", freeTest)
	}
	port, err := testutils.GetUnFreePort()
	if err != nil {
		t.Error(err)
	}
	if testutils.PortIsAvailable(port) {
		t.Errorf("port %d is available, should not be", port)
	}
}
