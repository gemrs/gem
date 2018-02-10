package protocol_os_163

import (
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/protocol/protocol_os"
)

type protocolImpl struct {
	protocol_os.ProtocolBase
}

var Protocol = protocolImpl{
	ProtocolBase: protocol_os.ProtocolBase{
		Revision: Revision,
	},
}

type OutboundPacketDefinition protocol_os.OutboundPacketDefinition

func (d OutboundPacketDefinition) Pack(e encoding.Encodable) encoding.Encodable {
	return protocol_os.OutboundPacketDefinition(d).Pack(e)
}

type InboundPacketDefinition protocol_os.InboundPacketDefinition

func (d InboundPacketDefinition) Pack(e encoding.Encodable) encoding.Encodable {
	return protocol_os.InboundPacketDefinition(d).Pack(e)
}

type PacketHeader protocol_os.PacketHeader

func (p PacketHeader) Encode(buf io.Writer, flags interface{}) {
	protocol_os.PacketHeader(p).Encode(buf, flags)
}

func (p *PacketHeader) Decode(buf io.Reader, flags interface{}) {
	(*protocol_os.PacketHeader)(p).Decode(buf, flags)
}

var SzFixed = protocol_os.SzFixed
var SzVar8 = protocol_os.SzVar8
var SzVar16 = protocol_os.SzVar16
