package protocol_os_157

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

// +gen define_inbound:"Pkt103,SzVar8"
type InboundCommand protocol.InboundCommand

func (struc *InboundCommand) Decode(buf io.Reader, flags interface{}) {
	(*encoding.JString)(&struc.Command).Decode(buf, 0)
}

// +gen define_inbound:"Pkt0,SzFixed"
type InboundPing protocol.InboundPing

func (struc *InboundPing) Decode(buf io.Reader, flags interface{}) {}

// +gen define_inbound:"Pkt122,SzFixed,InboundInventoryAction"
type InboundInventoryAction1 protocol.InboundInventoryAction

func (struc *InboundInventoryAction1) Decode(buf io.Reader, flags interface{}) {
	struc.Action = 0
	var tmp16 encoding.Uint16

	tmp16.Decode(buf, encoding.IntegerFlag(encoding.IntLittleEndian|encoding.IntOffset128))
	struc.InterfaceID = int(tmp16)

	tmp16.Decode(buf, encoding.IntegerFlag(encoding.IntOffset128))
	struc.Slot = int(tmp16)

	tmp16.Decode(buf, encoding.IntegerFlag(encoding.IntLittleEndian))
	struc.ItemID = int(tmp16)
}

// +gen define_inbound:"Pkt41,SzFixed,InboundInventoryAction"
type InboundInventoryAction2 protocol.InboundInventoryAction

func (struc *InboundInventoryAction2) Decode(buf io.Reader, flags interface{}) {
	struc.Action = 1
	var tmp16 encoding.Uint16

	tmp16.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	struc.ItemID = int(tmp16)

	tmp16.Decode(buf, encoding.IntegerFlag(encoding.IntOffset128))
	struc.Slot = int(tmp16)

	tmp16.Decode(buf, encoding.IntegerFlag(encoding.IntOffset128))
	struc.InterfaceID = int(tmp16)
}

// +gen define_inbound:"Pkt16,SzFixed,InboundInventoryAction"
type InboundInventoryAction3 protocol.InboundInventoryAction

func (struc *InboundInventoryAction3) Decode(buf io.Reader, flags interface{}) {
	struc.Action = 2
	var tmp16 encoding.Uint16

	tmp16.Decode(buf, encoding.IntegerFlag(encoding.IntOffset128))
	struc.ItemID = int(tmp16)

	tmp16.Decode(buf, encoding.IntegerFlag(encoding.IntLittleEndian|encoding.IntOffset128))
	struc.Slot = int(tmp16)

	tmp16.Decode(buf, encoding.IntegerFlag(encoding.IntLittleEndian|encoding.IntOffset128))
	struc.InterfaceID = int(tmp16)
}

// +gen define_inbound:"Pkt75,SzFixed,InboundInventoryAction"
type InboundInventoryAction4 protocol.InboundInventoryAction

func (struc *InboundInventoryAction4) Decode(buf io.Reader, flags interface{}) {
	struc.Action = 3
	var tmp16 encoding.Uint16

	tmp16.Decode(buf, encoding.IntegerFlag(encoding.IntLittleEndian|encoding.IntOffset128))
	struc.InterfaceID = int(tmp16)

	tmp16.Decode(buf, encoding.IntegerFlag(encoding.IntLittleEndian))
	struc.Slot = int(tmp16)

	tmp16.Decode(buf, encoding.IntegerFlag(encoding.IntOffset128))
	struc.ItemID = int(tmp16)
}

// +gen define_inbound:"Pkt87,SzFixed,InboundInventoryAction"
type InboundInventoryAction5 protocol.InboundInventoryAction

func (struc *InboundInventoryAction5) Decode(buf io.Reader, flags interface{}) {
	struc.Action = 4
	var tmp16 encoding.Uint16

	tmp16.Decode(buf, encoding.IntegerFlag(encoding.IntOffset128))
	struc.ItemID = int(tmp16)

	tmp16.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	struc.InterfaceID = int(tmp16)

	tmp16.Decode(buf, encoding.IntegerFlag(encoding.IntOffset128))
	struc.Slot = int(tmp16)
}

// +gen define_inbound:"Pkt214,SzFixed"
type InboundInventorySwapItem protocol.InboundInventorySwapItem

func (struc *InboundInventorySwapItem) Decode(buf io.Reader, flags interface{}) {
	var tmp16 encoding.Uint16
	var tmp8 encoding.Uint8

	tmp16.Decode(buf, encoding.IntegerFlag(encoding.IntLittleEndian|encoding.IntOffset128))
	struc.InterfaceID = int(tmp16)

	tmp8.Decode(buf, encoding.IntegerFlag(encoding.IntNegate))
	struc.Inserting = int(tmp8) == 1

	tmp16.Decode(buf, encoding.IntegerFlag(encoding.IntLittleEndian|encoding.IntOffset128))
	struc.From = int(tmp16)

	tmp16.Decode(buf, encoding.IntegerFlag(encoding.IntLittleEndian))
	struc.To = int(tmp16)
}

// +gen define_inbound:"Pkt236,SzFixed"
type InboundTakeGroundItem protocol.InboundTakeGroundItem

func (struc *InboundTakeGroundItem) Decode(buf io.Reader, flags interface{}) {
	var tmp16 encoding.Uint16

	tmp16.Decode(buf, encoding.IntegerFlag(encoding.IntLittleEndian))
	struc.Y = int(tmp16)

	tmp16.Decode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	struc.ItemID = int(tmp16)

	tmp16.Decode(buf, encoding.IntegerFlag(encoding.IntLittleEndian))
	struc.X = int(tmp16)
}

// +gen define_inbound:"Pkt4,SzVar8"
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
