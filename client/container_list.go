package client // import "github.com/ellcrys/docker/client"

import (
	"context"
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/ellcrys/docker/api/types"
	"github.com/ellcrys/docker/api/types/filters"
)

// ContainerList returns the list of containers in the docker host.
func (cli *Client) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
	query := url.Values{}

	if options.All {
		query.Set("all", "1")
	}

	if options.Limit != -1 {
		query.Set("limit", strconv.Itoa(options.Limit))
	}

	if options.Since != "" {
		query.Set("since", options.Since)
	}

	if options.Before != "" {
		query.Set("before", options.Before)
	}

	if options.Size {
		query.Set("size", "1")
	}

	if options.Filters.Len() > 0 {
		filterJSON, err := filters.ToParamWithVersion(cli.version, options.Filters)

		if err != nil {
			return nil, err
		}

		query.Set("filters", filterJSON)
	}

	resp, err := cli.get(ctx, "/containers/json", query, nil)
	if err != nil {
		return nil, err
	}

	var containers []types.Container
	err = json.NewDecoder(resp.body).Decode(&containers)
	ensureReaderClosed(resp)
	return containers, err
}
