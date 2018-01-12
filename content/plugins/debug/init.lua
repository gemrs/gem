local event = require("gem.event")
local game_event = require("gem.game.event")
local item = require("gem.game.item")
local debug_command = require("debug_command")

local function add_item(player, args)
   player:send_message("I'm giving you the item " .. tostring(args[1]))
   local definition = item.Definition(tonumber(args[1]))
   local stack = item.Stack(definition, 1)
   player:profile():inventory():add(stack)
end

debug_command.register_command("item", add_item)
