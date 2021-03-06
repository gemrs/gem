// Code generated by glua; DO NOT EDIT.
package event

import (
	"github.com/gemrs/gem/glua"
	lua "github.com/yuin/gopher-lua"
)

// Bindevent generates a lua binding for event
func Bindevent(L *lua.LState) {
	L.PreloadModule("gem.event", lBindevent)
}

// lBindevent generates the table for the event module
func lBindevent(L *lua.LState) int {
	mod := L.NewTable()

	lBindEvent(L, mod)

	lBindFunc(L, mod)

	L.Push(mod)
	return 1
}

func lBindEvent(L *lua.LState, mod *lua.LTable) {
	mt := L.NewTypeMetatable("event.Event")

	L.SetField(mt, "__call", L.NewFunction(lNewEvent))

	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), EventMethods))

	cls := L.NewUserData()
	L.SetField(mod, "Event", cls)
	L.SetMetatable(cls, mt)
	glua.RegisterType("event.Event", mt)
}

func lNewEvent(L *lua.LState) int {
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(string)
	L.Remove(1)
	retVal := NewEvent(arg0)
	L.Push(glua.ToLua(L, retVal))
	return 1

}

var EventMethods = map[string]lua.LGFunction{

	"key": lBindEventKey,

	"register": lBindEventRegister,

	"unregister": lBindEventUnregister,
}

func lBindEventKey(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Event)
	L.Remove(1)
	retVal := self.Key()
	L.Push(glua.ToLua(L, retVal))
	return 1

}

func lBindEventRegister(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Event)
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(ObserverIface)
	L.Remove(1)
	self.Register(arg0)
	return 0

}

func lBindEventUnregister(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Event)
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(ObserverIface)
	L.Remove(1)
	self.Unregister(arg0)
	return 0

}

func lBindFunc(L *lua.LState, mod *lua.LTable) {
	mt := L.NewTypeMetatable("event.Func")

	L.SetField(mt, "__call", L.NewFunction(lNewFunc))

	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), FuncMethods))

	cls := L.NewUserData()
	L.SetField(mod, "Func", cls)
	L.SetMetatable(cls, mt)
	glua.RegisterType("event.Func", mt)
}

func lNewFunc(L *lua.LState) int {
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(lua.LValue)
	L.Remove(1)
	retVal := NewFunc(L, arg0)
	L.Push(glua.ToLua(L, retVal))
	return 1

}

var FuncMethods = map[string]lua.LGFunction{}
