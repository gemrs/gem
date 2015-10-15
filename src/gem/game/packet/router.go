package packet

import (
	"reflect"

	"gem/encoding"
	"gem/game/server"
)

var routingTable = map[string]Handler{}

func registerHandler(packetType interface{}, handler Handler) {
	typeString := reflect.TypeOf(packetType).String()
	routingTable[typeString] = handler
}

func Dispatch(player server.Player, packet encoding.Decodable) {
	typeString := reflect.TypeOf(packet).String()
	if handler, ok := routingTable[typeString]; ok {
		handler(player, packet)
	} else {
		player.Log().Infof("Unhandled packet: %v", packet)
	}
}