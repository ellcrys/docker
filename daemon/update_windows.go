package daemon // import "github.com/ellcrys/docker/daemon"

import (
	"github.com/ellcrys/docker/api/types/container"
	"github.com/ellcrys/docker/libcontainerd"
)

func toContainerdResources(resources container.Resources) *libcontainerd.Resources {
	// We don't support update, so do nothing
	return nil
}
