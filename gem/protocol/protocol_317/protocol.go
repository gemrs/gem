package protocol_317

import (
	"github.com/gemrs/gem/gem/core/encoding"
)

type protocolImpl struct{}

var Protocol protocolImpl

type OutboundPacketDefinition struct {
	Number int
	Size   encoding.FrameSize
}

func (d OutboundPacketDefinition) Pack(e encoding.Encodable) encoding.Encodable {
	return encoding.PacketHeader{
		Number: d.Number,
		Size:   d.Size,
		Object: e,
	}
}

type InboundPacketDefinition struct {
	Type   encoding.Decodable
	Number int
	Size   encoding.FrameSize
}

func (d InboundPacketDefinition) Pack(e encoding.Encodable) encoding.Encodable {
	return encoding.PacketHeader{
		Number: d.Number,
		Size:   d.Size,
		Object: e,
	}
}
