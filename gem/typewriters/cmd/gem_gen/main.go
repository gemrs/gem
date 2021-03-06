package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/clipperhouse/typewriter"
	"github.com/gemrs/gem/gem/typewriters"
	"golang.org/x/tools/imports"
)

var outputFile = flag.String("output", "gem_generated.go", "The output file")

func main() {
	flag.Parse()

	app, err := typewriter.NewApp("+gen")
	if err != nil {
		panic(err)
	}

	for _, pkg := range app.Packages {
		var outputBuffer bytes.Buffer
		fmt.Fprintf(&outputBuffer, "// Code generated by gem_gen; DO NOT EDIT.\n")
		fmt.Fprintf(&outputBuffer, "package %v\n", pkg.Name())
		empty := true

		for _, t := range pkg.Types {
			for _, writer := range app.TypeWriters {
				if collector, ok := writer.(typewriters.TypeCollector); ok {
					collector.Visit(t)
				}

				beforeLength := outputBuffer.Len()
				writer.Write(&outputBuffer, t)
				afterLength := outputBuffer.Len()

				if afterLength > beforeLength {
					empty = false
				}
			}
		}

		for _, writer := range app.TypeWriters {
			if collector, ok := writer.(typewriters.TypeCollector); ok {
				beforeLength := outputBuffer.Len()
				collector.Collect(&outputBuffer)
				afterLength := outputBuffer.Len()

				if afterLength > beforeLength {
					empty = false
				}
			}
		}

		if empty {
			continue
		}

		formatted, err := imports.Process(*outputFile, outputBuffer.Bytes(), nil)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Wrote %v for package %v\n", *outputFile, pkg.Name())
		ioutil.WriteFile(*outputFile, formatted, 0744)
	}

}
