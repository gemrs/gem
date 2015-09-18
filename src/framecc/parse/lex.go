package parse

import (
	"io"
	"log"

	"framecc/ast"
)

/* wraps the nex generated Lexer */
type lexer struct {
	file *ast.File
	*Lexer
}

func newLexer(filename string, in io.Reader) *lexer {
	return &lexer{
		file: ast.NewFile(filename),
		Lexer: NewLexer(in),
	}
}

func (l *lexer) Error(e string) {
    log.Fatal(e)
}
