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
