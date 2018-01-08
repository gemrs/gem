cfg = {}

-- Listen addresses
cfg.game_server_listen = ":43594"
cfg.archive_server_listen = ":43595"

-- rsa key
cfg.rsa_key_path = "data/devel.key"

-- Game data paths
cfg.data_path = "./data"

cfg.game_data_file = "#{cfg.data_path}/main_file_cache.dat"
cfg.game_index_files = [ "#{cfg.data_path}/main_file_cache.idx#{i}" for i=0,4 ]

cfg.plugins = {
  "plugins.client_init"
}

cfg
