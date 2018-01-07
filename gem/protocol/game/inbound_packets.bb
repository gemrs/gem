/* Periodically sent to tell us that the client is still alive.. */
type InboundPing frame<0, Fixed> struct {}

/* The player has entered a public chat message */
type InboundChatMessage frame<4, Var8> struct {
    Effects uint8<IntOffset128, IntReverse>
    Colour uint8<IntOffset128, IntReverse>
    EncodedMessage byte[...]
}