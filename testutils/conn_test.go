package testutils_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/xplorfin/netutils/testutils"

	"github.com/stretchr/testify/assert"
)

func TestWaitForConnectDidTimeout(t *testing.T) {
	host := fmt.Sprintf(":%d", testutils.GetFreePort(t))
	didConnect := testutils.WaitForConnectTimeout(host, time.Millisecond)
	if didConnect {
		t.Errorf("expected connection to fail to %s", host)
	}
}

func TestWaitForConnect(t *testing.T) {
	host := fmt.Sprintf(":%d", testutils.GetUnfreePort(t))
	didConnect := testutils.WaitForConnect(host)
	if !didConnect {
		t.Errorf("expected to connect to %s", host)
	}
	testutils.AssertConnected(host, t)
}

func TestUrl(t *testing.T) {
	port := testutils.GetFreePort(t)
	go func() {
		err := testutils.MockHTTPServer(port)
		assert.Nil(t, err)
	}()
	testutils.AssertConnected(fmt.Sprintf("http://%s:%d", "localhost", port), t)
	testutils.AssertConnected("https://google.com/", t)
}
