// Code generated by gem_gen; DO NOT EDIT.
package protocol_317

import (
	"fmt"

	"github.com/gemrs/gem/gem/core/encoding"
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/protocol"
)

var InboundPlayerWalkMapDefinition = InboundPacketDefinition{
	Number: 248,
	Size:   SzVar8,
}

var OutboundLogoutDefinition = OutboundPacketDefinition{
	Number: 109,
	Size:   SzFixed,
}

var InboundCommandDefinition = InboundPacketDefinition{
	Number: 103,
	Size:   SzVar8,
}

var InboundTakeGroundItemDefinition = InboundPacketDefinition{
	Number: 236,
	Size:   SzFixed,
}

var InboundPlayerWalkDefinition = InboundPacketDefinition{
	Number: 164,
	Size:   SzVar8,
}

var PlayerUpdateDefinition = OutboundPacketDefinition{
	Number: 81,
	Size:   SzVar16,
}

var OutboundChatMessageDefinition = OutboundPacketDefinition{
	Number: 253,
	Size:   SzVar8,
}

var OutboundCreateGroundItemDefinition = OutboundPacketDefinition{
	Number: 44,
	Size:   SzFixed,
}

var OutboundPlayerInitDefinition = OutboundPacketDefinition{
	Number: 249,
	Size:   SzFixed,
}

var InboundPingDefinition = InboundPacketDefinition{
	Number: 0,
	Size:   SzFixed,
}

var InboundInventoryAction2Definition = InboundPacketDefinition{
	Number: 41,
	Size:   SzFixed,
}

var OutboundUpdateInventoryItemDefinition = OutboundPacketDefinition{
	Number: 34,
	Size:   SzVar16,
}

var OutboundSetTextDefinition = OutboundPacketDefinition{
	Number: 126,
	Size:   SzVar16,
}

var InboundInventoryAction4Definition = InboundPacketDefinition{
	Number: 75,
	Size:   SzFixed,
}

var InboundInventorySwapItemDefinition = InboundPacketDefinition{
	Number: 214,
	Size:   SzFixed,
}

var InboundPlayerWalkEntityDefinition = InboundPacketDefinition{
	Number: 98,
	Size:   SzVar8,
}

var OutboundSectorUpdateDefinition = OutboundPacketDefinition{
	Number: 85,
	Size:   SzFixed,
}

var OutboundSkillDefinition = OutboundPacketDefinition{
	Number: 134,
	Size:   SzFixed,
}

var OutboundCreateGlobalGroundItemDefinition = OutboundPacketDefinition{
	Number: 215,
	Size:   SzFixed,
}

var OutboundRemoveGroundItemDefinition = OutboundPacketDefinition{
	Number: 156,
	Size:   SzFixed,
}

var InboundInventoryAction5Definition = InboundPacketDefinition{
	Number: 87,
	Size:   SzFixed,
}

var InboundChatMessageDefinition = InboundPacketDefinition{
	Number: 4,
	Size:   SzVar8,
}

var OutboundTabInterfaceDefinition = OutboundPacketDefinition{
	Number: 71,
	Size:   SzFixed,
}

var OutboundRegionUpdateDefinition = OutboundPacketDefinition{
	Number: 73,
	Size:   SzFixed,
}

var InboundInventoryAction1Definition = InboundPacketDefinition{
	Number: 122,
	Size:   SzFixed,
}

var InboundInventoryAction3Definition = InboundPacketDefinition{
	Number: 16,
	Size:   SzFixed,
}

var inboundPacketBuilders = map[int]func() encoding.Decodable{

	248: func() encoding.Decodable {
		return &PacketHeader{
			Number: InboundPlayerWalkMapDefinition.Number,
			Size:   InboundPlayerWalkMapDefinition.Size,
			Object: new(InboundPlayerWalkMap),
		}
	},

	103: func() encoding.Decodable {
		return &PacketHeader{
			Number: InboundCommandDefinition.Number,
			Size:   InboundCommandDefinition.Size,
			Object: new(InboundCommand),
		}
	},

	236: func() encoding.Decodable {
		return &PacketHeader{
			Number: InboundTakeGroundItemDefinition.Number,
			Size:   InboundTakeGroundItemDefinition.Size,
			Object: new(InboundTakeGroundItem),
		}
	},

	164: func() encoding.Decodable {
		return &PacketHeader{
			Number: InboundPlayerWalkDefinition.Number,
			Size:   InboundPlayerWalkDefinition.Size,
			Object: new(InboundPlayerWalk),
		}
	},

	0: func() encoding.Decodable {
		return &PacketHeader{
			Number: InboundPingDefinition.Number,
			Size:   InboundPingDefinition.Size,
			Object: new(InboundPing),
		}
	},

	41: func() encoding.Decodable {
		return &PacketHeader{
			Number: InboundInventoryAction2Definition.Number,
			Size:   InboundInventoryAction2Definition.Size,
			Object: new(InboundInventoryAction2),
		}
	},

	75: func() encoding.Decodable {
		return &PacketHeader{
			Number: InboundInventoryAction4Definition.Number,
			Size:   InboundInventoryAction4Definition.Size,
			Object: new(InboundInventoryAction4),
		}
	},

	214: func() encoding.Decodable {
		return &PacketHeader{
			Number: InboundInventorySwapItemDefinition.Number,
			Size:   InboundInventorySwapItemDefinition.Size,
			Object: new(InboundInventorySwapItem),
		}
	},

	98: func() encoding.Decodable {
		return &PacketHeader{
			Number: InboundPlayerWalkEntityDefinition.Number,
			Size:   InboundPlayerWalkEntityDefinition.Size,
			Object: new(InboundPlayerWalkEntity),
		}
	},

	87: func() encoding.Decodable {
		return &PacketHeader{
			Number: InboundInventoryAction5Definition.Number,
			Size:   InboundInventoryAction5Definition.Size,
			Object: new(InboundInventoryAction5),
		}
	},

	4: func() encoding.Decodable {
		return &PacketHeader{
			Number: InboundChatMessageDefinition.Number,
			Size:   InboundChatMessageDefinition.Size,
			Object: new(InboundChatMessage),
		}
	},

	122: func() encoding.Decodable {
		return &PacketHeader{
			Number: InboundInventoryAction1Definition.Number,
			Size:   InboundInventoryAction1Definition.Size,
			Object: new(InboundInventoryAction1),
		}
	},

	16: func() encoding.Decodable {
		return &PacketHeader{
			Number: InboundInventoryAction3Definition.Number,
			Size:   InboundInventoryAction3Definition.Size,
			Object: new(InboundInventoryAction3),
		}
	},
}

func (p protocolImpl) Decode(message encoding.Decodable) server.Message {
	switch message := message.(type) {

	case *InboundPlayerWalkMap:
		return (*protocol.InboundPlayerWalk)(message)

	case *InboundCommand:
		return (*protocol.InboundCommand)(message)

	case *InboundTakeGroundItem:
		return (*protocol.InboundTakeGroundItem)(message)

	case *InboundPlayerWalk:
		return (*protocol.InboundPlayerWalk)(message)

	case *InboundPing:
		return (*protocol.InboundPing)(message)

	case *InboundInventoryAction2:
		return (*protocol.InboundInventoryAction)(message)

	case *InboundInventoryAction4:
		return (*protocol.InboundInventoryAction)(message)

	case *InboundInventorySwapItem:
		return (*protocol.InboundInventorySwapItem)(message)

	case *InboundPlayerWalkEntity:
		return (*protocol.InboundPlayerWalk)(message)

	case *InboundInventoryAction5:
		return (*protocol.InboundInventoryAction)(message)

	case *InboundChatMessage:
		return (*protocol.InboundChatMessage)(message)

	case *InboundInventoryAction1:
		return (*protocol.InboundInventoryAction)(message)

	case *InboundInventoryAction3:
		return (*protocol.InboundInventoryAction)(message)

	case *UnknownPacket:
		return (*protocol.UnknownPacket)(message)

	case *PacketHeader:
		return p.Decode(message.Object.(encoding.Decodable))
	}
	panic(fmt.Sprintf("cannot decode %T", message))
}

func (protocolImpl) Encode(message server.Message) encoding.Encodable {
	switch message := message.(type) {

	case protocol.OutboundLogout:
		return OutboundLogoutDefinition.Pack(OutboundLogout(message))

	case protocol.PlayerUpdate:
		return PlayerUpdateDefinition.Pack(PlayerUpdate(message))

	case protocol.OutboundChatMessage:
		return OutboundChatMessageDefinition.Pack(OutboundChatMessage(message))

	case protocol.OutboundCreateGroundItem:
		return OutboundCreateGroundItemDefinition.Pack(OutboundCreateGroundItem(message))

	case protocol.OutboundPlayerInit:
		return OutboundPlayerInitDefinition.Pack(OutboundPlayerInit(message))

	case protocol.OutboundUpdateInventoryItem:
		return OutboundUpdateInventoryItemDefinition.Pack(OutboundUpdateInventoryItem(message))

	case protocol.OutboundSetText:
		return OutboundSetTextDefinition.Pack(OutboundSetText(message))

	case protocol.OutboundSectorUpdate:
		return OutboundSectorUpdateDefinition.Pack(OutboundSectorUpdate(message))

	case protocol.OutboundSkill:
		return OutboundSkillDefinition.Pack(OutboundSkill(message))

	case protocol.OutboundCreateGlobalGroundItem:
		return OutboundCreateGlobalGroundItemDefinition.Pack(OutboundCreateGlobalGroundItem(message))

	case protocol.OutboundRemoveGroundItem:
		return OutboundRemoveGroundItemDefinition.Pack(OutboundRemoveGroundItem(message))

	case protocol.OutboundTabInterface:
		return OutboundTabInterfaceDefinition.Pack(OutboundTabInterface(message))

	case protocol.OutboundRegionUpdate:
		return OutboundRegionUpdateDefinition.Pack(OutboundRegionUpdate(message))

	case protocol.OutboundLoginResponse:
		return OutboundLoginResponse(message)

	}
	panic(fmt.Sprintf("cannot encode %T", message))
}
