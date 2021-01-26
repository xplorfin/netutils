package testutils

import (
	"fmt"
	"testing"
	"time"
)

func TestWaitForConnectDidTimeout(t *testing.T) {
	host := fmt.Sprintf(":%d", GetFreePort(t))
	didConnect := WaitForConnectTimeout(host, time.Millisecond)
	if didConnect {
		t.Errorf("expected connection to fail to %s", host)
	}
}

func TestWaitForConnect(t *testing.T) {
	host := fmt.Sprintf(":%d", GetUnfreePort(t))
	didConnect := WaitForConnect(host)
	if !didConnect {
		t.Errorf("expected to connect to %s", host)
	}
	AssertConnected(host, t)
}
