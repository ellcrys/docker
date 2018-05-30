package daemon // import "github.com/ellcrys/docker/daemon"

import (
	// Importing packages here only to make sure their init gets called and
	// therefore they register themselves to the logdriver factory.
	_ "github.com/ellcrys/docker/daemon/logger/awslogs"
	_ "github.com/ellcrys/docker/daemon/logger/fluentd"
	_ "github.com/ellcrys/docker/daemon/logger/gcplogs"
	_ "github.com/ellcrys/docker/daemon/logger/gelf"
	_ "github.com/ellcrys/docker/daemon/logger/journald"
	_ "github.com/ellcrys/docker/daemon/logger/jsonfilelog"
	_ "github.com/ellcrys/docker/daemon/logger/logentries"
	_ "github.com/ellcrys/docker/daemon/logger/splunk"
	_ "github.com/ellcrys/docker/daemon/logger/syslog"
)
