local event = require("gem.event")
local game_event = require("gem.game.event")
local item = require("gem.game.item")
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

debug_command.register_command("item", add_item)
