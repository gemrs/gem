package protocol_os_162

import (
	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/runite"
)

type protocolImpl struct{}

var Protocol protocolImpl

func (protocolImpl) NewUpdateService(runite *runite.Context) server.Service {
	return NewUpdateService(runite)
}

func (protocolImpl) GameServiceId() int {
	return GameServiceId
}

func (protocolImpl) UpdateServiceId() int {
	return UpdateServiceId
}

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
