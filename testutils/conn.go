package testutils

import (
	"github.com/jpillora/backoff"
	"net"
	"testing"
	"time"
)

// wait for a connect on a port progressively backing off
// returns false if we couldn't establish a connection by timeout
// after 10 timeouts
func WaitForConnectTimeout(host string, timeout time.Duration) bool {
	connected := false
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    timeout / 10,
		Max:    timeout,
	}
	for {
		b.Attempt()
		conn, err := net.Dial("tcp", host)
		if err != nil {
			d := b.Duration()
			if d == timeout && b.Attempt() > 10 {
				return false
			}
			time.Sleep(d)
			continue
		}
		connected = true
		b.Reset()
		conn.Close()
		break
	}
	return connected
}

// wait for a connect on a port progressively backing off
// returns false if we couldn't establish a connection
// uses default timeout of 5 seconds
func WaitForConnect(host string) bool {
	return WaitForConnectTimeout(host, time.Second*5)
}

// try to connect to host, if it fails, the test fails
func AssertConnected(host string, t *testing.T) {
	connected := WaitForConnect(host)
	if !connected {
		t.Errorf("could not connect to host %s", host)
	}
}
