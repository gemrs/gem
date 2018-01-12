local log = require("gem.log")
local runite = require("gem.runite")
local archive = require("gem.archive")
local engine = require("gem.engine")
local event = require("gem.event")
local game = require("gem.game")
local server = require("gem.game.server")
local engine_event = require("gem.engine.event")
local game_event = require("gem.game.event")
local item = require("gem.game.item")

local version_string = require("version")
local config = require("config")
local auth = require("auth")
require("interface_config")

local logger = log.Module("lua_main")
logger:info("Starting " .. version_string)

-- Load data files
item.load_definitions(config.item_definition_file)

local data_ctx = runite.Context()
data_ctx:unpack(config.game_data_file, config.game_index_files)

-- Create services
local archive_server = archive.Server()
local game_server = server.Server()
local game_service = game.GameService(data_ctx, config.rsa_key_path, auth.authenticate)
local update_service = game.UpdateService(data_ctx)

-- Start services on engine start
function startup_func()
   game_server:set_service(14, game_service)
   game_server:set_service(15, update_service)
   archive_server:start(config.archive_server_listen, data_ctx)
   game_server:start(config.game_server_listen)
end

engine_event.startup:register(event.Func(startup_func))

-- Stop services on engine stop
function shutdown_func()
   game_server:stop()
   archive_server:stop()
end

engine_event.shutdown:register(event.Func(shutdown_func))

-- Init plugins
function load_plugin(plugin)
   logger:notice("Loading plugin " .. plugin)
   return require(plugin)
end

for _, plugin in ipairs(config.plugins) do
   load_plugin(plugin)
end

-- Start and join the engine
engine = engine.Engine()
engine:start()
return engine:join()
