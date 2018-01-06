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

	lBindAuthResponse(L, mod)

	lBindFunc(L, mod)

	L.SetField(mod, "auth_attempts_exceeded", glua.ToLua(L, AuthAttemptsExceeded))

	L.SetField(mod, "auth_bad_session_id", glua.ToLua(L, AuthBadSessionId))

	L.SetField(mod, "auth_delay", glua.ToLua(L, AuthDelay))

	L.SetField(mod, "auth_disabled", glua.ToLua(L, AuthDisabled))

	L.SetField(mod, "auth_duplicate_session", glua.ToLua(L, AuthDuplicateSession))

	L.SetField(mod, "auth_end", glua.ToLua(L, AuthEnd))

	L.SetField(mod, "auth_incomplete", glua.ToLua(L, AuthIncomplete))

	L.SetField(mod, "auth_invalid_credentials", glua.ToLua(L, AuthInvalidCredentials))

	L.SetField(mod, "auth_invalid_login_server", glua.ToLua(L, AuthInvalidLoginServer))

	L.SetField(mod, "auth_invalid_transferring", glua.ToLua(L, AuthInvalidTransferring))

	L.SetField(mod, "auth_members_area", glua.ToLua(L, AuthMembersArea))

	L.SetField(mod, "auth_members_world", glua.ToLua(L, AuthMembersWorld))

	L.SetField(mod, "auth_no_login_server", glua.ToLua(L, AuthNoLoginServer))

	L.SetField(mod, "auth_okay", glua.ToLua(L, AuthOkay))

	L.SetField(mod, "auth_pending", glua.ToLua(L, AuthPending))

	L.SetField(mod, "auth_rejected", glua.ToLua(L, AuthRejected))

	L.SetField(mod, "auth_server_full", glua.ToLua(L, AuthServerFull))

	L.SetField(mod, "auth_too_many_connections", glua.ToLua(L, AuthTooManyConnections))

	L.SetField(mod, "auth_unknown", glua.ToLua(L, AuthUnknown))

	L.SetField(mod, "auth_updates", glua.ToLua(L, AuthUpdates))

	L.SetField(mod, "auth_updating", glua.ToLua(L, AuthUpdating))

	L.Push(mod)
	return 1
}

func lBindAuthResponse(L *lua.LState, mod *lua.LTable) {
	mt := L.NewTypeMetatable("auth.AuthResponse")
	L.SetField(mt, "__call", L.NewFunction(lNewAuthResponse))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), AuthResponseMethods))

	cls := L.NewUserData()
	L.SetField(mod, "AuthResponse", cls)
	L.SetMetatable(cls, mt)
	glua.RegisterType("auth.AuthResponse", mt)
}

func lNewAuthResponse(L *lua.LState) int {
	L.Remove(1)
	arg0Value := L.Get(1)
	arg0 := glua.FromLua(arg0Value).(int)
	L.Remove(1)
	retVal := NewAuthResponse(arg0)
	L.Push(glua.ToLua(L, retVal))
	return 1

}

var AuthResponseMethods = map[string]lua.LGFunction{}

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
	retVal := NewFunc(L, arg0)
	L.Push(glua.ToLua(L, retVal))
	return 1

}

var FuncMethods = map[string]lua.LGFunction{}
