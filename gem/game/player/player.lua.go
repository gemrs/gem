package player

import (
	"github.com/gemrs/gem/gem/game/position"
	"github.com/gemrs/gem/gem/game/server"
	"github.com/gemrs/gem/gem/game/world"
	"github.com/gemrs/gem/glua"
	lua "github.com/yuin/gopher-lua"
)

// Bindplayer generates a lua binding for player
func Bindplayer(L *lua.LState) {
	L.PreloadModule("gem.game.player", lBindplayer)
}

// lBindplayer generates the table for the player module
func lBindplayer(L *lua.LState) int {
	mod := L.NewTable()

	lBindPlayer(L, mod)

	lBindProfile(L, mod)

	L.Push(mod)
	return 1
}

func lBindPlayer(L *lua.LState, mod *lua.LTable) {
	mt := L.NewTypeMetatable("player.Player")
	L.SetField(mt, "__call", L.NewFunction(lNewPlayer))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), PlayerMethods))

	cls := L.NewUserData()
	L.SetField(mod, "Player", cls)
	L.SetMetatable(cls, mt)
	glua.RegisterType("player.Player", mt)
}

func lNewPlayer(L *lua.LState) int {
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(*server.Connection)
	L.Remove(1)
	arg1Value := L.Get(1)
	arg1 := glua.FromLua(arg1Value).(*world.Instance)
	L.Remove(1)
	retVal := NewPlayer(arg0, arg1)
	L.Push(glua.ToLua(L, retVal))
	return 1

}

var PlayerMethods = map[string]lua.LGFunction{

	"index": lBindPlayerIndex,

	"profile": lBindPlayerProfile,

	"send_message": lBindPlayerSendMessage,
}

func lBindPlayerIndex(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Player)
	L.Remove(1)
	retVal := self.Index()
	L.Push(glua.ToLua(L, retVal))
	return 1

}

func lBindPlayerProfile(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Player)
	L.Remove(1)
	retVal := self.Profile()
	L.Push(glua.ToLua(L, retVal))
	return 1

}

func lBindPlayerSendMessage(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Player)
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(string)
	L.Remove(1)
	self.SendMessage(arg0)
	return 0

}

func lBindProfile(L *lua.LState, mod *lua.LTable) {
	mt := L.NewTypeMetatable("player.Profile")
	L.SetField(mt, "__call", L.NewFunction(lNewProfile))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), ProfileMethods))

	cls := L.NewUserData()
	L.SetField(mod, "Profile", cls)
	L.SetMetatable(cls, mt)
	glua.RegisterType("player.Profile", mt)
}

func lNewProfile(L *lua.LState) int {
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(string)
	L.Remove(1)
	arg1Value := L.Get(1)
	arg1 := glua.FromLua(arg1Value).(string)
	L.Remove(1)
	retVal := NewProfile(arg0, arg1)
	L.Push(glua.ToLua(L, retVal))
	return 1

}

var ProfileMethods = map[string]lua.LGFunction{

	"password": lBindProfilePassword,

	"rights": lBindProfileRights,

	"username": lBindProfileUsername,

	"position": lBindPropProfilePosition,
}

func lBindProfilePassword(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Profile)
	L.Remove(1)
	retVal := self.Password()
	L.Push(glua.ToLua(L, retVal))
	return 1

}

func lBindProfileRights(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Profile)
	L.Remove(1)
	retVal := self.Rights()
	L.Push(glua.ToLua(L, retVal))
	return 1

}

func lBindProfileUsername(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Profile)
	L.Remove(1)
	retVal := self.Username()
	L.Push(glua.ToLua(L, retVal))
	return 1

}

func lBindPropProfilePosition(L *lua.LState) int {
	self := glua.FromLua(L.Get(1)).(*Profile)
	if L.GetTop() == 2 {
		val := glua.FromLua(L.Get(2)).(*position.Absolute)
		self.SetPosition(val)
		return 0
	}
	L.Push(glua.ToLua(L, self.Position()))
	return 1
}
