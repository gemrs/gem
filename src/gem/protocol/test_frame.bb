type EmbeddedStruct struct {
    A int32<IntLittleEndian>
    B int32<IntPDPEndian, IntOffset128>
    C int32<IntRPDPEndian, IntInverse128>
}

type TestFrame struct {
    Message  string[16]
    Values8  int8<IntNegate>[4]
    Values16 int16[2]
    Struc1   EmbeddedStruct
    Struc2   EmbeddedStruct[2]
    block    bitstruct {
        X bit[2]
        Y bit
    }
}
