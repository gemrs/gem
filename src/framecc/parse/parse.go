package parse

import (
	"bytes"
	"strings"
	"fmt"

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
	Decls map[string]ast.Node
	Errors ErrorList
}

func (res *result) AddError(err string) {
	res.Errors = res.Errors.Add(err)
}

func (l *Lexer) AddDecl(n ast.Node) {
	struc := n.(*ast.Struct)
	result := l.parseResult.(*result)
	result.Decls[struc.Name] = struc
}

func (l *Lexer) Ast() *ast.File {
	result := l.parseResult.(*result)
    return result.Ast
}

func (l *Lexer) Error(e string) {
	result := l.parseResult.(*result)
	result.AddError(e)
}

func (l *Lexer) resolveReferencesTo(name string, typ ast.Node, n ast.Node) {
	fmt.Printf("resolving reference to %v for node %v\n", name, n)
	switch n := n.(type) {
	case *ast.Scope:
		for _, decl := range n.S {
			l.resolveReferencesTo(name, typ, decl)
		}
	case *ast.DeclReference:
		fmt.Printf("found reference to %v on node %v\n", name, n)
		if name == n.DeclName {
			n.Object = typ
		}
	case *ast.Struct:
		l.resolveReferencesTo(name, typ, n.Scope)
	case *ast.DynamicLength:
		l.resolveReferencesTo(name, typ, n.Field)
	case *ast.ArrayType:
		l.resolveReferencesTo(name, typ, n.Object)
		l.resolveReferencesTo(name, typ, n.Length)
	case *ast.Field:
		l.resolveReferencesTo(name, typ, n.Type)
	case *ast.Frame:
		l.resolveReferencesTo(name, typ, n.Object)
	case *ast.IntegerType:
	case *ast.StringBaseType:
	case *ast.StaticLength:
	default:
		if n != nil {
			// nil can occur due to lex/parse errors.
			panic(fmt.Sprintf("couldn't do anything with %T\n", n))
		}
	}
}

func Parse(filename, source string) (*ast.File, ErrorList) {
	result := &result{
		Ast: ast.NewFile(filename),
		Decls: make(map[string]ast.Node),
		Errors: make(ErrorList, 0),
	}
	lexer := NewLexerWithInit(bytes.NewBufferString(source), func (lex *Lexer) {
		lex.parseResult = result
	})

	yyErrorVerbose = true
	_ = yyParse(lexer)
	for name, typ := range result.Decls {
		lexer.resolveReferencesTo(name, typ, result.Ast.Scope)
	}
	return result.Ast, result.Errors
}
