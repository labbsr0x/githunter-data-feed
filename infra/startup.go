package infra

import (
	"github.com/labbsr0x/githunter-repos/infra/server"
	"github.com/rs/zerolog/log"
)

// this Version value
var Version = "v1"

// Config is a function
func Config() {

	log.Warn().Msg("Starting HTTP server ...")
	server.Config(Version)
}
