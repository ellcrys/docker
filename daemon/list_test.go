package daemon

import (
	"testing"

	"github.com/ellcrys/docker/api/types"
	"github.com/ellcrys/docker/api/types/filters"
	"github.com/ellcrys/docker/container"
	"github.com/gotestyourself/gotestyourself/assert"
	is "github.com/gotestyourself/gotestyourself/assert/cmp"
)

func TestListInvalidFilter(t *testing.T) {
	db, err := container.NewViewDB()
	assert.Assert(t, err == nil)
	d := &Daemon{
		containersReplica: db,
	}

	f := filters.NewArgs(filters.Arg("invalid", "foo"))

	_, err = d.Containers(&types.ContainerListOptions{
		Filters: f,
	})
	assert.Assert(t, is.Error(err, "Invalid filter 'invalid'"))
}
