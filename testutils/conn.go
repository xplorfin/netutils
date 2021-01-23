package testutils

import (
	"github.com/jpillora/backoff"
	"net"
	"testing"
	"time"
)

// wait for a connect on a port progressively backing off
// returns false if we couldn't establish a connection
func WaitForConnect(host string) bool {
	connected := false
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    100 * time.Millisecond,
		Max:    5 * time.Second,
	}
	for {
		conn, err := net.Dial("tcp", host)
		if err != nil {
			d := b.Duration()
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

// try to connect to host, if it fails, the test fails
func AssertConnected(host string, t *testing.T) {
	connected := WaitForConnect(host)
	if !connected {
		t.Errorf("could not connect to host %s", host)
	}
}
