package parse

import (
	"bytes"
	"log"

	"framecc/ast"
)

//go:generate go tool yacc -o binbuf.y.go binbuf.y
//go:generate nex -e binbuf.nex

func (l *Lexer) Error(e string) {
    log.Fatal(e)
}

func Parse(filename, source string) *ast.File {
	lexer := NewLexerWithInit(bytes.NewBufferString(source), func (lex *Lexer) {
		lex.parseResult = ast.NewFile(filename)
	})
	_ = yyParse(lexer)
	return lexer.parseResult.(*ast.File)
}
