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
			passive:   2286, // Doomsday (hex_zone_keystone2800_)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(19),
				StatRolls:             map[uint32]uint32{0: 3},
			},
		},
		{
			jewel:     data.GloriousVanity,
			conqueror: data.Xibaqua,
			passive:   411, // Instability (maximum_power_charges742)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(67),
				StatRolls:             map[uint32]uint32{0: 5, 1: 22},
			},
		},
		{
			jewel:     data.GloriousVanity,
			conqueror: data.Xibaqua,
			passive:   519, // Intelligence (intelligence879)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(38),
				StatRolls:             map[uint32]uint32{0: 12},
			},
		},
		{
			jewel:     data.GloriousVanity,
			conqueror: data.Xibaqua,
			passive:   1190, // Attack Damage and Attack Speed with Shield (damage_while_using_shield1913)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(26),
				StatRolls:             map[uint32]uint32{0: 4},
			},
		},
		{
			jewel:     data.GloriousVanity,
			conqueror: data.Xibaqua,
			passive:   88, // Eagle Eye (eagle_eye199)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(77),
				StatRolls:             map[uint32]uint32{},
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(14),
						StatRolls:                map[uint32]uint32{0: 6},
					},
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(23),
						StatRolls:                map[uint32]uint32{0: 5},
					},
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(36),
						StatRolls:                map[uint32]uint32{0: 11},
					},
				},
			},
		},
		{
			jewel:     data.GloriousVanity,
			conqueror: data.Zerphi,
			passive:   2286, // Doomsday (hex_zone_keystone2800_)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(19),
				StatRolls:             map[uint32]uint32{0: 3},
			},
		},
		{
			jewel:     data.GloriousVanity,
			conqueror: data.Ahuana,
			passive:   2286, // Doomsday (hex_zone_keystone2800_)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(19),
				StatRolls:             map[uint32]uint32{0: 3},
			},
		},
		{
			jewel:     data.GloriousVanity,
			conqueror: data.Doryani,
			passive:   2286, // Doomsday (hex_zone_keystone2800_)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(19),
				StatRolls:             map[uint32]uint32{0: 3},
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
			passive:   2286, // Doomsday (hex_zone_keystone2800_)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(39),
						StatRolls:                map[uint32]uint32{0: 4},
					},
				},
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
						StatRolls:                map[uint32]uint32{0: 20},
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
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(38),
						StatRolls:                map[uint32]uint32{0: 2},
					},
				},
			},
		},
		{
			jewel:     data.LethalPride,
			conqueror: data.Kaom,
			passive:   1190, // Attack Damage and Attack Speed with Shield (damage_while_using_shield1913)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(39),
						StatRolls:                map[uint32]uint32{0: 4},
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
						StatRolls:                map[uint32]uint32{0: 12},
					},
				},
			},
		},
		{
			jewel:     data.LethalPride,
			conqueror: data.Rakiata,
			passive:   2286, // Doomsday (hex_zone_keystone2800_)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(39),
						StatRolls:                map[uint32]uint32{0: 4},
					},
				},
			},
		},
		{
			jewel:     data.LethalPride,
			conqueror: data.Kiloava,
			passive:   2286, // Doomsday (hex_zone_keystone2800_)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(39),
						StatRolls:                map[uint32]uint32{0: 4},
					},
				},
			},
		},
		{
			jewel:     data.LethalPride,
			conqueror: data.Akoya,
			passive:   2286, // Doomsday (hex_zone_keystone2800_)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(39),
						StatRolls:                map[uint32]uint32{0: 4},
					},
				},
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
			passive:   2286, // Doomsday (hex_zone_keystone2800_)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(66),
						StatRolls:                map[uint32]uint32{0: 4},
					},
				},
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
						StatRolls:                map[uint32]uint32{0: 10},
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
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(65),
						StatRolls:                map[uint32]uint32{0: 2},
					},
				},
			},
		},
		{
			jewel:     data.BrutalRestraint,
			conqueror: data.Deshret,
			passive:   1190, // Attack Damage and Attack Speed with Shield (damage_while_using_shield1913)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(66),
						StatRolls:                map[uint32]uint32{0: 4},
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
						StatRolls:                map[uint32]uint32{0: 20},
					},
				},
			},
		},
		{
			jewel:     data.BrutalRestraint,
			conqueror: data.Balbala,
			passive:   2286, // Doomsday (hex_zone_keystone2800_)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(66),
						StatRolls:                map[uint32]uint32{0: 4},
					},
				},
			},
		},
		{
			jewel:     data.BrutalRestraint,
			conqueror: data.Asenath,
			passive:   2286, // Doomsday (hex_zone_keystone2800_)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(66),
						StatRolls:                map[uint32]uint32{0: 4},
					},
				},
			},
		},
		{
			jewel:     data.BrutalRestraint,
			conqueror: data.Nasima,
			passive:   2286, // Doomsday (hex_zone_keystone2800_)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(66),
						StatRolls:                map[uint32]uint32{0: 4},
					},
				},
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
			passive:   2286, // Doomsday (hex_zone_keystone2800_)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(92),
						StatRolls:                map[uint32]uint32{0: 5},
					},
				},
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
						StatRolls:                map[uint32]uint32{0: 5},
					},
				},
			},
		},
		{
			jewel:     data.MilitantFaith,
			conqueror: data.Venarius,
			passive:   519, // Intelligence (intelligence879)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(90),
				StatRolls:             map[uint32]uint32{0: 10},
			},
		},
		{
			jewel:     data.MilitantFaith,
			conqueror: data.Venarius,
			passive:   1190, // Attack Damage and Attack Speed with Shield (damage_while_using_shield1913)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(92),
						StatRolls:                map[uint32]uint32{0: 5},
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
						StatRolls:                map[uint32]uint32{0: 5},
					},
				},
			},
		},
		{
			jewel:     data.MilitantFaith,
			conqueror: data.Maxarius,
			passive:   2286, // Doomsday (hex_zone_keystone2800_)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(92),
						StatRolls:                map[uint32]uint32{0: 5},
					},
				},
			},
		},
		{
			jewel:     data.MilitantFaith,
			conqueror: data.Dominus,
			passive:   2286, // Doomsday (hex_zone_keystone2800_)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(92),
						StatRolls:                map[uint32]uint32{0: 5},
					},
				},
			},
		},
		{
			jewel:     data.MilitantFaith,
			conqueror: data.Avarius,
			passive:   2286, // Doomsday (hex_zone_keystone2800_)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveAdditionInformations: []data.AlternatePassiveAdditionInformation{
					{
						AlternatePassiveAddition: data.GetAlternatePassiveAdditionByIndex(92),
						StatRolls:                map[uint32]uint32{0: 5},
					},
				},
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
			passive:   2286, // Doomsday (hex_zone_keystone2800_)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(109),
				StatRolls:             map[uint32]uint32{},
			},
		},
		{
			jewel:     data.ElegantHubris,
			conqueror: data.Cadiro,
			passive:   411, // Instability (maximum_power_charges742)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(123),
				StatRolls:             map[uint32]uint32{0: 30},
			},
		},
		{
			jewel:     data.ElegantHubris,
			conqueror: data.Cadiro,
			passive:   519, // Intelligence (intelligence879)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(109),
				StatRolls:             map[uint32]uint32{},
			},
		},
		{
			jewel:     data.ElegantHubris,
			conqueror: data.Cadiro,
			passive:   1190, // Attack Damage and Attack Speed with Shield (damage_while_using_shield1913)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(109),
				StatRolls:             map[uint32]uint32{},
			},
		},
		{
			jewel:     data.ElegantHubris,
			conqueror: data.Cadiro,
			passive:   88, // Eagle Eye (eagle_eye199)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(137),
				StatRolls:             map[uint32]uint32{0: 80},
			},
		},
		{
			jewel:     data.ElegantHubris,
			conqueror: data.Victario,
			passive:   2286, // Doomsday (hex_zone_keystone2800_)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(109),
				StatRolls:             map[uint32]uint32{},
			},
		},
		{
			jewel:     data.ElegantHubris,
			conqueror: data.Chitus,
			passive:   2286, // Doomsday (hex_zone_keystone2800_)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(109),
				StatRolls:             map[uint32]uint32{},
			},
		},
		{
			jewel:     data.ElegantHubris,
			conqueror: data.Caspiro,
			passive:   2286, // Doomsday (hex_zone_keystone2800_)
			result: data.AlternatePassiveSkillInformation{
				AlternatePassiveSkill: data.GetAlternatePassiveSkillByIndex(109),
				StatRolls:             map[uint32]uint32{},
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

	for i := 0; i < b.N; i++ {
		for jewelType := range data.TimelessJewelConquerors {
			var firstConqueror data.Conqueror
			for conqueror := range data.TimelessJewelConquerors[jewelType] {
				firstConqueror = conqueror
				break
			}

			min := data.TimelessJewelSeedRanges[jewelType].Min
			max := data.TimelessJewelSeedRanges[jewelType].Max

			if data.TimelessJewelSeedRanges[jewelType].Special {
				min /= 20
				max /= 20
			}

			for seed := min; seed <= max; seed++ {
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
