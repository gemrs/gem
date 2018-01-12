package glua

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func Gather(path string) ([]*lModule, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	modules := make([]*lModule, 0)
	for pkgName, pkg := range pkgs {
		modules = append(modules, gatherModule(pkgName, pkg))
	}

	return modules, nil
}

func gatherModule(name string, pkg *ast.Package) *lModule {
	module := &lModule{
		Name:      name,
		Types:     make(map[string]*lType),
		Fields:    make(map[string]*lField),
		Functions: make(map[string]*lFunction),
	}

	// Ew...
	// Gather types and fields
	for _, file := range pkg.Files {
		if file.Doc != nil && module.LuaName == "" {
			module.LuaName = extractAlias(file.Doc)
		}

		for _, decl := range file.Decls {
			if decl, ok := decl.(*ast.GenDecl); ok {
				for _, spec := range decl.Specs {
					if ok, _ := hasBindComment(decl.Doc); !ok {
						continue
					}

					if spec, ok := spec.(*ast.TypeSpec); ok {
						typ := gatherType(spec)
						module.Types[spec.Name.Name] = typ
					}

					if spec, ok := spec.(*ast.ValueSpec); ok {
						fields := gatherFields(spec)
						for _, field := range fields {
							if field.Name == "_" {
								continue
							}
							module.Fields[field.Name] = field
						}
					}
				}
			}
		}
	}

	// Gather methods
	for _, file := range pkg.Files {
		for _, decl := range file.Decls {
			if fn, ok := decl.(*ast.FuncDecl); ok {
				ok, bindArgs := hasBindComment(fn.Doc)
				if !ok {
					continue
				}

				receiver := receiverName(fn)
				if receiver == "" {
					if len(bindArgs) > 0 && bindArgs[0] == "constructor" {
						if len(bindArgs) != 2 {
							panic("expected type argument to constructor annotation")
						}

						typeName := bindArgs[1]
						typ, ok := module.Types[typeName]
						if !ok {
							panic(fmt.Sprintf("constructor for unbound type %v", typeName))
						}
						typ.Constructor = gatherFunction(fn)
						typ.Constructor.Recv = typeName
					} else {
						module.Functions[fn.Name.Name] = gatherFunction(fn)
					}
					continue
				}

				if typ, ok := module.Types[receiver]; ok {
					if len(bindArgs) > 0 && bindArgs[0] == "accessor" {
						typ.Accessors[fn.Name.Name] = fn.Type.Results.List[0].Type
					} else {
						typ.Methods[fn.Name.Name] = gatherFunction(fn)
						typ.Methods[fn.Name.Name].Recv = receiver
					}
				}
			}
		}
	}

	if module.LuaName == "" {
		module.LuaName = module.Name
	}

	return module
}

func gatherType(spec *ast.TypeSpec) *lType {
	typ := &lType{
		Name:      spec.Name.Name,
		Methods:   make(map[string]*lFunction),
		Accessors: make(map[string]ast.Expr),
	}
	return typ
}

func gatherFunction(fn *ast.FuncDecl) *lFunction {
	method := &lFunction{
		Name: fn.Name.Name,
	}

	paramList := fn.Type.Params.List
	if len(paramList) > 0 && printExpr(paramList[0].Type) == "*lua.LState" {
		method.PassState = true
		paramList = paramList[1:]
	}

	args := make([]ast.Expr, 0)
	for _, field := range paramList {
		for _, _ = range field.Names {
			args = append(args, field.Type)
		}
	}

	method.Args = args
	retList := fn.Type.Results
	switch retList.NumFields() {
	case 0:
	case 1:
		method.Ret = retList.List[0].Type
	default:
		panic("multiple ret values not supported")
	}

	return method
}

func gatherFields(spec *ast.ValueSpec) []*lField {
	fields := make([]*lField, len(spec.Names))
	for i, n := range spec.Names {
		fields[i] = &lField{
			Name: n.Name,
		}
	}
	return fields
}

func extractAlias(comment *ast.CommentGroup) string {
	for _, c := range comment.List {
		if !strings.HasPrefix(c.Text, "//glua:bind") {
			continue
		}
		args := strings.Split(c.Text, " ")[1:]
		if args[0] == "module" {
			return args[1]
		}
	}
	return ""
}

func hasBindComment(comment *ast.CommentGroup) (bool, []string) {
	if comment == nil {
		return false, nil
	}

	for _, c := range comment.List {
		if !strings.HasPrefix(c.Text, "//glua:bind") {
			continue
		}
		return true, strings.Split(c.Text, " ")[1:]
	}
	return false, nil
}

func receiverName(fn *ast.FuncDecl) string {
	if fn.Recv == nil {
		return ""
	}
	recv := fn.Recv.List[0].Type

	if star, ok := recv.(*ast.StarExpr); ok {
		recv = star.X
	}

	return fmt.Sprintf("%v", recv)
}
