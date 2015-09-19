/* Fixed length frame */
type ChatMessage frame<100, Var8> struct {
    Message string
}

/* Variable length (8 bit encoded) frame */
type VariableChatMessage struct {
    Length  int16<LittleEndian, Offset128>
    Message string[Length]
}

type ChatMessage1 frame<101, Var8>  VariableChatMessage
type ChatMessage2 frame<102, Var16> VariableChatMessage

/* Embedding structures */
type AppearanceBlock struct {
    Legs   int8
    Body   int8
    Helmet int8
}

type PositionBlock struct {
    X int16
    Y int16
    Z int16
}

type OtherPlayer struct {
    Name       string[32]
    Appearance AppearanceBlock
    Position   PositionBlock
}

type PlayerUpdate frame<200, Var16> struct {
	Magic       int8
    UpdateFlag  int8
    OthersCount int8
    Others      OtherPlayer[OthersCount]
    Appearance  AppearanceBlock
    Position    PositionBlock
    Skills      struct {
        Overall int16
        Skills  int8[20]
    }
}
