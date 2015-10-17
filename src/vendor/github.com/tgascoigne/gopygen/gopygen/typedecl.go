package gopygen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"text/template"
)

const classDefinitionStr = `
var {{.Ident}}Def = py.Class{
	Name: "{{.Ident}}",
	Flags:   py.TPFLAGS_BASETYPE,
	Pointer: (*{{.Ident}})(nil),
}
`

const objectAllocFunctionStr = `
// Alloc allocates an object for use in python land.
// Copies the member fields from this object to the newly allocated object
// Usage: obj := GoObject{X:1, Y: 2}.Alloc()
func New{{.Ident}}({{.Args.FuncParamList "arg_"}}) (*{{.Ident}}, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	// Allocate
	alloc_, err := {{.Ident}}Def.Alloc(0)
	if err != nil {
		return nil, err
	}
	alloc := alloc_.(*{{.Ident}})
	err = alloc.Init({{.Args.VarList "arg_"}})
	return alloc, err
}
`

const classRegisterStr = `
// Registers this type with a python module
func Register{{.Ident}}(module *py.Module) error {
	var err error
	var class *py.Type
	if class, err = {{.Ident}}Def.Create(); err != nil {
		return err
	}

	if err = module.AddObject("{{.Ident}}", class); err != nil {
		return err
	}

	return nil
}
`

const accessorsStr = `
{{$ident := .Ident}}
{{with .Field}}
  {{if not .Anonymous}}
func (obj *{{$ident}}) PyGet_{{.Name}}() (py.Object, error) {
	return gopygen.TypeConvOut(obj.{{.Name}}, "{{.Type}}")
}

func (obj *{{$ident}}) PySet_{{.Name}}(arg py.Object) error {
    arg.Incref()
	val, err := gopygen.TypeConvIn(arg, "{{.Type}}")
	if err != nil {
		return err
	}

    if _, ok := val.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
	    val.(py.Object).Incref()
    }
    arg.Decref()

    var tmp interface{}
    tmp = &obj.{{.Name}}
	obj.{{.Name}} = val.({{.Type}})

    if oldObj, ok := tmp.(py.Object); ok {
		// If we're not converting it from a python object, we should refcount it properly
        oldObj.Decref()
    }
	return nil
}
  {{end}}
{{end}}
`

const constructorWrapperStr = `{{$recv := .Recv}}
{{$num_results := len .Results.Fields}}
func (obj *{{$recv}}) PyInit(_args *py.Tuple, kwds *py.Dict) error {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != {{len .Params.Fields}} {
		return fmt.Errorf("({{$recv}}) PyInit: parameter length mismatch")
	}
{{with .Params.Fields}}
  {{range $i, $param := .}}
    args[{{$i}}].Incref()
	in_{{$i}}, err := gopygen.TypeConvIn(args[{{$i}}], "{{$param.Type}}")
	if err != nil {
		return err
	}
  {{end}}
{{end}}
    err = obj.Init({{.Params.ParamList "in_"}})
{{with .Params.Fields}}
  {{range $i, $param := .}}
    args[{{$i}}].Decref()
  {{end}}
{{end}}
    return err
}
`

var classDefinitionTmpl = template.Must(template.New("class_definition").Parse(classDefinitionStr))
var objectAllocFunctionTmpl = template.Must(template.New("object_alloc").Parse(objectAllocFunctionStr))
var classRegisterTmpl = template.Must(template.New("class_register").Parse(classRegisterStr))
var accessorsTmpl = template.Must(template.New("field_accessors").Parse(accessorsStr))
var constructorWrapperTmpl = template.Must(template.New("constructor_wrapper").Parse(constructorWrapperStr))

type TypeDeclData struct {
	Ident       Ident
	Fields      FieldList
	fieldFilter FilterFunc
	NewFunc     *FuncDecl
}

type TypeDecl struct {
	*TypeDeclData
	fileset *token.FileSet
}

func NewTypeDecl(fileset *token.FileSet, fieldFilter FilterFunc) TypeDecl {
	return TypeDecl{
		fileset: fileset,
		TypeDeclData: &TypeDeclData{
			Fields:      NewFieldList(fileset),
			fieldFilter: fieldFilter,
		},
	}
}

func (d TypeDecl) Visit(n ast.Node) ast.Visitor {
	switch node := n.(type) {
	case *ast.TypeSpec:
		d.Ident = NewIdent(d.fileset)
		ast.Walk(d.Ident, node.Name)
	case *ast.StructType:
		return d.Fields
	}
	return d
}

func (d *TypeDeclData) FilteredFieldName(name string) bool {
	return !d.fieldFilter(name)
}

func (d *TypeDeclData) FilteredFieldDecls() []FieldDecl {
	newDecls := []FieldDecl{}
	for _, decl := range d.Fields.Fields {
		if !d.FilteredFieldName(decl.Name.String()) {
			newDecls = append(newDecls, decl)
		}
	}
	return newDecls
}

func (d *TypeDeclData) AccessorFunctions() string {
	var buffer bytes.Buffer
	for _, f := range d.FilteredFieldDecls() {
		tmplData := struct {
			Ident Ident
			Field FieldDecl
		}{
			Ident: d.Ident,
			Field: f,
		}
		err := accessorsTmpl.Execute(&buffer, tmplData)
		if err != nil {
			panic(err)
		}
	}
	return buffer.String()
}

func (d *TypeDeclData) AllocateFunction() string {
	if d.NewFunc == nil {
		return ""
	}

	tmplData := struct {
		Args FieldList
		*TypeDeclData
	}{
		Args:         d.NewFunc.Params,
		TypeDeclData: d,
	}

	var buffer bytes.Buffer
	err := objectAllocFunctionTmpl.Execute(&buffer, tmplData)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

func (d *TypeDeclData) RegisterFunction() string {
	var buffer bytes.Buffer
	err := classRegisterTmpl.Execute(&buffer, d)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

func (d *TypeDeclData) Constructor() string {
	if d.NewFunc == nil {
		return ""
	}

	tmplData := struct {
		Recv Ident
		*FuncDecl
	}{
		Recv:     d.Ident,
		FuncDecl: d.NewFunc,
	}

	var buffer bytes.Buffer
	err := constructorWrapperTmpl.Execute(&buffer, tmplData)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

func (d *TypeDeclData) ClassDeclaration() string {
	var buffer bytes.Buffer
	err := classDefinitionTmpl.Execute(&buffer, d)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

func (d *TypeDeclData) String() string {
	var buffer bytes.Buffer
	fmt.Fprintf(&buffer, "%v", d.Ident.String())
	return buffer.String()
}
