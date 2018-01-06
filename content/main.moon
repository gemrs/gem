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

--archive_server = archive.Server!
--archive_server\start config.archive_server_listen, ctx

engine = engine.Engine!
engine\start!
engine\join!
