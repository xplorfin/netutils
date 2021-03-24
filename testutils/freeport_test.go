package testutils_test

import (
	"net"
	"strconv"
	"testing"

	"github.com/xplorfin/netutils/testutils"
)

func TestGetFreePort(t *testing.T) {
	port := testutils.GetFreePort(t)
	if port == 0 {
		t.Error("port == 0")
	}

	// Try to listen on the port
	l, err := net.Listen("tcp", "localhost"+":"+strconv.Itoa(port))
	if err != nil {
		t.Error(err)
	}
	err = l.Close()
	if err != nil {
		t.Error(err)
	}
}

func TestGetFreePorts(t *testing.T) {
	count := 3
	ports := testutils.GetFreePorts(count, t)
	if len(ports) == 0 {
		t.Error("len(ports) == 0")
	}
	for _, port := range ports {
		if port == 0 {
			t.Error("port == 0")
		}

		// Try to listen on the port
		l, err := net.Listen("tcp", "localhost"+":"+strconv.Itoa(port))
		if err != nil {
			t.Error(err)
		}
		_ = l.Close()
	}
}

func TestGetUnFreePort(t *testing.T) {
	port := testutils.GetUnfreePort(t)
	if testutils.PortIsAvailable(port) {
		t.Errorf("expected port %d to be unavailable", port)
	}
}

func TestGetUnfreePorts(t *testing.T) {
	ports := testutils.GetUnfreePorts(10, t)
	for _, port := range ports {
		if testutils.PortIsAvailable(port) {
			t.Errorf("expected port %d to be unavailable", port)
		}
	}
}
