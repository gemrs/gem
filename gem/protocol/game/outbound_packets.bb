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

/* Puts a chat message into the client's chat window. */
type OutboundPlayerChatMessage struct {
    Effects       uint8
    Colour        uint8
    Rights        uint8
    Length        uint8<IntNegate>
    PackedMessage byte[...]
}

/* Set tab interface */
type OutboundTabInterface frame<71, Fixed> struct {
	InterfaceID uint16
	Tab         uint8<IntOffset128>
}

/* Set skill */
type OutboundSkill frame<134, Fixed> struct {
	Skill      uint8
	Experience uint32<IntPDPEndian>
	Level      uint8
}

/* Force player disconnect */
type OutboundLogout frame<109, Fixed> struct {

}

/* Set interface text */
type OutboundSetText frame<126, Var16> struct {
	Text string
	Id uint16<IntLittleEndian, IntInverse128>
}

/* Tells the client about it's player on login */
type OutboundPlayerInit frame<249, Fixed> struct {
    Membership uint8<IntOffset128>
    Index      uint16<IntOffset128>
}
