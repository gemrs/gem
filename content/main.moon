log = require "gem.log"
runite = require "gem.runite"
archive = require "gem.archive"
engine = require "gem.engine"
event = require "gem.event"
game = require "gem.game"
player = require "gem.game.player"
position = require "gem.game.position"
auth = require "gem.game.auth"
server = require "gem.game.server"
engine_event = require "gem.engine.event"

import version_string from require("version")
config = require "config"

logger = log.Module "lua_main"
logger\info "Starting #{version_string}"

data_ctx = runite.Context!
data_ctx\unpack config.game_data_file, config.game_index_files

auth_func = auth.Func (name, password) ->
  profile = player.Profile(name, password)
  profile\position(position.Absolute(3200, 3200 ,0))
  if name == "x"
    return profile, auth.auth_invalid_credentials
  else
    return profile, auth.auth_okay

archive_server = archive.Server!
game_server = server.Server!
game_service = game.GameService data_ctx, config.rsa_key_path, auth_func
update_service = game.UpdateService data_ctx

engine_event.startup\register event.Func () ->
  game_server\set_service 14, game_service
  game_server\set_service 15, update_service

  archive_server\start config.archive_server_listen, data_ctx
  game_server\start config.game_server_listen

engine_event.shutdown\register event.Func () ->
  archive_server\stop!

engine = engine.Engine!
engine\start!
engine\join!
