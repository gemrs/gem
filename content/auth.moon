player = require "gem.game.player"
position = require "gem.game.position"
auth = require "gem.game.auth"

class DummyAuth
  authenticate: =>
    auth.Func (name, password) ->
      profile = player.Profile(name, password)
      profile\position(position.Absolute(3200, 3200, 0))

      if name == "x"
        return profile, auth.auth_invalid_credentials
      else
        return profile, auth.auth_okay

{:DummyAuth}
