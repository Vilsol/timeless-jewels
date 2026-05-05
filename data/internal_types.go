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

// StatRolls holds up to 4 rolled stat values for a replaced passive
// (or up to 2 for an addition). Indexed positionally by the corresponding
// StatsKeys slot.
type StatRolls [4]int32

type AlternatePassiveSkillInformation struct {
	AlternatePassiveSkill                *AlternatePassiveSkill
	StatRolls                            StatRolls
	AlternatePassiveAdditionInformations []AlternatePassiveAdditionInformation
}

type AlternatePassiveAdditionInformation struct {
	AlternatePassiveAddition *AlternatePassiveAddition
	StatRolls                StatRolls
}
