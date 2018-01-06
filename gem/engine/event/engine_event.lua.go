package engine_event

import (
	"github.com/gemrs/gem/glua"
	lua "github.com/yuin/gopher-lua"
)

// Bindengine_event generates a lua binding for engine_event
func Bindengine_event(L *lua.LState) {
	L.PreloadModule("gem.engine.event", lBindengine_event)
}

// lBindengine_event generates the table for the engine_event module
func lBindengine_event(L *lua.LState) int {
	mod := L.NewTable()

	L.SetField(mod, "post_tick", glua.ToLua(L, PostTick))

	L.SetField(mod, "pre_tick", glua.ToLua(L, PreTick))

	L.SetField(mod, "shutdown", glua.ToLua(L, Shutdown))

	L.SetField(mod, "startup", glua.ToLua(L, Startup))

	L.SetField(mod, "tick", glua.ToLua(L, Tick))

	L.Push(mod)
	return 1
}
