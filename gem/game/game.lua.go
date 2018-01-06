package game

import (
	"github.com/gemrs/gem/gem/auth"
	"github.com/gemrs/gem/gem/runite"
	"github.com/gemrs/gem/glua"
	lua "github.com/yuin/gopher-lua"
)

// Bindgame generates a lua binding for game
func Bindgame(L *lua.LState) {
	L.PreloadModule("gem.game", lBindgame)
}

// lBindgame generates the table for the game module
func lBindgame(L *lua.LState) int {
	mod := L.NewTable()

	lBindGameService(L, mod)

	lBindUpdateService(L, mod)

	L.Push(mod)
	return 1
}

func lBindGameService(L *lua.LState, mod *lua.LTable) {
	mt := L.NewTypeMetatable("game.GameService")
	L.SetField(mt, "__call", L.NewFunction(lNewGameService))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), GameServiceMethods))

	cls := L.NewUserData()
	L.SetField(mod, "GameService", cls)
	L.SetMetatable(cls, mt)
	glua.RegisterType("game.GameService", mt)
}

func lNewGameService(L *lua.LState) int {
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(*runite.Context)
	L.Remove(1)
	arg1Value := L.Get(1)
	arg1 := glua.FromLua(arg1Value).(string)
	L.Remove(1)
	arg2Value := L.Get(1)
	arg2 := glua.FromLua(arg2Value).(auth.Provider)
	L.Remove(1)
	retVal := NewGameService(arg0, arg1, arg2)
	L.Push(glua.ToLua(L, retVal))
	return 1

}

var GameServiceMethods = map[string]lua.LGFunction{}

func lBindUpdateService(L *lua.LState, mod *lua.LTable) {
	mt := L.NewTypeMetatable("game.UpdateService")
	L.SetField(mt, "__call", L.NewFunction(lNewUpdateService))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), UpdateServiceMethods))

	cls := L.NewUserData()
	L.SetField(mod, "UpdateService", cls)
	L.SetMetatable(cls, mt)
	glua.RegisterType("game.UpdateService", mt)
}

func lNewUpdateService(L *lua.LState) int {
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(*runite.Context)
	L.Remove(1)
	retVal := NewUpdateService(arg0)
	L.Push(glua.ToLua(L, retVal))
	return 1

}

var UpdateServiceMethods = map[string]lua.LGFunction{}
