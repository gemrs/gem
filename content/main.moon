import version_string from require("version")
log = require "log"
config = require "config"
runite = require "runite"
archive = require "archive"

logger = log.Module "lua_main"
logger\info "Starting #{version_string}"

ctx = runite.Context!
ctx\unpack config.game_data_file, config.game_index_files

archive_server = archive.Server!
archive_server\start config.archive_server_listen, ctx
