package testutils_test

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xplorfin/netutils/testutils"
)

// warning: this test will fail if the directory is not called netutils
// or if you are not in a git repo.
// TODO mock testing to test against failures (this is done in the parent package)
func TestGetGitRoot(t *testing.T) {
	res := testutils.GetGitRoot(t)
	_, repoName := path.Split(res)
	assert.Equal(t, repoName, "netutils")
}
