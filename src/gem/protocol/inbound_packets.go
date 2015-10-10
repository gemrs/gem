package protocol

import (
	"errors"
	"reflect"

	"gem/encoding"
)

var ErrUnknownPacket = errors.New("unknown packet")

var inboundTypes map[int]encoding.Codable

var definitions = []encoding.PacketHeader{
	InboundPingDefinition,
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
		return nil, ErrUnknownPacket
	}
	typ := reflect.TypeOf(typePtr)
	value := reflect.New(typ)
	return reflect.Indirect(value).Interface().(encoding.Codable), nil
}
