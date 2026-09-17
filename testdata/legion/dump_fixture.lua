-- Dumps Path of Building's legion LUT answers (jewel types 1-6) for legion_draws.txt.gz.
-- usage: cd <PathOfBuilding>/src &&
--   LUA_PATH="../runtime/lua/?.lua;../runtime/lua/?/init.lua;../.luarocks/share/lua/5.1/?.lua;;" \
--   LUA_CPATH="../.luarocks/lib/lua/5.1/?.so;;" \
--   luajit dump_fixture.lua <keyfile> <out> && gzip -9n <out>
-- keyfile lines are "<D|N> <jt> <seed> <node>"; each writes "<D|N> <jt> <seed> <node> <v,v,...>",
-- the raw data.readLUT bytes (ids already global). D = modulo and MSVC draws disagree on this
-- record, N = a neighbour where they agree. An empty answer is refused: it would grade nothing.
local keysPath, outPath = arg[1], arg[2]

dofile("HeadlessWrapper.lua")

local out = assert(io.open(outPath, "w"))
for line in io.lines(keysPath) do
	local tag, jt, seed, node = line:match("^([DN]) (%d+) (%d+) (%d+)$")
	assert(tag, "bad key line: " .. line)
	local r = data.readLUT(tonumber(seed), tonumber(node), tonumber(jt))
	assert(r and #r > 0, "empty LUT answer for " .. line)
	out:write(string.format("%s %s %s %s %s\n", tag, jt, seed, node, table.concat(r, ",")))
end
out:close()
