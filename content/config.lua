_ = require "underscore"

local cfg = { }

-- Listen addresses
cfg.game_server_listen = ":43594"
cfg.archive_server_listen = ":43595"

-- RSA Key
cfg.rsa_key_path = "data/devel.key"

-- Game data paths
cfg.data_path = "./data"
cfg.game_data_file = cfg.data_path .. "/cache/main_file_cache.dat2"
cfg.game_index_files = {}
for i=0,16 do
   cfg.game_index_files[i+1] = cfg.data_path .. "/cache/main_file_cache.idx" .. tostring(i)
end
cfg.item_definition_file = cfg.data_path .. "/item_definitions.json"
cfg.interface_id_file = cfg.data_path .. "/interface_ids.os157.properties"

-- Plugins to load
cfg.plugins = {
   "plugins.client_init",
   "plugins.debug"
}

return cfg
