// Code generated by glua; DO NOT EDIT.
package auth

import (
	"github.com/gemrs/gem/glua"
	lua "github.com/yuin/gopher-lua"
)

// Bindauth generates a lua binding for auth
func Bindauth(L *lua.LState) {
	L.PreloadModule("gem.game.auth", lBindauth)
}

// lBindauth generates the table for the auth module
func lBindauth(L *lua.LState) int {
	mod := L.NewTable()

	lBindFunc(L, mod)

	L.Push(mod)
	return 1
}

func lBindFunc(L *lua.LState, mod *lua.LTable) {
	mt := L.NewTypeMetatable("auth.Func")

	L.SetField(mt, "__call", L.NewFunction(lNewFunc))

	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), FuncMethods))

	cls := L.NewUserData()
	L.SetField(mod, "Func", cls)
	L.SetMetatable(cls, mt)
	glua.RegisterType("auth.Func", mt)
}

func lNewFunc(L *lua.LState) int {
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(lua.LValue)
	L.Remove(1)
	arg1Value := L.Get(1)
	arg1 := glua.FromLua(arg1Value).(lua.LValue)
	L.Remove(1)
	retVal := NewFunc(L, arg0, arg1)
	L.Push(glua.ToLua(L, retVal))
	return 1

}

var FuncMethods = map[string]lua.LGFunction{}
