package daemon // import "github.com/ellcrys/docker/daemon"

import (
	"context"

	"github.com/ellcrys/docker/api/types"
	"github.com/ellcrys/docker/dockerversion"
)

// AuthenticateToRegistry checks the validity of credentials in authConfig
func (daemon *Daemon) AuthenticateToRegistry(ctx context.Context, authConfig *types.AuthConfig) (string, string, error) {
	return daemon.RegistryService.Auth(ctx, authConfig, dockerversion.DockerUserAgent(ctx))
}
