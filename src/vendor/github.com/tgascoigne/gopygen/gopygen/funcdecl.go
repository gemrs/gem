package gopygen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"text/template"
)

const methodWrapperStr = `{{$recv := index .Recievers.Fields 0}}
{{$num_results := len .Results.Fields}}
func ({{$recv}}) Py_{{.Name}}(_args *py.Tuple, kwds *py.Dict) (py.Object, error) {
	lock := py.NewLock()
	defer lock.Unlock()

	var err error
	_ = err
	args := _args.Slice()
	if len(args) != {{len .Params.Fields}} {
		return nil, fmt.Errorf("Py_{{.Name}}: parameter length mismatch")
	}
    // Convert parameters
{{with .Params.Fields}}
  {{range $i, $param := .}}
    args[{{$i}}].Incref()
	in_{{$i}}, err := gopygen.TypeConvIn(args[{{$i}}], "{{$param.Type}}")
	if err != nil {
		return nil, err
	}
  {{end}}
{{end}}

    // Make the function call
{{with .Results}}
  {{if gt $num_results 0}}
    {{.VarList "res"}} := {{end}}{{end}}{{$recv.Name}}.{{.Name}}({{.Params.ParamList "in_"}})

    // Remove local references
{{with .Params.Fields}}
  {{range $i, $param := .}}
    args[{{$i}}].Decref()
  {{end}}
{{end}}

{{with .Results}}
  {{if eq $num_results 0}}
	py.None.Incref()
	return py.None, nil
  {{else}}
    {{range $i, $res := .Fields}}
	out_{{$i}}, err := gopygen.TypeConvOut(res{{$i}}, "{{.Type}}")
	if err != nil {
		return nil, err
	}
    out_{{$i}}.Incref()
    {{end}}
    {{if eq $num_results 1}}
	return out_0, nil
    {{else}}
	return py.PackTuple({{.VarList "out_"}})
    {{end}}
  {{end}}
{{end}}
}
`

type FuncDeclData struct {
	Name      Ident
	Recievers FieldList
	Params    FieldList
	Results   FieldList
}

type FuncDecl struct {
	*FuncDeclData
	fileset *token.FileSet
}

func NewFuncDecl(fileset *token.FileSet) FuncDecl {
	return FuncDecl{
		fileset: fileset,
		FuncDeclData: &FuncDeclData{
			Name:      NewIdent(fileset),
			Recievers: NewFieldList(fileset),
			Params:    NewFieldList(fileset),
			Results:   NewFieldList(fileset),
		},
	}

}

func (t FuncDecl) Visit(n ast.Node) ast.Visitor {
	switch node := n.(type) {
	case *ast.FuncDecl:
		ast.Walk(t.Name, node.Name)
		if node.Recv != nil {
			ast.Walk(t.Recievers, node.Recv)
		}

		if node.Type.Params != nil {
			ast.Walk(t.Params, node.Type.Params)
		}
		if node.Type.Results != nil {
			ast.Walk(t.Results, node.Type.Results)
		}

	}
	return nil
}

var methodWrapperTmpl = template.Must(template.New("method_wrap").Parse(methodWrapperStr))

func (t *FuncDeclData) Wrap() string {
	var buffer bytes.Buffer
	err := methodWrapperTmpl.Execute(&buffer, t)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

func (t *FuncDeclData) String() string {
	return fmt.Sprintf("func (%s) %s(%s) %s", t.Recievers, t.Name, t.Params, t.Results)
}
