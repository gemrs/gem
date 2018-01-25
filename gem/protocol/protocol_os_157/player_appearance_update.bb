/* The appearance block of the player update packet */
type OutboundPlayerAppearance struct {
    Gender   uint8
    HeadIcon uint8

    HelmModel       uint8
    CapeModel       uint8
    AmuletModel     uint8
    RightWieldModel uint8
    TorsoModel      uint16
    LeftWieldModel  uint8
    ArmsModel       uint16
    LegsModel       uint16
    HeadModel       uint16
    HandsModel      uint16
    FeetModel       uint16
    BeardModel      uint16

    HairColor  uint8
    TorsoColor uint8
    LegColor   uint8
    FeetColor  uint8
    SkinColor  uint8

   	AnimIdle       uint16
  	AnimSpotRotate uint16
  	AnimWalk       uint16
  	AnimRotate180  uint16
  	AnimRotateCCW  uint16
  	AnimRotateCW   uint16
   	AnimRun        uint16

    NameHash    uint64
    CombatLevel uint8
    SkillLevel  uint16
}
