package position

import (
	"github.com/gemrs/gem/glua"
	lua "github.com/yuin/gopher-lua"
)

// Bindposition generates a lua binding for position
func Bindposition(L *lua.LState) {
	L.PreloadModule("gem.game.position", lBindposition)
}

// lBindposition generates the table for the position module
func lBindposition(L *lua.LState) int {
	mod := L.NewTable()

	lBindAbsolute(L, mod)

	L.Push(mod)
	return 1
}

func lBindAbsolute(L *lua.LState, mod *lua.LTable) {
	mt := L.NewTypeMetatable("position.Absolute")
	L.SetField(mt, "__call", L.NewFunction(lNewAbsolute))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), AbsoluteMethods))

	cls := L.NewUserData()
	L.SetField(mod, "Absolute", cls)
	L.SetMetatable(cls, mt)
	glua.RegisterType("position.Absolute", mt)
}

func lNewAbsolute(L *lua.LState) int {
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(int)
	L.Remove(1)
	arg1Value := L.Get(1)
	arg1 := glua.FromLua(arg1Value).(int)
	L.Remove(1)
	arg2Value := L.Get(1)
	arg2 := glua.FromLua(arg2Value).(int)
	L.Remove(1)
	retVal := NewAbsolute(arg0, arg1, arg2)
	L.Push(glua.ToLua(L, retVal))
	return 1

}

var AbsoluteMethods = map[string]lua.LGFunction{

	"x": lBindAbsoluteX,

	"y": lBindAbsoluteY,

	"z": lBindAbsoluteZ,
}

func lBindAbsoluteX(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Absolute)
	L.Remove(1)
	retVal := self.X()
	L.Push(glua.ToLua(L, retVal))
	return 1

}

func lBindAbsoluteY(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Absolute)
	L.Remove(1)
	retVal := self.Y()
	L.Push(glua.ToLua(L, retVal))
	return 1

}

func lBindAbsoluteZ(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Absolute)
	L.Remove(1)
	retVal := self.Z()
	L.Push(glua.ToLua(L, retVal))
	return 1

}
