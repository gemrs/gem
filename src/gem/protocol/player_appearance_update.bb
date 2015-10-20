/* The appearance block of the player update packet */
type OutboundPlayerAppearance struct {
    Gender   int8
    HeadIcon int8

    HelmModel       int8
    CapeModel       int8
    AmuletModel     int8
    RightWieldModel int8
    TorsoModel      int16
    LeftWieldModel  int8
    ArmsModel       int16
    LegsModel       int16
    HeadModel       int16
    HandsModel      int16
    FeetModel       int16
    BeardModel      int16

    HairColor  int8
    TorsoColor int8
    LegColor   int8
    FeetColor  int8
    SkinColor  int8

   	AnimIdle       int16
  	AnimSpotRotate int16
  	AnimWalk       int16
  	AnimRotate180  int16
  	AnimRotateCCW  int16
  	AnimRotateCW   int16
   	AnimRun        int16

    NameHash    int64
    CombatLevel int8
    SkillLevel  int16
}
