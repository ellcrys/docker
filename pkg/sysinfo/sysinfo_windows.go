package sysinfo // import "github.com/ellcrys/docker/pkg/sysinfo"

// New returns an empty SysInfo for windows for now.
func New(quiet bool) *SysInfo {
	sysInfo := &SysInfo{}
	return sysInfo
}
