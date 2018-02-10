package protocol_os_162

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

// +gen define_inbound:"Pkt93,SzVar8"
type InboundCommand protocol.InboundCommand

func (struc *InboundCommand) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Command = buf.GetStringZ()
}

// +gen define_inbound:"Pkt86,SzFixed,InboundInventoryAction"
type InboundInventoryAction1 protocol.InboundInventoryAction

func (struc *InboundInventoryAction1) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Action = 0
	struc.InterfaceID = buf.GetU32() >> 16
	struc.Slot = buf.GetU16(encoding.IntOffset128, encoding.IntLittleEndian)
	struc.ItemID = buf.GetU16(encoding.IntLittleEndian)
}

// +gen define_inbound:"Pkt56,SzFixed,InboundInventoryAction"
type InboundInventoryAction2 protocol.InboundInventoryAction

func (struc *InboundInventoryAction2) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Action = 1
	struc.ItemID = buf.GetU16(encoding.IntOffset128, encoding.IntLittleEndian)
	struc.Slot = buf.GetU16(encoding.IntLittleEndian)
	struc.InterfaceID = buf.GetU32() >> 16
}

// +gen define_inbound:"Pkt5,SzFixed,InboundInventoryAction"
type InboundInventoryAction3 protocol.InboundInventoryAction

func (struc *InboundInventoryAction3) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Action = 2
	struc.ItemID = buf.GetU16(encoding.IntOffset128, encoding.IntLittleEndian)
	struc.InterfaceID = buf.GetU32() >> 16
	struc.Slot = buf.GetU16(encoding.IntLittleEndian)
}

// +gen define_inbound:"Pkt85,SzFixed,InboundInventoryAction"
type InboundInventoryAction4 protocol.InboundInventoryAction

func (struc *InboundInventoryAction4) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Action = 3
	struc.Slot = buf.GetU16(encoding.IntLittleEndian)
	struc.InterfaceID = buf.GetU32() >> 16
	struc.ItemID = buf.GetU16(encoding.IntOffset128)
}

// +gen define_inbound:"Pkt14,SzFixed,InboundInventoryAction"
type InboundInventoryAction5 protocol.InboundInventoryAction

func (struc *InboundInventoryAction5) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Action = 4
	struc.Slot = buf.GetU16(encoding.IntLittleEndian, encoding.IntOffset128)
	struc.ItemID = buf.GetU16(encoding.IntLittleEndian)
	struc.InterfaceID = buf.GetU32(encoding.IntLittleEndian) >> 16
}

// +gen define_inbound:"Pkt57,SzFixed"
type InboundInventorySwapItem protocol.InboundInventorySwapItem

func (struc *InboundInventorySwapItem) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.InterfaceID = buf.GetU32(encoding.IntRPDPEndian)
	struc.Inserting = buf.GetU8(encoding.IntOffset128) == 1
	struc.To = buf.GetU16(encoding.IntLittleEndian, encoding.IntOffset128)
	struc.From = buf.GetU16(encoding.IntOffset128)
}

// +gen define_inbound:"Pkt42,SzFixed"
type InboundTakeGroundItem protocol.InboundTakeGroundItem

func (struc *InboundTakeGroundItem) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.ItemID = buf.GetU16(encoding.IntOffset128)
	struc.X = buf.GetU16(encoding.IntOffset128)
	// shift pressed?
	buf.GetU8(encoding.IntOffset128)
	struc.Y = buf.GetU16(encoding.IntOffset128)
}

// +gen define_inbound:"Pkt17,SzVar8"
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
