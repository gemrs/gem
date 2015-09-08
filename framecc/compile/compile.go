package compile

import (
	"bytes"
	"fmt"
	"strconv"
	"text/template"

	"github.com/sinusoids/gem/framecc/ast"
	"github.com/sinusoids/gem/framecc/parse"
)

var packageTmpl = template.Must(template.New("package").Parse(`package {{.Package}}
import (
	"github.com/sinusoids/gem/gem/encode"
)

{{range .Types}}{{.}}

{{end}}
`))

var typeDefTmpl = template.Must(template.New("typedef").Parse(`type {{.Name}} struct {
{{range .Fields}}{{.}}
{{end}}
}

{{.EncodeFuncs}}
`))

var frameDefTmpl = template.Must(template.New("framedef").Parse(`type {{.Identifier}} {{.Object.Identifier}}`))

var fieldFuncTmpl = template.Must(template.New("fieldfunc").Parse(`err = struc.{{.Name}}.{{.Operation}}(buf, {{.Flags}})
if err != nil {
	return err
}`))

var fieldFuncArrayTmpl = template.Must(template.New("fieldfunc").Parse(`for i := 0; i < {{.Size}}; i++ {
	err = struc.{{.Name}}[i].{{.Operation}}(buf, {{.Flags}})
	if err != nil {
		return err
	}
}`))

var encodeFuncsTmpl = template.Must(template.New("encode").Parse(`func (struc *{{.Type}}) Encode(buf *bytes.Buffer, flags_ interface{}) (err error) {
{{range .EncodeFields}}{{.}}
{{end}}
}

func (struc *{{.Type}}) Decode(buf *bytes.Buffer, flags_ interface{}) (err error) {
{{range .DecodeFields}}{{.}}
{{end}}
}`))

type context struct {
	types map[string]string
}

func Compile(filename, pkg, input string) (string, error) {
	ast, errors := parse.Parse(filename, input)
	if len(errors) > 0 {
		return "", fmt.Errorf("parse errors\n%v", errors)
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
		Package: pkg,
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
		array := typ
		switch typ := typ.Object.(type) {
		case *ast.StringBaseType:
			return "encode.JString"
		default:
			baseType := c.goType(typ)
			switch arrayLength := array.Length.(type) {
			case *ast.StaticLength:
				return "[" + strconv.Itoa(arrayLength.Length) + "]" + baseType
			}
			return "[]" + baseType
		}
	case *ast.DeclReference:
		return c.goType(typ.Object)
	case *ast.IntegerType:
		return fmt.Sprintf("encode.Int%v", typ.Bitsize)
	case *ast.Struct:
		return typ.Name
	}
	panic(fmt.Sprintf("unknown type: %T", typ))
}

func (c *context) generateTypes(n ast.Node) error {
	switch n := n.(type) {
	case *ast.Struct:
		if err := c.doGenerateType(n); err != nil {
			return err
		}
	case *ast.Scope:
		for _, decl := range n.S {
			if err := c.generateTypes(decl); err != nil {
				return err
			}
		}
	case *ast.DynamicLength:
		if err := c.generateTypes(n.Field); err != nil {
			return err
		}
	case *ast.ArrayType:
		if err := c.generateTypes(n.Object); err != nil {
			return err
		}
		if err := c.generateTypes(n.Length); err != nil {
			return err
		}
	case *ast.Field:
		if err := c.generateTypes(n.Type); err != nil {
			return err
		}
	case *ast.Frame:
		if err := c.generateTypes(n.Object); err != nil {
			return err
		}

		if err := c.doGenerateTypeDef(n); err != nil {
			return err
		}
	case *ast.DeclReference:
		if n.Object == nil {
			panic("unresolved reference at compile time")
		}
		if err := c.generateTypes(n.Object); err != nil {
			return err
		}
	case *ast.IntegerType:
	case *ast.StringBaseType:
	case *ast.StaticLength:
	default:
		panic(fmt.Sprintf("couldn't do anything with %T\n", n))
	}
	return nil
}

func (c *context) doGenerateTypeDef(frame *ast.Frame) error {
	if _, ok := c.types[frame.Identifier()]; ok {
		fmt.Printf("Already generated type for frame %v\n", frame.Identifier())
		return nil
	}

	fmt.Printf("Generating type for %v\n", frame.Identifier())

	typeStr, err := bufferTemplate(frameDefTmpl, frame)
	if err != nil {
		return err
	}
	c.types[frame.Identifier()] = typeStr
	return nil
}

func (c *context) doGenerateType(strct *ast.Struct) error {
	if _, ok := c.types[strct.Identifier()]; ok {
		fmt.Printf("Already generated type for structure %v\n", strct.Identifier())
		return nil
	}

	fmt.Printf("Generating type for %v\n", strct.Identifier())

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

	funcs, err := c.generateEncodeFuncs(strct)
	if err != nil {
		return err
	}

	tmplData := struct {
		Name        string
		Fields      []string
		EncodeFuncs string
	}{
		Name:        strct.Identifier(),
		Fields:      fields,
		EncodeFuncs: funcs,
	}

	typeStr, err := bufferTemplate(typeDefTmpl, tmplData)
	if err != nil {
		return err
	}
	c.types[strct.Identifier()] = typeStr
	return nil
}

func (c *context) generateArrayLength(length ast.LengthSpec) string {
	switch length := length.(type) {
	case *ast.StaticLength:
		return strconv.Itoa(length.Length)
	case *ast.DynamicLength:
		panic("dynamic array length not implemented")
	default:
		panic(fmt.Sprintf("unknown array length spec: %T", length))
	}
}

func (c *context) generateFieldFunc(operation string, field *ast.Field) (string, error) {
	type fieldFuncTmplData struct {
		Name      string
		Operation string
		Flags     string
		Size      string
	}
	var tmpl *template.Template
	var tmplData fieldFuncTmplData

outer:
	switch typ := field.Type.(type) {
	case *ast.ArrayType:
		// Strings are a special case
		switch typ.Object.(type) {
		case *ast.StringBaseType:
			break outer
		}

		tmpl = fieldFuncArrayTmpl
		tmplData = fieldFuncTmplData{
			Name:      field.Identifier(),
			Operation: operation,
			Flags:     c.generateEncodeFlags(typ.Object),
			Size:      c.generateArrayLength(typ.Length),
		}
	}

	if tmpl == nil {
		tmpl = fieldFuncTmpl
		tmplData = fieldFuncTmplData{
			Name:      field.Identifier(),
			Operation: operation,
			Flags:     c.generateEncodeFlags(field.Type),
		}
	}

	fieldFunc, err := bufferTemplate(tmpl, tmplData)
	if err != nil {
		return "", err
	}
	return fieldFunc, nil
}

func (c *context) generateEncodeFuncs(strct *ast.Struct) (string, error) {
	encodeFields, decodeFields := make([]string, 0), make([]string, 0)

	for _, field := range strct.Scope.S {
		switch field := field.(type) {
		case *ast.Field:
			encode, err := c.generateFieldFunc("Encode", field)
			if err != nil {
				return "", err
			}

			decode, err := c.generateFieldFunc("Decode", field)
			if err != nil {
				return "", err
			}

			encodeFields = append(encodeFields, encode)
			decodeFields = append(decodeFields, decode)
		default:
			panic("non-field in struct scope")
		}
	}

	tmplData := struct {
		Type         string
		EncodeFields []string
		DecodeFields []string
	}{
		Type:         strct.Identifier(),
		EncodeFields: encodeFields,
		DecodeFields: decodeFields,
	}

	return bufferTemplate(encodeFuncsTmpl, tmplData)
}

func (c *context) generateEncodeFlags(typ ast.Node) string {
	switch typ := typ.(type) {
	case *ast.IntegerType:
		return fmt.Sprintf("encode.IntegerFlag(%v)", typ.Modifiers)
	case *ast.ArrayType:
		// For strings, pass the expected length as a flag
		switch typ.Object.(type) {
		case *ast.StringBaseType:
			return c.generateArrayLength(typ.Length)
		}

		return fmt.Sprintf("encode.NilFlags")
	case *ast.DeclReference:
		return fmt.Sprintf("encode.NilFlags")
	default:
		panic(fmt.Errorf("couldn't do anything with type %T", typ))
	}
}

func bufferTemplate(tmpl *template.Template, data interface{}) (string, error) {
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
