// +build !exclude_graphdriver_overlay2,linux

package register // import "github.com/ellcrys/docker/daemon/graphdriver/register"

import (
	// register the overlay2 graphdriver
	_ "github.com/ellcrys/docker/daemon/graphdriver/overlay2"
)
