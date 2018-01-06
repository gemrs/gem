package log

import (
	"github.com/gemrs/gem/glua"
	lua "github.com/yuin/gopher-lua"
)

// Bindlog generates a lua binding for log
func Bindlog(L *lua.LState) {
	L.PreloadModule("gem.log", lBindlog)
}

// lBindlog generates the table for the log module
func lBindlog(L *lua.LState) int {
	mod := L.NewTable()

	lBindModule(L, mod)

	L.Push(mod)
	return 1
}

func lBindModule(L *lua.LState, mod *lua.LTable) {
	mt := L.NewTypeMetatable("log.Module")
	L.SetField(mt, "__call", L.NewFunction(lNewModule))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), ModuleMethods))

	cls := L.NewUserData()
	L.SetField(mod, "Module", cls)
	L.SetMetatable(cls, mt)
	glua.RegisterType("log.Module", mt)
}

func lNewModule(L *lua.LState) int {
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(string)
	L.Remove(1)
	retVal := NewModule(arg0)
	L.Push(glua.ToLua(L, retVal))
	return 1

}

/*
func lNewModule(L *lua.LState) int {
	// FIXME only works for structs, no custom constructor..
	obj := &Module{}
	ud := L.NewUserData()
	ud.Value = obj
	L.SetMetatable(ud, L.GetTypeMetatable("log.Module"))
	L.Push(ud)
	return 1
}
*/
var ModuleMethods = map[string]lua.LGFunction{

	"debug": lBindModuleDebug,

	"error": lBindModuleError,

	"info": lBindModuleInfo,

	"notice": lBindModuleNotice,
}

func lBindModuleDebug(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Module)
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(string)
	L.Remove(1)
	self.Debug(arg0)
	return 0

}

func lBindModuleError(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Module)
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(string)
	L.Remove(1)
	self.Error(arg0)
	return 0

}

func lBindModuleInfo(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Module)
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(string)
	L.Remove(1)
	self.Info(arg0)
	return 0

}

func lBindModuleNotice(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Module)
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(string)
	L.Remove(1)
	self.Notice(arg0)
	return 0

}
