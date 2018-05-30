package daemon // import "github.com/ellcrys/docker/daemon"

import (
	"github.com/ellcrys/docker/container"
	"github.com/ellcrys/docker/pkg/archive"
)

func (daemon *Daemon) tarCopyOptions(container *container.Container, noOverwriteDirNonDir bool) (*archive.TarOptions, error) {
	return daemon.defaultTarCopyOptions(noOverwriteDirNonDir), nil
}
