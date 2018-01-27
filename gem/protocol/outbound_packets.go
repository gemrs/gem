package protocol

import "github.com/gemrs/gem/gem/game/position"

// +gen pack_outgoing
type OutboundLoginResponse struct {
	Index    int
	Response AuthResponse
	Rights   int
	Flagged  bool
}

// +gen pack_outgoing
type OutboundChatMessage struct {
	Message string
}

// +gen pack_outgoing
type OutboundRegionUpdate struct {
	Position    *position.Absolute
	PlayerIndex int
}

// +gen pack_outgoing
type OutboundSkill struct {
	Skill      int
	Experience int
	Level      int
}

// +gen pack_outgoing
type OutboundTabInterface struct {
	ProtoData interface{}

	Tab         InterfaceTab
	InterfaceID int
}

// +gen pack_outgoing
type OutboundLogout struct{}

// +gen pack_outgoing
type OutboundPlayerChatMessage struct {
	Effects       int
	Colour        int
	Rights        int
	Length        int
	PackedMessage []byte
}

// +gen pack_outgoing
type OutboundSetText struct {
	Text string
	Id   int
}

// +gen pack_outgoing
type OutboundPlayerInit struct {
	Membership int
	Index      int
}

// +gen pack_outgoing
type OutboundUpdateInventoryItem struct {
	InventoryID int
	Slot        int
	ItemID      int
	Count       int
}

// +gen pack_outgoing
type OutboundCreateGlobalGroundItem struct {
	ItemID         int
	PositionOffset int
	Index          int
	Count          int
}

// +gen pack_outgoing
type OutboundCreateGroundItem struct {
	ItemID         int
	Count          int
	PositionOffset int
}

// +gen pack_outgoing
type OutboundRemoveGroundItem struct {
	PositionOffset int
	ItemID         int
}

// +gen pack_outgoing
type OutboundSectorUpdate struct {
	PositionX int
	PositionY int
}

// +gen pack_outgoing
type OutboundResetCamera struct {
}

// +gen pack_outgoing
type OutboundDnsLookup struct {
	Host string
}

// +gen pack_outgoing
type OutboundSetRootInterface struct {
	Frame FrameType
}

// +gen pack_outgoing
type OutboundSetInterface struct {
	RootID      int
	ChildID     int
	InterfaceID int
	Clickable   bool
}

// +gen pack_outgoing
type OutboundScriptEvent struct {
	ScriptID int
	Params   []interface{}
}

// +gen pack_outgoing
type OutboundInitInterface struct {
	ProtoData interface{}
}
