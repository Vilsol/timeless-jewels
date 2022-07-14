package data

import "strconv"

func GetApplicableAlternatePassiveAdditions(passiveSkill *PassiveSkill, timelessJewel TimelessJewel) []*AlternatePassiveAddition {
	return reverseAlternatePassiveAdditions[GetPassiveSkillType(passiveSkill)][timelessJewel.AlternateTreeVersion.Index]
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

func GetAlternatePassiveSkillKeyStone(timelessJewel TimelessJewel) *AlternatePassiveSkill {
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

func GetApplicableAlternatePassiveSkills(passiveSkill *PassiveSkill, timelessJewel TimelessJewel) []*AlternatePassiveSkill {
	return reverseAlternatePassiveSkills[GetPassiveSkillType(passiveSkill)][timelessJewel.AlternateTreeVersion.Index]
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

func GetApplicablePassives() []*PassiveSkill {
	applicable := make([]*PassiveSkill, 0)
	for _, skill := range PassiveSkills {
		if skill.Name == "" {
			continue
		}

		if skill.IsJewelSocket {
			continue
		}

		if node, ok := SkillTreeData.Nodes[strconv.Itoa(int(skill.PassiveSkillGraphID))]; ok {
			if node.AscendancyName != nil {
				continue
			}

			if node.IsProxy != nil && *node.IsProxy {
				continue
			}

			if node.IsBlighted != nil && *node.IsBlighted {
				continue
			}

			if node.IsMastery != nil && *node.IsMastery {
				continue
			}

			applicable = append(applicable, skill)
		}
	}
	return applicable
}
