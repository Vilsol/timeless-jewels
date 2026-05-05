package main

import (
	"strconv"
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
		passive   uint32
		result    data.AlternatePassiveSkillInformation
	}{
		{
			jewel:     data.GloriousVanity,
			conqueror: data.Xibaqua,
			passive:   2286, // Supreme Ego (supreme_ego_keystone2696)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(0),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.GloriousVanity,
			conqueror: data.Xibaqua,
			passive:   411, // Instability (maximum_power_charges742)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(67),
				StatRolls:             data.StatRolls{8, 22},
			},
		},
		{
			jewel:     data.GloriousVanity,
			conqueror: data.Xibaqua,
			passive:   519, // Intelligence (intelligence879)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(38),
				StatRolls:             data.StatRolls{12},
			},
		},
		{
			jewel:     data.GloriousVanity,
			conqueror: data.Xibaqua,
			passive:   1190, // Elemental Damage (elemental_damage1906)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(21),
				StatRolls:             data.StatRolls{3},
			},
		},
		{
			jewel:     data.GloriousVanity,
			conqueror: data.Xibaqua,
			passive:   88, // Eagle Eye (eagle_eye199)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(77),
				StatRolls:             data.StatRolls{},
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(14),
						StatRolls:                data.StatRolls{6},
					},
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(23),
						StatRolls:                data.StatRolls{5},
					},
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(36),
						StatRolls:                data.StatRolls{11},
					},
				},
			},
		},
		{
			jewel:     data.GloriousVanity,
			conqueror: data.Zerphi,
			passive:   2286, // Supreme Ego (supreme_ego_keystone2696)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(1),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.GloriousVanity,
			conqueror: data.Ahuana,
			passive:   2286, // Supreme Ego (supreme_ego_keystone2696)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(2),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.GloriousVanity,
			conqueror: data.Doryani,
			passive:   2286, // Supreme Ego (supreme_ego_keystone2696)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(3),
				StatRolls:             data.StatRolls{1},
			},
		},
	}

	for _, test := range tests {
		t.Run(string(test.conqueror), func(t *testing.T) {
			t.Run(strconv.Itoa(int(test.passive)), func(t *testing.T) {
				testza.AssertEqual(t, test.result, calculator.Calculate(test.passive, seed, test.jewel, test.conqueror))
			})
		})
	}
}

func TestLethalPride(t *testing.T) {
	const seed = 12000

	tests := []struct {
		jewel     data.JewelType
		conqueror data.Conqueror
		passive   uint32
		result    data.AlternatePassiveSkillInformation
	}{
		{
			jewel:     data.LethalPride,
			conqueror: data.Kaom,
			passive:   2286, // Supreme Ego (supreme_ego_keystone2696)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(78),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.LethalPride,
			conqueror: data.Kaom,
			passive:   411, // Instability (maximum_power_charges742)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(42),
						StatRolls:                data.StatRolls{20},
					},
				},
			},
		},
		{
			jewel:     data.LethalPride,
			conqueror: data.Kaom,
			passive:   519, // Intelligence (intelligence879)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(39),
						StatRolls:                data.StatRolls{4},
					},
				},
			},
		},
		{
			jewel:     data.LethalPride,
			conqueror: data.Kaom,
			passive:   1190, // Elemental Damage (elemental_damage1906)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(39),
						StatRolls:                data.StatRolls{4},
					},
				},
			},
		},
		{
			jewel:     data.LethalPride,
			conqueror: data.Kaom,
			passive:   88, // Eagle Eye (eagle_eye199)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(57),
						StatRolls:                data.StatRolls{12},
					},
				},
			},
		},
		{
			jewel:     data.LethalPride,
			conqueror: data.Rakiata,
			passive:   2286, // Supreme Ego (supreme_ego_keystone2696)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(79),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.LethalPride,
			conqueror: data.Kiloava,
			passive:   2286, // Supreme Ego (supreme_ego_keystone2696)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(80),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.LethalPride,
			conqueror: data.Akoya,
			passive:   2286, // Supreme Ego (supreme_ego_keystone2696)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(81),
				StatRolls:             data.StatRolls{1},
			},
		},
	}

	for _, test := range tests {
		t.Run(string(test.conqueror), func(t *testing.T) {
			t.Run(strconv.Itoa(int(test.passive)), func(t *testing.T) {
				testza.AssertEqual(t, test.result, calculator.Calculate(test.passive, seed, test.jewel, test.conqueror))
			})
		})
	}
}

func TestBrutalRestraint(t *testing.T) {
	const seed = 2000

	tests := []struct {
		jewel     data.JewelType
		conqueror data.Conqueror
		passive   uint32
		result    data.AlternatePassiveSkillInformation
	}{
		{
			jewel:     data.BrutalRestraint,
			conqueror: data.Deshret,
			passive:   2286, // Supreme Ego (supreme_ego_keystone2696)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(82),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.BrutalRestraint,
			conqueror: data.Deshret,
			passive:   411, // Instability (maximum_power_charges742)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(70),
						StatRolls:                data.StatRolls{10},
					},
				},
			},
		},
		{
			jewel:     data.BrutalRestraint,
			conqueror: data.Deshret,
			passive:   519, // Intelligence (intelligence879)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(66),
						StatRolls:                data.StatRolls{4},
					},
				},
			},
		},
		{
			jewel:     data.BrutalRestraint,
			conqueror: data.Deshret,
			passive:   1190, // Elemental Damage (elemental_damage1906)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(66),
						StatRolls:                data.StatRolls{4},
					},
				},
			},
		},
		{
			jewel:     data.BrutalRestraint,
			conqueror: data.Deshret,
			passive:   88, // Eagle Eye (eagle_eye199)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(76),
						StatRolls:                data.StatRolls{20},
					},
				},
			},
		},
		{
			jewel:     data.BrutalRestraint,
			conqueror: data.Balbala,
			passive:   2286, // Supreme Ego (supreme_ego_keystone2696)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(83),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.BrutalRestraint,
			conqueror: data.Asenath,
			passive:   2286, // Supreme Ego (supreme_ego_keystone2696)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(84),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.BrutalRestraint,
			conqueror: data.Nasima,
			passive:   2286, // Supreme Ego (supreme_ego_keystone2696)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(85),
				StatRolls:             data.StatRolls{1},
			},
		},
	}

	for _, test := range tests {
		t.Run(string(test.conqueror), func(t *testing.T) {
			t.Run(strconv.Itoa(int(test.passive)), func(t *testing.T) {
				testza.AssertEqual(t, test.result, calculator.Calculate(test.passive, seed, test.jewel, test.conqueror))
			})
		})
	}
}

func TestMilitantFaith(t *testing.T) {
	const seed = 2000

	tests := []struct {
		jewel     data.JewelType
		conqueror data.Conqueror
		passive   uint32
		result    data.AlternatePassiveSkillInformation
	}{
		{
			jewel:     data.MilitantFaith,
			conqueror: data.Venarius,
			passive:   2286, // Supreme Ego (supreme_ego_keystone2696)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(86),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.MilitantFaith,
			conqueror: data.Venarius,
			passive:   411, // Instability (maximum_power_charges742)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(93),
						StatRolls:                data.StatRolls{5},
					},
				},
			},
		},
		{
			jewel:     data.MilitantFaith,
			conqueror: data.Venarius,
			passive:   519, // Intelligence (intelligence879)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(92),
						StatRolls:                data.StatRolls{5},
					},
				},
			},
		},
		{
			jewel:     data.MilitantFaith,
			conqueror: data.Venarius,
			passive:   1190, // Elemental Damage (elemental_damage1906)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(92),
						StatRolls:                data.StatRolls{5},
					},
				},
			},
		},
		{
			jewel:     data.MilitantFaith,
			conqueror: data.Venarius,
			passive:   88, // Eagle Eye (eagle_eye199)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(93),
						StatRolls:                data.StatRolls{5},
					},
				},
			},
		},
		{
			jewel:     data.MilitantFaith,
			conqueror: data.Maxarius,
			passive:   2286, // Supreme Ego (supreme_ego_keystone2696)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(87),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.MilitantFaith,
			conqueror: data.Dominus,
			passive:   2286, // Supreme Ego (supreme_ego_keystone2696)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(88),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.MilitantFaith,
			conqueror: data.Avarius,
			passive:   2286, // Supreme Ego (supreme_ego_keystone2696)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(89),
				StatRolls:             data.StatRolls{1},
			},
		},
	}

	for _, test := range tests {
		t.Run(string(test.conqueror), func(t *testing.T) {
			t.Run(strconv.Itoa(int(test.passive)), func(t *testing.T) {
				testza.AssertEqual(t, test.result, calculator.Calculate(test.passive, seed, test.jewel, test.conqueror))
			})
		})
	}
}

func TestElegantHubris(t *testing.T) {
	const seed = 2000

	tests := []struct {
		jewel     data.JewelType
		conqueror data.Conqueror
		passive   uint32
		result    data.AlternatePassiveSkillInformation
	}{
		{
			jewel:     data.ElegantHubris,
			conqueror: data.Cadiro,
			passive:   2286, // Supreme Ego (supreme_ego_keystone2696)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(105),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.ElegantHubris,
			conqueror: data.Cadiro,
			passive:   411, // Instability (maximum_power_charges742)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(123),
				StatRolls:             data.StatRolls{30},
			},
		},
		{
			jewel:     data.ElegantHubris,
			conqueror: data.Cadiro,
			passive:   519, // Intelligence (intelligence879)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(109),
				StatRolls:             data.StatRolls{},
			},
		},
		{
			jewel:     data.ElegantHubris,
			conqueror: data.Cadiro,
			passive:   1190, // Elemental Damage (elemental_damage1906)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(109),
				StatRolls:             data.StatRolls{},
			},
		},
		{
			jewel:     data.ElegantHubris,
			conqueror: data.Cadiro,
			passive:   88, // Eagle Eye (eagle_eye199)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(137),
				StatRolls:             data.StatRolls{80},
			},
		},
		{
			jewel:     data.ElegantHubris,
			conqueror: data.Victario,
			passive:   2286, // Supreme Ego (supreme_ego_keystone2696)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(106),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.ElegantHubris,
			conqueror: data.Chitus,
			passive:   2286, // Supreme Ego (supreme_ego_keystone2696)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(107),
				StatRolls:             data.StatRolls{1},
			},
		},
		{
			jewel:     data.ElegantHubris,
			conqueror: data.Caspiro,
			passive:   2286, // Supreme Ego (supreme_ego_keystone2696)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(108),
				StatRolls:             data.StatRolls{1},
			},
		},
	}

	for _, test := range tests {
		t.Run(string(test.conqueror), func(t *testing.T) {
			t.Run(strconv.Itoa(int(test.passive)), func(t *testing.T) {
				testza.AssertEqual(t, test.result, calculator.Calculate(test.passive, seed, test.jewel, test.conqueror))
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
