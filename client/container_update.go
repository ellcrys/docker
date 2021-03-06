package client // import "github.com/ellcrys/docker/client"

import (
	"context"
	"encoding/json"

	"github.com/ellcrys/docker/api/types/container"
)

// ContainerUpdate updates resources of a container
func (cli *Client) ContainerUpdate(ctx context.Context, containerID string, updateConfig container.UpdateConfig) (container.ContainerUpdateOKBody, error) {
	var response container.ContainerUpdateOKBody
	serverResp, err := cli.post(ctx, "/containers/"+containerID+"/update", nil, updateConfig, nil)
	if err != nil {
		return response, err
	}

	err = json.NewDecoder(serverResp.body).Decode(&response)

	ensureReaderClosed(serverResp)
	return response, err
}
