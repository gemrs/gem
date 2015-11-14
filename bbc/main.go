package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gemrs/gem/bbc/compile"

	"golang.org/x/tools/imports"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("bbc: binary buffer compiler")
		fmt.Println("Usage: bbc input.bb output.bb.go")
		os.Exit(1)
	}

	pkg := os.Getenv("GOPACKAGE")
	if pkg == "" {
		fmt.Printf("error: $GOPACKAGE unset")
		os.Exit(1)
	}

	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	output, err := compile.Compile(os.Args[1], pkg, string(input))
	if err != nil {
		fmt.Printf("error during compilation: %v\n", err)
		os.Exit(1)
	}

	output, err = imports.Process(os.Args[1], output, nil)
	if err != nil {
		panic(fmt.Sprintf("generated invalid go: %v", err))
	}

	err = ioutil.WriteFile(os.Args[2], []byte(output), 0644)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("bbc: wrote output: %v\n", os.Args[2])
}
