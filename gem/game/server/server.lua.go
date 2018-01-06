package server

import (
	"github.com/gemrs/gem/glua"
	lua "github.com/yuin/gopher-lua"
)

// Bindserver generates a lua binding for server
func Bindserver(L *lua.LState) {
	L.PreloadModule("gem.game.server", lBindserver)
}

// lBindserver generates the table for the server module
func lBindserver(L *lua.LState) int {
	mod := L.NewTable()

	lBindServer(L, mod)

	L.Push(mod)
	return 1
}

func lBindServer(L *lua.LState, mod *lua.LTable) {
	mt := L.NewTypeMetatable("server.Server")
	L.SetField(mt, "__call", L.NewFunction(lNewServer))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), ServerMethods))

	cls := L.NewUserData()
	L.SetField(mod, "Server", cls)
	L.SetMetatable(cls, mt)
	glua.RegisterType("server.Server", mt)
}

func lNewServer(L *lua.LState) int {
	L.Remove(1)
	retVal := NewServer()
	L.Push(glua.ToLua(L, retVal))
	return 1

}

var ServerMethods = map[string]lua.LGFunction{

	"set_service": lBindServerSetService,

	"start": lBindServerStart,

	"stop": lBindServerStop,
}

func lBindServerSetService(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Server)
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(int)
	L.Remove(1)
	arg1Value := L.Get(1)
	arg1 := glua.FromLua(arg1Value).(Service)
	L.Remove(1)
	self.SetService(arg0, arg1)
	return 0

}

func lBindServerStart(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Server)
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(string)
	L.Remove(1)
	retVal := self.Start(arg0)
	L.Push(glua.ToLua(L, retVal))
	return 1

}

func lBindServerStop(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Server)
	L.Remove(1)
	retVal := self.Stop()
	L.Push(glua.ToLua(L, retVal))
	return 1

}