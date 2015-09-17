package parse

import (
	"framecc/ast"
	"log"
)

type token struct {
	tok int
	val yySymType
}

type dummyLexer struct {
	file *ast.File
	tokens []token
}

func (l *dummyLexer) Lex(lval *yySymType) int {
    if len(l.tokens) == 0 {
        return 0
    }

    v := l.tokens[0]
    l.tokens = l.tokens[1:]
    *lval = v.val
    return v.tok
}

func (l *dummyLexer) Error(e string) {
    log.Fatal(e)
}
