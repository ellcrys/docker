// +build windows

package runconfig // import "github.com/ellcrys/docker/runconfig"

import (
	"testing"

	"github.com/ellcrys/docker/api/types/container"
)

func TestValidatePrivileged(t *testing.T) {
	expected := "Windows does not support privileged mode"
	err := validatePrivileged(&container.HostConfig{Privileged: true})
	if err == nil || err.Error() != expected {
		t.Fatalf("Expected %s", expected)
	}
}
