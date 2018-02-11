package protocol

type InboundPing struct{}

type InboundMouseMovement struct{}

type InboundMouseClick struct{}

type InboundWindowFocus struct{}

type InboundKeyPress struct{}

type InboundCameraMovement struct{}

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

type InboundWidgetAction struct {
	Action      int
	InterfaceID int
	WidgetID    int
	Param       int
	ItemID      int
}

type InboundGroundItemAction struct {
	Action int
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
	X       int
	Y       int
	Running bool
}

type UnknownPacket struct {
	Number   int
	SizeType int
	Size     int
	Payload  []byte
}
