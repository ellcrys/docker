package container // import "github.com/ellcrys/docker/integration/container"

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	"github.com/ellcrys/docker/api/types"
	"github.com/ellcrys/docker/api/types/filters"
	"github.com/ellcrys/docker/integration/internal/container"
	"github.com/ellcrys/docker/internal/test/request"
	"github.com/gotestyourself/gotestyourself/assert"
	is "github.com/gotestyourself/gotestyourself/assert/cmp"
	"github.com/gotestyourself/gotestyourself/skip"
)

func TestLinksEtcHostsContentMatch(t *testing.T) {
	skip.If(t, testEnv.IsRemoteDaemon())

	hosts, err := ioutil.ReadFile("/etc/hosts")
	skip.If(t, os.IsNotExist(err))

	defer setupTest(t)()
	client := request.NewAPIClient(t)
	ctx := context.Background()

	cID := container.Run(t, ctx, client, container.WithNetworkMode("host"))
	res, err := container.Exec(ctx, client, cID, []string{"cat", "/etc/hosts"})
	assert.NilError(t, err)
	assert.Assert(t, is.Len(res.Stderr(), 0))
	assert.Equal(t, 0, res.ExitCode)

	assert.Check(t, is.Equal(string(hosts), res.Stdout()))
}

func TestLinksContainerNames(t *testing.T) {
	skip.If(t, testEnv.DaemonInfo.OSType != "linux")

	defer setupTest(t)()
	client := request.NewAPIClient(t)
	ctx := context.Background()

	container.Run(t, ctx, client, container.WithName("first"))
	container.Run(t, ctx, client, container.WithName("second"), container.WithLinks("first:first"))

	f := filters.NewArgs(filters.Arg("name", "first"))

	containers, err := client.ContainerList(ctx, types.ContainerListOptions{
		Filters: f,
	})
	assert.NilError(t, err)
	assert.Check(t, is.Equal(1, len(containers)))
	assert.Check(t, is.DeepEqual([]string{"/first", "/second/first"}, containers[0].Names))
}
