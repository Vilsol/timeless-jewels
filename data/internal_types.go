package data

type TimelessJewel struct {
	Seed                   uint32
	AlternateTreeVersion   *AlternateTreeVersion
	TimelessJewelConqueror *TimelessJewelConqueror
}

func (t TimelessJewel) GetSeed() uint32 {
	if t.AlternateTreeVersion.Index != uint32(ElegantHubris) {
		return t.Seed
	}
	return t.Seed / 20
}

type TimelessJewelConqueror struct {
	Index   uint32
	Version uint32
}

type PassiveSkillType uint32

const (
	None = PassiveSkillType(iota)
	SmallAttribute
	SmallNormal
	Notable
	KeyStone
	JewelSocket
)

type AlternatePassiveSkillInformation struct {
	AlternatePassiveSkill                *AlternatePassiveSkill
	StatRolls                            map[uint32]uint32
	AlternatePassiveAdditionInformations []AlternatePassiveAdditionInformation
}

type AlternatePassiveAdditionInformation struct {
	AlternatePassiveAddition *AlternatePassiveAddition
	StatRolls                map[uint32]uint32
}
