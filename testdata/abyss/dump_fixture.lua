-- Dumps Path of Building's Abyss LUT answers for abyss_lut.txt.gz.
-- usage: cd <PathOfBuilding>/src && luajit dump_fixture.lua <out> 100,4242,8000 && gzip -9 <out>
--   Z <seed> <node> <comps>        every ABYN (Zorath) block; an empty modification is "-"
--   Z <seed> <node> -              every tree keystone, which has no ABYN block
--   S <jt> <seed> <node> <comps>   ABYS (7-10), union over sockets; asserts sockets agree
-- comps: type:id:roll,roll;...  type 1 replaces (id = alternate skill + 337), type 2 adds
local outPath = arg[1]
local seeds = {}
for s in arg[2]:gmatch("%d+") do seeds[#seeds+1] = tonumber(s) end

dofile("HeadlessWrapper.lua")

local function enc(m)
	if not next(m) then return "-" end
	local t = {}
	for i, c in ipairs(m) do t[i] = c.type .. ":" .. c.id .. ":" .. table.concat(c.rolls, ",") end
	return table.concat(t, ";")
end
local function sortedKeys(tbl)
	local keys = {}
	for k in pairs(tbl) do keys[#keys+1] = k end
	table.sort(keys, function(a, b) return tostring(a) < tostring(b) end)
	return keys
end

local out = io.open(outPath, "w")
for _, seed in ipairs(seeds) do
	data.readAbyssJewelLUT(seed, 0, 11, {})
	local lut = data.timelessJewelLUTs[11]
	local nodes = {}
	for k in pairs(lut.blockOffsets) do nodes[#nodes+1] = k end
	table.sort(nodes)
	-- one node per call so an empty modification is visible (the reader drops it)
	for _, n in ipairs(nodes) do
		local res = data.readAbyssJewelLUT(seed, n, 11, { [n] = true }, "__none__")
		out:write(string.format("Z %d %d %s\n", seed, n, enc(res[n] or {})))
	end
	local keystones = {}
	for id, node in pairs(build.spec.tree.nodes) do
		if type(id) == "number" and node.type == "Keystone" and not lut.blockOffsets[id] then keystones[#keystones+1] = id end
	end
	table.sort(keystones)
	for _, n in ipairs(keystones) do
		local res = data.readAbyssJewelLUT(seed, n, 11, { [n] = true }, "__none__")
		assert(not res[n], "keystone has a modification")
		out:write(string.format("Z %d %d -\n", seed, n))
	end
	for jt = 7, 10 do
		data.readAbyssJewelLUT(seed, 0, jt, nil)
		local l = data.timelessJewelLUTs[jt]
		local sockets = {}
		for k in pairs(l.blockOffsets) do sockets[#sockets+1] = k end
		table.sort(sockets)
		local seen = {}
		for _, s in ipairs(sockets) do
			for n, m in pairs(data.readAbyssJewelLUT(seed, s, jt, nil)) do
				local e = enc(m)
				assert(not seen[n] or seen[n] == e, "socket conflict")
				seen[n] = e
			end
		end
		local ns = {}
		for n in pairs(seen) do ns[#ns+1] = n end
		table.sort(ns)
		for _, n in ipairs(ns) do out:write(string.format("S %d %d %d %s\n", jt, seed, n, seen[n])) end
	end
end
out:close()
