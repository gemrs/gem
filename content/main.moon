runite = require "runite"

ctx = runite.Context!
ctx\Unpack "./data/main_file_cache.dat", [ "./data/main_file_cache.idx#{i}" for i=0,4 ]
