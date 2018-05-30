// +build !linux

package vfs // import "github.com/ellcrys/docker/daemon/graphdriver/vfs"

import "github.com/ellcrys/docker/pkg/chrootarchive"

func dirCopy(srcDir, dstDir string) error {
	return chrootarchive.NewArchiver(nil).CopyWithTar(srcDir, dstDir)
}
