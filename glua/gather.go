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
		Name:  name,
		Types: make(map[string]*lType),
	}

	// Ew...
	// Gather types
	for _, file := range pkg.Files {
		for _, decl := range file.Decls {
			if decl, ok := decl.(*ast.GenDecl); ok {
				for _, spec := range decl.Specs {
					if spec, ok := spec.(*ast.TypeSpec); ok {
						if hasBindComment(decl.Doc) {
							typ := gatherType(spec)
							module.Types[spec.Name.Name] = typ
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
				if !hasBindComment(fn.Doc) {
					continue
				}

				receiver := receiverName(fn)
				if receiver == "" {
					// gatherFunction(fn)
					continue
				}

				if typ, ok := module.Types[receiver]; ok {
					typ.Methods[fn.Name.Name] = gatherMethod(fn)
					typ.Methods[fn.Name.Name].Recv = receiver
				}
			}
		}
	}

	return module
}

func gatherType(spec *ast.TypeSpec) *lType {
	typ := &lType{
		Name:    spec.Name.Name,
		Methods: make(map[string]*lMethod),
	}
	return typ
}

func gatherMethod(fn *ast.FuncDecl) *lMethod {
	method := &lMethod{
		Name: fn.Name.Name,
	}

	args := make([]ast.Expr, fn.Type.Params.NumFields())
	for i, field := range fn.Type.Params.List {
		args[i] = field.Type
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

func hasBindComment(comment *ast.CommentGroup) bool {
	return strings.Contains(comment.Text(), "glua:bind")
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
