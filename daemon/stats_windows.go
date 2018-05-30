package daemon // import "github.com/ellcrys/docker/daemon"

import (
	"github.com/ellcrys/docker/api/types"
	"github.com/ellcrys/docker/container"
)

// Windows network stats are obtained directly through HCS, hence this is a no-op.
func (daemon *Daemon) getNetworkStats(c *container.Container) (map[string]types.NetworkStats, error) {
	return make(map[string]types.NetworkStats), nil
}
