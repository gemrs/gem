package packet

import (
	"github.com/sinusoids/gem/gem/encoding"
	"github.com/sinusoids/gem/gem/game/interface/player"
)

type Handler func(player.Player, encoding.Decodable)
