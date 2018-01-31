package packet

import (
	"reflect"

	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/protocol"
)

var routingTable = map[string]Handler{}

func registerHandler(packetType interface{}, handler Handler) {
	typeString := reflect.TypeOf(packetType).String()
	routingTable[typeString] = handler
}

func Dispatch(player protocol.Player, message server.Message) {
	typeString := reflect.TypeOf(message).String()
	if handler, ok := routingTable[typeString]; ok {
		handler(player, message)
	} else {
		player.Log().Debug("Unhandled message of type %v: %v", typeString, message)
	}
}
