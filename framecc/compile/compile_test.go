package compile

import (
	"testing"
)

type compileTestCase struct {
	input  string
	output string
}

//TODO: Proper test cases
var compileTestCases = []compileTestCase{
	compileTestCase{
		input: `/* Variable length (8 bit encoded) frame */
VariableChatMessage struct {
    Length  int16<LittleEndian, Offset128>
    Message string[Length]
}`,
		output: "",
	},
}

func TestCompile(t *testing.T) {
	for i, tc := range compileTestCases {
		t.Logf("-- Test case %v\n", i)
		compiled, err := Compile("test.frame", tc.input)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%v", string(compiled))
	}
}
