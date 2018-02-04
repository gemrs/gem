package protocol_os_162

import (
	"fmt"
	"io"

	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/game/data"
	"github.com/gemrs/gem/gem/protocol"
)

// +gen define_outbound:"Pkt52,SzVar8"
type OutboundChatMessage protocol.OutboundChatMessage

func (o OutboundChatMessage) Encode(w io.Writer, flags interface{}) {
	buf := encoding.WrapWriter(w)
	buf.PutU16(0, encoding.IntPacked)
	buf.PutU8(0)
	buf.PutStringZ(o.Message)
}

// +gen define_outbound:"Unimplemented"
type OutboundCreateGlobalGroundItem protocol.OutboundCreateGlobalGroundItem

func (struc OutboundCreateGlobalGroundItem) Encode(w io.Writer, flags interface{}) {
}

// +gen define_outbound:"Pkt46,SzFixed"
type OutboundCreateGroundItem protocol.OutboundCreateGroundItem

func (struc OutboundCreateGroundItem) Encode(w io.Writer, flags interface{}) {
	buf := encoding.WrapWriter(w)
	buf.PutU16(struc.Count, encoding.IntLittleEndian)
	buf.PutU8(struc.PositionOffset)
	buf.PutU16(struc.ItemID, encoding.IntLittleEndian)
}

// +gen define_outbound:"Pkt35,SzFixed"
type OutboundRemoveGroundItem protocol.OutboundRemoveGroundItem

func (struc OutboundRemoveGroundItem) Encode(w io.Writer, flags interface{}) {
	buf := encoding.WrapWriter(w)
	buf.PutU8(struc.PositionOffset, encoding.IntInverse128)
	buf.PutU16(struc.ItemID, encoding.IntLittleEndian, encoding.IntOffset128)
}

// +gen define_outbound
type OutboundTabInterface protocol.OutboundTabInterface

func (struc OutboundTabInterface) Encode(w io.Writer, flags interface{}) {
	frame := getPlayerData(struc.ProtoData).frame
	packet := OutboundSetInterfaceDefinition.Pack(OutboundSetInterface{
		RootID:      frame.Root,
		ChildID:     frame.Tabs[struc.Tab],
		InterfaceID: struc.InterfaceID,
		Clickable:   true,
	})
	packet.Encode(w, flags)
}

// +gen define_outbound:"Unimplemented"
type OutboundSetText protocol.OutboundSetText

func (struc OutboundSetText) Encode(w io.Writer, flags interface{}) {
}

// +gen define_outbound:"Pkt13,SzVar16"
type OutboundUpdateInventoryItem protocol.OutboundUpdateInventoryItem

func (struc OutboundUpdateInventoryItem) Encode(w io.Writer, flags interface{}) {
	buf := encoding.WrapWriter(w)

	inventory := struc.Container
	root, child, iface := inventory.InterfaceLocation()

	buf.PutU32(root<<16 | child)
	buf.PutU16(iface)

	buf.PutU16(struc.Slot, encoding.IntPacked)

	if !inventory.SlotPopulated(struc.Slot) {
		buf.PutU16(0)
	} else {
		stack := inventory.Slot(struc.Slot)
		buf.PutU16(stack.Definition().Id() + 1)

		count := stack.Count()
		if count > 255 {
			buf.PutU8(255)
			buf.PutU32(count)
		} else if count > 0 {
			buf.PutU8(count)
		}
	}
}

// +gen define_outbound:"Pkt62,SzVar16"
type OutboundUpdateAllInventoryItems protocol.OutboundUpdateAllInventoryItems

func (struc OutboundUpdateAllInventoryItems) Encode(w io.Writer, flags interface{}) {
	buf := encoding.WrapWriter(w)

	inventory := struc.Container
	root, child, iface := inventory.InterfaceLocation()

	buf.PutU32(root<<16 | child)
	buf.PutU16(iface)

	cap := inventory.Capacity()
	buf.PutU16(cap, encoding.IntPacked)

	for i := 0; i < cap; i++ {
		if !inventory.SlotPopulated(i) {
			buf.PutU8(0)
			buf.PutU16(0, encoding.IntLittleEndian)
		} else {
			stack := inventory.Slot(i)
			count := stack.Count()
			if count > 255 {
				buf.PutU8(255)
				buf.PutU32(count, encoding.IntPDPEndian)
			} else if count > 0 {
				buf.PutU8(count)
			}

			buf.PutU16(stack.Definition().Id()+1, encoding.IntOffset128)
		}
	}
}

// +gen define_outbound:"Unimplemented"
type OutboundLogout protocol.OutboundLogout

func (struc OutboundLogout) Encode(w io.Writer, flags interface{}) {

}

// +gen define_outbound:"Unimplemented"
type OutboundPlayerInit protocol.OutboundPlayerInit

func (struc OutboundPlayerInit) Encode(w io.Writer, flags interface{}) {
}

// +gen define_outbound:"Pkt39,SzVar16"
type OutboundRegionUpdate protocol.OutboundRegionUpdate

func (struc OutboundRegionUpdate) Encode(w io.Writer, flags interface{}) {
	buf := encoding.WrapWriter(w)

	pdata := getPlayerData(struc.ProtoData)
	pos := struc.Player.Position()
	playerIndex := struc.Player.Index()

	if !pdata.playerIndexInitialized {
		pdata.playerIndexInitialized = true

		compressedPos := pos.Y() + (pos.X() << 14) + (pos.Z() << 28)
		bitBuf := encoding.NewBitBuffer(buf)
		bitBuf.Write(30, uint32(compressedPos))

		pdata.localPlayers[pdata.localPlayerCount] = playerIndex
		pdata.localPlayerCount++

		for i := 1; i < protocol.MaxPlayers; i++ {
			if i != playerIndex {
				bitBuf.Write(18, 0)
				pdata.externalPlayers[pdata.externalPlayerCount] = i
				pdata.externalPlayerCount++
			}
		}

		bitBuf.Close()
	}

	sector := pos.Sector()
	sectorX := sector.X()
	sectorY := sector.Y()
	buf.PutU16(sectorY, encoding.IntLittleEndian)
	buf.PutU16(sectorX, encoding.IntOffset128, encoding.IntLittleEndian)

	regionX, regionY := sectorX/8, sectorY/8
	tutorialIsland := false
	if (regionX == 48 || regionX == 49) && regionY == 48 {
		tutorialIsland = true
	}

	if regionX == 48 && regionY == 148 {
		tutorialIsland = true
	}

	count := 0
	allKeys := make([]int, 0)
	for x := (sectorX - 6) / 8; x <= (sectorX+6)/8; x++ {
		for y := (sectorY - 6) / 8; y <= (sectorY+6)/8; y++ {
			if !tutorialIsland || y != 49 && y != 149 && y != 147 && x != 50 && (x != 49 || y != 47) {
				region := y + (x << 8)
				keys, ok := data.GetMapKeys(region)
				if !ok {
					panic(fmt.Errorf("don't have map keys for region %v", region))
				}
				for _, key := range keys {
					allKeys = append(allKeys, key)
				}
				count++
			}
		}
	}

	buf.PutU16(count)
	for _, key := range allKeys {
		buf.PutU32(key)
	}
}

// +gen define_outbound:"Pkt0,SzFixed"
type OutboundSetUpdatingTile protocol.OutboundSetUpdatingTile

func (struc OutboundSetUpdatingTile) Encode(w io.Writer, flags interface{}) {
	buf := encoding.WrapWriter(w)
	buf.PutU8(struc.PositionY, encoding.IntInverse128)
	buf.PutU8(struc.PositionX, encoding.IntNegate)
}

// +gen define_outbound:"Pkt37,SzFixed"
type OutboundSkill protocol.OutboundSkill

func (struc OutboundSkill) Encode(w io.Writer, flags interface{}) {
	buf := encoding.WrapWriter(w)
	buf.PutU8(struc.Skill, encoding.IntNegate)
	buf.PutU8(struc.Level)
	buf.PutU32(struc.Experience, encoding.IntRPDPEndian)
}

// +gen define_outbound:"Unimplemented"
type OutboundResetCamera protocol.OutboundResetCamera

func (struc OutboundResetCamera) Encode(w io.Writer, flags interface{}) {
}

// +gen define_outbound:"Unimplemented"
type OutboundDnsLookup protocol.OutboundDnsLookup

func (struc OutboundDnsLookup) Encode(w io.Writer, flags interface{}) {
	buf := encoding.WrapWriter(w)
	// FIXME: encode address properly
	buf.PutU32(0)
}

// +gen define_outbound:"Pkt2,SzFixed"
type OutboundSetRootInterface protocol.OutboundSetRootInterface

func (struc OutboundSetRootInterface) Encode(w io.Writer, flags interface{}) {
	buf := encoding.WrapWriter(w)
	frame := struc.Frame.(FrameType)
	buf.PutU16(frame.Root)
}

// +gen define_outbound:"Pkt29,SzFixed"
type OutboundSetInterface protocol.OutboundSetInterface

func (struc OutboundSetInterface) Encode(w io.Writer, flags interface{}) {
	buf := encoding.WrapWriter(w)
	clickable := 0
	if struc.Clickable {
		clickable = 1
	}

	buf.PutU8(clickable, encoding.IntOffset128)
	buf.PutU32((struc.RootID<<16)|struc.ChildID, encoding.IntRPDPEndian)
	buf.PutU16(struc.InterfaceID)
}

// +gen define_outbound:"Pkt18,SzVar16"
type OutboundScriptEvent protocol.OutboundScriptEvent

func (struc OutboundScriptEvent) Encode(w io.Writer, flags interface{}) {
	buf := encoding.WrapWriter(w)
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

	buf.PutStringZ(formatString)
	for i := len(struc.Params) - 1; i >= 0; i-- {
		p := struc.Params[i]
		switch p := p.(type) {
		case int:
			buf.PutU32(p)

		case string:
			buf.PutStringZ(p)

		}
	}

	buf.PutU32(struc.ScriptID)
}

// +gen define_outbound
type OutboundInitInterface protocol.OutboundInitInterface

func (struc OutboundInitInterface) Encode(w io.Writer, flags interface{}) {
	frame := getPlayerData(struc.ProtoData).frame

	OutboundSetRootInterfaceDefinition.Pack(OutboundSetRootInterface{
		Frame: frame,
	}).Encode(w, flags)

	OutboundSetInterfaceDefinition.Pack(OutboundSetInterface{
		RootID:      frame.Root,
		ChildID:     frame.ChatBox,
		InterfaceID: 162,
		Clickable:   true,
	}).Encode(w, flags)

	OutboundSetInterfaceDefinition.Pack(OutboundSetInterface{
		RootID:      frame.Root,
		ChildID:     frame.ExpDisplay,
		InterfaceID: 163,
		Clickable:   true,
	}).Encode(w, flags)

	OutboundSetInterfaceDefinition.Pack(OutboundSetInterface{
		RootID:      frame.Root,
		ChildID:     frame.DataOrbs,
		InterfaceID: 160,
		Clickable:   true,
	}).Encode(w, flags)

	OutboundSetInterfaceDefinition.Pack(OutboundSetInterface{
		RootID:      frame.Root,
		ChildID:     frame.PrivateChat,
		InterfaceID: 122,
		Clickable:   true,
	}).Encode(w, flags)

	//	struc.Inventory.SetInterfaceContainer()
}
