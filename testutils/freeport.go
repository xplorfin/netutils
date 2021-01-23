package testutils

import (
	"github.com/phayes/freeport"
	"testing"
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
func GetUnfreePort(t *testing.T) (port int) {
	port, err := GetUnFreePort()
	if err != nil {
		t.Error(err)
	}
	return port
}
