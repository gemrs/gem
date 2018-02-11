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
cfg.game_metadata_file = cfg.data_path .. "/cache/main_file_cache.idx255"
cfg.game_index_files = {}
for i=0,16 do
   cfg.game_index_files[i+1] = cfg.data_path .. "/cache/main_file_cache.idx" .. tostring(i)
end

cfg.id_files = {
   cfg.data_path .. "/widgets.toml",
   cfg.data_path .. "/inventories.toml",
}
cfg.map_key_file = cfg.data_path .. "/map_keys.json"
cfg.equipment_data_file = cfg.data_path .. "/equipment.json"
cfg.weapon_data_file = cfg.data_path .. "/weapons.json"

cfg.profile_path = "./profiles"

-- Plugins to load
cfg.plugins = {
   "plugins.client_init",
   "plugins.debug"
}

return cfg
