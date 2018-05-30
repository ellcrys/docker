// +build !exclude_graphdriver_zfs,linux !exclude_graphdriver_zfs,freebsd

package register // import "github.com/ellcrys/docker/daemon/graphdriver/register"

import (
	// register the zfs driver
	_ "github.com/ellcrys/docker/daemon/graphdriver/zfs"
)
