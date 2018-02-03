package protocol_os_162

import (
	"errors"
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
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

// + gen define_inbound:"Pkt93,SzVar8"
// +gen define_inbound:"Unimplemented"
type InboundCommand protocol.InboundCommand

func (struc *InboundCommand) Decode(buf io.Reader, flags interface{}) {
	(*encoding.String)(&struc.Command).Decode(buf, 0)
}

// + gen define_inbound:"Pkt88,SzVar8"
// +gen define_inbound:"Unimplemented"
type InboundMouseMovement protocol.InboundMouseMovement

func (struc *InboundMouseMovement) Decode(buf io.Reader, flags interface{}) {}

// + gen define_inbound:"Pkt3,SzFixed"
// +gen define_inbound:"Unimplemented"
type InboundMouseClick protocol.InboundMouseClick

func (struc *InboundMouseClick) Decode(buf io.Reader, flags interface{}) {}

// + gen define_inbound:"Pkt63,SzFixed"
// +gen define_inbound:"Unimplemented"
type InboundPing protocol.InboundPing

func (struc *InboundPing) Decode(buf io.Reader, flags interface{}) {}

// + gen define_inbound:"Pkt41,SzFixed"
// +gen define_inbound:"Unimplemented"
type InboundWindowFocus protocol.InboundWindowFocus

func (struc *InboundWindowFocus) Decode(buf io.Reader, flags interface{}) {}

// + gen define_inbound:"Pkt71,SzVar16"
// +gen define_inbound:"Unimplemented"
type InboundKeyPress protocol.InboundKeyPress

func (struc *InboundKeyPress) Decode(buf io.Reader, flags interface{}) {}

// + gen define_inbound:"Pkt84,SzFixed"
// +gen define_inbound:"Unimplemented"
type InboundCameraMovement protocol.InboundCameraMovement

func (struc *InboundCameraMovement) Decode(buf io.Reader, flags interface{}) {}

// + gen define_inbound:"Pkt43,SzFixed,InboundInventoryAction"
// +gen define_inbound:"Unimplemented"
type InboundInventoryAction1 protocol.InboundInventoryAction

func (struc *InboundInventoryAction1) Decode(buf io.Reader, flags interface{}) {
	struc.Action = 0
	var tmp16 encoding.Uint16
	var tmp32 encoding.Uint32

	tmp16.Decode(buf, encoding.IntLittleEndian)
	struc.ItemID = int(tmp16)

	tmp32.Decode(buf, encoding.IntLittleEndian)
	struc.InterfaceID = int(tmp32) >> 16

	tmp16.Decode(buf, encoding.IntOffset128)
	struc.Slot = int(tmp16)
}

// + gen define_inbound:"Pkt95,SzFixed,InboundInventoryAction"
// +gen define_inbound:"Unimplemented"
type InboundInventoryAction2 protocol.InboundInventoryAction

func (struc *InboundInventoryAction2) Decode(buf io.Reader, flags interface{}) {
	struc.Action = 1
	var tmp16 encoding.Uint16
	var tmp32 encoding.Uint32

	tmp16.Decode(buf, encoding.IntOffset128)
	struc.ItemID = int(tmp16)

	tmp32.Decode(buf, encoding.IntRPDPEndian)
	struc.InterfaceID = int(tmp32) >> 16

	tmp16.Decode(buf, encoding.IntOffset128)
	struc.Slot = int(tmp16)
}

// + gen define_inbound:"Pkt57,SzFixed,InboundInventoryAction"
// +gen define_inbound:"Unimplemented"
type InboundInventoryAction3 protocol.InboundInventoryAction

func (struc *InboundInventoryAction3) Decode(buf io.Reader, flags interface{}) {
	struc.Action = 2
	var tmp16 encoding.Uint16
	var tmp32 encoding.Uint32

	tmp32.Decode(buf, encoding.IntPDPEndian)
	struc.InterfaceID = int(tmp32) >> 16

	tmp16.Decode(buf, encoding.IntOffset128)
	struc.Slot = int(tmp16)

	tmp16.Decode(buf, encoding.IntNilFlag)
	struc.ItemID = int(tmp16)
}

// + gen define_inbound:"Pkt76,SzFixed,InboundInventoryAction"
// +gen define_inbound:"Unimplemented"
type InboundInventoryAction4 protocol.InboundInventoryAction

func (struc *InboundInventoryAction4) Decode(buf io.Reader, flags interface{}) {
	struc.Action = 3
	var tmp16 encoding.Uint16
	var tmp32 encoding.Uint32

	tmp16.Decode(buf, encoding.IntNilFlag)
	struc.Slot = int(tmp16)

	tmp16.Decode(buf, encoding.IntLittleEndian)
	struc.ItemID = int(tmp16)

	tmp32.Decode(buf, encoding.IntPDPEndian)
	struc.InterfaceID = int(tmp32) >> 16
}

// + gen define_inbound:"Pkt14,SzFixed,InboundInventoryAction"
// +gen define_inbound:"Unimplemented"
type InboundInventoryAction5 protocol.InboundInventoryAction

func (struc *InboundInventoryAction5) Decode(buf io.Reader, flags interface{}) {
	struc.Action = 4
	var tmp16 encoding.Uint16
	var tmp32 encoding.Uint32

	tmp16.Decode(buf, encoding.IntLittleEndian)
	struc.Slot = int(tmp16)

	tmp16.Decode(buf, encoding.IntNilFlag)
	struc.ItemID = int(tmp16)

	tmp32.Decode(buf, encoding.IntNilFlag)
	struc.InterfaceID = int(tmp32) >> 16
}

// + gen define_inbound:"Pkt2,SzFixed"
// +gen define_inbound:"Unimplemented"
type InboundInventorySwapItem protocol.InboundInventorySwapItem

func (struc *InboundInventorySwapItem) Decode(buf io.Reader, flags interface{}) {
	var tmp32 encoding.Uint32
	var tmp16 encoding.Uint16
	var tmp8 encoding.Uint8

	tmp16.Decode(buf, encoding.IntNilFlag)
	struc.To = int(tmp16)

	tmp8.Decode(buf, encoding.IntNegate)
	struc.Inserting = int(tmp8) == 1

	tmp32.Decode(buf, encoding.IntRPDPEndian)
	struc.InterfaceID = int(tmp32) >> 16

	tmp16.Decode(buf, encoding.IntNilFlag)
	struc.From = int(tmp16)
}

// + gen define_inbound:"Pkt24,SzFixed"
// +gen define_inbound:"Unimplemented"
type InboundTakeGroundItem protocol.InboundTakeGroundItem

func (struc *InboundTakeGroundItem) Decode(buf io.Reader, flags interface{}) {
	var tmp16 encoding.Uint16
	var tmp8 encoding.Uint8

	// shift pressed?
	tmp8.Decode(buf, encoding.IntNilFlag)

	tmp16.Decode(buf, encoding.IntOffset128)
	struc.ItemID = int(tmp16)

	tmp16.Decode(buf, encoding.IntNilFlag)
	struc.Y = int(tmp16)

	tmp16.Decode(buf, encoding.IntOffset128)
	struc.X = int(tmp16)
}

// +gen define_inbound:"Unimplemented"
type InboundChatMessage protocol.InboundChatMessage

func (struc *InboundChatMessage) Decode(buf io.Reader, flags interface{}) {
	var tmp8 encoding.Uint8
	var message encoding.Bytes

	tmp8.Decode(buf, encoding.IntegerFlag(encoding.IntOffset128|encoding.IntReverse))
	struc.Effects = int(tmp8)

	tmp8.Decode(buf, encoding.IntegerFlag(encoding.IntOffset128|encoding.IntReverse))
	struc.Colour = int(tmp8)

	message.Decode(buf, nil)
	data := []byte(message)
	size := len(data)
	decoded := make([]byte, size)
	for i, _ := range data {
		decoded[i] = byte(data[size-i-1] - 128)
	}
	struc.Message = encoding.ChatTextSanitize(encoding.ChatTextUnpack(decoded))
	struc.PackedMessage = encoding.ChatTextPack(struc.Message)
}
