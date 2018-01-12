local event = require("gem.event")
local game_event = require("gem.game.event")
local game_player = require("gem.game.player")
local interfaces = require("id.interface")

function player_init(event, player)
   player:send_message("Welcome to Gielinor!")

   local config = player:client_config()
   config:set_tab_interface(game_player.tab_attack, interfaces.tab_attack)
   config:set_tab_interface(game_player.tab_skills, interfaces.tab_skills)
   config:set_tab_interface(game_player.tab_quests, interfaces.tab_quests)
   config:set_tab_interface(game_player.tab_inventory, interfaces.tab_inventory)
   config:set_tab_interface(game_player.tab_equipment, interfaces.tab_equipment)
   config:set_tab_interface(game_player.tab_prayer, interfaces.tab_prayer)
   config:set_tab_interface(game_player.tab_magic, interfaces.tab_magic)
   config:set_tab_interface(game_player.tab_friends, interfaces.tab_friends)
   config:set_tab_interface(game_player.tab_ignore, interfaces.tab_ignore)
   config:set_tab_interface(game_player.tab_logout, interfaces.tab_logout)
   config:set_tab_interface(game_player.tab_settings, interfaces.tab_settings_lowmem)
   config:set_tab_interface(game_player.tab_run, interfaces.tab_run)
   config:set_tab_interface(game_player.tab_music, interfaces.tab_music_lowmem)

   local hitpoints = player:profile():skills():skill(game_player.skill_hitpoints)
   hitpoints:experience(1154)
end

game_event.player_login:register(event.Func(player_init))

function item_action(event, player, stack, action)
   player:send_message("item " .. tostring(stack:definition():id()) .. " action " .. tostring(action))
end

game_event.player_inventory_action:register(event.Func(item_action))
