local event = require("gem.event")
local game_event = require("gem.game.event")
local game_impl = require("gem.game.impl")
local protocol = require("gem.game.protocol")
local id = require("id")

function player_init(event, player)
   player:send_message("Welcome to Gielinor!")

   local config = player:client_config()
   config:set_tab_interface(protocol.tab_attack, id.widget.attack_group_id)
   config:set_tab_interface(protocol.tab_skills, id.widget.skills_group_id)
   config:set_tab_interface(protocol.tab_quests, id.widget.journal_group_id)
   config:set_tab_interface(protocol.tab_inventory, id.widget.inventory_group_id)
   config:set_tab_interface(protocol.tab_equipment, id.widget.equipment_group_id)
   config:set_tab_interface(protocol.tab_prayer, id.widget.prayer_group_id)
   config:set_tab_interface(protocol.tab_magic, id.widget.spellbook_group_id)
   config:set_tab_interface(protocol.tab_friends, id.widget.friends_group_id)
   config:set_tab_interface(protocol.tab_ignore, id.widget.ignore_group_id)
   config:set_tab_interface(protocol.tab_logout, id.widget.logout_group_id)
   config:set_tab_interface(protocol.tab_settings, id.widget.settings_group_id)
   config:set_tab_interface(protocol.tab_run, id.widget.run_group_id)
   config:set_tab_interface(protocol.tab_music, id.widget.music_group_id)
   config:set_tab_interface(protocol.tab_item_bag, id.widget.item_bag_group_id)
   config:set_tab_interface(protocol.tab_clan_chat, id.widget.clan_chat_group_id)
   config:set_tab_interface(protocol.tab_emotes, id.widget.emotes_group_id)
end

game_event.player_login:register(event.Func(player_init))

function item_action(event, player, stack, slot, action)
   player:send_message("item " .. tostring(stack:definition():id()) .. " action " .. tostring(action))
end

game_event.player_inventory_action:register(event.Func(item_action))
