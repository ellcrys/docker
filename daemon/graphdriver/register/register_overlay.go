// +build !exclude_graphdriver_overlay,linux

package register // import "github.com/ellcrys/docker/daemon/graphdriver/register"

import (
	// register the overlay graphdriver
	_ "github.com/ellcrys/docker/daemon/graphdriver/overlay"
)
