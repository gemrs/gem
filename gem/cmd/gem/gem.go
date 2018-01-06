package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/gemrs/gem/gem/archive"
	"github.com/gemrs/gem/gem/log"
	"github.com/gemrs/gem/gem/runite"
	willow "github.com/gemrs/willow/log"
	lua "github.com/yuin/gopher-lua"
)

var contentDir = flag.String("content", "content", "the content directory")
var compiledDir = flag.String("content_out", "content_out", "the compiled content directory")
var noCompile = flag.Bool("no-compile", false, "skip lua compilation")

func buildMoonScript(dir string) {
	fmt.Println("Compiling content directory:", dir)
	out, err := exec.Command("moonc", "-t", *compiledDir, dir).CombinedOutput()
	fmt.Println(string(out))
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()

	/* Setup default targets */
	stdoutTarget := willow.NewTextTarget(os.Stdout)
	bufferingTarget := willow.NewBufferingTarget(stdoutTarget)
	willow.Targets["stdout"] = bufferingTarget

	if !*noCompile {
		buildMoonScript(*contentDir)
	}
	finalDir := *compiledDir + "/" + *contentDir
	luaPath := fmt.Sprintf("%v/?.lua;%v", finalDir, lua.LuaPathDefault)

	os.Setenv(lua.LuaPath, luaPath)
	mainFile := finalDir + "/main.lua"

	L := lua.NewState()
	defer L.Close()

	runite.Bindrunite(L)
	log.Bindlog(L)
	archive.Bindarchive(L)

	if fn, err := L.LoadFile(mainFile); err != nil {
		panic(err)
	} else {
		// Specifically not using PCall, as that hides panic traces
		L.Push(fn)
		L.Call(0, 0)
	}
}
