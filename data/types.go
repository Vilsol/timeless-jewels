package data

type Stat struct {
	Index uint32 `json:"_rid"`
	ID    string `json:"Id"`
	Text  string `json:"Text"`
}

type PassiveSkill struct {
	Index               uint32   `json:"_rid"`
	ID                  string   `json:"Id"`
	StatIndices         []uint32 `json:"Stats"`
	PassiveSkillGraphID uint32   `json:"PassiveSkillGraphId"`
	Name                string   `json:"Name"`
	IsKeystone          bool     `json:"IsKeystone"`
	IsNotable           bool     `json:"IsNotable"`
	IsJewelSocket       bool     `json:"IsJewelSocket"`
}

type AlternateTreeVersion struct {
	Index                                  uint32 `json:"_rid"`
	ID                                     string `json:"Id"`
	AreSmallAttributePassiveSkillsReplaced bool   `json:"Unknown2"`
	AreSmallNormalPassiveSkillsReplaced    bool   `json:"Unknown3"`
	MinimumAdditions                       uint32 `json:"Unknown6"`
	MaximumAdditions                       uint32 `json:"Unknown7"`
	NotableReplacementSpawnWeight          uint32 `json:"Unknown10"`
}

type AlternatePassiveSkill struct {
	Index                    uint32             `json:"_rid"`
	ID                       string             `json:"Id"`
	AlternateTreeVersionsKey uint32             `json:"AlternateTreeVersionsKey"`
	Name                     string             `json:"Name"`
	PassiveType              []PassiveSkillType `json:"PassiveType"`
	StatsKeys                []uint32           `json:"StatsKeys"`
	Stat1Min                 uint32             `json:"Stat1Min"`
	Stat1Max                 uint32             `json:"Stat1Max"`
	Stat2Min                 uint32             `json:"Stat2Min"`
	Stat2Max                 uint32             `json:"Stat2Max"`
	Stat3Min                 uint32             `json:"Unknown10"`
	Stat3Max                 uint32             `json:"Unknown11"`
	Stat4Min                 uint32             `json:"Unknown12"`
	Stat4Max                 uint32             `json:"Unknown13"`
	SpawnWeight              uint32             `json:"SpawnWeight"`
	ConquerorIndex           uint32             `json:"Unknown19"`
	RandomMin                uint32             `json:"RandomMin"`
	RandomMax                uint32             `json:"RandomMax"`
	ConquerorVersion         uint32             `json:"Unknown25"`
}

type AlternatePassiveAddition struct {
	Index                    uint32             `json:"_rid"`
	ID                       string             `json:"Id"`
	AlternateTreeVersionsKey uint32             `json:"AlternateTreeVersionsKey"`
	SpawnWeight              uint32             `json:"SpawnWeight"`
	StatsKeys                []uint32           `json:"StatsKeys"`
	Stat1Min                 uint32             `json:"Stat1Min"`
	Stat1Max                 uint32             `json:"Stat1Max"`
	Stat2Min                 uint32             `json:"Unknown7"`
	Stat2Max                 uint32             `json:"Unknown8"`
	PassiveType              []PassiveSkillType `json:"PassiveType"`
}
