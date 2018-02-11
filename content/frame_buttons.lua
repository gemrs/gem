local log = require("gem.log")
local event = require("gem.event")
local game_event = require("gem.game.event")
local id = require("id")

local logger = log.Module("frame_buttons")

local button_handlers = {}

local function button_dispatch(event, player, action, iface, widget)
   if button_handlers[iface] == nil then
      button_handlers[iface] = {}
   end

   if button_handlers[iface][widget] ~= nil then
      button_handlers[iface][widget](player, action)
   else
      player:send_message("you clicked "..iface.." "..widget)
   end
end

game_event.player_widget_action:register(event.Func(button_dispatch))

local function add_button_handler(iface, widget, callback)
   if button_handlers[iface] == nil then
      button_handlers[iface] = {}
   end

   if button_handlers[iface][widget] ~= nil then
      error("Duplicate button handlers for "..iface..", "..widget)
   end

   button_handlers[iface][widget] = callback
end


local function logout_button(player, action)
   player:disconnect()
end

add_button_handler(id.widget.logout_group_id,
                   id.widget.logout_panel.button,
                   logout_button)
