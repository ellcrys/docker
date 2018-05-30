// +build linux,cgo,!static_build

package devicemapper // import "github.com/ellcrys/docker/pkg/devicemapper"

// #cgo pkg-config: devmapper
import "C"
