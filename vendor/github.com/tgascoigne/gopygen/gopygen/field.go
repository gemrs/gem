package gopygen

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

type FieldDeclData struct {
	Name      Ident
	Type      Type
	Directive Directive
}

type FieldDecl struct {
	*FieldDeclData
	fileset *token.FileSet
}

type FieldListData struct {
	Fields []FieldDecl
}

type FieldList struct {
	*FieldListData
	fileset *token.FileSet
}

func NewFieldDecl(fileset *token.FileSet) FieldDecl {
	return FieldDecl{
		fileset: fileset,
		FieldDeclData: &FieldDeclData{
			Type:      NewType(fileset),
			Name:      NewIdent(fileset),
			Directive: NewDirective(fileset),
		},
	}
}

func NewFieldList(fileset *token.FileSet) FieldList {
	return FieldList{
		fileset: fileset,
		FieldListData: &FieldListData{
			Fields: []FieldDecl{},
		},
	}
}

func (f FieldList) Visit(n ast.Node) ast.Visitor {
	switch node := n.(type) {
	case *ast.Field:
		// BUG(tom): ignore embedded fields
		fieldType := NewType(f.fileset)
		ast.Walk(fieldType, node.Type)

		fieldDirective := NewDirective(f.fileset)
		if node.Comment != nil {
			ast.Walk(fieldDirective, node.Comment)
		}

		if len(node.Names) == 0 {
			// Anonymous field
			field := NewFieldDecl(f.fileset)
			field.Type = fieldType
			field.Directive = fieldDirective
			f.Fields = append(f.Fields, field)
		}

		for _, name := range node.Names {
			field := NewFieldDecl(f.fileset)
			field.Type = fieldType
			field.Directive = fieldDirective
			ast.Walk(field.Name, name)
			f.Fields = append(f.Fields, field)
		}

		return f
	}

	return f
}

func (f *FieldDeclData) String() string {
	return fmt.Sprintf("%s %s", f.Name.String(), f.Type.String())
}

func (f *FieldDeclData) Anonymous() bool {
	return f.Name.String() == ""
}

func (f *FieldListData) Empty() bool {
	return len(f.Fields) == 0
}

func (f *FieldListData) VarList(prefix string) string {
	fieldStrs := []string{}
	for i, _ := range f.Fields {
		fieldName := fmt.Sprintf("%v", i)

		fieldStrs = append(fieldStrs, fmt.Sprintf("%s%s", prefix, fieldName))
	}
	return strings.Join(fieldStrs, ", ")
}

func (f *FieldListData) ParamList(prefix string) string {
	fieldStrs := []string{}
	for i, field := range f.Fields {
		fieldName := fmt.Sprintf("%v", i)

		fieldStrs = append(fieldStrs, fmt.Sprintf("%s%s.(%s)", prefix, fieldName, field.Type.String()))
	}
	return strings.Join(fieldStrs, ", ")
}

func (f *FieldListData) FuncParamList(prefix string) string {
	fieldStrs := []string{}
	for i, field := range f.Fields {
		fieldName := fmt.Sprintf("%v", i)

		fieldStrs = append(fieldStrs, fmt.Sprintf("%s%s %s", prefix, fieldName, field.Type.String()))
	}
	return strings.Join(fieldStrs, ", ")
}

func (f *FieldListData) String() string {
	fieldStrs := []string{}
	for _, field := range f.Fields {
		fieldStrs = append(fieldStrs, field.String())
	}
	return strings.Join(fieldStrs, ", ")
}
