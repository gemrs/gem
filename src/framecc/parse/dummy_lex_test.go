package parse

import (
	"testing"
	"encoding/json"

	"framecc/ast"
)

var lexers = []*dummyLexer{
	&dummyLexer{
		file: &ast.File{Name:"test_1"},
		tokens: []token{
			{tIdentifier, yySymType{sval:"SomeStruct"}},
			{tStruct, yySymType{sval:""}},
			{'{', yySymType{sval:""}},
			{tIdentifier, yySymType{sval:"SomeInt"}},
			{tIntegerType, yySymType{sval:"int8"}},
			{'}', yySymType{sval:""}},
		},
	},
}

func TestParser(t *testing.T) {
	yyDebug = 1
	yyErrorVerbose = true
	for _, lexer := range lexers {
		_ = yyParse(lexer)
		astStr, _ := json.Marshal(lexer.file)
		t.Logf("%v", string(astStr))
	}
}
