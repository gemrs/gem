package gopygen

import (
	"go/ast"
	"go/token"
)

type IdentData struct {
	Identifier string
}

type Ident struct {
	*IdentData
	fileset *token.FileSet
}

func NewIdent(fileset *token.FileSet) Ident {
	return Ident{
		fileset:   fileset,
		IdentData: &IdentData{},
	}
}

func (i Ident) String() string {
	return i.Identifier
}

func (f Ident) Visit(n ast.Node) ast.Visitor {
	switch node := n.(type) {
	case *ast.Ident:
		f.Identifier = node.Name
		return nil
	}
	return f
}
