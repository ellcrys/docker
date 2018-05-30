package httputils // import "github.com/ellcrys/docker/api/server/httputils"

import (
	"io"

	"github.com/ellcrys/docker/api/types/container"
	"github.com/ellcrys/docker/api/types/network"
)

// ContainerDecoder specifies how
// to translate an io.Reader into
// container configuration.
type ContainerDecoder interface {
	DecodeConfig(src io.Reader) (*container.Config, *container.HostConfig, *network.NetworkingConfig, error)
	DecodeHostConfig(src io.Reader) (*container.HostConfig, error)
}
