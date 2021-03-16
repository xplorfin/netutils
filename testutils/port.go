package testutils

import (
	"fmt"
	"net"
	"strconv"

	"github.com/phayes/freeport"
)

// GetUnFreePort gets a port and start an http server on it to mock a taken port
func GetUnFreePort() (port int, err error) {
	port, err = freeport.GetFreePort()
	if err != nil {
		return 0, err
	}
	go func() {
		err = MockHTTPServer(port)
	}()
	host := fmt.Sprintf("localhost:%d", port)
	connected := WaitForConnect(host)
	if !connected {
		return port, fmt.Errorf("could not connect to host %s", host)
	}
	return port, err
}

// PortIsAvailable will determine if a port is available
func PortIsAvailable(port int) bool {
	// Concatenate a colon and the port
	host := ":" + strconv.Itoa(port)

	// Try to create a server with the port
	server, err := net.Listen("tcp", host)

	// if it fails then the port is likely taken
	if err != nil {
		return false
	}

	// close the server
	_ = server.Close()

	// we successfully used and closed the port
	// so it's now available to be used again
	return true
}
