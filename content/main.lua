local runite = require("runite")
local ctx = runite.Context()
return ctx:Unpack("./data/main_file_cachea.dat", (function()
  local _accum_0 = { }
  local _len_0 = 1
  for i = 0, 4 do
    _accum_0[_len_0] = "./data/main_file_cache.idx" .. tostring(i)
    _len_0 = _len_0 + 1
  end
  return _accum_0
end)())
