package protocol

type InboundPing struct{}

type InboundCommand struct {
	Command string
}

type InboundInventoryAction struct {
	Action      int
	ItemID      int
	InterfaceID int
	Slot        int
}

type InboundChatMessage struct {
	Effects       int
	Colour        int
	Message       string
	PackedMessage []byte
}

type InboundTakeGroundItem struct {
	X, Y   int
	ItemID int
}

type InboundInventorySwapItem struct {
	InterfaceID int
	Inserting   bool
	From, To    int
}

type WalkWaypoint struct {
	X, Y int
}

type InboundPlayerWalk struct {
	OriginX   int
	OriginY   int
	Waypoints []WalkWaypoint
	Running   bool
}

type UnknownPacket struct {
	Number   int
	SizeType int
	Size     int
	Payload  []byte
}
