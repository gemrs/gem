package parse

import (
	"bytes"

	"framecc/ast"
)

//go:generate go tool yacc -o binbuf.y.go binbuf.y
//go:generate nex -e binbuf.nex

func Parse(filename, source string) *ast.File {
	lexer := newLexer(filename, bytes.NewBufferString(source))
	_ = yyParse(lexer)
	return lexer.file
}
