package gopygen

import (
	"fmt"
	"go/ast"
	"go/token"
)

type Stringable interface {
	String() string
}

type SelectorExprType struct {
	X   Type
	Sel Type
}

func (s SelectorExprType) String() string {
	return s.X.String() + "." + s.Sel.String()
}

type SliceType struct {
	Base Type
}

func (s SliceType) String() string {
	return "[]" + s.Base.String()
}

type InterfaceType struct{}

func (i InterfaceType) String() string {
	//BUG(tom): Don't support non-empty interface types
	return "interface{}"
}

type PtrType struct {
	Base Type
}

func (t PtrType) String() string {
	return "*" + t.Base.String()
}

type MapType struct {
	Key   Type
	Value Type
}

func (t MapType) String() string {
	return fmt.Sprintf("map[%s]%s", t.Key, t.Value)
}

type TypeData struct {
	typeString Stringable
}

type Type struct {
	*TypeData
	fileset *token.FileSet
}

func NewType(fileset *token.FileSet) Type {
	return Type{
		fileset:  fileset,
		TypeData: &TypeData{},
	}
}

func (t Type) Visit(n ast.Node) ast.Visitor {
	switch node := n.(type) {
	case *ast.StarExpr:
		ptrType := &PtrType{
			Base: NewType(t.fileset),
		}
		ast.Walk(ptrType.Base, node.X)
		t.typeString = ptrType
		return nil
	case *ast.MapType:
		mapType := &MapType{
			Key:   NewType(t.fileset),
			Value: NewType(t.fileset),
		}
		ast.Walk(mapType.Key, node.Key)
		ast.Walk(mapType.Value, node.Value)
		t.typeString = mapType
		return nil
	case *ast.Ident:
		plainType := NewIdent(t.fileset)
		ast.Walk(plainType, node)
		t.typeString = plainType
		return nil
	case *ast.InterfaceType:
		t.typeString = &InterfaceType{}
		return nil
	case *ast.ArrayType:
		sliceType := &SliceType{
			Base: NewType(t.fileset),
		}
		ast.Walk(sliceType.Base, node.Elt)
		t.typeString = sliceType
		return nil
	case *ast.SelectorExpr:
		selectorType := &SelectorExprType{
			X:   NewType(t.fileset),
			Sel: NewType(t.fileset),
		}
		ast.Walk(selectorType.X, node.X)
		ast.Walk(selectorType.Sel, node.Sel)
		t.typeString = selectorType
	}
	return nil
}

func (t Type) String() string {
	if t.typeString == nil {
		return "INVALID!!"
	}
	return t.typeString.String()
}

func (t Type) BaseType() string {
	switch typ := t.typeString.(type) {
	case *PtrType:
		return typ.Base.String()
	default:
		return typ.String()
	}
}
