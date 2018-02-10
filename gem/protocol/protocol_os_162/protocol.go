package protocol_os_162

import (
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/protocol/protocol_os"
)

type protocolImpl struct {
	protocol_os.ProtocolBase
}

var Protocol protocolImpl

func init() {
	protocol_os.Revision = Revision
	protocol_os.InboundSizes = inboundPacketLengths
}

type OutboundPacketDefinition struct {
	Number int
	Size   protocol_os.FrameSize
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
	Size   protocol_os.FrameSize
}

func (d InboundPacketDefinition) Pack(e encoding.Encodable) encoding.Encodable {
	return PacketHeader{
		Number: d.Number,
		Size:   d.Size,
		Object: e,
	}
}

type PacketHeader protocol_os.PacketHeader

func (p PacketHeader) Encode(buf io.Writer, flags interface{}) {
	protocol_os.PacketHeader(p).Encode(buf, flags)
}

func (p *PacketHeader) Decode(buf io.Reader, flags interface{}) {
	(*protocol_os.PacketHeader)(p).Decode(buf, flags)
}

type UnknownPacket protocol_os.UnknownPacket

func (p *UnknownPacket) String() string {
	return (*protocol_os.UnknownPacket)(p).String()
}

func (p *UnknownPacket) Decode(buf io.Reader, flags interface{}) {
	(*protocol_os.UnknownPacket)(p).Decode(buf, flags)
}

var SzFixed = protocol_os.SzFixed
var SzVar8 = protocol_os.SzVar8
var SzVar16 = protocol_os.SzVar16
