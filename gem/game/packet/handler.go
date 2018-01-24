package packet

import (
	"github.com/gemrs/gem/gem/game/player"
	"github.com/gemrs/gem/gem/game/server"
)

type Handler func(*player.Player, server.Message)
