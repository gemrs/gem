package protocol_os_163

import (
	"errors"
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/game/data"
	"github.com/gemrs/gem/gem/protocol"
)

var ErrUnknownPacket = errors.New("unknown packet")

// NewInboundPacket accepts a packet id and returns a Decodable which can decode it
func (protocolImpl) NewInboundPacket(id int) (encoding.Decodable, error) {
	definition, ok := inboundPacketBuilders[id]
	if !ok {
		return new(UnknownPacket), nil
	}
	return definition(), nil
}

// +gen define_inbound:"Pkt13,SzVar8"
type InboundCommand protocol.InboundCommand

func (struc *InboundCommand) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Command = buf.GetStringZ()
}

// +gen define_inbound:"Pkt38,SzFixed"
type InboundInventorySwapItem protocol.InboundInventorySwapItem

func (struc *InboundInventorySwapItem) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.From = buf.GetU16(encoding.IntOffset128)
	struc.InterfaceID = buf.GetU32(encoding.IntRPDPEndian)
	struc.To = buf.GetU16(encoding.IntOffset128)
	struc.Inserting = buf.GetU8(encoding.IntOffset128) == 1
}

// +gen define_inbound:"Pkt14,SzVar8"
type InboundChatMessage protocol.InboundChatMessage

func (struc *InboundChatMessage) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	buf.GetU8()
	struc.Colour = buf.GetU8()
	struc.Effects = buf.GetU8()
	decompressedSize := buf.GetU16(encoding.IntPacked)

	message := buf.GetBytes(-1)
	compressed := []byte(message)
	decompressed := data.Huffman.Decompress(compressed, decompressedSize)
	struc.Message = string(decompressed)
	struc.PackedMessage = data.Huffman.Compress([]byte(struc.Message))
}
