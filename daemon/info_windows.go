package daemon // import "github.com/ellcrys/docker/daemon"

import (
	"github.com/ellcrys/docker/api/types"
	"github.com/ellcrys/docker/pkg/sysinfo"
)

// FillPlatformInfo fills the platform related info.
func (daemon *Daemon) FillPlatformInfo(v *types.Info, sysInfo *sysinfo.SysInfo) {
}
