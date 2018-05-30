// +build !linux,!windows

package daemon // import "github.com/ellcrys/docker/daemon"

func secretsSupported() bool {
	return false
}
