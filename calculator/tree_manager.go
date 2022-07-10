package calculator

import (
	"github.com/Vilsol/timeless-jewels/data"
	"github.com/Vilsol/timeless-jewels/random"
)

type AlternateTreeManager struct {
	PassiveSkill  *data.PassiveSkill
	TimelessJewel *data.TimelessJewel
}

func (a *AlternateTreeManager) IsPassiveSkillReplaced() bool {
	if a.PassiveSkill.IsKeystone {
		return true
	}

	if a.PassiveSkill.IsNotable {
		if a.TimelessJewel.AlternateTreeVersion.NotableReplacementSpawnWeight >= 100 {
			return true
		}

		if a.TimelessJewel.AlternateTreeVersion.NotableReplacementSpawnWeight == 0 {
			return false
		}

		rng := random.NewRNG(a.PassiveSkill, a.TimelessJewel)
		return rng.Generate(0, 100) < a.TimelessJewel.AlternateTreeVersion.NotableReplacementSpawnWeight
	}

	if len(a.PassiveSkill.StatIndices) == 1 && data.IsSmallAttribute(a.PassiveSkill.StatIndices[0]) {
		return a.TimelessJewel.AlternateTreeVersion.AreSmallAttributePassiveSkillsReplaced
	}

	return a.TimelessJewel.AlternateTreeVersion.AreSmallNormalPassiveSkillsReplaced
}

func (a *AlternateTreeManager) AugmentPassiveSkill() []data.AlternatePassiveAdditionInformation {
	rng := random.NewRNG(a.PassiveSkill, a.TimelessJewel)
	passiveSkillType := data.GetPassiveSkillType(a.PassiveSkill)

	if passiveSkillType == data.Notable {
		rng.Generate(0, 100)
	}

	minAdditions := a.TimelessJewel.AlternateTreeVersion.MinimumAdditions
	maxAdditions := a.TimelessJewel.AlternateTreeVersion.MaximumAdditions

	alternatePassiveAdditionInformations := a.RollAdditions(minAdditions, maxAdditions, rng)

	return alternatePassiveAdditionInformations
}

func (a *AlternateTreeManager) RollAlternatePassiveAddition(rng *random.NumberGenerator) *data.AlternatePassiveAddition {
	applicableAlternatePassiveAdditions := data.GetApplicableAlternatePassiveAdditions(a.PassiveSkill, a.TimelessJewel)

	totalSpawnWeight := uint32(0)
	for _, addition := range applicableAlternatePassiveAdditions {
		totalSpawnWeight += addition.SpawnWeight
	}

	additionRoll := rng.GenerateSingle(totalSpawnWeight)

	for _, addition := range applicableAlternatePassiveAdditions {
		if addition.SpawnWeight > additionRoll {
			return addition
		}

		additionRoll -= addition.SpawnWeight
	}

	return nil
}

func (a *AlternateTreeManager) ReplacePassiveSkill() data.AlternatePassiveSkillInformation {
	if a.PassiveSkill.IsKeystone {
		alternatePassiveSkillKeyStone := data.GetAlternatePassiveSkillKeyStone(a.TimelessJewel)

		alternatePassiveSkillKeyStoneStatRolls := map[uint32]uint32{
			0: alternatePassiveSkillKeyStone.Stat1Min,
		}

		return data.AlternatePassiveSkillInformation{
			AlternatePassiveSkill: alternatePassiveSkillKeyStone,
			StatRolls:             alternatePassiveSkillKeyStoneStatRolls,
		}
	}

	applicableAlternatePassiveSkills := data.GetApplicableAlternatePassiveSkills(a.PassiveSkill, a.TimelessJewel)

	var rolledAlternatePassiveSkill *data.AlternatePassiveSkill
	rng := random.NewRNG(a.PassiveSkill, a.TimelessJewel)

	currentSpawnWeight := uint32(0)

	if data.GetPassiveSkillType(a.PassiveSkill) == data.Notable {
		rng.Generate(0, 100)
	}

	for _, skill := range applicableAlternatePassiveSkills {
		currentSpawnWeight += skill.SpawnWeight

		roll := rng.GenerateSingle(currentSpawnWeight)

		if roll < skill.SpawnWeight {
			rolledAlternatePassiveSkill = skill
		}
	}

	alternatePassiveSkillStatRollRanges := map[uint32]data.Range{
		0: {
			Min: rolledAlternatePassiveSkill.Stat1Min,
			Max: rolledAlternatePassiveSkill.Stat1Max,
		},
		1: {
			Min: rolledAlternatePassiveSkill.Stat2Min,
			Max: rolledAlternatePassiveSkill.Stat2Max,
		},
		2: {
			Min: rolledAlternatePassiveSkill.Stat3Min,
			Max: rolledAlternatePassiveSkill.Stat3Max,
		},
		3: {
			Min: rolledAlternatePassiveSkill.Stat4Min,
			Max: rolledAlternatePassiveSkill.Stat4Max,
		},
	}

	alternatePassiveSkillStatRolls := make(map[uint32]uint32)
	for i := uint32(0); i < min(uint32(len(rolledAlternatePassiveSkill.StatsKeys)), 4); i++ {
		alternatePassiveSkillStatRoll := alternatePassiveSkillStatRollRanges[i].Min

		if alternatePassiveSkillStatRollRanges[i].Max > alternatePassiveSkillStatRollRanges[i].Min {
			alternatePassiveSkillStatRoll = rng.Generate(alternatePassiveSkillStatRollRanges[i].Min, alternatePassiveSkillStatRollRanges[i].Max)
		}

		alternatePassiveSkillStatRolls[i] = alternatePassiveSkillStatRoll
	}

	if (rolledAlternatePassiveSkill.RandomMin == 0) && (rolledAlternatePassiveSkill.RandomMax == 0) {
		return data.AlternatePassiveSkillInformation{
			AlternatePassiveSkill: rolledAlternatePassiveSkill,
			StatRolls:             alternatePassiveSkillStatRolls,
		}
	}

	minAdditions := a.TimelessJewel.AlternateTreeVersion.MinimumAdditions + rolledAlternatePassiveSkill.RandomMin
	maxAdditions := a.TimelessJewel.AlternateTreeVersion.MaximumAdditions + rolledAlternatePassiveSkill.RandomMax

	alternatePassiveAdditionInformations := a.RollAdditions(minAdditions, maxAdditions, rng)

	return data.AlternatePassiveSkillInformation{
		AlternatePassiveSkill:                rolledAlternatePassiveSkill,
		StatRolls:                            alternatePassiveSkillStatRolls,
		AlternatePassiveAdditionInformations: alternatePassiveAdditionInformations,
	}
}

func (a *AlternateTreeManager) RollAdditions(minimumAdditions uint32, maximumAdditions uint32, rng *random.NumberGenerator) []data.AlternatePassiveAdditionInformation {
	additionCountRoll := minimumAdditions

	if maximumAdditions > minimumAdditions {
		additionCountRoll = rng.Generate(minimumAdditions, maximumAdditions)
	}

	alternatePassiveAdditionInformations := make([]data.AlternatePassiveAdditionInformation, 0)

	for i := uint32(0); i < additionCountRoll; i++ {
		var rolledAlternatePassiveAddition *data.AlternatePassiveAddition

		for rolledAlternatePassiveAddition == nil {
			rolledAlternatePassiveAddition = a.RollAlternatePassiveAddition(rng)
		}

		alternatePassiveAdditionStatRollRanges := map[uint32]data.Range{
			0: {
				Min: rolledAlternatePassiveAddition.Stat1Min,
				Max: rolledAlternatePassiveAddition.Stat1Max,
			},
			1: {
				Min: rolledAlternatePassiveAddition.Stat2Min,
				Max: rolledAlternatePassiveAddition.Stat2Max,
			},
		}

		alternatePassiveAdditionStatRolls := make(map[uint32]uint32)

		for j := uint32(0); j < min(uint32(len(rolledAlternatePassiveAddition.StatsKeys)), 2); j++ {
			alternatePassiveAdditionStatRoll := alternatePassiveAdditionStatRollRanges[j].Min

			if alternatePassiveAdditionStatRollRanges[j].Max > alternatePassiveAdditionStatRollRanges[j].Min {
				alternatePassiveAdditionStatRoll = rng.Generate(alternatePassiveAdditionStatRollRanges[j].Min, alternatePassiveAdditionStatRollRanges[j].Max)
			}

			alternatePassiveAdditionStatRolls[j] = alternatePassiveAdditionStatRoll
		}

		alternatePassiveAdditionInformations = append(alternatePassiveAdditionInformations, data.AlternatePassiveAdditionInformation{
			AlternatePassiveAddition: rolledAlternatePassiveAddition,
			StatRolls:                alternatePassiveAdditionStatRolls,
		})
	}

	return alternatePassiveAdditionInformations
}

func min(a uint32, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
}
