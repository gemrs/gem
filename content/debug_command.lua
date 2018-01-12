local stringx = require "pl.stringx"

local event = require("gem.event")
local game_event = require("gem.game.event")

local commands = {}

function register_command(command, callback)
   commands[command] = callback
end

function dispatch_command(event, player, command)
   local parts = stringx.split(command, " ")
   local command = parts[1]
   local args = _.rest(parts, 2)
   local handler = commands[command]
   if handler ~= nil then
      handler(player, args)
   end
end

game_event.player_command:register(event.Func(dispatch_command))

return {
   register_command = register_command
}
