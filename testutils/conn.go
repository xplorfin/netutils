package testutils

import (
	"testing"
	"time"

	netutils "github.com/xplorfin/netutils"
)

// WaitForConnectTimeout will wait for a connection
// on a port progressively backing off
// returns false if we couldn't establish a connection by timeout
// after 10 timeouts
// deprecated: use WaitForConnectTimeout in main package
func WaitForConnectTimeout(host string, timeout time.Duration) bool {
	return netutils.WaitForConnectTimeout(host, timeout)
}

// WaitForConnect on a port progressively backing off
// returns false if we couldn't establish a connection
// uses default timeout of 5 seconds
// deprecated: use WaitForConnect in main package
func WaitForConnect(host string) bool {
	return netutils.WaitForConnect(host)
}

// AssertConnected tries to connect to a given host several times
// fails if unable to connect
func AssertConnected(host string, t *testing.T) {
	connected := netutils.WaitForConnect(host)
	if !connected {
		t.Errorf("could not connect to host %s", host)
	}
}
