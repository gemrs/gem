/* Puts a message into the client's chat window.
   This is not the player chat message */
type OutboundChatMessage frame<253, Var8> struct {
    Message string
}

/* Loads the region centered at a given sector */
type OutboundRegionUpdate frame<73, Fixed> struct {
    SectorX int16<IntOffset128>
    SectorY int16
}