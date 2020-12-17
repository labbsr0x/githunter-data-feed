package infra

import (
	"github.com/labbsr0x/githunter-data-feed/infra/env"
	"github.com/labbsr0x/githunter-data-feed/infra/server"
	"github.com/sirupsen/logrus"
)

// this Version value
var Version = "v1"

// Config is a function
func Config() {
	logrus.Info("Starting environment and logRUS configuration...")
	env.Config()

	logrus.Info("Starting HTTP server ...")
	server.Config(Version)
}
