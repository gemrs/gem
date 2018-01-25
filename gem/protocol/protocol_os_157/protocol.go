package protocol_os_157

import (
	"github.com/gemrs/gem/gem/core/encoding"
)

type protocolImpl struct{}

var Protocol protocolImpl

type OutboundPacketDefinition struct {
	Number int
	Size   FrameSize
}

func (d OutboundPacketDefinition) Pack(e encoding.Encodable) encoding.Encodable {
	return PacketHeader{
		Number: d.Number,
		Size:   d.Size,
		Object: e,
	}
}

type InboundPacketDefinition struct {
	Type   encoding.Decodable
	Number int
	Size   FrameSize
}

func (d InboundPacketDefinition) Pack(e encoding.Encodable) encoding.Encodable {
	return PacketHeader{
		Number: d.Number,
		Size:   d.Size,
		Object: e,
	}
}
