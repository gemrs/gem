package runite

import (
	"github.com/gemrs/gem/glua"
	lua "github.com/yuin/gopher-lua"
)

// Bindrunite generates a lua binding for runite
func Bindrunite(L *lua.LState) {
	L.PreloadModule("runite", lBindrunite)
}

// lBindrunite generates the table for the runite module
func lBindrunite(L *lua.LState) int {
	mod := L.NewTable()

	lBindContext(L, mod)

	L.Push(mod)
	return 1
}

func lBindContext(L *lua.LState, mod *lua.LTable) {
	mt := L.NewTypeMetatable("Context")
	L.SetField(mt, "__call", L.NewFunction(lNewContext))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), ContextMethods))

	cls := L.NewUserData()
	L.SetField(mod, "Context", cls)
	L.SetMetatable(cls, mt)
}

func lNewContext(L *lua.LState) int {
	// FIXME only works for structs, no custom constructor..
	obj := &Context{}
	ud := L.NewUserData()
	ud.Value = obj
	L.SetMetatable(ud, L.GetTypeMetatable("Context"))
	L.Push(ud)
	return 1
}

var ContextMethods = map[string]lua.LGFunction{

	"Unpack": lBindContextUnpack,
}

func lBindContextUnpack(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Context)
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(string)
	L.Remove(1)
	arg1Value := L.Get(1)
	arg1ValueArray := arg1Value.(*lua.LTable)
	arg1 := make([]string, arg1ValueArray.Len())
	arg1ValueArray.ForEach(func(k lua.LValue, val lua.LValue) {
		i := int(lua.LVAsNumber(k)) - 1
		arg1[i] = glua.FromLua(val).(string)
	})
	L.Remove(1)
	retVal := self.Unpack(arg0, arg1)
	L.Push(glua.ToLua(L, retVal))
	return 1

}
