// +build !linux,!darwin,!freebsd,!windows

package daemon // import "github.com/ellcrys/docker/daemon"

func (d *Daemon) setupDumpStackTrap(_ string) {
	return
}
