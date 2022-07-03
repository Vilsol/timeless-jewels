package data

func GetApplicableAlternatePassiveAdditions(passiveSkill *PassiveSkill, timelessJewel *TimelessJewel) []*AlternatePassiveAddition {
	applicableAlternatePassiveAdditions := make([]*AlternatePassiveAddition, 0)
	skillType := GetPassiveSkillType(passiveSkill)
	for _, addition := range AlternatePassiveAdditions {
		isSkillType := false
		for _, passiveType := range addition.PassiveType {
			if passiveType == skillType {
				isSkillType = true
				break
			}
		}

		if addition.AlternateTreeVersionsKey != timelessJewel.AlternateTreeVersion.Index || !isSkillType {
			continue
		}

		applicableAlternatePassiveAdditions = append(applicableAlternatePassiveAdditions, addition)
	}
	return applicableAlternatePassiveAdditions
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
	applicableAlternatePassiveSkills := make([]*AlternatePassiveSkill, 0)
	skillType := GetPassiveSkillType(passiveSkill)
	for _, addition := range AlternatePassiveSkills {
		isSkillType := false
		for _, passiveType := range addition.PassiveType {
			if passiveType == skillType {
				isSkillType = true
				break
			}
		}

		if addition.AlternateTreeVersionsKey != timelessJewel.AlternateTreeVersion.Index || !isSkillType {
			continue
		}

		applicableAlternatePassiveSkills = append(applicableAlternatePassiveSkills, addition)
	}
	return applicableAlternatePassiveSkills
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

func GetStatByIndex(index uint32) *Stat {
	for _, stat := range Stats {
		if stat.Index == index {
			return stat
		}
	}
	return nil
}

func GetAlternatePassiveSkillByIndex(index uint32) *AlternatePassiveSkill {
	for _, skill := range AlternatePassiveSkills {
		if skill.Index == index {
			return skill
		}
	}
	return nil
}

func GetAlternatePassiveAdditionByIndex(index uint32) *AlternatePassiveAddition {
	for _, addition := range AlternatePassiveAdditions {
		if addition.Index == index {
			return addition
		}
	}
	return nil
}
