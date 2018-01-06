package game_events

import (
	"github.com/gemrs/gem/glua"
	lua "github.com/yuin/gopher-lua"
)

// Bindgame_event generates a lua binding for game_event
func Bindgame_event(L *lua.LState) {
	L.PreloadModule("gem.game.event", lBindgame_event)
}

// lBindgame_event generates the table for the game_event module
func lBindgame_event(L *lua.LState) int {
	mod := L.NewTable()

	L.SetField(mod, "entity_region_change", glua.ToLua(L, EntityRegionChange))

	L.SetField(mod, "entity_sector_change", glua.ToLua(L, EntitySectorChange))

	L.SetField(mod, "player_appearance_update", glua.ToLua(L, PlayerAppearanceUpdate))

	L.SetField(mod, "player_finish_login", glua.ToLua(L, PlayerFinishLogin))

	L.SetField(mod, "player_load_profile", glua.ToLua(L, PlayerLoadProfile))

	L.SetField(mod, "player_login", glua.ToLua(L, PlayerLogin))

	L.SetField(mod, "player_logout", glua.ToLua(L, PlayerLogout))

	L.Push(mod)
	return 1
}
