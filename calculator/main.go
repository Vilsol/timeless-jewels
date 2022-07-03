package calculator

import "timeless-jewels/data"

func Calculate(passiveID uint32, seed uint32, timelessJewelType data.JewelType, conqueror data.Conqueror) *data.AlternatePassiveSkillInformation {
	var passiveSkill *data.PassiveSkill

	for _, skill := range data.PassiveSkills {
		if skill.Index == passiveID {
			passiveSkill = skill
			break
		}
	}

	if !data.IsPassiveSkillValidForAlteration(passiveSkill) {
		return nil
	}

	alternateTreeVersionIndex := uint32(timelessJewelType)
	var alternateTreeVersion *data.AlternateTreeVersion

	for _, version := range data.AlternateTreeVersions {
		if version.Index == alternateTreeVersionIndex {
			alternateTreeVersion = version
			break
		}
	}

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
