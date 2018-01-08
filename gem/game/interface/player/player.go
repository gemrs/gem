package player

import (
	"github.com/gemrs/gem/gem/game/interface/entity"
	"github.com/gemrs/gem/gem/game/position"
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/willow/log"
)

// Player is an Entity representing a player
type Player interface {
	entity.Movable
	Profile() Profile
	Animations() Animations
	ClientConfig() ClientConfig

	Log() log.Log
	Conn() *server.Connection

	LoadedRegion() *position.Region
	VisibleEntities() *entity.Collection

	ChatMessageQueue() []*ChatMessage
	AppendChatMessage(*ChatMessage)
}

type ChatMessage struct {
	Effects       uint8
	Colour        uint8
	Message       string
	PackedMessage []byte
}
