package event

import (
	"github.com/gemrs/gem/gem/util/expire"
	"github.com/gemrs/gem/glua"
	lua "github.com/yuin/gopher-lua"
)

//glua:bind
type Func struct {
	*Observer
}

//glua:bind constructor Func
func NewFunc(L *lua.LState, fn lua.LValue) *Func {
	cb := func(event *Event, argsIface ...interface{}) {
		args := make([]lua.LValue, len(argsIface))
		for i, a := range argsIface {
			args[i] = glua.ToLua(L, a)
		}
		args = append([]lua.LValue{glua.ToLua(L, event)}, args...)
		if err := L.CallByParam(lua.P{
			Fn:      fn,
			NRet:    0,
			Protect: true,
		}, args...); err != nil {
			panic(err)
		}
	}

	return &Func{NewObserver(expire.NewNonExpirable(), cb)}
}
