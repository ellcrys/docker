package client // import "github.com/ellcrys/docker/client"

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/ellcrys/docker/api/types"
	"github.com/ellcrys/docker/api/types/filters"
)

// PluginList returns the installed plugins
func (cli *Client) PluginList(ctx context.Context, filter filters.Args) (types.PluginsListResponse, error) {
	var plugins types.PluginsListResponse
	query := url.Values{}

	if filter.Len() > 0 {
		filterJSON, err := filters.ToParamWithVersion(cli.version, filter)
		if err != nil {
			return plugins, err
		}
		query.Set("filters", filterJSON)
	}
	resp, err := cli.get(ctx, "/plugins", query, nil)
	if err != nil {
		return plugins, wrapResponseError(err, resp, "plugin", "")
	}

	err = json.NewDecoder(resp.body).Decode(&plugins)
	ensureReaderClosed(resp)
	return plugins, err
}
