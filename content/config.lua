_ = require 'underscore'

local cfg = { }

-- Listen addresses
cfg.game_server_listen = ":43594"
cfg.archive_server_listen = ":43595"

-- RSA Key
cfg.rsa_key_path = "data/devel.key"

-- Game data paths
cfg.data_path = "./data"
cfg.game_data_file = cfg.data_path .. "/main_file_cache.dat"
cfg.game_index_files = _.map({1,2,3,4}, function(i)
      return cfg.data_path .. "/main_file_cache.idx" .. tostring(i-1)
end)
cfg.item_definition_file = cfg.data_path .. "/item_definitions.json"

-- Plugins to load
cfg.plugins = {
   "plugins.client_init",
   "plugins.debug"
}

return cfg
