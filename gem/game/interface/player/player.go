package player

import (
	"github.com/sinusoids/gem/gem/game/interface/entity"
	"github.com/sinusoids/gem/gem/log"
)

// Player is an Entity representing a player
type Player interface {
	entity.Movable
	Profile() Profile
	Session() Session
	Log() *log.Module
}
