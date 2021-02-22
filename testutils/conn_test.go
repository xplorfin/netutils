package testutils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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

func TestUrl(t *testing.T) {
	port := GetFreePort(t)
	go func() {
		err := MockHttpServer(port)
		assert.Nil(t, err)
	}()
	AssertConnected(fmt.Sprintf("http://%s:%d", "localhost", port), t)
	AssertConnected("https://google.com/", t)
}
