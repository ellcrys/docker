// +build !exclude_graphdriver_btrfs,linux

package register // import "github.com/ellcrys/docker/daemon/graphdriver/register"

import (
	// register the btrfs graphdriver
	_ "github.com/ellcrys/docker/daemon/graphdriver/btrfs"
)
