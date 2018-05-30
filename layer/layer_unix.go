// +build linux freebsd darwin openbsd

package layer // import "github.com/ellcrys/docker/layer"

import "github.com/ellcrys/docker/pkg/stringid"

func (ls *layerStore) mountID(name string) string {
	return stringid.GenerateRandomID()
}
