package libnetwork

import "github.com/ellcrys/docker/pkg/reexec"

type setKeyData struct {
	ContainerID string
	Key         string
}

func init() {
	reexec.Register("libnetwork-setkey", processSetKeyReexec)
}
