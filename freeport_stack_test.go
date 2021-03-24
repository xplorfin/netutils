package netutils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xplorfin/netutils"
	"github.com/xplorfin/netutils/testutils"
)

func TestGetFreeportStack(t *testing.T) {
	stack := netutils.NewFreeportStack()

	for i := 0; i < 10; i++ {
		port := stack.GetPort()
		assert.True(t, testutils.PortIsAvailable(port))
	}
}

// will assert no-panic
func TestFreeportExamples(t *testing.T) {
	ExampleFreePortStack_GetFreePort()
}

// Freeport stack gets multiple freeports that are non overlapping
func ExampleFreePortStack_GetFreePort() {
	// create a port stack
	stack := netutils.NewFreeportStack()

	// stack.GetPort will gurantee the ports have not been used before in the stack
	for _, port := range []int{stack.GetPort(), stack.GetPort()} {
		port := port
		go func() {
			err := testutils.MockHTTPServer(port)
			if err != nil {
				panic(err)
			}
		}()
	}
}
