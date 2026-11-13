package main

import (
	"testing"

	"github.com/MarvinJWendt/testza"
	"github.com/Vilsol/timeless-jewels/calculator"
	"github.com/Vilsol/timeless-jewels/data"
)

func TestGloriousVanity(t *testing.T) {
	const seed = 2000

	tests := []struct {
		jewel     data.JewelType
		conqueror data.Conqueror
		passive   string
		result    data.AlternatePassiveSkillInformation
	}{
		{
			jewel:     data.GloriousVanity,
			conqueror: data.Xibaqua,
			passive:   "supreme_ego_keystone2696", // Supreme Ego
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "vaal_keystone_1"),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.GloriousVanity,
			conqueror: data.Xibaqua,
			passive:   "maximum_power_charges742", // Instability
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "vaal_notable_block_2"),
				StatRolls:             data.StatRolls{8, 22},
			},
		},
		{
			jewel:     data.GloriousVanity,
			conqueror: data.Xibaqua,
			passive:   "intelligence879", // Intelligence
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "vaal_small_fire_resistance"),
				StatRolls:             data.StatRolls{12},
			},
		},
		{
			jewel:     data.GloriousVanity,
			conqueror: data.Xibaqua,
			passive:   "elemental_damage1906", // Elemental Damage
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "vaal_small_movement_speed"),
				StatRolls:             data.StatRolls{3},
			},
		},
		{
			jewel:     data.GloriousVanity,
			conqueror: data.Xibaqua,
			passive:   "eagle_eye199", // Eagle Eye
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "vaal_notable_random_defence"),
				StatRolls:             data.StatRolls{},
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: alternatePassiveAddition(t, "vaal_small_critical_strike_multiplier"),
						StatRolls:                data.StatRolls{6},
					},
					{
						AlternatePassiveAddition: alternatePassiveAddition(t, "vaal_small_mana"),
						StatRolls:                data.StatRolls{5},
					},
					{
						AlternatePassiveAddition: alternatePassiveAddition(t, "vaal_small_lightning_resistance"),
						StatRolls:                data.StatRolls{11},
					},
				},
			},
		},
		{
			jewel:     data.GloriousVanity,
			conqueror: data.Zerphi,
			passive:   "supreme_ego_keystone2696", // Supreme Ego
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "vaal_keystone_2"),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.GloriousVanity,
			conqueror: data.Ahuana,
			passive:   "supreme_ego_keystone2696", // Supreme Ego
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "vaal_keystone_2_v2"),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.GloriousVanity,
			conqueror: data.Doryani,
			passive:   "supreme_ego_keystone2696", // Supreme Ego
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "vaal_keystone_3"),
				StatRolls:             data.StatRolls{1},
			},
		},
	}

	for _, test := range tests {
		t.Run(string(test.conqueror), func(t *testing.T) {
			t.Run(test.passive, func(t *testing.T) {
				testza.AssertEqual(t, test.result, calculator.Calculate(passiveIndex(t, test.passive), seed, test.jewel, test.conqueror))
			})
		})
	}
}

func TestLethalPride(t *testing.T) {
	const seed = 12000

	tests := []struct {
		jewel     data.JewelType
		conqueror data.Conqueror
		passive   string
		result    data.AlternatePassiveSkillInformation
	}{
		{
			jewel:     data.LethalPride,
			conqueror: data.Kaom,
			passive:   "supreme_ego_keystone2696", // Supreme Ego
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "karui_keystone_1"),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.LethalPride,
			conqueror: data.Kaom,
			passive:   "maximum_power_charges742", // Instability
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: alternatePassiveAddition(t, "karui_notable_add_armour"),
						StatRolls:                data.StatRolls{20},
					},
				},
			},
		},
		{
			jewel:     data.LethalPride,
			conqueror: data.Kaom,
			passive:   "intelligence879", // Intelligence
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: alternatePassiveAddition(t, "karui_small_strength"),
						StatRolls:                data.StatRolls{4},
					},
				},
			},
		},
		{
			jewel:     data.LethalPride,
			conqueror: data.Kaom,
			passive:   "elemental_damage1906", // Elemental Damage
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: alternatePassiveAddition(t, "karui_small_strength"),
						StatRolls:                data.StatRolls{4},
					},
				},
			},
		},
		{
			jewel:     data.LethalPride,
			conqueror: data.Kaom,
			passive:   "eagle_eye199", // Eagle Eye
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: alternatePassiveAddition(t, "karui_notable_add_totem_placement_speed"),
						StatRolls:                data.StatRolls{12},
					},
				},
			},
		},
		{
			jewel:     data.LethalPride,
			conqueror: data.Rakiata,
			passive:   "supreme_ego_keystone2696", // Supreme Ego
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "karui_keystone_2"),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.LethalPride,
			conqueror: data.Kiloava,
			passive:   "supreme_ego_keystone2696", // Supreme Ego
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "karui_keystone_3"),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.LethalPride,
			conqueror: data.Akoya,
			passive:   "supreme_ego_keystone2696", // Supreme Ego
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "karui_keystone_3_v2"),
				StatRolls:             data.StatRolls{1},
			},
		},
	}

	for _, test := range tests {
		t.Run(string(test.conqueror), func(t *testing.T) {
			t.Run(test.passive, func(t *testing.T) {
				testza.AssertEqual(t, test.result, calculator.Calculate(passiveIndex(t, test.passive), seed, test.jewel, test.conqueror))
			})
		})
	}
}

func TestBrutalRestraint(t *testing.T) {
	const seed = 2000

	tests := []struct {
		jewel     data.JewelType
		conqueror data.Conqueror
		passive   string
		result    data.AlternatePassiveSkillInformation
	}{
		{
			jewel:     data.BrutalRestraint,
			conqueror: data.Deshret,
			passive:   "supreme_ego_keystone2696", // Supreme Ego
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "maraketh_keystone_1"),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.BrutalRestraint,
			conqueror: data.Deshret,
			passive:   "maximum_power_charges742", // Instability
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: alternatePassiveAddition(t, "maraketh_notable_add_flask_charges"),
						StatRolls:                data.StatRolls{10},
					},
				},
			},
		},
		{
			jewel:     data.BrutalRestraint,
			conqueror: data.Deshret,
			passive:   "intelligence879", // Intelligence
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: alternatePassiveAddition(t, "maraketh_small_dex"),
						StatRolls:                data.StatRolls{4},
					},
				},
			},
		},
		{
			jewel:     data.BrutalRestraint,
			conqueror: data.Deshret,
			passive:   "elemental_damage1906", // Elemental Damage
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: alternatePassiveAddition(t, "maraketh_small_dex"),
						StatRolls:                data.StatRolls{4},
					},
				},
			},
		},
		{
			jewel:     data.BrutalRestraint,
			conqueror: data.Deshret,
			passive:   "eagle_eye199", // Eagle Eye
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: alternatePassiveAddition(t, "maraketh_notable_add_projectile_damage"),
						StatRolls:                data.StatRolls{20},
					},
				},
			},
		},
		{
			jewel:     data.BrutalRestraint,
			conqueror: data.Balbala,
			passive:   "supreme_ego_keystone2696", // Supreme Ego
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "maraketh_keystone_1_v2"),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.BrutalRestraint,
			conqueror: data.Asenath,
			passive:   "supreme_ego_keystone2696", // Supreme Ego
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "maraketh_keystone_2"),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.BrutalRestraint,
			conqueror: data.Nasima,
			passive:   "supreme_ego_keystone2696", // Supreme Ego
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "maraketh_keystone_3"),
				StatRolls:             data.StatRolls{1},
			},
		},
	}

	for _, test := range tests {
		t.Run(string(test.conqueror), func(t *testing.T) {
			t.Run(test.passive, func(t *testing.T) {
				testza.AssertEqual(t, test.result, calculator.Calculate(passiveIndex(t, test.passive), seed, test.jewel, test.conqueror))
			})
		})
	}
}

func TestMilitantFaith(t *testing.T) {
	const seed = 2000

	tests := []struct {
		jewel     data.JewelType
		conqueror data.Conqueror
		passive   string
		result    data.AlternatePassiveSkillInformation
	}{
		{
			jewel:     data.MilitantFaith,
			conqueror: data.Venarius,
			passive:   "supreme_ego_keystone2696", // Supreme Ego
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "templar_keystone_1"),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.MilitantFaith,
			conqueror: data.Venarius,
			passive:   "maximum_power_charges742", // Instability
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: alternatePassiveAddition(t, "templar_notable_devotion"),
						StatRolls:                data.StatRolls{5},
					},
				},
			},
		},
		{
			jewel:     data.MilitantFaith,
			conqueror: data.Venarius,
			passive:   "intelligence879", // Intelligence
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: alternatePassiveAddition(t, "templar_small_devotion"),
						StatRolls:                data.StatRolls{5},
					},
				},
			},
		},
		{
			jewel:     data.MilitantFaith,
			conqueror: data.Venarius,
			passive:   "elemental_damage1906", // Elemental Damage
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: alternatePassiveAddition(t, "templar_small_devotion"),
						StatRolls:                data.StatRolls{5},
					},
				},
			},
		},
		{
			jewel:     data.MilitantFaith,
			conqueror: data.Venarius,
			passive:   "eagle_eye199", // Eagle Eye
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: alternatePassiveAddition(t, "templar_notable_devotion"),
						StatRolls:                data.StatRolls{5},
					},
				},
			},
		},
		{
			jewel:     data.MilitantFaith,
			conqueror: data.Maxarius,
			passive:   "supreme_ego_keystone2696", // Supreme Ego
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "templar_keystone_1_v2"),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.MilitantFaith,
			conqueror: data.Dominus,
			passive:   "supreme_ego_keystone2696", // Supreme Ego
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "templar_keystone_2"),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.MilitantFaith,
			conqueror: data.Avarius,
			passive:   "supreme_ego_keystone2696", // Supreme Ego
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "templar_keystone_3"),
				StatRolls:             data.StatRolls{1},
			},
		},
	}

	for _, test := range tests {
		t.Run(string(test.conqueror), func(t *testing.T) {
			t.Run(test.passive, func(t *testing.T) {
				testza.AssertEqual(t, test.result, calculator.Calculate(passiveIndex(t, test.passive), seed, test.jewel, test.conqueror))
			})
		})
	}
}

func TestElegantHubris(t *testing.T) {
	const seed = 2000

	tests := []struct {
		jewel     data.JewelType
		conqueror data.Conqueror
		passive   string
		result    data.AlternatePassiveSkillInformation
	}{
		{
			jewel:     data.ElegantHubris,
			conqueror: data.Cadiro,
			passive:   "supreme_ego_keystone2696", // Supreme Ego
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "eternal_keystone_1"),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.ElegantHubris,
			conqueror: data.Cadiro,
			passive:   "maximum_power_charges742", // Instability
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "eternal_notable_shock_1"),
				StatRolls:             data.StatRolls{30},
			},
		},
		{
			jewel:     data.ElegantHubris,
			conqueror: data.Cadiro,
			passive:   "intelligence879", // Intelligence
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "eternal_small_blank"),
				StatRolls:             data.StatRolls{},
			},
		},
		{
			jewel:     data.ElegantHubris,
			conqueror: data.Cadiro,
			passive:   "elemental_damage1906", // Elemental Damage
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "eternal_small_blank"),
				StatRolls:             data.StatRolls{},
			},
		},
		{
			jewel:     data.ElegantHubris,
			conqueror: data.Cadiro,
			passive:   "eagle_eye199", // Eagle Eye
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "eternal_notable_physical_damage_1"),
				StatRolls:             data.StatRolls{80},
			},
		},
		{
			jewel:     data.ElegantHubris,
			conqueror: data.Victario,
			passive:   "supreme_ego_keystone2696", // Supreme Ego
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "eternal_keystone_2"),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.ElegantHubris,
			conqueror: data.Chitus,
			passive:   "supreme_ego_keystone2696", // Supreme Ego
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "eternal_keystone_3"),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.ElegantHubris,
			conqueror: data.Caspiro,
			passive:   "supreme_ego_keystone2696", // Supreme Ego
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "eternal_keystone_3_v2"),
				StatRolls:             data.StatRolls{1},
			},
		},
	}

	for _, test := range tests {
		t.Run(string(test.conqueror), func(t *testing.T) {
			t.Run(test.passive, func(t *testing.T) {
				testza.AssertEqual(t, test.result, calculator.Calculate(passiveIndex(t, test.passive), seed, test.jewel, test.conqueror))
			})
		})
	}
}

func TestHeroicTragedy(t *testing.T) {
	const seed = 2000

	tests := []struct {
		jewel     data.JewelType
		conqueror data.Conqueror
		passive   string
		result    data.AlternatePassiveSkillInformation
	}{
		{
			jewel:     data.HeroicTragedy,
			conqueror: data.Vorana,
			passive:   "supreme_ego_keystone2696", // Supreme Ego
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "kalguur_keystone_1"),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.HeroicTragedy,
			conqueror: data.Vorana,
			passive:   "maximum_power_charges742", // Instability
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "kalguur_notable_19"),
				StatRolls:             data.StatRolls{12, 8},
			},
		},
		{
			jewel:     data.HeroicTragedy,
			conqueror: data.Vorana,
			passive:   "intelligence879", // Intelligence
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: alternatePassiveAddition(t, "kalguuran_small_ward"),
						StatRolls:                data.StatRolls{2},
					},
				},
			},
		},
		{
			jewel:     data.HeroicTragedy,
			conqueror: data.Vorana,
			passive:   "elemental_damage1906", // Elemental Damage
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: alternatePassiveAddition(t, "kalguuran_small_ward"),
						StatRolls:                data.StatRolls{2},
					},
				},
			},
		},
		{
			jewel:     data.HeroicTragedy,
			conqueror: data.Vorana,
			passive:   "eagle_eye199", // Eagle Eye
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "kalguur_notable_22"),
				StatRolls:             data.StatRolls{20, 20},
			},
		},
		{
			jewel:     data.HeroicTragedy,
			conqueror: data.Uhtred,
			passive:   "supreme_ego_keystone2696", // Supreme Ego
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "kalguur_keystone_2"),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.HeroicTragedy,
			conqueror: data.Medved,
			passive:   "supreme_ego_keystone2696", // Supreme Ego
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: alternatePassiveSkill(t, "kalguur_keystone_3"),
				StatRolls:             data.StatRolls{1},
			},
		},
	}

	for _, test := range tests {
		t.Run(string(test.conqueror), func(t *testing.T) {
			t.Run(test.passive, func(t *testing.T) {
				testza.AssertEqual(t, test.result, calculator.Calculate(passiveIndex(t, test.passive), seed, test.jewel, test.conqueror))
			})
		})
	}
}

func BenchmarkAll(b *testing.B) {
	applicable := data.GetApplicablePassives()

	b.ReportAllocs()
	b.ResetTimer()

	for range b.N {
		for jewelType := range data.TimelessJewelConquerors {
			var firstConqueror data.Conqueror
			for conqueror := range data.TimelessJewelConquerors[jewelType] {
				firstConqueror = conqueror
				break
			}

			seedMin := data.TimelessJewelSeedRanges[jewelType].Min
			seedMax := data.TimelessJewelSeedRanges[jewelType].Max

			if data.TimelessJewelSeedRanges[jewelType].Special {
				seedMin /= 20
				seedMax /= 20
			}

			for seed := seedMin; seed <= seedMax; seed++ {
				realSeed := seed
				if data.TimelessJewelSeedRanges[jewelType].Special {
					realSeed *= 20
				}

				for _, skill := range applicable {
					calculator.Calculate(skill.Index, realSeed, jewelType, firstConqueror)
				}
			}
		}
	}
}
