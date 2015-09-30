package gopygen

import (
	"bytes"
	"go/ast"
	"go/token"
	"strings"
	"text/template"
)

const headerStr = `// Generated by gopygen; DO NOT EDIT
package {{.PackageName}}

import (
	"fmt"

	"github.com/qur/gopy/lib"
	"github.com/tgascoigne/gopygen/gopygen"
)

// Sometimes we might generate code which doesn't use some of the above imports
// Use them here just in case
var _ = fmt.Sprintf("")
var _ = gopygen.Dummy

{{range .FilteredTypeDecls}}{{.ClassDeclaration}}
{{.RegisterFunction}}
{{.AllocateFunction}}
{{.AccessorFunctions}}
{{end}}
{{range .FilteredFuncDecls}}{{.Wrap}}
{{end}}
`

var headerTmpl = template.Must(template.New("header").Parse(headerStr))

var Dummy = ""

type FileData struct {
	PackageName string
	TypeDecls   []TypeDecl
	FuncDecls   []FuncDecl
	types       []string
	funcFilter  FilterFunc
	fieldFilter FilterFunc
}

type FilterFunc func(string) bool

type File struct {
	*FileData
	fileset *token.FileSet
}

func NewFile(fileset *token.FileSet, types []string, funcFilter, fieldFilter FilterFunc) File {
	return File{
		FileData: &FileData{
			types:       types,
			funcFilter:  funcFilter,
			fieldFilter: fieldFilter,
		},
		fileset: fileset,
	}
}

func (f *FileData) FilteredType(typ string) bool {
	for _, t := range f.types {
		if t == typ {
			return false
		}
	}
	return true
}

func (f *FileData) FilteredFuncName(name string) bool {
	if strings.HasPrefix(name, "Py") {
		return true
	}

	include := f.funcFilter(name)
	return !include
}

func (f *FileData) FilteredTypeDecls() []TypeDecl {
	newDecls := []TypeDecl{}
	for _, decl := range f.TypeDecls {
		if !f.FilteredType(decl.Ident.String()) {
			newDecls = append(newDecls, decl)
		}
	}
	return newDecls
}

func (f *FileData) FilteredFuncDecls() []FuncDecl {
	newDecls := []FuncDecl{}
	for _, decl := range f.FuncDecls {
		if decl.Recievers.Empty() {
			continue
		}

		recv := decl.Recievers.Fields[0]

		if f.FilteredFuncName(decl.Name.String()) {
			continue
		}

		if f.FilteredType(recv.Type.BaseType()) {
			continue
		}

		newDecls = append(newDecls, decl)
	}
	return newDecls
}

func (f File) Visit(n ast.Node) ast.Visitor {
	switch node := n.(type) {
	case *ast.File:
		f.PackageName = node.Name.String()
	case *ast.GenDecl:
		switch node.Tok {
		case token.TYPE:
			f.TypeDecls = append(f.TypeDecls, NewTypeDecl(f.fileset, f.fieldFilter))
			return f.TypeDecls[len(f.TypeDecls)-1]
		}
	case *ast.FuncDecl:
		newFunc := NewFuncDecl(f.fileset)
		ast.Walk(newFunc, n)
		f.FuncDecls = append(f.FuncDecls, newFunc)
		return nil
	}
	return f
}

func (f *FileData) String() string {
	var buffer bytes.Buffer
	err := headerTmpl.Execute(&buffer, f)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}