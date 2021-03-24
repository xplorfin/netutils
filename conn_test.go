package netutils_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/xplorfin/netutils"
	"github.com/xplorfin/netutils/testutils"
)

func TestWaitForConnectDidTimeout(t *testing.T) {
	host := fmt.Sprintf(":%d", testutils.GetFreePort(t))
	didConnect := netutils.WaitForConnectTimeout(host, time.Millisecond)
	if didConnect {
		t.Errorf("expected connection to fail to %s", host)
	}
}

func TestWaitForConnect(t *testing.T) {
	host := fmt.Sprintf(":%d", testutils.GetUnfreePort(t))
	didConnect := netutils.WaitForConnect(host)
	if !didConnect {
		t.Errorf("expected to connect to %s", host)
	}
	testutils.AssertConnected(host, t)
}

func TestConnectExamples(t *testing.T) {
	ExampleWaitForConnect()
	ExampleWaitForConnectTimeout()
}

// WaitForConnect will attempt to connect and timeout after 10 retries of up to 5 seconds
func ExampleWaitForConnect() {
	// will attempt to connect and not timeout since url is valid
	succesful := netutils.WaitForConnect("https://entropy.rocks")
	if succesful {
		fmt.Println("connected to entropy!")
	}
}

// WaitForConnectTimeout will attempt to connect and timeout after 5 seconds
func ExampleWaitForConnectTimeout() {
	// will attempt to connect and not timeout since url is valid
	succesful := netutils.WaitForConnectTimeout("https://entropy.rocks", 5*time.Second)
	if succesful {
		fmt.Println("connected to entropy!")
	}

	succesful = netutils.WaitForConnectTimeout("https://entropy.rockssocks", time.Millisecond)
	if !succesful {
		fmt.Println("could not connect to non-existent domain")
	}
}
