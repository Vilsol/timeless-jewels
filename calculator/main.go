package calculator

import "timeless-jewels/data"

type UpdateFunc func(seed uint32)

func Calculate(passiveID uint32, seed uint32, timelessJewelType data.JewelType, conqueror data.Conqueror) *data.AlternatePassiveSkillInformation {
	passiveSkill := data.GetPassiveSkillByIndex(passiveID)

	if !data.IsPassiveSkillValidForAlteration(passiveSkill) {
		return nil
	}

	alternateTreeVersion := data.GetAlternateTreeVersionIndex(uint32(timelessJewelType))

	timelessJewelConqueror := data.TimelessJewelConquerors[timelessJewelType][conqueror]

	timelessJewel := &data.TimelessJewel{
		Seed:                   seed,
		AlternateTreeVersion:   alternateTreeVersion,
		TimelessJewelConqueror: timelessJewelConqueror,
	}

	alternateTreeManager := AlternateTreeManager{
		PassiveSkill:  passiveSkill,
		TimelessJewel: timelessJewel,
	}

	replace := alternateTreeManager.IsPassiveSkillReplaced()

	if replace {
		information := alternateTreeManager.ReplacePassiveSkill()
		return &information
	}

	return &data.AlternatePassiveSkillInformation{
		AlternatePassiveAdditionInformations: alternateTreeManager.AugmentPassiveSkill(),
	}
}

func ReverseSearch(passiveIDs []uint32, statIDs []uint32, timelessJewelType data.JewelType, conqueror data.Conqueror, updates UpdateFunc) map[uint32]map[uint32]map[uint32]uint32 {
	passiveSkills := make(map[uint32]*data.PassiveSkill)
	for _, id := range passiveIDs {
		skill := data.GetPassiveSkillByIndex(id)
		if data.IsPassiveSkillValidForAlteration(skill) {
			passiveSkills[id] = skill
		}
	}

	alternateTreeVersion := data.GetAlternateTreeVersionIndex(uint32(timelessJewelType))

	timelessJewelConqueror := data.TimelessJewelConquerors[timelessJewelType][conqueror]

	timelessJewel := &data.TimelessJewel{
		AlternateTreeVersion:   alternateTreeVersion,
		TimelessJewelConqueror: timelessJewelConqueror,
	}

	alternateTreeManager := &AlternateTreeManager{
		TimelessJewel: timelessJewel,
	}

	statMap := make(map[uint32]bool)
	for _, id := range statIDs {
		statMap[id] = true
	}

	results := make(map[uint32]map[uint32]map[uint32]uint32)
	for seed := data.TimelessJewelSeedRanges[timelessJewelType].Min; seed <= data.TimelessJewelSeedRanges[timelessJewelType].Max; seed++ {
		timelessJewel.Seed = seed

		if seed%10 == 0 && updates != nil {
			updates(seed)
		}

		for _, skill := range passiveSkills {
			alternateTreeManager.PassiveSkill = skill

			if alternateTreeManager.IsPassiveSkillReplaced() {
				replacement := alternateTreeManager.ReplacePassiveSkill()
				if replacement.AlternatePassiveSkill != nil {
					for i, key := range replacement.AlternatePassiveSkill.StatsKeys {
						if _, ok := statMap[key]; ok {
							if _, ok := results[seed]; !ok {
								results[seed] = make(map[uint32]map[uint32]uint32)
							}

							if _, ok := results[seed][skill.Index]; !ok {
								results[seed][skill.Index] = make(map[uint32]uint32)
							}

							if replacement.StatRolls != nil {
								results[seed][skill.Index][key] = replacement.StatRolls[uint32(i)]
							}
						}
					}
				}
			} else {
				augments := alternateTreeManager.AugmentPassiveSkill()
				for _, augment := range augments {
					if augment.AlternatePassiveAddition != nil {
						for i, key := range augment.AlternatePassiveAddition.StatsKeys {
							if _, ok := statMap[key]; ok {
								if _, ok := results[seed]; !ok {
									results[seed] = make(map[uint32]map[uint32]uint32)
								}

								if _, ok := results[seed][skill.Index]; !ok {
									results[seed][skill.Index] = make(map[uint32]uint32)
								}

								if augment.StatRolls != nil {
									results[seed][skill.Index][key] = augment.StatRolls[uint32(i)]
								}
							}
						}
					}
				}
			}
		}
	}

	return results
}
