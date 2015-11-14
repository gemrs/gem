package packet

import (
	"github.com/gemrs/gem/gem/encoding"
	"github.com/gemrs/gem/gem/game/interface/player"
)

type Handler func(player.Player, encoding.Decodable)
