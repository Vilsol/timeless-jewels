package data

type SkillTree struct {
	Tree            string                `json:"tree"`
	Classes         []Class               `json:"classes"`
	Groups          map[string]Group      `json:"groups"`
	Nodes           map[string]Node       `json:"nodes"`
	ExtraImages     map[string]ExtraImage `json:"extraImages"`
	JewelSlots      []int64               `json:"jewelSlots"`
	MinX            int64                 `json:"min_x"`
	MinY            int64                 `json:"min_y"`
	MaxX            int64                 `json:"max_x"`
	MaxY            int64                 `json:"max_y"`
	Assets          Assets                `json:"assets"`
	Constants       Constants             `json:"constants"`
	Sprites         Sprites               `json:"sprites"`
	ImageZoomLevels []float64             `json:"imageZoomLevels"`
	Points          Points                `json:"points"`
}

type Assets struct {
	PSSkillFrame                            Asset       `json:"PSSkillFrame"`
	PSSkillFrameHighlighted                 Asset       `json:"PSSkillFrameHighlighted"`
	PSSkillFrameActive                      Asset       `json:"PSSkillFrameActive"`
	PSGroupBackground1                      Asset       `json:"PSGroupBackground1"`
	PSGroupBackground2                      Asset       `json:"PSGroupBackground2"`
	PSGroupBackground3                      Asset       `json:"PSGroupBackground3"`
	NotableFrameUnallocated                 Asset       `json:"NotableFrameUnallocated"`
	NotableFrameCanAllocate                 Asset       `json:"NotableFrameCanAllocate"`
	NotableFrameAllocated                   Asset       `json:"NotableFrameAllocated"`
	KeystoneFrameUnallocated                Asset       `json:"KeystoneFrameUnallocated"`
	KeystoneFrameCanAllocate                Asset       `json:"KeystoneFrameCanAllocate"`
	KeystoneFrameAllocated                  Asset       `json:"KeystoneFrameAllocated"`
	Orbit1Normal                            Asset       `json:"Orbit1Normal"`
	Orbit1Intermediate                      Asset       `json:"Orbit1Intermediate"`
	Orbit1Active                            Asset       `json:"Orbit1Active"`
	Orbit2Normal                            Asset       `json:"Orbit2Normal"`
	Orbit2Intermediate                      Asset       `json:"Orbit2Intermediate"`
	Orbit2Active                            Asset       `json:"Orbit2Active"`
	Orbit3Normal                            Asset       `json:"Orbit3Normal"`
	Orbit3Intermediate                      Asset       `json:"Orbit3Intermediate"`
	Orbit3Active                            Asset       `json:"Orbit3Active"`
	Orbit4Normal                            Asset       `json:"Orbit4Normal"`
	Orbit4Intermediate                      Asset       `json:"Orbit4Intermediate"`
	Orbit4Active                            Asset       `json:"Orbit4Active"`
	Orbit5Normal                            Asset       `json:"Orbit5Normal"`
	Orbit5Intermediate                      Asset       `json:"Orbit5Intermediate"`
	Orbit5Active                            Asset       `json:"Orbit5Active"`
	Orbit6Normal                            Asset       `json:"Orbit6Normal"`
	Orbit6Intermediate                      Asset       `json:"Orbit6Intermediate"`
	Orbit6Active                            Asset       `json:"Orbit6Active"`
	LineConnectorNormal                     Asset       `json:"LineConnectorNormal"`
	LineConnectorIntermediate               Asset       `json:"LineConnectorIntermediate"`
	LineConnectorActive                     Asset       `json:"LineConnectorActive"`
	PSLineDeco                              Asset       `json:"PSLineDeco"`
	PSLineDecoHighlighted                   Asset       `json:"PSLineDecoHighlighted"`
	PSPointsFrame                           JewelCircle `json:"PSPointsFrame"`
	Background2                             Asset       `json:"Background2"`
	ImgPSFadeCorner                         JewelCircle `json:"imgPSFadeCorner"`
	ImgPSFadeSide                           JewelCircle `json:"imgPSFadeSide"`
	GroupBackgroundSmallAlt                 Asset       `json:"GroupBackgroundSmallAlt"`
	GroupBackgroundMediumAlt                Asset       `json:"GroupBackgroundMediumAlt"`
	GroupBackgroundLargeHalfAlt             Asset       `json:"GroupBackgroundLargeHalfAlt"`
	PSStartNodeBackgroundInactive           Asset       `json:"PSStartNodeBackgroundInactive"`
	Centerduelist                           Asset       `json:"centerduelist"`
	Centermarauder                          Asset       `json:"centermarauder"`
	Centerranger                            Asset       `json:"centerranger"`
	Centershadow                            Asset       `json:"centershadow"`
	Centertemplar                           Asset       `json:"centertemplar"`
	Centerwitch                             Asset       `json:"centerwitch"`
	Centerscion                             Asset       `json:"centerscion"`
	BlightedNotableFrameUnallocated         Asset       `json:"BlightedNotableFrameUnallocated"`
	BlightedNotableFrameCanAllocate         Asset       `json:"BlightedNotableFrameCanAllocate"`
	BlightedNotableFrameAllocated           Asset       `json:"BlightedNotableFrameAllocated"`
	JewelFrameUnallocated                   Asset       `json:"JewelFrameUnallocated"`
	JewelFrameCanAllocate                   Asset       `json:"JewelFrameCanAllocate"`
	JewelFrameAllocated                     Asset       `json:"JewelFrameAllocated"`
	JewelSocketActiveBlue                   Asset       `json:"JewelSocketActiveBlue"`
	JewelSocketActiveGreen                  Asset       `json:"JewelSocketActiveGreen"`
	JewelSocketActiveRed                    Asset       `json:"JewelSocketActiveRed"`
	JewelSocketActivePrismatic              Asset       `json:"JewelSocketActivePrismatic"`
	JewelSocketActiveAbyss                  Asset       `json:"JewelSocketActiveAbyss"`
	JewelSocketActiveLegion                 Asset       `json:"JewelSocketActiveLegion"`
	JewelSocketActiveAltRed                 Asset       `json:"JewelSocketActiveAltRed"`
	JewelSocketActiveAltBlue                Asset       `json:"JewelSocketActiveAltBlue"`
	JewelSocketActiveAltPurple              Asset       `json:"JewelSocketActiveAltPurple"`
	JewelCircle1                            JewelCircle `json:"JewelCircle1"`
	JewelCircle1Inverse                     JewelCircle `json:"JewelCircle1Inverse"`
	VaalJewelCircle1                        JewelCircle `json:"VaalJewelCircle1"`
	VaalJewelCircle2                        JewelCircle `json:"VaalJewelCircle2"`
	KaruiJewelCircle1                       JewelCircle `json:"KaruiJewelCircle1"`
	KaruiJewelCircle2                       JewelCircle `json:"KaruiJewelCircle2"`
	MarakethJewelCircle1                    JewelCircle `json:"MarakethJewelCircle1"`
	MarakethJewelCircle2                    JewelCircle `json:"MarakethJewelCircle2"`
	TemplarJewelCircle1                     JewelCircle `json:"TemplarJewelCircle1"`
	TemplarJewelCircle2                     JewelCircle `json:"TemplarJewelCircle2"`
	EternalEmpireJewelCircle1               JewelCircle `json:"EternalEmpireJewelCircle1"`
	EternalEmpireJewelCircle2               JewelCircle `json:"EternalEmpireJewelCircle2"`
	JewelSocketAltNormal                    Asset       `json:"JewelSocketAltNormal"`
	JewelSocketAltCanAllocate               Asset       `json:"JewelSocketAltCanAllocate"`
	JewelSocketAltActive                    Asset       `json:"JewelSocketAltActive"`
	JewelSocketActiveBlueAlt                Asset       `json:"JewelSocketActiveBlueAlt"`
	JewelSocketActiveGreenAlt               Asset       `json:"JewelSocketActiveGreenAlt"`
	JewelSocketActiveRedAlt                 Asset       `json:"JewelSocketActiveRedAlt"`
	JewelSocketActivePrismaticAlt           Asset       `json:"JewelSocketActivePrismaticAlt"`
	JewelSocketActiveAbyssAlt               Asset       `json:"JewelSocketActiveAbyssAlt"`
	JewelSocketActiveLegionAlt              Asset       `json:"JewelSocketActiveLegionAlt"`
	JewelSocketClusterAltNormal1Small       Asset       `json:"JewelSocketClusterAltNormal1Small"`
	JewelSocketClusterAltCanAllocate1Small  Asset       `json:"JewelSocketClusterAltCanAllocate1Small"`
	JewelSocketClusterAltNormal1Medium      Asset       `json:"JewelSocketClusterAltNormal1Medium"`
	JewelSocketClusterAltCanAllocate1Medium Asset       `json:"JewelSocketClusterAltCanAllocate1Medium"`
	JewelSocketClusterAltNormal1Large       Asset       `json:"JewelSocketClusterAltNormal1Large"`
	JewelSocketClusterAltCanAllocate1Large  Asset       `json:"JewelSocketClusterAltCanAllocate1Large"`
	AscendancyButton                        Asset       `json:"AscendancyButton"`
	AscendancyButtonHighlight               Asset       `json:"AscendancyButtonHighlight"`
	AscendancyButtonPressed                 Asset       `json:"AscendancyButtonPressed"`
	AscendancyFrameLargeAllocated           Asset       `json:"AscendancyFrameLargeAllocated"`
	AscendancyFrameLargeCanAllocate         Asset       `json:"AscendancyFrameLargeCanAllocate"`
	AscendancyFrameLargeNormal              Asset       `json:"AscendancyFrameLargeNormal"`
	AscendancyFrameSmallAllocated           Asset       `json:"AscendancyFrameSmallAllocated"`
	AscendancyFrameSmallCanAllocate         Asset       `json:"AscendancyFrameSmallCanAllocate"`
	AscendancyFrameSmallNormal              Asset       `json:"AscendancyFrameSmallNormal"`
	AscendancyMiddle                        Asset       `json:"AscendancyMiddle"`
	ClassesAscendant                        Asset       `json:"ClassesAscendant"`
	ClassesJuggernaut                       Asset       `json:"ClassesJuggernaut"`
	ClassesBerserker                        Asset       `json:"ClassesBerserker"`
	ClassesChieftain                        Asset       `json:"ClassesChieftain"`
	ClassesRaider                           Asset       `json:"ClassesRaider"`
	ClassesDeadeye                          Asset       `json:"ClassesDeadeye"`
	ClassesPathfinder                       Asset       `json:"ClassesPathfinder"`
	ClassesOccultist                        Asset       `json:"ClassesOccultist"`
	ClassesElementalist                     Asset       `json:"ClassesElementalist"`
	ClassesNecromancer                      Asset       `json:"ClassesNecromancer"`
	ClassesSlayer                           Asset       `json:"ClassesSlayer"`
	ClassesGladiator                        Asset       `json:"ClassesGladiator"`
	ClassesChampion                         Asset       `json:"ClassesChampion"`
	ClassesInquisitor                       Asset       `json:"ClassesInquisitor"`
	ClassesHierophant                       Asset       `json:"ClassesHierophant"`
	ClassesGuardian                         Asset       `json:"ClassesGuardian"`
	ClassesAssassin                         Asset       `json:"ClassesAssassin"`
	ClassesTrickster                        Asset       `json:"ClassesTrickster"`
	ClassesSaboteur                         Asset       `json:"ClassesSaboteur"`
	BackgroundDex                           Asset       `json:"BackgroundDex"`
	BackgroundDexInt                        Asset       `json:"BackgroundDexInt"`
	BackgroundInt                           Asset       `json:"BackgroundInt"`
	BackgroundStr                           Asset       `json:"BackgroundStr"`
	BackgroundStrDex                        Asset       `json:"BackgroundStrDex"`
	BackgroundStrInt                        Asset       `json:"BackgroundStrInt"`
	PassiveMasteryConnectedButton           Asset       `json:"PassiveMasteryConnectedButton"`
}

type Asset struct {
	The01246 string `json:"0.1246"`
	The02109 string `json:"0.2109"`
	The02972 string `json:"0.2972"`
	The03835 string `json:"0.3835"`
}

type JewelCircle struct {
	The1 string `json:"1"`
}

type Class struct {
	Name         string       `json:"name"`
	BaseStr      int64        `json:"base_str"`
	BaseDex      int64        `json:"base_dex"`
	BaseInt      int64        `json:"base_int"`
	Ascendancies []Ascendancy `json:"ascendancies"`
}

type Ascendancy struct {
	ID                string           `json:"id"`
	Name              string           `json:"name"`
	FlavourText       *string          `json:"flavourText,omitempty"`
	FlavourTextColour *string          `json:"flavourTextColour,omitempty"`
	FlavourTextRect   *FlavourTextRect `json:"flavourTextRect,omitempty"`
}

type FlavourTextRect struct {
	X      int64 `json:"x"`
	Y      int64 `json:"y"`
	Width  int64 `json:"width"`
	Height int64 `json:"height"`
}

type Constants struct {
	Classes              Classes             `json:"classes"`
	CharacterAttributes  CharacterAttributes `json:"characterAttributes"`
	PSSCentreInnerRadius int64               `json:"PSSCentreInnerRadius"`
	SkillsPerOrbit       []int64             `json:"skillsPerOrbit"`
	OrbitRadii           []int64             `json:"orbitRadii"`
}

type CharacterAttributes struct {
	Strength     int64 `json:"Strength"`
	Dexterity    int64 `json:"Dexterity"`
	Intelligence int64 `json:"Intelligence"`
}

type Classes struct {
	StrDexIntClass int64 `json:"StrDexIntClass"`
	StrClass       int64 `json:"StrClass"`
	DexClass       int64 `json:"DexClass"`
	IntClass       int64 `json:"IntClass"`
	StrDexClass    int64 `json:"StrDexClass"`
	StrIntClass    int64 `json:"StrIntClass"`
	DexIntClass    int64 `json:"DexIntClass"`
}

type ExtraImage struct {
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
	Image string  `json:"image"`
}

type Group struct {
	X       float64  `json:"x"`
	Y       float64  `json:"y"`
	Orbits  []int64  `json:"orbits"`
	Nodes   []string `json:"nodes"`
	IsProxy *bool    `json:"isProxy,omitempty"`
}

type Node struct {
	Skill                  *int64          `json:"skill,omitempty"`
	Name                   *string         `json:"name,omitempty"`
	Icon                   *string         `json:"icon,omitempty"`
	IsNotable              *bool           `json:"isNotable,omitempty"`
	Recipe                 []Recipe        `json:"recipe,omitempty"`
	Stats                  []string        `json:"stats,omitempty"`
	Group                  *int64          `json:"group,omitempty"`
	Orbit                  *int64          `json:"orbit,omitempty"`
	OrbitIndex             *int64          `json:"orbitIndex,omitempty"`
	Out                    []string        `json:"out,omitempty"`
	In                     []string        `json:"in,omitempty"`
	ReminderText           []string        `json:"reminderText,omitempty"`
	IsMastery              *bool           `json:"isMastery,omitempty"`
	InactiveIcon           *string         `json:"inactiveIcon,omitempty"`
	ActiveIcon             *string         `json:"activeIcon,omitempty"`
	ActiveEffectImage      *string         `json:"activeEffectImage,omitempty"`
	MasteryEffects         []MasteryEffect `json:"masteryEffects,omitempty"`
	GrantedStrength        *int64          `json:"grantedStrength,omitempty"`
	AscendancyName         *string         `json:"ascendancyName,omitempty"`
	GrantedDexterity       *int64          `json:"grantedDexterity,omitempty"`
	IsAscendancyStart      *bool           `json:"isAscendancyStart,omitempty"`
	IsMultipleChoice       *bool           `json:"isMultipleChoice,omitempty"`
	GrantedIntelligence    *int64          `json:"grantedIntelligence,omitempty"`
	IsJewelSocket          *bool           `json:"isJewelSocket,omitempty"`
	ExpansionJewel         *ExpansionJewel `json:"expansionJewel,omitempty"`
	GrantedPassivePoints   *int64          `json:"grantedPassivePoints,omitempty"`
	IsKeystone             *bool           `json:"isKeystone,omitempty"`
	FlavourText            []string        `json:"flavourText,omitempty"`
	IsProxy                *bool           `json:"isProxy,omitempty"`
	IsMultipleChoiceOption *bool           `json:"isMultipleChoiceOption,omitempty"`
	IsBlighted             *bool           `json:"isBlighted,omitempty"`
	ClassStartIndex        *int64          `json:"classStartIndex,omitempty"`
}

type ExpansionJewel struct {
	Size   int64   `json:"size"`
	Index  int64   `json:"index"`
	Proxy  string  `json:"proxy"`
	Parent *string `json:"parent,omitempty"`
}

type MasteryEffect struct {
	Effect       int64    `json:"effect"`
	Stats        []string `json:"stats"`
	ReminderText []string `json:"reminderText,omitempty"`
}

type Points struct {
	TotalPoints      int64 `json:"totalPoints"`
	AscendancyPoints int64 `json:"ascendancyPoints"`
}

type Sprites struct {
	NormalActive          map[string]Sprite `json:"normalActive"`
	NotableActive         map[string]Sprite `json:"notableActive"`
	KeystoneActive        map[string]Sprite `json:"keystoneActive"`
	NormalInactive        map[string]Sprite `json:"normalInactive"`
	NotableInactive       map[string]Sprite `json:"notableInactive"`
	KeystoneInactive      map[string]Sprite `json:"keystoneInactive"`
	Mastery               map[string]Sprite `json:"mastery"`
	MasteryConnected      map[string]Sprite `json:"masteryConnected"`
	MasteryActiveSelected map[string]Sprite `json:"masteryActiveSelected"`
	MasteryInactive       map[string]Sprite `json:"masteryInactive"`
	MasteryActiveEffect   map[string]Sprite `json:"masteryActiveEffect"`
	GroupBackground       map[string]Sprite `json:"groupBackground"`
	Frame                 map[string]Sprite `json:"frame"`
}

type Sprite struct {
	Filename string           `json:"filename"`
	Coords   map[string]Coord `json:"coords"`
}

type Coord struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
	W int64 `json:"w"`
	H int64 `json:"h"`
}

type Recipe string

const (
	AmberOil      Recipe = "AmberOil"
	AzureOil      Recipe = "AzureOil"
	BlackOil      Recipe = "BlackOil"
	ClearOil      Recipe = "ClearOil"
	CrimsonOil    Recipe = "CrimsonOil"
	GoldenOil     Recipe = "GoldenOil"
	IndigoOil     Recipe = "IndigoOil"
	OpalescentOil Recipe = "OpalescentOil"
	SepiaOil      Recipe = "SepiaOil"
	SilverOil     Recipe = "SilverOil"
	TealOil       Recipe = "TealOil"
	VerdantOil    Recipe = "VerdantOil"
	VioletOil     Recipe = "VioletOil"
)
