// +build !linux,!windows

package daemon // import "github.com/ellcrys/docker/daemon"

func configsSupported() bool {
	return false
}
