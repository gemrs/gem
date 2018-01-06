import auth

# Listen addresses
game_server_listen = ":43594"
archive_server_listen = ":43595"

# rsa key
rsa_key_path = "data/devel.key"

# Game data paths
data_path = "./data"
game_data = {
    'data_file': data_path+"/main_file_cache.dat",
    'index_files': [data_path+"/main_file_cache.idx{0}".format(i) for i in range(0, 5)]
}

# Authenticator
auth_provider = auth.DummyProvider()
