package testutils

import (
	"net"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	testutils "github.com/xplorfin/netutils"
)

func TestGetFreePort(t *testing.T) {
	port := GetFreePort(t)
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
	ports := GetFreePorts(count, t)
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
	port := GetUnfreePort(t)
	if PortIsAvailable(port) {
		t.Errorf("expected port %d to be unavailable", port)
	}
}

func TestGetUnfreePorts(t *testing.T) {
	ports := GetUnfreePorts(10, t)
	for _, port := range ports {
		if PortIsAvailable(port) {
			t.Errorf("expected port %d to be unavailable", port)
		}
	}
}

func TestGetFreeportStack(t *testing.T) {
	stack := testutils.NewFreeportStack()
	for i := 0; i < 10; i++ {
		port := stack.GetPort()
		assert.True(t, PortIsAvailable(port))
	}
}