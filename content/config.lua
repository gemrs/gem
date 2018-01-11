local cfg = { }

-- Listen addresses
cfg.game_server_listen = ":43594"
cfg.archive_server_listen = ":43595"

-- RSA Key
cfg.rsa_key_path = "data/devel.key"

-- Game data paths
cfg.data_path = "./data"
cfg.game_data_file = cfg.data_path .. "/main_file_cache.dat"
cfg.game_index_files = {}
for i = 1, 4 do
   cfg.game_index_files[i] = cfg.data_path .. "/main_file_cache.idx" .. tostring(i-1)
end

-- Plugins to load
cfg.plugins = {
  "plugins.client_init"
}

return cfg
