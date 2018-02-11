local log = require("gem.log")
local impl = require("gem.game.impl")
local position = require("gem.game.position")
local auth = require("gem.game.auth")
local protocol = require("gem.game.protocol")
local config = require("config")

local logger = log.Module("auth")

local function create_new_profile(profile)
   profile:position(position.Absolute(3200, 3200, 0))
   local hitpoints = profile:skills():skill(protocol.skill_hitpoints)
   hitpoints:experience(1154)
end

local function profile_path(profile)
   return config.profile_path .. "/" .. profile:username() .. ".json"
end

local function save_profile(profile)
   profile_data = profile:save()
   file = io.open(profile_path(profile), "w")
   file:write(profile_data)
   file:close()
end

local function profile_exists(profile)
   local f=io.open(profile_path(profile), "r")
   if f ~= nil then
      io.close(f)
      return true
   else
      return false
   end
end

local function load_profile(name, password)
   local profile = impl.Profile(name, password)
   if not profile_exists(profile) then
      create_new_profile(profile)
      save_profile(profile)
   else
      file = io.open(profile_path(profile), "r")
      local data = file:read()
      profile:load(data)
      file:close()
   end

   if name == "x" then
      return profile, protocol.auth_invalid_credentials
   else
      return profile, protocol.auth_okay
   end
end

return {
   authenticate = auth.Func(load_profile, save_profile)
}
