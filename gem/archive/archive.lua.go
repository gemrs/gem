package archive

import (
	"github.com/gemrs/gem/gem/runite"
	"github.com/gemrs/gem/glua"
	lua "github.com/yuin/gopher-lua"
)

// Bindarchive generates a lua binding for archive
func Bindarchive(L *lua.LState) {
	L.PreloadModule("archive", lBindarchive)
}

// lBindarchive generates the table for the archive module
func lBindarchive(L *lua.LState) int {
	mod := L.NewTable()

	lBindServer(L, mod)

	L.Push(mod)
	return 1
}

func lBindServer(L *lua.LState, mod *lua.LTable) {
	mt := L.NewTypeMetatable("archive.Server")
	L.SetField(mt, "__call", L.NewFunction(lNewServer))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), ServerMethods))

	cls := L.NewUserData()
	L.SetField(mod, "Server", cls)
	L.SetMetatable(cls, mt)
}

func lNewServer(L *lua.LState) int {
	L.Remove(1)
	retVal := NewServer()
	L.Push(glua.ToLua(L, retVal))
	return 1

}

/*
func lNewServer(L *lua.LState) int {
	// FIXME only works for structs, no custom constructor..
	obj := &Server{}
	ud := L.NewUserData()
	ud.Value = obj
	L.SetMetatable(ud, L.GetTypeMetatable("archive.Server"))
	L.Push(ud)
	return 1
}
*/
var ServerMethods = map[string]lua.LGFunction{

	"start": lBindServerStart,

	"stop": lBindServerStop,
}

func lBindServerStart(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Server)
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(string)
	L.Remove(1)
	arg1Value := L.Get(1)
	arg1 := glua.FromLua(arg1Value).(*runite.Context)
	L.Remove(1)
	retVal := self.Start(arg0, arg1)
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
