log = require "gem.log"
runite = require "gem.runite"
archive = require "gem.archive"
engine = require "gem.engine"
event = require "gem.event"
engine_event = require "gem.engine.event"

import version_string from require("version")
config = require "config"

logger = log.Module "lua_main"
logger\info "Starting #{version_string}"

data_ctx = runite.Context!
data_ctx\unpack config.game_data_file, config.game_index_files

archive_server = archive.Server!

startup_func = event.Func () ->
  archive_server\start config.archive_server_listen, data_ctx

shutdown_func = event.Func () ->
  archive_server\stop!

engine_event.startup\register startup_func
engine_event.shutdown\register shutdown_func


engine = engine.Engine!
engine\start!
engine\join!
