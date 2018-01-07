package game

import (
	"errors"
	"reflect"

	"github.com/gemrs/gem/gem/encoding"
)

var ErrUnknownPacket = errors.New("unknown packet")

var inboundTypes map[int]encoding.Codable

var definitions = []encoding.PacketHeader{
	InboundPingDefinition,
	InboundChatMessageDefinition,
	InboundPlayerWalkDefinition,
	InboundPlayerWalkMapDefinition,
}

func init() {
	inboundTypes = make(map[int]encoding.Codable)
	for _, p := range definitions {
		inboundTypes[p.Number] = p.Type
	}
}

// NewInboundPacket accepts a packet id and returns a Codable which can decode it
func NewInboundPacket(id int) (encoding.Codable, error) {
	typePtr, ok := inboundTypes[id]
	if !ok {
		return new(UnknownPacket), nil
	}
	typ := reflect.TypeOf(typePtr).Elem()
	value := reflect.New(typ)
	return value.Interface().(encoding.Codable), nil
}
