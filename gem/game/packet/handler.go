package packet

import (
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/protocol"
)

type Handler func(protocol.Player, server.Message)
