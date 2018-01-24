local player = require("gem.game.player")
local position = require("gem.game.position")
local auth = require("gem.game.auth")
local protocol = require("gem.game.protocol")

local function do_auth(name, password)
   local profile = player.Profile(name, password)
   profile:position(position.Absolute(3200, 3200, 0))
   if name == "x" then
      return profile, protocol.auth_invalid_credentials
   else
      return profile, protocol.auth_okay
   end
end

return {
   authenticate = auth.Func(do_auth)
}
