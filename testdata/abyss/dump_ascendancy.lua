-- Dumps Path of Building's Zorath ascendancy picks (the LUT's ASCS section) for
-- abyss_ascendancy.txt.gz.
-- usage: cd <PathOfBuilding>/src && luajit dump_ascendancy.lua <out> [step] [full,...] && gzip -9n <out>
--   the committed fixture: luajit dump_ascendancy.lua <out> 97 Ascendant,Deadeye,Guardian
--   <ascendancy> <seed> <node,node,...>   "-" when the seed selects nothing
-- Every step-th seed (default 1) plus the last, and every seed for the ascendancies listed in full.
local outPath = arg[1]
local step = tonumber(arg[2] or "1")
local full = {}
for name in (arg[3] or ""):gmatch("[^,]+") do full[name] = true end

dofile("HeadlessWrapper.lua")

data.readAbyssJewelLUT(100, 0, 11, {})
local lut = data.timelessJewelLUTs[11]
assert(lut and lut.ascendancyOffsets, "no Zorath ascendancy section")

local names = {}
for name in pairs(lut.ascendancyOffsets) do names[#names+1] = name end
table.sort(names)

local out = io.open(outPath, "w")
for _, name in ipairs(names) do
	for seed = lut.seedMinimum, lut.seedMaximum, lut.seedIncrement do
		if full[name] or (seed - lut.seedMinimum) % step == 0 or seed == lut.seedMaximum then
			local nodes = {}
			for id in pairs(data.readAbyssJewelLUT(seed, 0, 11, {}, name)) do nodes[#nodes+1] = id end
			table.sort(nodes)
			out:write(string.format("%s %d %s\n", name, seed, #nodes > 0 and table.concat(nodes, ",") or "-"))
		end
	end
end
out:close()
