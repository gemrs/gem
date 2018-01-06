log = require "gem.log"
runite = require "gem.runite"
archive = require "gem.archive"
engine = require "gem.engine"
event = require "gem.event"
game = require "gem.game"
server = require "gem.game.server"
engine_event = require "gem.engine.event"

import version_string from require("version")
config = require "config"
auth = require "auth"

logger = log.Module "lua_main"
logger\info "Starting #{version_string}"

-- Load data files
data_ctx = runite.Context!
data_ctx\unpack config.game_data_file, config.game_index_files

-- Setup authentication
authenticator = auth.DummyAuth!

-- Create services
archive_server = archive.Server!
game_server = server.Server!
game_service = game.GameService data_ctx, config.rsa_key_path, authenticator\authenticate!
update_service = game.UpdateService data_ctx

-- Start services on engine start
engine_event.startup\register event.Func () ->
  game_server\set_service 14, game_service
  game_server\set_service 15, update_service

  archive_server\start config.archive_server_listen, data_ctx
  game_server\start config.game_server_listen

-- Stop services on engine stop
engine_event.shutdown\register event.Func () ->
  game_server\stop!
  archive_server\stop!

-- Start and join the engine
engine = engine.Engine!
engine\start!
engine\join!
