package client // import "github.com/ellcrys/docker/client"

import "context"

// ServiceRemove kills and removes a service.
func (cli *Client) ServiceRemove(ctx context.Context, serviceID string) error {
	resp, err := cli.delete(ctx, "/services/"+serviceID, nil, nil)
	ensureReaderClosed(resp)
	return wrapResponseError(err, resp, "service", serviceID)
}
