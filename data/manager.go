package data

func GetApplicableAlternatePassiveAdditions(passiveSkill *PassiveSkill, timelessJewel *TimelessJewel) []*AlternatePassiveAddition {
	skillType := GetPassiveSkillType(passiveSkill)
	return reverseAlternatePassiveAdditions[skillType][timelessJewel.AlternateTreeVersion.Index]
}

func GetPassiveSkillType(passiveSkill *PassiveSkill) PassiveSkillType {
	if passiveSkill.IsJewelSocket {
		return JewelSocket
	} else if passiveSkill.IsKeystone {
		return KeyStone
	} else if passiveSkill.IsNotable {
		return Notable
	} else if len(passiveSkill.StatIndices) == 1 && IsSmallAttribute(passiveSkill.StatIndices[0]) {
		return SmallAttribute
	}

	return SmallNormal
}

func GetAlternatePassiveSkillKeyStone(timelessJewel *TimelessJewel) *AlternatePassiveSkill {
	var alternatePassiveSkillKeyStone *AlternatePassiveSkill
	for _, skill := range AlternatePassiveSkills {
		if skill.AlternateTreeVersionsKey != timelessJewel.AlternateTreeVersion.Index {
			continue
		}

		if skill.ConquerorIndex != timelessJewel.TimelessJewelConqueror.Index {
			continue
		}

		if skill.ConquerorVersion != timelessJewel.TimelessJewelConqueror.Version {
			continue
		}

		alternatePassiveSkillKeyStone = skill
		break
	}

	if alternatePassiveSkillKeyStone == nil {
		return nil
	}

	hasApplicablePassives := false
	for _, passiveType := range alternatePassiveSkillKeyStone.PassiveType {
		if passiveType == KeyStone {
			hasApplicablePassives = true
			break
		}
	}

	if !hasApplicablePassives {
		return nil
	}

	return alternatePassiveSkillKeyStone
}

func GetApplicableAlternatePassiveSkills(passiveSkill *PassiveSkill, timelessJewel *TimelessJewel) []*AlternatePassiveSkill {
	skillType := GetPassiveSkillType(passiveSkill)
	return reverseAlternatePassiveSkills[skillType][timelessJewel.AlternateTreeVersion.Index]
}

func IsSmallAttribute(stat uint32) bool {
	bitPosition := (stat + 1) - 574
	if bitPosition <= 6 && (0x49&(1<<(bitPosition))) != 0 {
		return true
	}
	return false
}

func IsPassiveSkillValidForAlteration(passiveSkill *PassiveSkill) bool {
	passiveSkillType := GetPassiveSkillType(passiveSkill)
	return (passiveSkillType != None) && (passiveSkillType != JewelSocket)
}

func GetPassiveSkillByIndex(index uint32) *PassiveSkill {
	return idToPassiveSkill[index]
}

func GetStatByIndex(index uint32) *Stat {
	return idToStat[index]
}

func GetAlternatePassiveSkillByIndex(index uint32) *AlternatePassiveSkill {
	return idToAlternatePassiveSkill[index]
}

func GetAlternatePassiveAdditionByIndex(index uint32) *AlternatePassiveAddition {
	return idToAlternatePassiveAddition[index]
}

func GetAlternateTreeVersionIndex(index uint32) *AlternateTreeVersion {
	return idToAlternateTreeVersion[index]
}

func FindSkillsWithStats(stat *Stat) []*AlternatePassiveSkill {
	skills := make([]*AlternatePassiveSkill, 0)
	for _, skill := range AlternatePassiveSkills {
		for _, key := range skill.StatsKeys {
			if key == stat.Index {
				skills = append(skills, skill)
				break
			}
		}
	}
	return skills
}

func FindAlternatePassiveAdditionsWithStats(stat *Stat) []*AlternatePassiveAddition {
	additions := make([]*AlternatePassiveAddition, 0)
	for _, addition := range AlternatePassiveAdditions {
		for _, key := range addition.StatsKeys {
			if key == stat.Index {
				additions = append(additions, addition)
				break
			}
		}
	}
	return additions
}

func FindSkillsMatchingAdditions(additions []*AlternatePassiveAddition) []*PassiveSkill {
	additionPassiveTypes := make(map[PassiveSkillType]bool)
	for _, addition := range additions {
		for _, skillType := range addition.PassiveType {
			additionPassiveTypes[skillType] = true
		}
	}

	skills := make([]*PassiveSkill, 0)
	for _, skill := range PassiveSkills {
		skillType := GetPassiveSkillType(skill)
		if _, ok := additionPassiveTypes[skillType]; ok {
			skills = append(skills, skill)
		}
	}

	return skills
}
