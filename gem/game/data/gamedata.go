//glua:bind module gem.game.data
package data

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/gemrs/willow/log"
	lua "github.com/yuin/gopher-lua"
)

var logger = log.New("data", nil)

//go:generate glua .

var gamedata = map[string]interface{}{}
var gamedataMap = map[string]interface{}{}

//glua:bind
func Load(dataFiles []string) {
	for _, dataFile := range dataFiles {
		_, err := toml.DecodeFile(dataFile, &gamedataMap)
		if err != nil {
			panic(err)
		}
	}
	flattenDataTable("", gamedataMap)
}

//glua:bind
func Get(key string) interface{} {
	if v, ok := gamedata[key]; ok {
		return v
	}
	panic(fmt.Errorf("no such gamedata key: %v", key))
}

//glua:bind
func Int(key string) int {
	if v, ok := Get(key).(int); ok {
		return v
	}
	panic(fmt.Errorf("expected integer type: %v", key))
}

//glua:bind
func AsTable(L *lua.LState) lua.LValue {
	result := L.NewTable()
	toLuaTable(L, result, gamedataMap)
	return result
}

func toLuaTable(L *lua.LState, table *lua.LTable, source map[string]interface{}) {
	for k, v := range source {
		var value lua.LValue
		switch v := v.(type) {
		case string:
			value = lua.LString(v)

		case int64:
			value = lua.LNumber(float64(v))

		case map[string]interface{}:
			subtable := L.NewTable()
			toLuaTable(L, subtable, v)
			value = subtable

		default:
			panic(fmt.Errorf("don't know how to convert type %T to LValue", v))
		}
		L.SetField(table, k, value)
	}
}

func flattenDataTable(namespace string, m map[string]interface{}) {
	for k, v := range m {
		switch v := v.(type) {
		case string:
			gamedata[namespace+k] = v
		case int64:
			gamedata[namespace+k] = int(v)
		case map[string]interface{}:
			flattenDataTable(namespace+k+".", v)
		default:
			panic(fmt.Errorf("don't know how to flatten type %T", v))
		}
	}
}
