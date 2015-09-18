package parse

import (
	"testing"
	"encoding/json"

	_ "framecc/ast"
)

type testCase struct {
	source string
}

var tests = []testCase{
	{
	source: `SomeStruct struct {
	SomeInt int8
}`,
	},
}

func TestParser(t *testing.T) {
	yyDebug = 1
	yyErrorVerbose = true
	for _, tc := range tests {
		ast := Parse("in_file", tc.source)
		astStr, _ := json.Marshal(ast)
		t.Logf("%v", string(astStr))
	}
}
