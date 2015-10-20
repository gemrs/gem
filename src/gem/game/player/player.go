package player

import (
	"gem/game/entity"
	"gem/log"
)

// Player is an Entity representing a player
type Player interface {
	entity.Mob
	Profile() Profile
	Session() Session
	Log() *log.Module
}
