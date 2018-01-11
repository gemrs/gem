local event = require("gem.event")
local game_event = require("gem.game.event")
local game_player = require("gem.game.player")

function player_init(event, player)
   player:send_message("Welcome to Gielinor!")

   local config = player:client_config()
   config:set_tab_interface(game_player.tab_attack, 2423)
   config:set_tab_interface(game_player.tab_skills, 3917)
   config:set_tab_interface(game_player.tab_quests, 638)
   config:set_tab_interface(game_player.tab_inventory, 3213)
   config:set_tab_interface(game_player.tab_equipment, 1644)
   config:set_tab_interface(game_player.tab_prayer, 5608)
   config:set_tab_interface(game_player.tab_magic, 1151)
   config:set_tab_interface(game_player.tab_friends, 5065)
   config:set_tab_interface(game_player.tab_ignore, 5715)
   config:set_tab_interface(game_player.tab_logout, 2449)
   config:set_tab_interface(game_player.tab_settings, 4445)
   config:set_tab_interface(game_player.tab_run, 147)
   config:set_tab_interface(game_player.tab_music, 6299)

   local hitpoints = player:profile():skills():skill(game_player.skill_hitpoints)
   hitpoints:experience(1154)
end

game_event.player_login:register(event.Func(player_init))
