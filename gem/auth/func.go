package auth

import (
	"github.com/gemrs/gem/gem/game/player"
	"github.com/gemrs/gem/gem/protocol"
	"github.com/gemrs/gem/glua"
	lua "github.com/yuin/gopher-lua"
)

//go:generate glua .

//glua:bind
type Func struct {
	fn func(name, password string) (*player.Profile, protocol.AuthResponse)
}

//glua:bind constructor Func
func NewFunc(L *lua.LState, cb lua.LValue) *Func {
	fn := func(name, password string) (*player.Profile, protocol.AuthResponse) {
		if err := L.CallByParam(lua.P{
			Fn:      cb,
			NRet:    2,
			Protect: true,
		}, glua.ToLua(L, name), glua.ToLua(L, password)); err != nil {
			panic(err)
		}
		profile := glua.FromLua(L.Get(1)).(*player.Profile)
		response := glua.FromLua(L.Get(2)).(protocol.AuthResponse)
		L.Pop(2)
		return profile, response
	}
	return &Func{fn}
}

func (f *Func) LookupProfile(name, password string) (*player.Profile, protocol.AuthResponse) {
	return f.fn(name, password)
}
