package parse

import (
	"bytes"
	"strings"

	"framecc/ast"
)

//go:generate go tool yacc -o binbuf.y.go binbuf.y
//go:generate nex -e binbuf.nex

type ErrorList []string

func (list ErrorList) Add(err string) ErrorList {
	return append(list, err)
}

func (list ErrorList) String() string {
	return strings.Join(list, "\n")
}

type result struct {
	Ast *ast.File
	Errors ErrorList
}

func (res *result) AddError(err string) {
	res.Errors = res.Errors.Add(err)
}

func (l *Lexer) Ast() *ast.File {
	result := l.parseResult.(*result)
    return result.Ast
}

func (l *Lexer) Error(e string) {
	result := l.parseResult.(*result)
	result.AddError(e)
}

func Parse(filename, source string) (*ast.File, ErrorList) {
	result := &result{
		Ast: ast.NewFile(filename),
		Errors: make(ErrorList, 0),
	}
	lexer := NewLexerWithInit(bytes.NewBufferString(source), func (lex *Lexer) {
		lex.parseResult = result
	})

	yyErrorVerbose = true
	_ = yyParse(lexer)
	return result.Ast, result.Errors
}
