package entity

import (
	"gem/game/player"
	"gem/game/position"
	"gem/game/server"
)

// An Entity is a 'thing' within the world, with a position, and an index.
type Entity interface {
	position.Locatable
}

// Player is an Entity representing a player
type Player interface {
	Entity
	server.Client
	Profile() *player.Profile
	Session() *player.Session
}
