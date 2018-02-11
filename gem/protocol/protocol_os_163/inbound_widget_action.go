package protocol_os_163

import (
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/protocol"
)

// +gen define_inbound:"Pkt76,SzFixed,InboundWidgetAction"
type InboundWidgetAction1 protocol.InboundWidgetAction

func (struc *InboundWidgetAction1) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Action = 0
	locator := buf.Get32()
	struc.InterfaceID = locator >> 16
	struc.WidgetID = locator & 0xFFFF
	struc.Param = buf.Get16()
	struc.ItemID = buf.Get16()
}

// +gen define_inbound:"Pkt41,SzFixed,InboundWidgetAction"
type InboundWidgetAction2 protocol.InboundWidgetAction

func (struc *InboundWidgetAction2) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Action = 1
	locator := buf.Get32()
	struc.InterfaceID = locator >> 16
	struc.WidgetID = locator & 0xFFFF
	struc.Param = buf.Get16()
	struc.ItemID = buf.Get16()
}

// +gen define_inbound:"Pkt16,SzFixed,InboundWidgetAction"
type InboundWidgetAction3 protocol.InboundWidgetAction

func (struc *InboundWidgetAction3) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Action = 2
	locator := buf.Get32()
	struc.InterfaceID = locator >> 16
	struc.WidgetID = locator & 0xFFFF
	struc.Param = buf.Get16()
	struc.ItemID = buf.Get16()
}

// +gen define_inbound:"Pkt17,SzFixed,InboundWidgetAction"
type InboundWidgetAction4 protocol.InboundWidgetAction

func (struc *InboundWidgetAction4) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Action = 3
	locator := buf.Get32()
	struc.InterfaceID = locator >> 16
	struc.WidgetID = locator & 0xFFFF
	struc.Param = buf.Get16()
	struc.ItemID = buf.Get16()
}

// +gen define_inbound:"Pkt62,SzFixed,InboundWidgetAction"
type InboundWidgetAction5 protocol.InboundWidgetAction

func (struc *InboundWidgetAction5) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Action = 4
	locator := buf.Get32()
	struc.InterfaceID = locator >> 16
	struc.WidgetID = locator & 0xFFFF
	struc.Param = buf.Get16()
	struc.ItemID = buf.Get16()
}

// +gen define_inbound:"Pkt34,SzFixed,InboundWidgetAction"
type InboundWidgetAction6 protocol.InboundWidgetAction

func (struc *InboundWidgetAction6) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Action = 5
	locator := buf.Get32()
	struc.InterfaceID = locator >> 16
	struc.WidgetID = locator & 0xFFFF
	struc.Param = buf.Get16()
	struc.ItemID = buf.Get16()
}

// +gen define_inbound:"Pkt85,SzFixed,InboundWidgetAction"
type InboundWidgetAction7 protocol.InboundWidgetAction

func (struc *InboundWidgetAction7) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Action = 6
	locator := buf.Get32()
	struc.InterfaceID = locator >> 16
	struc.WidgetID = locator & 0xFFFF
	struc.Param = buf.Get16()
	struc.ItemID = buf.Get16()
}

// +gen define_inbound:"Pkt0,SzFixed,InboundWidgetAction"
type InboundWidgetAction8 protocol.InboundWidgetAction

func (struc *InboundWidgetAction8) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Action = 7
	locator := buf.Get32()
	struc.InterfaceID = locator >> 16
	struc.WidgetID = locator & 0xFFFF
	struc.Param = buf.Get16()
	struc.ItemID = buf.Get16()
}

// +gen define_inbound:"Pkt57,SzFixed,InboundWidgetAction"
type InboundWidgetAction9 protocol.InboundWidgetAction

func (struc *InboundWidgetAction9) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Action = 8
	locator := buf.Get32()
	struc.InterfaceID = locator >> 16
	struc.WidgetID = locator & 0xFFFF
	struc.Param = buf.Get16()
	struc.ItemID = buf.Get16()
}

// +gen define_inbound:"Pkt39,SzFixed,InboundWidgetAction"
type InboundWidgetAction10 protocol.InboundWidgetAction

func (struc *InboundWidgetAction10) Decode(r io.Reader, flags interface{}) {
	buf := encoding.WrapReader(r)
	struc.Action = 9
	locator := buf.Get32()
	struc.InterfaceID = locator >> 16
	struc.WidgetID = locator & 0xFFFF
	struc.Param = buf.Get16()
	struc.ItemID = buf.Get16()
}
