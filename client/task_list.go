package client // import "github.com/ellcrys/docker/client"

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/ellcrys/docker/api/types"
	"github.com/ellcrys/docker/api/types/filters"
	"github.com/ellcrys/docker/api/types/swarm"
)

// TaskList returns the list of tasks.
func (cli *Client) TaskList(ctx context.Context, options types.TaskListOptions) ([]swarm.Task, error) {
	query := url.Values{}

	if options.Filters.Len() > 0 {
		filterJSON, err := filters.ToJSON(options.Filters)
		if err != nil {
			return nil, err
		}

		query.Set("filters", filterJSON)
	}

	resp, err := cli.get(ctx, "/tasks", query, nil)
	if err != nil {
		return nil, err
	}

	var tasks []swarm.Task
	err = json.NewDecoder(resp.body).Decode(&tasks)
	ensureReaderClosed(resp)
	return tasks, err
}
