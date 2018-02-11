package protocol_os_163

import (
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/protocol"
)

// +gen define_inbound:"Pkt48,SzFixed,InboundInventoryAction"
type InboundInventoryAction1 protocol.InboundInventoryAction

func (struc *InboundInventoryAction1) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Action = 0
	struc.InterfaceID = buf.GetU32(encoding.IntLittleEndian) >> 16
	struc.Slot = buf.GetU16(encoding.IntOffset128)
	struc.ItemID = buf.GetU16(encoding.IntLittleEndian)
}

// +gen define_inbound:"Pkt91,SzFixed,InboundInventoryAction"
type InboundInventoryAction2 protocol.InboundInventoryAction

func (struc *InboundInventoryAction2) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Action = 1
	struc.Slot = buf.GetU16()
	struc.ItemID = buf.GetU16(encoding.IntOffset128)
	struc.InterfaceID = buf.GetU32() >> 16
}

// +gen define_inbound:"Pkt42,SzFixed,InboundInventoryAction"
type InboundInventoryAction3 protocol.InboundInventoryAction

func (struc *InboundInventoryAction3) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Action = 2
	struc.Slot = buf.GetU16()
	struc.InterfaceID = buf.GetU32(encoding.IntLittleEndian) >> 16
	struc.ItemID = buf.GetU16(encoding.IntLittleEndian)
}

// +gen define_inbound:"Pkt58,SzFixed,InboundInventoryAction"
type InboundInventoryAction4 protocol.InboundInventoryAction

func (struc *InboundInventoryAction4) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Action = 3
	struc.Slot = buf.GetU16(encoding.IntLittleEndian)
	struc.InterfaceID = buf.GetU32() >> 16
	struc.ItemID = buf.GetU16()
}

// +gen define_inbound:"Pkt59,SzFixed,InboundInventoryAction"
type InboundInventoryAction5 protocol.InboundInventoryAction

func (struc *InboundInventoryAction5) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Action = 4
	struc.Slot = buf.GetU16(encoding.IntLittleEndian, encoding.IntOffset128)
	struc.ItemID = buf.GetU16()
	struc.InterfaceID = buf.GetU32(encoding.IntLittleEndian) >> 16
}
