package parse

import (
	"testing"
	"encoding/json"

	"github.com/sinusoids/gem/framecc/ast"
)

type parseTestCase struct {
	input string
	ast *ast.File
}

var parseTestCases = []parseTestCase{
	parseTestCase{
		input: `/* Variable length (8 bit encoded) frame */
VariableChatMessage struct {
    Length  int16<LittleEndian, Offset128>
    Message string[Length]
}`,
		ast: nil,
	},
	parseTestCase{
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
		ast: nil,
	},
}

func TestParse(t *testing.T) {
	for i, tc := range parseTestCases {
		t.Logf("-- Test case %v\n", i)
		ast, errors := Parse("test.frame", tc.input)
		astStr, _ := json.Marshal(ast)
		t.Logf("%v", string(astStr))
		t.Log(errors)
	}
}
