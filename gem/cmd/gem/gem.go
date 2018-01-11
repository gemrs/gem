package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gemrs/gem/gem/archive"
	"github.com/gemrs/gem/gem/auth"
	"github.com/gemrs/gem/gem/engine"
	engine_event "github.com/gemrs/gem/gem/engine/event"
	"github.com/gemrs/gem/gem/event"
	"github.com/gemrs/gem/gem/game"
	game_event "github.com/gemrs/gem/gem/game/event"
	"github.com/gemrs/gem/gem/game/item"
	"github.com/gemrs/gem/gem/game/player"
	"github.com/gemrs/gem/gem/game/position"
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/log"
	"github.com/gemrs/gem/gem/runite"
	willow "github.com/gemrs/willow/log"
	lua "github.com/yuin/gopher-lua"
)

var contentDir = flag.String("content", "content_out", "the content directory")
var unsafeLua = flag.Bool("lua-unsafe", false, "invoke lua main without pcall")

func main() {
	flag.Parse()

	/* Setup default targets */
	stdoutTarget := willow.NewTextTarget(os.Stdout)
	bufferingTarget := willow.NewBufferingTarget(stdoutTarget)
	willow.Targets["stdout"] = bufferingTarget

	luaPath := fmt.Sprintf("%v/?.lua;%v/?/init.lua;%v", *contentDir, *contentDir, lua.LuaPathDefault)

	os.Setenv(lua.LuaPath, luaPath)
	mainFile := *contentDir + "/main.lua"

	L := lua.NewState()
	defer L.Close()

	runite.Bindrunite(L)
	log.Bindlog(L)
	archive.Bindarchive(L)
	engine.Bindengine(L)
	event.Bindevent(L)
	server.Bindserver(L)
	game.Bindgame(L)
	auth.Bindauth(L)
	player.Bindplayer(L)
	position.Bindposition(L)
	engine_event.Bindengine_event(L)
	game_event.Bindgame_event(L)
	item.Binditem(L)

	if *unsafeLua {
		if fn, err := L.LoadFile(mainFile); err != nil {
			panic(err)
		} else {
			// Specifically not using PCall, as that hides panic traces
			L.Push(fn)
			L.Call(0, 0)
		}
	} else {
		err := L.DoFile(mainFile)
		if err != nil {
			panic(err)
		}
	}
}
