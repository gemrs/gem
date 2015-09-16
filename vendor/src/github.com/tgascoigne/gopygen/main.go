package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/tgascoigne/gopygen/gopygen"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("gopygen: ")
	flag.Parse()

	// BUG(tom): For now, just accept a single file as an argument
	args := flag.Args()
	filename := args[0]
	types := args[1:]

	src, err := gopygen.Process(filename, types)
	if err != nil {
		log.Fatalf("processing input: %s", err)
	}

	basename := strings.TrimSuffix(filename, filepath.Ext(filename))
	outputName := fmt.Sprintf("%v_gopy.go", basename)

	err = ioutil.WriteFile(outputName, []byte(src), 0644)
	if err != nil {
		log.Fatalf("writing output: %s", err)
	}

	log.Printf("wrote output: %v", outputName)
}
