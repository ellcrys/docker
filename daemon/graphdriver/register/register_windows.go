package register // import "github.com/ellcrys/docker/daemon/graphdriver/register"

import (
	// register the windows graph drivers
	_ "github.com/ellcrys/docker/daemon/graphdriver/lcow"
	_ "github.com/ellcrys/docker/daemon/graphdriver/windows"
)
