// +build windows

package daemon // import "github.com/ellcrys/docker/daemon"

import "github.com/ellcrys/docker/pkg/plugingetter"

func registerMetricsPluginCallback(getter plugingetter.PluginGetter, sockPath string) {
}

func (daemon *Daemon) listenMetricsSock() (string, error) {
	return "", nil
}
