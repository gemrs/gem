package player

import (
	"github.com/sinusoids/gem/gem/game/entity"
	"github.com/sinusoids/gem/gem/log"
)

// Player is an Entity representing a player
type Player interface {
	entity.Mob
	Profile() Profile
	Session() Session
	Log() *log.Module
}
