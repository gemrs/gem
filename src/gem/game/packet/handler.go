package packet

import (
	"gem/encoding"
	"gem/game/entity"
)

type Handler func(entity.Player, encoding.Decodable)
