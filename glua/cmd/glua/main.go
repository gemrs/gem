package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gemrs/gem/glua"
	"golang.org/x/tools/imports"
)

func main() {
	path := os.Args[1]
	modules, err := glua.Gather(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	for _, module := range modules {
		srcName := module.Name + ".lua.go"
		src, err := imports.Process(srcName, []byte(fmt.Sprint(module)), nil)
		if err != nil {
			fmt.Print(module)
			panic(err)
		}

		ioutil.WriteFile(path+"/"+srcName, src, 0744)
	}
}
