local player = require("gem.game.player")
local position = require("gem.game.position")
local auth = require("gem.game.auth")

local function do_auth(name, password)
   local profile = player.Profile(name, password)
   profile:position(position.Absolute(3200, 3200, 0))
   if name == "x" then
      return profile, auth.auth_invalid_credentials
   else
      return profile, auth.auth_okay
   end
end

return {
   authenticate = auth.Func(do_auth)
}
