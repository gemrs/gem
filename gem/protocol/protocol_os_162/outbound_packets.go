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

func (o OutboundChatMessage) Encode(buf io.Writer, flags interface{}) {
	encoding.Uint16(0).Encode(buf, encoding.IntPacked)
	encoding.Uint8(0).Encode(buf, encoding.IntNilFlag)
	encoding.String(o.Message).Encode(buf, 0)
}

// +gen define_outbound:"Unimplemented"
type OutboundCreateGlobalGroundItem protocol.OutboundCreateGlobalGroundItem

func (struc OutboundCreateGlobalGroundItem) Encode(buf io.Writer, flags interface{}) {
}

// +gen define_outbound:"Pkt46,SzFixed"
type OutboundCreateGroundItem protocol.OutboundCreateGroundItem

func (struc OutboundCreateGroundItem) Encode(buf io.Writer, flags interface{}) {
	encoding.Uint16(struc.Count).Encode(buf, encoding.IntLittleEndian)
	encoding.Uint8(struc.PositionOffset).Encode(buf, encoding.IntNilFlag)
	encoding.Uint16(struc.ItemID).Encode(buf, encoding.IntLittleEndian)
}

// +gen define_outbound:"Pkt35,SzFixed"
type OutboundRemoveGroundItem protocol.OutboundRemoveGroundItem

func (struc OutboundRemoveGroundItem) Encode(buf io.Writer, flags interface{}) {
	encoding.Uint8(struc.PositionOffset).Encode(buf, encoding.IntInverse128)
	encoding.Uint16(struc.ItemID).Encode(buf, encoding.IntLittleEndian|encoding.IntOffset128)
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

// +gen define_outbound:"Pkt13,SzVar16"
type OutboundUpdateInventoryItem protocol.OutboundUpdateInventoryItem

func (struc OutboundUpdateInventoryItem) Encode(buf io.Writer, flags interface{}) {
	inventory := struc.Container
	root, child, iface := inventory.InterfaceLocation()

	encoding.Uint32(root<<16|child).Encode(buf, encoding.IntNilFlag)
	encoding.Uint16(iface).Encode(buf, encoding.IntNilFlag)

	encoding.Uint16(struc.Slot).Encode(buf, encoding.IntPacked)
	if !inventory.SlotPopulated(struc.Slot) {
		encoding.Uint16(0).Encode(buf, encoding.IntNilFlag)
	} else {
		stack := inventory.Slot(struc.Slot)
		encoding.Uint16(stack.Definition().Id()+1).Encode(buf, encoding.IntNilFlag)

		count := stack.Count()
		if count > 255 {
			encoding.Uint8(255).Encode(buf, encoding.IntNilFlag)
			encoding.Uint32(count).Encode(buf, encoding.IntNilFlag)
		} else if count > 0 {
			encoding.Uint8(count).Encode(buf, encoding.IntNilFlag)
		}
	}
}

// +gen define_outbound:"Pkt62,SzVar16"
type OutboundUpdateAllInventoryItems protocol.OutboundUpdateAllInventoryItems

func (struc OutboundUpdateAllInventoryItems) Encode(buf io.Writer, flags interface{}) {
	inventory := struc.Container
	root, child, iface := inventory.InterfaceLocation()

	encoding.Uint32(root<<16|child).Encode(buf, encoding.IntNilFlag)
	encoding.Uint16(iface).Encode(buf, encoding.IntNilFlag)

	cap := inventory.Capacity()
	encoding.Uint16(cap).Encode(buf, encoding.IntNilFlag)

	for i := 0; i < cap; i++ {
		if !inventory.SlotPopulated(i) {
			encoding.Uint8(0).Encode(buf, encoding.IntNilFlag)
			encoding.Uint16(0).Encode(buf, encoding.IntLittleEndian)
		} else {
			stack := inventory.Slot(i)
			count := stack.Count()
			if count > 255 {
				encoding.Uint8(255).Encode(buf, encoding.IntNilFlag)
				encoding.Uint32(count).Encode(buf, encoding.IntPDPEndian)
			} else if count > 0 {
				encoding.Uint8(count).Encode(buf, encoding.IntNilFlag)
			}

			encoding.Uint16(stack.Definition().Id()+1).Encode(buf, encoding.IntOffset128)
		}
	}
}

// +gen define_outbound:"Unimplemented"
type OutboundLogout protocol.OutboundLogout

func (struc OutboundLogout) Encode(buf io.Writer, flags interface{}) {

}

// +gen define_outbound:"Unimplemented"
type OutboundPlayerInit protocol.OutboundPlayerInit

func (struc OutboundPlayerInit) Encode(buf io.Writer, flags interface{}) {
}

// +gen define_outbound:"Pkt39,SzVar16"
type OutboundRegionUpdate protocol.OutboundRegionUpdate

func (struc OutboundRegionUpdate) Encode(buf io.Writer, flags interface{}) {
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
	encoding.Uint16(sectorY).Encode(buf, encoding.IntLittleEndian)
	encoding.Uint16(sectorX).Encode(buf, encoding.IntOffset128|encoding.IntLittleEndian)

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

	encoding.Uint16(count).Encode(buf, encoding.IntNilFlag)
	for _, key := range allKeys {
		encoding.Uint32(key).Encode(buf, encoding.IntNilFlag)
	}
}

// +gen define_outbound:"Pkt0,SzFixed"
type OutboundSetUpdatingTile protocol.OutboundSetUpdatingTile

func (struc OutboundSetUpdatingTile) Encode(buf io.Writer, flags interface{}) {
	encoding.Uint8(struc.PositionY).Encode(buf, encoding.IntInverse128)
	encoding.Uint8(struc.PositionX).Encode(buf, encoding.IntNegate)
}

// +gen define_outbound:"Pkt37,SzFixed"
type OutboundSkill protocol.OutboundSkill

func (struc OutboundSkill) Encode(buf io.Writer, flags interface{}) {
	encoding.Uint8(struc.Skill).Encode(buf, encoding.IntNegate)
	encoding.Uint8(struc.Level).Encode(buf, encoding.IntNilFlag)
	encoding.Uint32(struc.Experience).Encode(buf, encoding.IntRPDPEndian)
}

// +gen define_outbound:"Unimplemented"
type OutboundResetCamera protocol.OutboundResetCamera

func (struc OutboundResetCamera) Encode(buf io.Writer, flags interface{}) {
}

// +gen define_outbound:"Unimplemented"
type OutboundDnsLookup protocol.OutboundDnsLookup

func (struc OutboundDnsLookup) Encode(buf io.Writer, flags interface{}) {
	// FIXME: encode address properly
	encoding.Uint32(0).Encode(buf, encoding.IntNilFlag)
}

// +gen define_outbound:"Pkt2,SzFixed"
type OutboundSetRootInterface protocol.OutboundSetRootInterface

func (struc OutboundSetRootInterface) Encode(buf io.Writer, flags interface{}) {
	frame := struc.Frame.(FrameType)
	encoding.Uint16(frame.Root).Encode(buf, encoding.IntNilFlag)
}

// +gen define_outbound:"Pkt29,SzFixed"
type OutboundSetInterface protocol.OutboundSetInterface

func (struc OutboundSetInterface) Encode(buf io.Writer, flags interface{}) {
	clickable := 0
	if struc.Clickable {
		clickable = 1
	}

	encoding.Uint8(clickable).Encode(buf, encoding.IntOffset128)
	encoding.Uint32((struc.RootID<<16)|struc.ChildID).Encode(buf, encoding.IntRPDPEndian)
	encoding.Uint16(struc.InterfaceID).Encode(buf, encoding.IntNilFlag)
}

// +gen define_outbound:"Pkt18,SzVar16"
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

	//	struc.Inventory.SetInterfaceContainer()
}
