package protocol_os

import (
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/runite"
)

type ProtocolBase struct{}

func (ProtocolBase) NewUpdateService(runite *runite.Context) server.Service {
	return NewUpdateService(runite)
}

func (ProtocolBase) GameServiceId() int {
	return GameServiceId
}

func (ProtocolBase) UpdateServiceId() int {
	return UpdateServiceId
}

func (p ProtocolBase) ProtocolRevision() int {
	return Revision
}
