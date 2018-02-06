// Code generated by glua; DO NOT EDIT.
package data

import (
	"github.com/gemrs/gem/gem/runite"
	"github.com/gemrs/gem/glua"
	lua "github.com/yuin/gopher-lua"
)

// Binddata generates a lua binding for data
func Binddata(L *lua.LState) {
	L.PreloadModule("gem.game.data", lBinddata)
}

// lBinddata generates the table for the data module
func lBinddata(L *lua.LState) int {
	mod := L.NewTable()

	L.SetField(mod, "as_table", L.NewFunction(lBindAsTable))

	L.SetField(mod, "get", L.NewFunction(lBindGet))

	L.SetField(mod, "int", L.NewFunction(lBindInt))

	L.SetField(mod, "load", L.NewFunction(lBindLoad))

	L.SetField(mod, "load_config", L.NewFunction(lBindLoadConfig))

	L.SetField(mod, "load_huffman_table", L.NewFunction(lBindLoadHuffmanTable))

	L.SetField(mod, "load_map", L.NewFunction(lBindLoadMap))

	L.SetField(mod, "load_map_keys", L.NewFunction(lBindLoadMapKeys))

	L.Push(mod)
	return 1
}

func lBindAsTable(L *lua.LState) int {
	retVal := AsTable(L)
	L.Push(glua.ToLua(L, retVal))
	return 1

}

func lBindGet(L *lua.LState) int {
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(string)
	L.Remove(1)
	retVal := Get(arg0)
	L.Push(glua.ToLua(L, retVal))
	return 1

}

func lBindInt(L *lua.LState) int {
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(string)
	L.Remove(1)
	retVal := Int(arg0)
	L.Push(glua.ToLua(L, retVal))
	return 1

}

func lBindLoad(L *lua.LState) int {
	arg0Value := L.Get(1)
	arg0ValueArray := arg0Value.(*lua.LTable)
	arg0 := make([]string, arg0ValueArray.Len())
	arg0ValueArray.ForEach(func(k lua.LValue, val lua.LValue) {
		i := int(lua.LVAsNumber(k)) - 1
		arg0[i] = glua.FromLua(val).(string)
	})
	L.Remove(1)
	Load(arg0)
	return 0

}

func lBindLoadConfig(L *lua.LState) int {
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(*runite.Context)
	L.Remove(1)
	retVal := LoadConfig(arg0)
	L.Push(glua.ToLua(L, retVal))
	return 1

}

func lBindLoadHuffmanTable(L *lua.LState) int {
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(*runite.Context)
	L.Remove(1)
	retVal := LoadHuffmanTable(arg0)
	L.Push(glua.ToLua(L, retVal))
	return 1

}

func lBindLoadMap(L *lua.LState) int {
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(*runite.Context)
	L.Remove(1)
	retVal := LoadMap(arg0)
	L.Push(glua.ToLua(L, retVal))
	return 1

}

func lBindLoadMapKeys(L *lua.LState) int {
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(string)
	L.Remove(1)
	LoadMapKeys(arg0)
	return 0

}
