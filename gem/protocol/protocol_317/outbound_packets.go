package protocol_317

import (
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/protocol"
)

// +gen define_outbound:"Pkt253,SzVar8"
type OutboundChatMessage protocol.OutboundChatMessage

func (o OutboundChatMessage) Encode(buf io.Writer, flags interface{}) {
	encoding.JString(o.Message).Encode(buf, 0)
}

// +gen define_outbound:"Pkt215,SzFixed"
type OutboundCreateGlobalGroundItem protocol.OutboundCreateGlobalGroundItem

func (struc OutboundCreateGlobalGroundItem) Encode(buf io.Writer, flags interface{}) {
	encoding.Uint16(struc.ItemID).Encode(buf, encoding.IntegerFlag(encoding.IntOffset128))
	encoding.Uint8(struc.PositionOffset).Encode(buf, encoding.IntegerFlag(encoding.IntInverse128))
	encoding.Uint16(struc.Index).Encode(buf, encoding.IntegerFlag(encoding.IntOffset128))
	encoding.Uint16(struc.Count).Encode(buf, encoding.IntegerFlag(encoding.IntOffset128))
}

// +gen define_outbound:"Pkt44,SzFixed"
type OutboundCreateGroundItem protocol.OutboundCreateGroundItem

func (struc OutboundCreateGroundItem) Encode(buf io.Writer, flags interface{}) {
	encoding.Uint16(struc.ItemID).Encode(buf, encoding.IntegerFlag(encoding.IntLittleEndian|encoding.IntOffset128))
	encoding.Uint16(struc.Count).Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	encoding.Uint8(struc.PositionOffset).Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
}

// +gen define_outbound:"Pkt156,SzFixed"
type OutboundRemoveGroundItem protocol.OutboundRemoveGroundItem

func (struc OutboundRemoveGroundItem) Encode(buf io.Writer, flags interface{}) {
	encoding.Uint8(struc.PositionOffset).Encode(buf, encoding.IntegerFlag(encoding.IntOffset128))
	encoding.Uint16(struc.ItemID).Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
}

// +gen define_outbound:"Pkt71,SzFixed"
type OutboundTabInterface protocol.OutboundTabInterface

func (struc OutboundTabInterface) Encode(buf io.Writer, flags interface{}) {
	encoding.Uint16(struc.InterfaceID).Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	encoding.Uint8(struc.Tab).Encode(buf, encoding.IntegerFlag(encoding.IntOffset128))
}

// +gen define_outbound:"Pkt126,SzVar16"
type OutboundSetText protocol.OutboundSetText

func (struc OutboundSetText) Encode(buf io.Writer, flags interface{}) {
	encoding.JString(struc.Text).Encode(buf, 0)
	encoding.Uint16(struc.Id).Encode(buf, encoding.IntegerFlag(encoding.IntLittleEndian|encoding.IntInverse128))
}

// +gen define_outbound:"Pkt34,SzVar16"
type OutboundUpdateInventoryItem protocol.OutboundUpdateInventoryItem

func (struc OutboundUpdateInventoryItem) Encode(buf io.Writer, flags interface{}) {
	encoding.Uint16(struc.InventoryID).Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	encoding.Uint8(struc.Slot).Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	encoding.Uint16(struc.ItemID).Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	encoding.Uint8(struc.Count).Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
}

// +gen define_outbound:"Pkt109,SzFixed"
type OutboundLogout protocol.OutboundLogout

func (struc OutboundLogout) Encode(buf io.Writer, flags interface{}) {

}

// +gen define_outbound:"Pkt249,SzFixed"
type OutboundPlayerInit protocol.OutboundPlayerInit

func (struc OutboundPlayerInit) Encode(buf io.Writer, flags interface{}) {
	encoding.Uint8(struc.Membership).Encode(buf, encoding.IntegerFlag(encoding.IntOffset128))
	encoding.Uint16(struc.Index).Encode(buf, encoding.IntegerFlag(encoding.IntOffset128))
}

// +gen define_outbound:"Pkt73,SzFixed"
type OutboundRegionUpdate protocol.OutboundRegionUpdate

func (struc OutboundRegionUpdate) Encode(buf io.Writer, flags interface{}) {
	encoding.Uint16(struc.SectorX).Encode(buf, encoding.IntegerFlag(encoding.IntOffset128))
	encoding.Uint16(struc.SectorY).Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
}

// +gen define_outbound:"Pkt85,SzFixed"
type OutboundSectorUpdate protocol.OutboundSectorUpdate

func (struc OutboundSectorUpdate) Encode(buf io.Writer, flags interface{}) {
	encoding.Uint8(struc.PositionY).Encode(buf, encoding.IntegerFlag(encoding.IntNegate))
	encoding.Uint8(struc.PositionX).Encode(buf, encoding.IntegerFlag(encoding.IntNegate))
}

// +gen define_outbound:"Pkt134,SzFixed"
type OutboundSkill protocol.OutboundSkill

func (struc OutboundSkill) Encode(buf io.Writer, flags interface{}) {
	encoding.Uint8(struc.Skill).Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
	encoding.Uint32(struc.Experience).Encode(buf, encoding.IntegerFlag(encoding.IntPDPEndian))
	encoding.Uint8(struc.Level).Encode(buf, encoding.IntegerFlag(encoding.IntNilFlag))
}
