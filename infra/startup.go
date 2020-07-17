package infra

import (
	"github.com/labbsr0x/githunter-api/infra/env"
	"github.com/labbsr0x/githunter-api/infra/server"
	"github.com/sirupsen/logrus"
)

// this Version value
var Version = "v1"

// Config is a function
func Config() {

	env.Config()

	logrus.Debugf("Starting HTTP server ...")
	server.Config(Version)
}
