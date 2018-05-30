// +build linux,!seccomp

package seccomp // import "github.com/ellcrys/docker/profiles/seccomp"

import (
	"github.com/ellcrys/docker/api/types"
)

// DefaultProfile returns a nil pointer on unsupported systems.
func DefaultProfile() *types.Seccomp {
	return nil
}
