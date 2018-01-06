package glua

import (
	"fmt"
	"reflect"

	lua "github.com/yuin/gopher-lua"
)

var typeRegistry = map[string]lua.LValue{}

func RegisterType(name string, metatable lua.LValue) {
	if _, ok := typeRegistry[name]; ok {
		panic(fmt.Sprintf("duplicate type registered: %v", name))
	}

	typeRegistry[name] = metatable
}

func FromLua(lv lua.LValue) interface{} {
	switch lv := lv.(type) {
	case lua.LString:
		return lua.LVAsString(lv)
	case lua.LBool:
		return lua.LVAsBool(lv)
	case lua.LNumber:
		//FIXME wont work for float
		return int(lua.LVAsNumber(lv))
	case *lua.LUserData:
		return lv.Value
	case *lua.LFunction:
		return lv
	case *lua.LNilType:
		return nil
	default:
		panic(fmt.Sprintf("don't know how to convert %v to native type", lv))
	}
}

func ToLua(L *lua.LState, v interface{}) lua.LValue {
	switch v := v.(type) {
	case string:
		return lua.LString(v)
	case bool:
		if v {
			return lua.LTrue
		} else {
			return lua.LFalse
		}
	case int:
		return lua.LNumber(v)
	case int8:
		return lua.LNumber(v)
	case int16:
		return lua.LNumber(v)
	case int32:
		return lua.LNumber(v)
	case int64:
		return lua.LNumber(v)
	case uint:
		return lua.LNumber(v)
	case uint8:
		return lua.LNumber(v)
	case uint16:
		return lua.LNumber(v)
	case uint32:
		return lua.LNumber(v)
	case uint64:
		return lua.LNumber(v)
	case float32:
		return lua.LNumber(v)
	case float64:
		return lua.LNumber(v)
	default:
		if v == nil {
			return lua.LNil
		}

		if err, ok := v.(error); ok {
			if v != nil {
				L.Error(lua.LString(err.Error()), 0)
			}
			return lua.LNil
		}

		rType := reflect.TypeOf(v)
		if rType.Kind() == reflect.Ptr {
			rType = rType.Elem()
		}
		typ := rType.String()
		ud := L.NewUserData()
		ud.Value = v
		mt, ok := typeRegistry[typ]
		if ok == false {
			panic(fmt.Sprintf("can't find metatable for %v, not a bound type?", typ))
		}
		L.SetMetatable(ud, mt)
		return ud
	}
}
