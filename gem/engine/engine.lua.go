package engine

import (
	"github.com/gemrs/gem/glua"
	lua "github.com/yuin/gopher-lua"
)

// Bindengine generates a lua binding for engine
func Bindengine(L *lua.LState) {
	L.PreloadModule("gem.engine", lBindengine)
}

// lBindengine generates the table for the engine module
func lBindengine(L *lua.LState) int {
	mod := L.NewTable()

	lBindEngine(L, mod)

	L.Push(mod)
	return 1
}

func lBindEngine(L *lua.LState, mod *lua.LTable) {
	mt := L.NewTypeMetatable("engine.Engine")
	L.SetField(mt, "__call", L.NewFunction(lNewEngine))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), EngineMethods))

	cls := L.NewUserData()
	L.SetField(mod, "Engine", cls)
	L.SetMetatable(cls, mt)
	glua.RegisterType("engine.Engine", mt)
}

func lNewEngine(L *lua.LState) int {
	L.Remove(1)
	retVal := NewEngine()
	L.Push(glua.ToLua(L, retVal))
	return 1

}

var EngineMethods = map[string]lua.LGFunction{

	"join": lBindEngineJoin,

	"start": lBindEngineStart,

	"stop": lBindEngineStop,
}

func lBindEngineJoin(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Engine)
	L.Remove(1)
	self.Join()
	return 0

}

func lBindEngineStart(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Engine)
	L.Remove(1)
	self.Start()
	return 0

}

func lBindEngineStop(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Engine)
	L.Remove(1)
	self.Stop()
	return 0

}
