package parse

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/gemrs/gem/bbc/ast"
)

//go:generate nex -e binbuf.nex
//go:generate goyacc -o binbuf.y.go binbuf.y

type ErrorList []string

func (list ErrorList) Add(err string) ErrorList {
	return append(list, err)
}

func (list ErrorList) String() string {
	return strings.Join(list, "\n")
}

type result struct {
	Ast     *ast.File
	Decls   map[string]ast.Node
	Errors  ErrorList
	AnonIdx int
}

func (res *result) AddError(err string) {
	res.Errors = res.Errors.Add(err)
}

func (l *Lexer) AddDecl(n ast.Node) {
	struc := n
	result := l.parseResult.(*result)
	result.Decls[struc.Identifier()] = struc
}

func (l *Lexer) Ast() *ast.File {
	result := l.parseResult.(*result)
	return result.Ast
}

func (l *Lexer) Error(e string) {
	result := l.parseResult.(*result)
	result.AddError(e)
}

func (l *Lexer) NameAnonStruct() string {
	// convert filename to something suitable for a variable name
	filename := l.Ast().Name
	filename = strings.Map(func(r rune) rune {
		if match, _ := regexp.MatchString("[a-zA-Z0-9_]", string(r)); match != true {
			return '_'
		}
		return r
	}, filename)

	idx := l.parseResult.(*result).AnonIdx
	l.parseResult.(*result).AnonIdx++
	return fmt.Sprintf("anonymous_%v_%v", filename, idx)
}

func (l *Lexer) resolveReferencesTo(name string, typ ast.Node, n ast.Node) {
	switch n := n.(type) {
	case *ast.Scope:
		for _, decl := range n.S {
			l.resolveReferencesTo(name, typ, decl)
		}
	case *ast.DeclReference:
		if name == n.DeclName {
			n.Object = typ
		}
	case *ast.Struct:
		l.resolveReferencesTo(name, typ, n.Scope)
	case *ast.BitStruct:
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
	case *ast.BitsType:
	case *ast.StringBaseType:
	case *ast.ByteBaseType:
	case *ast.StaticLength:
	case *ast.RemainingLength:
	default:
		if n != nil {
			// nil can occur due to lex/parse errors.
			panic(fmt.Sprintf("couldn't do anything with %T\n", n))
		}
	}
}

func Parse(filename, source string) (*ast.File, ErrorList) {
	result := &result{
		Ast:    ast.NewFile(filename),
		Decls:  make(map[string]ast.Node),
		Errors: make(ErrorList, 0),
	}
	lexer := NewLexerWithInit(bytes.NewBufferString(source), func(lex *Lexer) {
		lex.parseResult = result
	})

	yyErrorVerbose = true
	_ = yyParse(lexer)
	for name, typ := range result.Decls {
		lexer.resolveReferencesTo(name, typ, result.Ast.Scope)
	}
	return result.Ast, result.Errors
}
