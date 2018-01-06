runite = require "runite"
log = require "log"

logger = log.Module "lua_main"
logger\Info "Starting Runite"

ctx = runite.Context!
ctx\Unpack "./data/main_file_cache.dat", [ "./data/main_file_cache.idx#{i}" for i=0,4 ]
