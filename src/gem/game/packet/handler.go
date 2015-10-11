package packet

import (
	"gem/encoding"
	"gem/game/server"
)

type Handler func(server.Player, encoding.Decodable)
