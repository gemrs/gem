package auth

import (
	"github.com/gemrs/gem/gem/protocol"
	"github.com/gemrs/gem/glua"
	lua "github.com/yuin/gopher-lua"
)

//go:generate glua .

//glua:bind
type Func struct {
	load func(name, password string) (protocol.Profile, protocol.AuthResponse)
	save func(profile protocol.Profile)
}

//glua:bind constructor Func
func NewFunc(L *lua.LState, loadCallback, saveCallback lua.LValue) *Func {
	loadFunc := func(name, password string) (protocol.Profile, protocol.AuthResponse) {
		if err := L.CallByParam(lua.P{
			Fn:      loadCallback,
			NRet:    2,
			Protect: true,
		}, glua.ToLua(L, name), glua.ToLua(L, password)); err != nil {
			panic(err)
		}
		profile := glua.FromLua(L.Get(1)).(protocol.Profile)
		response := glua.FromLua(L.Get(2)).(protocol.AuthResponse)
		L.Pop(2)
		return profile, response
	}

	saveFunc := func(profile protocol.Profile) {
		if err := L.CallByParam(lua.P{
			Fn:      saveCallback,
			NRet:    0,
			Protect: true,
		}, glua.ToLua(L, profile)); err != nil {
			panic(err)
		}
	}

	return &Func{loadFunc, saveFunc}
}

func (f *Func) LoadProfile(name, password string) (protocol.Profile, protocol.AuthResponse) {
	return f.load(name, password)
}

func (f *Func) SaveProfile(profile protocol.Profile) {
	f.save(profile)
}
