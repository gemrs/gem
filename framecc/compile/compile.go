package compile

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/sinusoids/gem/framecc/ast"
	"github.com/sinusoids/gem/framecc/parse"
)

var packageTmpl = template.Must(template.New("package").Parse(`package {{.Package}}

{{range .Types}}{{.}}

{{end}}
`))

var typeDefTmpl = template.Must(template.New("package").Parse(`type {{.Name}} struct {
{{range .Fields}}{{.}}
{{end}}
}`))

type context struct {
	types map[string]string
}

func Compile(filename string, input string) (string, error) {
	ast, errors := parse.Parse(filename, input)
	if len(errors) > 0 {
		return "", fmt.Errorf("parse errors")
	}
	ctx := &context{make(map[string]string)}

	err := ctx.generateTypes(ast.Scope)
	if err != nil {
		return "", err
	}

	types := make([]string, 0)
	for _, v := range ctx.types {
		types = append(types, v)
	}

	tmplData := struct {
		Package string
		Types   []string
	}{
		Package: "compile",
		Types:   types,
	}

	var buf bytes.Buffer
	err = packageTmpl.Execute(&buf, tmplData)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (c *context) goType(typ ast.Node) string {
	switch typ := typ.(type) {
	case *ast.ArrayType:
		switch typ := typ.Object.(type) {
		case *ast.StringBaseType:
			return "ast.JString"
		default:
			return "[]" + c.goType(typ)
		}
	case *ast.IntegerType:
		return fmt.Sprintf("ast.Int%v", typ.Bitsize)
	case *ast.Struct:
		return typ.Name
	}
	return "unknown"
}

func (c *context) generateTypes(n ast.Node) error {
	switch n := n.(type) {
	case *ast.Struct:
		err := c.doGenerateType(n)
		if err != nil {
			return err
		}
	case *ast.Scope:
		for _, decl := range n.S {
			c.generateTypes(decl)
		}
	case *ast.DynamicLength:
		c.generateTypes(n.Field)
	case *ast.ArrayType:
		c.generateTypes(n.Object)
		c.generateTypes(n.Length)
	case *ast.Field:
		c.generateTypes(n.Type)
	case *ast.Frame:
		c.generateTypes(n.Object)
	case *ast.IntegerType:
	case *ast.StringBaseType:
	case *ast.StaticLength:
	default:
		panic(fmt.Sprintf("couldn't do anything with %T\n", n))
	}
	return nil
}

func (c *context) doGenerateType(strct *ast.Struct) error {
	if _, ok := c.types[strct.Identifier()]; ok {
		fmt.Printf("Already generated type for structure %v\n", strct.Identifier())
		return nil
	}

	fields := make([]string, 0)
	for _, f := range strct.Scope.S {
		switch f := f.(type) {
		case *ast.Field:
			fieldStr := fmt.Sprintf("%v %v", f.Name, c.goType(f.Type))
			fields = append(fields, fieldStr)
		default:
			panic("declaration in structure is not a valid field")
		}
	}

	tmplData := struct {
		Name   string
		Fields []string
	}{
		Name:   strct.Identifier(),
		Fields: fields,
	}

	var buf bytes.Buffer
	err := typeDefTmpl.Execute(&buf, tmplData)
	if err != nil {
		return err
	}
	c.types[strct.Identifier()] = buf.String()
	return nil
}
