local reader = require "pl.config"
local config = require "config"

return reader.read(config.interface_id_file)
