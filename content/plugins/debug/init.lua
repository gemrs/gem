local event = require("gem.event")
local game_event = require("gem.game.event")
local item = require("gem.game.item")
local position = require("gem.game.position")
local debug_command = require("debug_command")

local function add_item(player, args)
   local size = 1
   if #args == 2 then
      size = tonumber(args[2])
   end
   player:send_message("I'm giving you " .. tostring(size) .. " of the item " .. tostring(args[1]))
   local definition = item.Definition(tonumber(args[1]))
   if definition:stackable() then
      local stack = item.Stack(definition, size)
      player:profile():inventory():add(stack)
   else
      for i=1,size do
         local stack = item.Stack(definition, 1)
         player:profile():inventory():add(stack)
      end
   end
end

local function walk_to(player, args)
   local current = player:profile():position()
   local destination = position.Absolute(tonumber(args[1]), tonumber(args[2]), current:z())
   local success = player:set_walk_destination(destination)
   if success == true then
      player:send_message("Path created")
   else
      player:send_message("No path")
   end
end

local function get_pos(player, args)
   local pos = player:profile():position()
   player:send_message("You're at "..pos:x()..","..pos:y()..","..pos:z())
end

debug_command.register_command("item", add_item)
debug_command.register_command("walk", walk_to)
debug_command.register_command("pos", get_pos)
