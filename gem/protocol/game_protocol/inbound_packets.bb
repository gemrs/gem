/* Periodically sent to tell us that the client is still alive.. */
type InboundPing frame<0, Fixed> struct {}

/* The player has entered a public chat message */
type InboundChatMessage frame<4, Var8> struct {
    Effects uint8<IntOffset128, IntReverse>
    Colour uint8<IntOffset128, IntReverse>
    EncodedMessage byte[...]
}

/* The player has entered a command using the '::' syntax */
type InboundCommand frame<103, Var8> struct {
    Command string
}

/* The player has swapped two slots in an inventory interface */
type InboundInventorySwapItem frame<214, Fixed> struct {
	InterfaceID uint16<IntLittleEndian, IntOffset128>
	Inserting uint8<IntNegate>
	FromSlot uint16<IntLittleEndian, IntOffset128>
	ToSlot uint16<IntLittleEndian>
}
