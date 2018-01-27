package protocol_os_157

import (
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/protocol"
)

// +gen define_outbound:"Unimplemented"
type OutboundChatMessage protocol.OutboundChatMessage

func (o OutboundChatMessage) Encode(buf io.Writer, flags interface{}) {
}

// +gen define_outbound:"Unimplemented"
type OutboundCreateGlobalGroundItem protocol.OutboundCreateGlobalGroundItem

func (struc OutboundCreateGlobalGroundItem) Encode(buf io.Writer, flags interface{}) {
}

// +gen define_outbound:"Unimplemented"
type OutboundCreateGroundItem protocol.OutboundCreateGroundItem

func (struc OutboundCreateGroundItem) Encode(buf io.Writer, flags interface{}) {
}

// +gen define_outbound:"Unimplemented"
type OutboundRemoveGroundItem protocol.OutboundRemoveGroundItem

func (struc OutboundRemoveGroundItem) Encode(buf io.Writer, flags interface{}) {
}

// +gen define_outbound
type OutboundTabInterface protocol.OutboundTabInterface

func (struc OutboundTabInterface) Encode(buf io.Writer, flags interface{}) {
	frame := getPlayerData(struc.ProtoData).frame
	packet := OutboundSetInterfaceDefinition.Pack(OutboundSetInterface{
		RootID:      frame.Root,
		ChildID:     frame.Tabs[struc.Tab],
		InterfaceID: struc.InterfaceID,
		Clickable:   true,
	})
	packet.Encode(buf, flags)
}

// +gen define_outbound:"Unimplemented"
type OutboundSetText protocol.OutboundSetText

func (struc OutboundSetText) Encode(buf io.Writer, flags interface{}) {
}

// +gen define_outbound:"Unimplemented"
type OutboundUpdateInventoryItem protocol.OutboundUpdateInventoryItem

func (struc OutboundUpdateInventoryItem) Encode(buf io.Writer, flags interface{}) {
}

// +gen define_outbound:"Unimplemented"
type OutboundLogout protocol.OutboundLogout

func (struc OutboundLogout) Encode(buf io.Writer, flags interface{}) {

}

// +gen define_outbound:"Unimplemented"
type OutboundPlayerInit protocol.OutboundPlayerInit

func (struc OutboundPlayerInit) Encode(buf io.Writer, flags interface{}) {
}

// +gen define_outbound:"Pkt0,SzVar16"
type OutboundRegionUpdate protocol.OutboundRegionUpdate

func (struc OutboundRegionUpdate) Encode(buf io.Writer, flags interface{}) {
	pos := struc.Position

	compressedPos := pos.Y() + (pos.X() << 14) + (pos.Z() << 28)
	bitBuf := encoding.NewBitBuffer(buf)
	bitBuf.Write(30, uint32(compressedPos))

	for i := 1; i < 2048; i++ {
		if i != struc.PlayerIndex {
			bitBuf.Write(18, 0)
		}
	}

	bitBuf.Close()

	sector := pos.Sector()
	sectorX := sector.X()
	sectorY := sector.Y()
	encoding.Uint16(sectorX).Encode(buf, encoding.IntLittleEndian)
	encoding.Uint16(sectorY).Encode(buf, encoding.IntOffset128)
	encoding.Uint16(9).Encode(buf, encoding.IntNilFlag)

	for x := (sectorX - 6) / 8; x <= (6+sectorX)/8; x++ {
		for y := (sectorY - 6) / 8; y <= (6+sectorY)/8; y++ {
			region := y + (x << 8)
			keys := mapKeys[region]
			for _, key := range keys {
				encoding.Uint32(key).Encode(buf, encoding.IntNilFlag)
			}
		}
	}
}

// +gen define_outbound:"Unimplemented"
type OutboundSectorUpdate protocol.OutboundSectorUpdate

func (struc OutboundSectorUpdate) Encode(buf io.Writer, flags interface{}) {
}

// +gen define_outbound:"Unimplemented"
type OutboundSkill protocol.OutboundSkill

func (struc OutboundSkill) Encode(buf io.Writer, flags interface{}) {
}

// +gen define_outbound:"Pkt45,SzFixed"
type OutboundResetCamera protocol.OutboundResetCamera

func (struc OutboundResetCamera) Encode(buf io.Writer, flags interface{}) {
}

// +gen define_outbound:"Pkt5,SzFixed"
type OutboundDnsLookup protocol.OutboundDnsLookup

func (struc OutboundDnsLookup) Encode(buf io.Writer, flags interface{}) {
	// FIXME: encode address properly
	encoding.Uint32(0).Encode(buf, encoding.IntNilFlag)
}

// +gen define_outbound:"Pkt25,SzFixed"
type OutboundSetRootInterface protocol.OutboundSetRootInterface

func (struc OutboundSetRootInterface) Encode(buf io.Writer, flags interface{}) {
	frame := struc.Frame.(FrameType)
	encoding.Uint16(frame.Root).Encode(buf, encoding.IntOffset128)
}

// +gen define_outbound:"Pkt24,SzFixed"
type OutboundSetInterface protocol.OutboundSetInterface

func (struc OutboundSetInterface) Encode(buf io.Writer, flags interface{}) {
	clickable := 0
	if struc.Clickable {
		clickable = 1
	}

	encoding.Uint8(clickable).Encode(buf, encoding.IntOffset128)
	encoding.Uint16(struc.InterfaceID).Encode(buf, encoding.IntLittleEndian)
	encoding.Uint32((struc.RootID<<16)|struc.ChildID).Encode(buf, encoding.IntPDPEndian)
}

// +gen define_outbound:"Pkt26,SzVar16"
type OutboundScriptEvent protocol.OutboundScriptEvent

func (struc OutboundScriptEvent) Encode(buf io.Writer, flags interface{}) {
	formatString := ""
	for _, p := range struc.Params {
		switch p.(type) {
		case int:
			formatString += "i"
		case string:
			formatString += "s"
		default:
			panic("invalid script parameter type")
		}
	}

	encoding.String(formatString).Encode(buf, 0)
	for i := len(struc.Params) - 1; i >= 0; i-- {
		p := struc.Params[i]
		switch p := p.(type) {
		case int:
			encoding.Uint32(p).Encode(buf, encoding.IntNilFlag)

		case string:
			encoding.String(p).Encode(buf, 0)

		}
	}

	encoding.Uint32(struc.ScriptID).Encode(buf, encoding.IntNilFlag)
}

// +gen define_outbound
type OutboundInitInterface protocol.OutboundInitInterface

func (struc OutboundInitInterface) Encode(buf io.Writer, flags interface{}) {
	frame := getPlayerData(struc.ProtoData).frame

	OutboundSetRootInterfaceDefinition.Pack(OutboundSetRootInterface{
		Frame: frame,
	}).Encode(buf, flags)

	OutboundSetInterfaceDefinition.Pack(OutboundSetInterface{
		RootID:      frame.Root,
		ChildID:     frame.ChatBox,
		InterfaceID: 162,
		Clickable:   true,
	}).Encode(buf, flags)

	OutboundSetInterfaceDefinition.Pack(OutboundSetInterface{
		RootID:      frame.Root,
		ChildID:     frame.ExpDisplay,
		InterfaceID: 163,
		Clickable:   true,
	}).Encode(buf, flags)

	OutboundSetInterfaceDefinition.Pack(OutboundSetInterface{
		RootID:      frame.Root,
		ChildID:     frame.DataOrbs,
		InterfaceID: 160,
		Clickable:   true,
	}).Encode(buf, flags)

	OutboundSetInterfaceDefinition.Pack(OutboundSetInterface{
		RootID:      frame.Root,
		ChildID:     frame.PrivateChat,
		InterfaceID: 122,
		Clickable:   true,
	}).Encode(buf, flags)
}
