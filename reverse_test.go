package main

import (
	"testing"

	"github.com/MarvinJWendt/testza"
	"github.com/Vilsol/timeless-jewels/calculator"
	"github.com/Vilsol/timeless-jewels/data"
)

// Every passive in radius of the jewel socket the reverse search sweeps.
var reversePassives = []string{
	"mana_spell_block2610",
	"shield_defences1063",
	"super_strength2438",
	"dexterity973",
	"mace_damage_and_armour1568",
	"armour170",
	"tincture_notable_2966",
	"life_es1766",
	"mastery_impale15",
	"tincture_node_2920",
	"strength810",
	"staff_damage_block2572",
	"evasion_phasing2371",
	"strength999",
	"evasion_and_energy_shield1426",
	"weapon_damage1884",
	"mana_curse_notable2595",
	"weapon_damage1883",
	"shield_mastery34",
	"staff_damage_block2571",
	"elemental_damage1725",
	"staff_damage_block2582",
	"minion_damage1127",
	"life_reduced_mana_cost1936",
	"strength808",
	"elemental_damage_notable1727",
	"tincture_node_2957",
	"strength809",
	"spell_damage1868",
	"additional_minions1123",
	"evasion_and_energy_shield1424",
	"life_reduced_mana_cost1934",
	"agility876",
	"strength1980",
	"reduced_mana_cost135",
	"strength795",
	"super_strength2439",
	"evasion_and_energy_shield1425",
	"strength801",
	"reduced_mana_cost137",
	"super_dexterity2436",
	"life_reduced_mana_cost1935",
	"minion_damage284",
	"mana_spell_block_notable2611",
	"intelligence945",
	"strength786",
	"strength827",
	"staff_damage_block2570",
	"flask_generic_notable2844_",
	"power_charges2604",
	"savant878",
	"strength1002",
	"strength998",
	"ailmentres_lightning2762",
	"minion_damage1126",
	"evasion_phasing_notable2372",
	"impale_node_2907",
	"block_notable_2934",
	"shield_block_and_block_recovery29",
	"power_charges2603",
	"shield_block32",
}

func TestReverseGloriousVanity(t *testing.T) {
	statIDs := []uint32{statIndex(t, "spell_damage_+%")}
	result := calculator.ReverseSearch(passiveIndices(t, reversePassives), statIDs, data.GloriousVanity, data.Xibaqua, nil)
	testza.AssertLen(t, result, 7294)
	testza.AssertLen(t, result[1001], 2)
	testza.AssertEqual(t, int32(7), result[1001][passiveIndex(t, "weapon_damage1884")][statIDs[0]])
}

func TestReverseElegantHubris(t *testing.T) {
	statIDs := []uint32{statIndex(t, "spell_damage_+%")}
	result := calculator.ReverseSearch(passiveIndices(t, reversePassives), statIDs, data.ElegantHubris, data.Cadiro, nil)
	testza.AssertLen(t, result, 2375)
	testza.AssertLen(t, result[45060], 1)
	testza.AssertEqual(t, int32(80), result[45060][passiveIndex(t, "tincture_notable_2966")][statIDs[0]])
}

func BenchmarkGloriousVanity(b *testing.B) {
	passives := passiveIndices(b, reversePassives)
	stats := []uint32{statIndex(b, "monster_is_beyond_boss")}
	b.ReportAllocs()
	b.Run("cached", func(b *testing.B) {
		calculator.ReverseSearch(passives, stats, data.GloriousVanity, data.Xibaqua, nil)
		b.ResetTimer()
		for range b.N {
			calculator.ReverseSearch(passives, stats, data.GloriousVanity, data.Xibaqua, nil)
		}
	})

	b.Run("uncached", func(b *testing.B) {
		for range b.N {
			calculator.ClearCache()
			calculator.ReverseSearch(passives, stats, data.GloriousVanity, data.Xibaqua, nil)
		}
	})
}

func BenchmarkElegantHubris(b *testing.B) {
	passives := passiveIndices(b, reversePassives)
	stats := []uint32{statIndex(b, "spell_damage_+%")}
	b.ReportAllocs()
	b.Run("cached", func(b *testing.B) {
		calculator.ReverseSearch(passives, stats, data.ElegantHubris, data.Cadiro, nil)
		b.ResetTimer()
		for range b.N {
			calculator.ReverseSearch(passives, stats, data.ElegantHubris, data.Cadiro, nil)
		}
	})

	b.Run("uncached", func(b *testing.B) {
		for range b.N {
			calculator.ClearCache()
			calculator.ReverseSearch(passives, stats, data.ElegantHubris, data.Cadiro, nil)
		}
	})
}
