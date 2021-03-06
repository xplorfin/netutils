package testutils

import (
	"testing"

	"github.com/phayes/freeport"
)

// GetFreePort returns a freeport, throw error if not available
func GetFreePort(t *testing.T) int {
	port, err := freeport.GetFreePort()
	if err != nil {
		t.Error(err)
	}
	return port
}

// GetFreePorts returns a freeport, throw error into test object if not available
func GetFreePorts(count int, t *testing.T) []int {
	ports, err := freeport.GetFreePorts(count)
	if err != nil {
		t.Error(err)
	}
	return ports
}

// GetUnfreePort gets a free port and start an http server on it
// so its taken.
func GetUnfreePort(t *testing.T) (port int) {
	port, err := GetUnFreePort()
	if err != nil {
		t.Error(err)
	}
	return port
}

// GetUnfreePorts gets a list of ports that are taken
func GetUnfreePorts(count int, t *testing.T) (unfreePorts []int) {
	for i := 0; i < count; i++ {
		unfreePorts = append(unfreePorts, GetUnfreePort(t))
	}
	return unfreePorts
}
