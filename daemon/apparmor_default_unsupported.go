// +build !linux

package daemon // import "github.com/ellcrys/docker/daemon"

func ensureDefaultAppArmorProfile() error {
	return nil
}
