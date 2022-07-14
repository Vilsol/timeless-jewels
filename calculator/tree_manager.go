package calculator

import (
	"github.com/Vilsol/timeless-jewels/data"
	"github.com/Vilsol/timeless-jewels/random"
)

type AlternateTreeManager struct {
	PassiveSkill  *data.PassiveSkill
	TimelessJewel data.TimelessJewel
}

func (a *AlternateTreeManager) IsPassiveSkillReplaced(rng *random.NumberGenerator) bool {
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

		rng.Reset(a.PassiveSkill, a.TimelessJewel)
		return rng.Generate(0, 100) < a.TimelessJewel.AlternateTreeVersion.NotableReplacementSpawnWeight
	}

	if len(a.PassiveSkill.StatIndices) == 1 && data.IsSmallAttribute(a.PassiveSkill.StatIndices[0]) {
		return a.TimelessJewel.AlternateTreeVersion.AreSmallAttributePassiveSkillsReplaced
	}

	return a.TimelessJewel.AlternateTreeVersion.AreSmallNormalPassiveSkillsReplaced
}

func (a *AlternateTreeManager) AugmentPassiveSkill(rng *random.NumberGenerator) []data.AlternatePassiveAdditionInformation {
	rng.Reset(a.PassiveSkill, a.TimelessJewel)

	if data.GetPassiveSkillType(a.PassiveSkill) == data.Notable {
		rng.Generate(0, 100)
	}

	return a.RollAdditions(a.TimelessJewel.AlternateTreeVersion.MinimumAdditions, a.TimelessJewel.AlternateTreeVersion.MaximumAdditions, rng)
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

func (a *AlternateTreeManager) ReplacePassiveSkill(rng *random.NumberGenerator) data.AlternatePassiveSkillInformation {
	if a.PassiveSkill.IsKeystone {
		alternatePassiveSkillKeyStone := data.GetAlternatePassiveSkillKeyStone(a.TimelessJewel)
		return data.AlternatePassiveSkillInformation{
			AlternatePassiveSkill: alternatePassiveSkillKeyStone,
			StatRolls: map[uint32]uint32{
				0: alternatePassiveSkillKeyStone.Stat1Min,
			},
		}
	}

	applicableAlternatePassiveSkills := data.GetApplicableAlternatePassiveSkills(a.PassiveSkill, a.TimelessJewel)

	var rolledAlternatePassiveSkill *data.AlternatePassiveSkill
	rng.Reset(a.PassiveSkill, a.TimelessJewel)

	if data.GetPassiveSkillType(a.PassiveSkill) == data.Notable {
		rng.Generate(0, 100)
	}

	currentSpawnWeight := uint32(0)
	for _, skill := range applicableAlternatePassiveSkills {
		currentSpawnWeight += skill.SpawnWeight
		if rng.GenerateSingle(currentSpawnWeight) < skill.SpawnWeight {
			rolledAlternatePassiveSkill = skill
		}
	}

	elements := min(uint32(len(rolledAlternatePassiveSkill.StatsKeys)), 4)
	alternatePassiveSkillStatRolls := make(map[uint32]uint32, elements)
	for i := uint32(0); i < elements; i++ {
		alternatePassiveSkillStatRolls[i] = rolledAlternatePassiveSkill.GetStatMinMax(true, i)

		if rolledAlternatePassiveSkill.GetStatMinMax(false, i) > rolledAlternatePassiveSkill.GetStatMinMax(true, i) {
			alternatePassiveSkillStatRolls[i] = rng.Generate(rolledAlternatePassiveSkill.GetStatMinMax(true, i), rolledAlternatePassiveSkill.GetStatMinMax(false, i))
		}
	}

	if (rolledAlternatePassiveSkill.RandomMin == 0) && (rolledAlternatePassiveSkill.RandomMax == 0) {
		return data.AlternatePassiveSkillInformation{
			AlternatePassiveSkill: rolledAlternatePassiveSkill,
			StatRolls:             alternatePassiveSkillStatRolls,
		}
	}

	minAdditions := a.TimelessJewel.AlternateTreeVersion.MinimumAdditions + rolledAlternatePassiveSkill.RandomMin
	maxAdditions := a.TimelessJewel.AlternateTreeVersion.MaximumAdditions + rolledAlternatePassiveSkill.RandomMax

	return data.AlternatePassiveSkillInformation{
		AlternatePassiveSkill:                rolledAlternatePassiveSkill,
		StatRolls:                            alternatePassiveSkillStatRolls,
		AlternatePassiveAdditionInformations: a.RollAdditions(minAdditions, maxAdditions, rng),
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

		elements := min(uint32(len(rolledAlternatePassiveAddition.StatsKeys)), 2)
		alternatePassiveAdditionStatRolls := make(map[uint32]uint32, elements)
		for j := uint32(0); j < elements; j++ {
			alternatePassiveAdditionStatRolls[j] = rolledAlternatePassiveAddition.GetStatMinMax(true, j)

			if rolledAlternatePassiveAddition.GetStatMinMax(false, j) > rolledAlternatePassiveAddition.GetStatMinMax(true, j) {
				alternatePassiveAdditionStatRolls[j] = rng.Generate(rolledAlternatePassiveAddition.GetStatMinMax(true, j), rolledAlternatePassiveAddition.GetStatMinMax(false, j))
			}
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
