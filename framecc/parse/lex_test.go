package parse

import (
	"testing"
)

type lexTestCase struct {
	input string
	items []item
}

var lexTestCases = []lexTestCase{
	lexTestCase{
		input: `ChatMessage frame<100, Var8> struct`,
		items: []item{
			item{itemIdentifier, "ChatMessage", 0},
			item{itemWhiteSpace, " ", 0},
			item{itemFrame, "frame", 0},
			item{itemLeftAngleBrack, "<", 0},
			item{itemNumber, "100", 0},
			item{itemComma, ",", 0},
			item{itemWhiteSpace, " ", 0},
			item{itemLenSpec, "Var8", 0},
			item{itemRightAngleBrack, ">", 0},
			item{itemWhiteSpace, " ", 0},
			item{itemStruct, "struct", 0},
			item{itemEOF, "", 0},
		},
	},

	lexTestCase{
		input: `PlayerUpdate frame<200, Var16> struct {
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
}`,
		items: []item{
			item{itemIdentifier, "PlayerUpdate", 0},
			item{itemWhiteSpace, "", 0},
			item{itemFrame, "frame", 0},
			item{itemLeftAngleBrack, "<", 0},
			item{itemNumber, "200", 0},
			item{itemComma, ",", 0},
			item{itemWhiteSpace, "", 0},
			item{itemLenSpec, "Var16", 0},
			item{itemRightAngleBrack, ">", 0},
			item{itemWhiteSpace, "", 0},
			item{itemStruct, "struct", 0},
			item{itemWhiteSpace, "", 0},
			item{itemLeftBrack, "{", 0},
			item{itemEOL, "", 0},
			item{itemWhiteSpace, "", 0},
			item{itemIdentifier, "Magic", 0},
			item{itemWhiteSpace, "", 0},
			item{itemIntType, "int8", 0},
			item{itemEOL, "", 0},
			item{itemWhiteSpace, "", 0},
			item{itemIdentifier, "UpdateFlag", 0},
			item{itemWhiteSpace, "", 0},
			item{itemIntType, "int8", 0},
			item{itemEOL, "", 0},
			item{itemWhiteSpace, "", 0},
			item{itemIdentifier, "OthersCount", 0},
			item{itemWhiteSpace, "", 0},
			item{itemIntType, "int8", 0},
			item{itemEOL, "", 0},
			item{itemWhiteSpace, "", 0},
			item{itemIdentifier, "Others", 0},
			item{itemWhiteSpace, "", 0},
			item{itemIdentifier, "OtherPlayer", 0},
			item{itemLeftSquareBrack, "[", 0},
			item{itemIdentifier, "OthersCount", 0},
			item{itemRightSquareBrack, "]", 0},
			item{itemEOL, "", 0},
			item{itemWhiteSpace, "", 0},
			item{itemIdentifier, "Appearance", 0},
			item{itemWhiteSpace, "", 0},
			item{itemIdentifier, "AppearanceBlock", 0},
			item{itemEOL, "", 0},
			item{itemWhiteSpace, "", 0},
			item{itemIdentifier, "Position", 0},
			item{itemWhiteSpace, "", 0},
			item{itemIdentifier, "PositionBlock", 0},
			item{itemEOL, "", 0},
			item{itemWhiteSpace, "", 0},
			item{itemIdentifier, "Skills", 0},
			item{itemWhiteSpace, "", 0},
			item{itemStruct, "struct", 0},
			item{itemWhiteSpace, "", 0},
			item{itemLeftBrack, "{", 0},
			item{itemEOL, "", 0},
			item{itemWhiteSpace, "", 0},
			item{itemIdentifier, "Overall", 0},
			item{itemWhiteSpace, "", 0},
			item{itemIntType, "int16", 0},
			item{itemEOL, "", 0},
			item{itemWhiteSpace, "", 0},
			item{itemIdentifier, "Skills", 0},
			item{itemWhiteSpace, "", 0},
			item{itemIntType, "int8", 0},
			item{itemLeftSquareBrack, "[", 0},
			item{itemNumber, "20", 0},
			item{itemRightSquareBrack, "]", 0},
			item{itemEOL, "", 0},
			item{itemWhiteSpace, "", 0},
			item{itemRightBrack, "}", 0},
			item{itemEOL, "", 0},
			item{itemRightBrack, "}", 0},
			item{itemEOF, "", 0},
		},
	},

	lexTestCase{
		input: `/* Variable length (8 bit encoded) frame */
VariableChatMessage struct {
    Length  int16<LittleEndian, Offset128>
    Message string[Length]
}`,
		items: []item{
			item{itemComment, "/* Variable length (8 bit encoded) frame */", 0},
			item{itemWhiteSpace, "", 0},
			item{itemIdentifier, "VariableChatMessage", 0},
			item{itemWhiteSpace, "", 0},
			item{itemStruct, "struct", 0},
			item{itemWhiteSpace, "", 0},
			item{itemLeftBrack, "{", 0},
			item{itemEOL, "", 0},
			item{itemWhiteSpace, "", 0},
			item{itemIdentifier, "Length", 0},
			item{itemWhiteSpace, "", 0},
			item{itemIntType, "int16", 0},
			item{itemLeftAngleBrack, "<", 0},
			item{itemFlag, "LittleEndian", 0},
			item{itemComma, ",", 0},
			item{itemWhiteSpace, "", 0},
			item{itemFlag, "Offset128", 0},
			item{itemRightAngleBrack, ">", 0},
			item{itemEOL, "", 0},
			item{itemWhiteSpace, "", 0},
			item{itemIdentifier, "Message", 0},
			item{itemWhiteSpace, "", 0},
			item{itemStringType, "string", 0},
			item{itemLeftSquareBrack, "[", 0},
			item{itemIdentifier, "Length", 0},
			item{itemRightSquareBrack, "]", 0},
			item{itemEOL, "", 0},
			item{itemRightBrack, "}", 0},
			item{itemEOF, "", 0},
		},
	},
}

func TestLex(t *testing.T) {
	for i, tc := range lexTestCases {
		t.Logf("-- Test case %v", i)
		doTestLex(t, tc.input, tc.items)
	}
}

func doTestLex(t *testing.T, input string, expected []item) {
	lexer := lex("test.frame", input)
	i := 0
	for item := range lexer.items {
		if i > len(expected) {
			t.Errorf("not enough tokens")
			return
		}
		exp := expected[i]
		i = i + 1

		if item.typ != exp.typ {
			t.Logf("%v\n", item)
			t.Errorf("incorrect type. got %v, expected %v", item.typ, exp.typ)
		}
		if item.typ != itemEOL && item.typ != itemWhiteSpace && item.val != exp.val {
			t.Logf("%v\n", item)
			t.Errorf("incorrect value. got %v, expected %v", item.val, exp.val)
		}
	}
}
