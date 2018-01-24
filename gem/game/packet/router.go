package packet

import (
	"reflect"

	"github.com/gemrs/gem/gem/game/player"
	"github.com/gemrs/gem/gem/game/server"
)

var routingTable = map[string]Handler{}

func registerHandler(packetType interface{}, handler Handler) {
	typeString := reflect.TypeOf(packetType).String()
	routingTable[typeString] = handler
}

func Dispatch(player *player.Player, message server.Message) {
	typeString := reflect.TypeOf(message).String()
	if handler, ok := routingTable[typeString]; ok {
		handler(player, message)
	} else {
		player.Log().Info("Unhandled message of type %v: %v", typeString, message)
	}
}
