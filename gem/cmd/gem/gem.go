package main

import (
	"os"

	"github.com/gemrs/gem/gem/runite"
	lua "github.com/yuin/gopher-lua"
)

func main() {
	L := lua.NewState()
	defer L.Close()

	runite.Bindrunite(L)

	if err := L.DoFile(os.Args[1]); err != nil {
		panic(err)
	}
}
