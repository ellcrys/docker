// +build !exclude_graphdriver_aufs,linux

package register // import "github.com/ellcrys/docker/daemon/graphdriver/register"

import (
	// register the aufs graphdriver
	_ "github.com/ellcrys/docker/daemon/graphdriver/aufs"
)
