package dockerfile // import "github.com/ellcrys/docker/builder/dockerfile"

import "github.com/ellcrys/docker/pkg/idtools"

func parseChownFlag(chown, ctrRootPath string, idMappings *idtools.IDMappings) (idtools.IDPair, error) {
	return idMappings.RootPair(), nil
}
