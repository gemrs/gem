package protocol_os_163

import (
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/protocol"
)

// +gen define_inbound:"Pkt51,SzFixed,InboundGroundItemAction"
type InboundGroundItemAction1 protocol.InboundGroundItemAction

func (struc *InboundGroundItemAction1) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Action = 0
	buf.GetU8(encoding.IntNegate)
	struc.X = buf.GetU16(encoding.IntLittleEndian)
	struc.ItemID = buf.GetU16(encoding.IntLittleEndian)
	struc.Y = buf.GetU16(encoding.IntOffset128)
}

// +gen define_inbound:"Pkt25,SzFixed,InboundGroundItemAction"
type InboundGroundItemAction2 protocol.InboundGroundItemAction

func (struc *InboundGroundItemAction2) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Action = 1
	struc.ItemID = buf.GetU16(encoding.IntLittleEndian)
	buf.GetU8(encoding.IntNegate)
	struc.X = buf.GetU16(encoding.IntLittleEndian)
	struc.Y = buf.GetU16()
}

// +gen define_inbound:"Pkt33,SzFixed,InboundGroundItemAction"
type InboundGroundItemAction3 protocol.InboundGroundItemAction

func (struc *InboundGroundItemAction3) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Action = 2
	struc.ItemID = buf.GetU16()
	struc.Y = buf.GetU16(encoding.IntLittleEndian, encoding.IntOffset128)
	struc.X = buf.GetU16(encoding.IntOffset128)
	buf.GetU8(encoding.IntOffset128)
}

// +gen define_inbound:"Pkt26,SzFixed,InboundGroundItemAction"
type InboundGroundItemAction4 protocol.InboundGroundItemAction

func (struc *InboundGroundItemAction4) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Action = 3
	struc.ItemID = buf.GetU16(encoding.IntOffset128)
	struc.X = buf.GetU16(encoding.IntLittleEndian)
	struc.Y = buf.GetU16()
	buf.GetU8(encoding.IntNegate)
}

// +gen define_inbound:"Pkt2,SzFixed,InboundGroundItemAction"
type InboundGroundItemAction5 protocol.InboundGroundItemAction

func (struc *InboundGroundItemAction5) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Action = 4
	buf.GetU8(encoding.IntOffset128)
	struc.Y = buf.GetU16()
	struc.X = buf.GetU16()
	struc.ItemID = buf.GetU16()
}
