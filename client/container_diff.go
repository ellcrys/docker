package client // import "github.com/ellcrys/docker/client"

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/ellcrys/docker/api/types/container"
)

// ContainerDiff shows differences in a container filesystem since it was started.
func (cli *Client) ContainerDiff(ctx context.Context, containerID string) ([]container.ContainerChangeResponseItem, error) {
	var changes []container.ContainerChangeResponseItem

	serverResp, err := cli.get(ctx, "/containers/"+containerID+"/changes", url.Values{}, nil)
	if err != nil {
		return changes, err
	}

	err = json.NewDecoder(serverResp.body).Decode(&changes)
	ensureReaderClosed(serverResp)
	return changes, err
}
