/* Loads the region centered at a given sector */
type OutboundRegionUpdate frame<73, Fixed> struct {
    SectorX uint16<IntOffset128>
    SectorY uint16
}

/* Puts a message into the client's chat window.
   This is not the player chat message */
type OutboundChatMessage frame<253, Var8> struct {
    Message string
}

/* Tells the client about it's player on login */
type OutboundPlayerInit frame<249, Fixed> struct {
    Membership uint8<IntOffset128>
    Index      uint16<IntOffset128>
}
