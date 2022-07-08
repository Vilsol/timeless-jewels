package data

type JewelType uint

const (
	GloriousVanity = JewelType(iota + 1)
	LethalPride
	BrutalRestraint
	MilitantFaith
	ElegantHubris
)

func (t JewelType) String() string {
	switch t {
	case GloriousVanity:
		return "Glorious Vanity"
	case LethalPride:
		return "Lethal Pride"
	case BrutalRestraint:
		return "Brutal Restraint"
	case MilitantFaith:
		return "Militant Faith"
	case ElegantHubris:
		return "Elegant Hubris"
	default:
		return "N/A"
	}
}

type Conqueror string

const (
	Xibaqua = Conqueror("Xibaqua")
	Zerphi  = Conqueror("Zerphi")
	Ahuana  = Conqueror("Ahuana")
	Doryani = Conqueror("Doryani")

	Kaom    = Conqueror("Kaom")
	Rakiata = Conqueror("Rakiata")
	Kiloava = Conqueror("Kiloava")
	Akoya   = Conqueror("Akoya")

	Deshret = Conqueror("Deshret")
	Balbala = Conqueror("Balbala")
	Asenath = Conqueror("Asenath")
	Nasima  = Conqueror("Nasima")

	Venarius = Conqueror("Venarius")
	Maxarius = Conqueror("Maxarius")
	Dominus  = Conqueror("Dominus")
	Avarius  = Conqueror("Avarius")

	Cadiro   = Conqueror("Cadiro")
	Victario = Conqueror("Victario")
	Chitus   = Conqueror("Chitus")
	Caspiro  = Conqueror("Caspiro")
)

var TimelessJewelConquerors = map[JewelType]map[Conqueror]*TimelessJewelConqueror{
	GloriousVanity: {
		Xibaqua: &TimelessJewelConqueror{
			Index:   1,
			Version: 0,
		},
		Zerphi: &TimelessJewelConqueror{
			Index:   2,
			Version: 0,
		},
		Ahuana: &TimelessJewelConqueror{
			Index:   2,
			Version: 1,
		},
		Doryani: &TimelessJewelConqueror{
			Index:   3,
			Version: 0,
		},
	},
	LethalPride: {
		Kaom: &TimelessJewelConqueror{
			Index:   1,
			Version: 0,
		},
		Rakiata: &TimelessJewelConqueror{
			Index:   2,
			Version: 0,
		},
		Kiloava: &TimelessJewelConqueror{
			Index:   3,
			Version: 0,
		},
		Akoya: &TimelessJewelConqueror{
			Index:   3,
			Version: 1,
		},
	},
	BrutalRestraint: {
		Deshret: &TimelessJewelConqueror{
			Index:   1,
			Version: 0,
		},
		Balbala: &TimelessJewelConqueror{
			Index:   1,
			Version: 1,
		},
		Asenath: &TimelessJewelConqueror{
			Index:   2,
			Version: 0,
		},
		Nasima: &TimelessJewelConqueror{
			Index:   3,
			Version: 0,
		},
	},
	MilitantFaith: {
		Venarius: &TimelessJewelConqueror{
			Index:   1,
			Version: 0,
		},
		Maxarius: &TimelessJewelConqueror{
			Index:   1,
			Version: 1,
		},
		Dominus: &TimelessJewelConqueror{
			Index:   2,
			Version: 0,
		},
		Avarius: &TimelessJewelConqueror{
			Index:   3,
			Version: 0,
		},
	},
	ElegantHubris: {
		Cadiro: &TimelessJewelConqueror{
			Index:   1,
			Version: 0,
		},
		Victario: &TimelessJewelConqueror{
			Index:   2,
			Version: 0,
		},
		Chitus: &TimelessJewelConqueror{
			Index:   3,
			Version: 0,
		},
		Caspiro: &TimelessJewelConqueror{
			Index:   3,
			Version: 1,
		},
	},
}

type Range struct {
	Min     uint32
	Max     uint32
	Special bool
}

var TimelessJewelSeedRanges = map[JewelType]Range{
	GloriousVanity: {
		Min: 100,
		Max: 8000,
	},
	LethalPride: {
		Min: 10000,
		Max: 18000,
	},
	BrutalRestraint: {
		Min: 500,
		Max: 8000,
	},
	MilitantFaith: {
		Min: 2000,
		Max: 10000,
	},
	ElegantHubris: {
		Min:     2000,
		Max:     160000,
		Special: true,
	},
}
