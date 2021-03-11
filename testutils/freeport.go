package testutils

import (
	"testing"

	"github.com/phayes/freeport"
)

// return a freeport, throw error if not available
func GetFreePort(t *testing.T) int {
	port, err := freeport.GetFreePort()
	if err != nil {
		t.Error(err)
	}
	return port
}

// retunr a freeport, throw error into test object if not available
func GetFreePorts(count int, t *testing.T) []int {
	ports, err := freeport.GetFreePorts(count)
	if err != nil {
		t.Error(err)
	}
	return ports
}

// get a port and start an http server on it
// so its taken.
func GetUnfreePort(t *testing.T) (port int) {
	port, err := GetUnFreePort()
	if err != nil {
		t.Error(err)
	}
	return port
}

// get a list of ports tthat are taken
func GetUnfreePorts(count int, t *testing.T) (unfreePorts []int) {
	for i := 0; i < count; i++ {
		unfreePorts = append(unfreePorts, GetUnfreePort(t))
	}
	return unfreePorts
}
