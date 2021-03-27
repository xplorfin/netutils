package testutils

import (
	"testing"

	"github.com/integralist/go-findroot/find"
)

// GetGitRoot gets the root directory of the current package
// this is useful for testing file generation, especially for
// automated swagger-like documentation
func GetGitRoot(t *testing.T) string {
	root, err := find.Repo()
	if err != nil {
		t.Error(err)
	}
	return root.Path
}
