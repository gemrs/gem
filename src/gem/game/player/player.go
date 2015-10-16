package player

import (
	"gem/game/server"
)

// Player is a Client which has a Profile and a Session
type Player interface {
	server.Client
	Profile() *Profile
	Session() *Session
}
