package packet

import (
	"gem/encoding"
	"gem/game/player"
)

type Handler func(player.Player, encoding.Decodable)
