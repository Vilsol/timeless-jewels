package calculator

import "timeless-jewels/data"

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
