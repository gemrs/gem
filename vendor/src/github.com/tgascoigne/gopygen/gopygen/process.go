package gopygen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
)

func Process(filename string, types []string) (string, error) {
	// Create the AST
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return "", err
	}

	defer func() {
		if r := recover(); r != nil {
			//			ast.Print(fset, f)
			panic(r)
		}
	}()

	// Begin our buffer and print the header
	var filebuffer bytes.Buffer

	file := NewFile(fset, types)
	ast.Walk(file, f)
	fmt.Fprintf(&filebuffer, "%s", file)

	return doFormat(filebuffer.String()), nil
}

func doFormat(in string) string {
	// The callee doesn't recieve errors, because we want to write bad output for debugging
	src, err := format.Source([]byte(in))
	if err != nil {
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		return in
	}
	return string(src)
}
