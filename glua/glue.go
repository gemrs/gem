package glua

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"html/template"
	"strings"

	"github.com/fatih/camelcase"
)

var funcMap = template.FuncMap{
	"underscore": underscore,
	"printExpr":  printExpr,
}

var modTemplate = template.Must(template.New("").Funcs(funcMap).Parse(`package {{.Name}}

import lua "github.com/yuin/gopher-lua"

// Bind{{.Name}} generates a lua binding for {{.Name}}
func Bind{{.Name}}(L *lua.LState) {
	L.PreloadModule("{{.LuaName}}", lBind{{.Name}})
}

// lBind{{.Name}} generates the table for the {{.Name}} module
func lBind{{.Name}}(L *lua.LState) int {
	mod := L.NewTable()
{{range $name, $typ := .Types}}
	lBind{{$name}}(L, mod)
{{end}}

{{range .Fields}}
	L.SetField(mod, "{{underscore .Name}}", glua.ToLua(L, {{.Name}}))
{{end}}
	L.Push(mod)
	return 1
}

{{$modName := .Name}}

{{range $typeName, $typ := .Types}}
func lBind{{$typeName}}(L *lua.LState, mod *lua.LTable) {
	mt := L.NewTypeMetatable("{{$modName}}.{{$typeName}}")
	L.SetField(mt, "__call", L.NewFunction(lNew{{$typeName}}))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), {{$typeName}}Methods))

	cls := L.NewUserData()
	L.SetField(mod, "{{$typeName}}", cls)
	L.SetMetatable(cls, mt)
	glua.RegisterType("{{$modName}}.{{$typeName}}", mt)
}

{{with $typ.Constructor}}
func lNew{{$typeName}}(L *lua.LState) int {
{{.GenerateCtor}}
}
{{end}}

var {{$typeName}}Methods = map[string]lua.LGFunction{
{{range $methodName, $method := $typ.Methods}}
	"{{underscore $methodName}}": lBind{{$typeName}}{{$methodName}},
{{end}}
{{range $propName, $propType := $typ.Accessors}}
	"{{underscore $propName}}": lBindProp{{$typeName}}{{$propName}},
{{end}}
}

{{range $methodName, $method := $typ.Methods}}
func lBind{{$typeName}}{{$methodName}}(L *lua.LState) int {
{{$method.Generate}}
}
{{end}}

{{range $propName, $propType := $typ.Accessors}}
func lBindProp{{$typeName}}{{$propName}}(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*{{$typeName}})
    if L.GetTop() == 2 {
		val := glua.FromLua(L.Get(2)).({{printExpr $propType}})
		self.Set{{$propName}}(val)
        return 0
    }
    L.Push(glua.ToLua(L, self.{{$propName}}()))
    return 1
}
{{end}}

{{end}}
`))

type lModule struct {
	Name    string
	LuaName string
	Types   map[string]*lType
	Fields  map[string]*lField
}

func (mod *lModule) String() string {
	var buf bytes.Buffer

	modTemplate.Execute(&buf, mod)
	return string(buf.Bytes())
}

func printExpr(typ ast.Expr) string {
	var buf bytes.Buffer
	printer.Fprint(&buf, token.NewFileSet(), typ)
	return buf.String()
}

func fromLua(src, dest string, typ ast.Expr) string {
	if array, ok := typ.(*ast.ArrayType); ok {
		var buf bytes.Buffer
		fmt.Fprintf(&buf, "%vArray := %v.(*lua.LTable)\n", src, src)
		fmt.Fprintf(&buf, "%v := make([]%v, %vArray.Len())\n", dest, array.Elt, src)
		fmt.Fprintf(&buf, "%vArray.ForEach(func (k lua.LValue, val lua.LValue) {\n", src)
		fmt.Fprintf(&buf, "i := int(lua.LVAsNumber(k)) - 1\n")
		fmt.Fprintf(&buf, "%v[i] = glua.FromLua(val).(%v)\n", dest, printExpr(array.Elt))
		fmt.Fprintf(&buf, "})\n")
		return buf.String()
	}

	return fmt.Sprintf("%v := glua.FromLua(%v).(%v)\n", dest, src, printExpr(typ))
}

func (method *lFunction) Generate() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "self := glua.FromLua(L.Get(1)).(*%v)\n", method.Recv)
	fmt.Fprintf(&buf, "L.Remove(1)\n")

	argNames := make([]string, len(method.Args))
	for i, arg := range method.Args {
		fmt.Fprintf(&buf, "arg%vValue := L.Get(1)\n", i)
		fmt.Fprintf(&buf, "%v", fromLua(fmt.Sprintf("arg%vValue", i), fmt.Sprintf("arg%v", i), arg))
		fmt.Fprintf(&buf, "L.Remove(1)\n")
		argNames[i] = fmt.Sprintf("arg%v", i)
	}

	args := strings.Join(argNames, ", ")
	if method.Ret != nil {
		fmt.Fprintf(&buf, "retVal := self.%v(%v)\n", method.Name, args)
		fmt.Fprintf(&buf, "L.Push(glua.ToLua(L, retVal))\n")
		fmt.Fprintf(&buf, "return 1\n")
	} else {
		fmt.Fprintf(&buf, "self.%v(%v)\n", method.Name, args)
		fmt.Fprintf(&buf, "return 0\n")
	}
	return buf.String()
}

func (method *lFunction) GenerateCtor() string {
	var buf bytes.Buffer
	argNames := make([]string, len(method.Args))

	// Remove the cls
	fmt.Fprintf(&buf, "L.Remove(1)\n")

	for i, arg := range method.Args {
		fmt.Fprintf(&buf, "arg%vValue := L.Get(1)\n", i)
		fmt.Fprintf(&buf, "%v", fromLua(fmt.Sprintf("arg%vValue", i), fmt.Sprintf("arg%v", i), arg))
		fmt.Fprintf(&buf, "L.Remove(1)\n")
		argNames[i] = fmt.Sprintf("arg%v", i)
	}

	if method.PassState {
		argNames = append([]string{"L"}, argNames...)
	}

	args := strings.Join(argNames, ", ")
	if method.Ret != nil {
		fmt.Fprintf(&buf, "retVal := %v(%v)\n", method.Name, args)
		fmt.Fprintf(&buf, "L.Push(glua.ToLua(L, retVal))\n")
		fmt.Fprintf(&buf, "return 1\n")
	} else {
		fmt.Fprintf(&buf, "self.%v(%v)\n", method.Name, args)
		fmt.Fprintf(&buf, "return 0\n")
	}
	return buf.String()
}

type lType struct {
	Name        string
	Constructor *lFunction
	Methods     map[string]*lFunction
	Accessors   map[string]ast.Expr
}

type lFunction struct {
	Name      string
	Recv      string
	Args      []ast.Expr
	Ret       ast.Expr
	PassState bool
}

type lField struct {
	Name string
}

func underscore(ident string) string {
	parts := camelcase.Split(ident)
	for i, s := range parts {
		parts[i] = strings.ToLower(s)
	}
	return strings.Join(parts, "_")
}
