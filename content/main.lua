local log = require("gem.log")
local runite = require("gem.runite")
local engine = require("gem.engine")
local event = require("gem.event")
local game = require("gem.game")
local server = require("gem.game.server")
local engine_event = require("gem.engine.event")
local game_event = require("gem.game.event")
local item = require("gem.game.item")
local data = require("gem.game.data")

local version_string = require("version")
local config = require("config")
local auth = require("auth")

local logger = log.Module("lua_main")
logger:info("Starting " .. version_string)

-- Load data files
local data_ctx = runite.Context()
data_ctx:unpack(config.game_data_file, config.game_index_files, config.game_metadata_file)
data.load(config.id_files)
data.load_config(data_ctx)
data.load_equipment_data(config.equipment_data_file)
data.load_weapon_data(config.weapon_data_file)
data.load_map_keys(config.map_key_file)
data.load_map(data_ctx)
data.load_huffman_table(data_ctx)

-- Create services
local game_server = server.Server()
game.register_services(game_server, data_ctx, config.rsa_key_path, auth.authenticate)

-- Start services on engine start
function startup_func()
   game_server:start(config.game_server_listen)
end

engine_event.startup:register(event.Func(startup_func))

-- Stop services on engine stop
function shutdown_func()
   game_server:stop()
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

-- Load some event handlers
require("frame_buttons")

-- Start and join the engine
engine = engine.Engine()
engine:start()
return engine:join()
