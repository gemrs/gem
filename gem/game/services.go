package game

import (
	"github.com/gemrs/gem/gem/auth"
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/runite"
)

//glua:bind
func RegisterServices(target *server.Server, runite *runite.Context, rsaKeyPath string, auth auth.Provider) {
	target.SetService(14, NewGameService(runite, rsaKeyPath, auth))
	target.SetService(15, server.Proto.NewUpdateService(runite))
}
