package plugingetter // import "github.com/ellcrys/docker/pkg/plugingetter"

import (
	"net"
	"time"

	"github.com/ellcrys/docker/pkg/plugins"
)

const (
	// Lookup doesn't update RefCount
	Lookup = 0
	// Acquire increments RefCount
	Acquire = 1
	// Release decrements RefCount
	Release = -1
)

// CompatPlugin is an abstraction to handle both v2(new) and v1(legacy) plugins.
type CompatPlugin interface {
	Client() *plugins.Client
	Name() string
	ScopedPath(string) string
	IsV1() bool
}

// PluginAddr is a plugin that exposes the socket address for creating custom clients rather than the built-in `*plugins.Client`
type PluginAddr interface {
	CompatPlugin
	Addr() net.Addr
	Timeout() time.Duration
	Protocol() string
}

// CountedPlugin is a plugin which is reference counted.
type CountedPlugin interface {
	Acquire()
	Release()
	CompatPlugin
}

// PluginGetter is the interface implemented by Store
type PluginGetter interface {
	Get(name, capability string, mode int) (CompatPlugin, error)
	GetAllByCap(capability string) ([]CompatPlugin, error)
	GetAllManagedPluginsByCap(capability string) []CompatPlugin
	Handle(capability string, callback func(string, *plugins.Client))
}
