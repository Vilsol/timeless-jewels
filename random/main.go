package random

import (
	"github.com/Vilsol/timeless-jewels/data"
)

const (
	InitialStateConstant0 = 0x40336050
	InitialStateConstant1 = 0xCFA3723C
	InitialStateConstant2 = 0x3CAC5F6F
	InitialStateConstant3 = 0x3793FDFF

	TinyMT32SH0   = 1
	TinyMT32SH1   = 10
	TinyMT32Mask  = 0x7FFFFFFF
	TinyMT32Alpha = 0x19660D
	TinyMT32Bravo = 0x5D588B65
)

type NumberGenerator struct {
	state [4]uint32

	// The seeded state is a pure function of (graph id, jewel seed), and the calculator asks
	// for the same pair twice in a row on every notable: IsPassiveSkillReplaced reseeds to
	// roll the replacement chance, then ReplacePassiveSkill reseeds to the identical state to
	// roll it again. Reset is the most expensive operation here — 11 mixing rounds plus 8 full
	// state generations, ~35% of Calculate — so the second one is memoised rather than redone.
	seededKey   [2]uint32
	seededState [4]uint32
	seeded      bool

	// Abyss jewels draw bounded values like MSVC's _Rng_from_urng (rejecting the biased tail);
	// legion jewels keep plain modulo, which is what PoB's legion LUTs were generated with.
	msvcDraw bool
}

func NewRNG() *NumberGenerator {
	return &NumberGenerator{state: [4]uint32{}}
}

func (g *NumberGenerator) Reset(passiveSkill *data.PassiveSkill, timelessJewel data.TimelessJewel) {
	key := [2]uint32{passiveSkill.PassiveSkillGraphID, timelessJewel.GetSeed()}

	// Before the memo check: the key omits the jewel type, so a hit may come from the other mode.
	g.msvcDraw = data.JewelType(timelessJewel.AlternateTreeVersion.Index) >= data.AbyssTecrod

	if g.seeded && g.seededKey == key {
		g.state = g.seededState
		return
	}

	g.state[0] = InitialStateConstant0
	g.state[1] = InitialStateConstant1
	g.state[2] = InitialStateConstant2
	g.state[3] = InitialStateConstant3

	g.Initialize([]uint32{key[0], key[1]})

	g.seededKey = key
	g.seededState = g.state
	g.seeded = true
}

func (g *NumberGenerator) Initialize(seeds []uint32) {
	index := uint32(1)

	for _, seed := range seeds {
		roundState := ManipulateAlpha(
			g.state[(index%4)] ^
				g.state[((index+1)%4)] ^
				g.state[(((index+4)-1)%4)])

		g.state[((index + 1) % 4)] += roundState

		roundState += seed + index

		g.state[(((index + 1) + 1) % 4)] += roundState
		g.state[(index % 4)] = roundState

		index = (index + 1) % 4
	}

	for range 5 {
		roundState := ManipulateAlpha(
			g.state[(index%4)] ^
				g.state[((index+1)%4)] ^
				g.state[(((index+4)-1)%4)])

		g.state[((index + 1) % 4)] += roundState

		roundState += index

		g.state[(((index + 1) + 1) % 4)] += roundState
		g.state[(index % 4)] = roundState

		index = (index + 1) % 4
	}

	for range 4 {
		roundState := ManipulateBravo(
			g.state[(index%4)] +
				g.state[((index+1)%4)] +
				g.state[(((index+4)-1)%4)])

		g.state[((index + 1) % 4)] ^= roundState

		roundState -= index

		g.state[(((index + 1) + 1) % 4)] ^= roundState
		g.state[(index % 4)] = roundState

		index = (index + 1) % 4
	}

	for range 8 {
		g.GenerateNextState()
	}
}

func (g *NumberGenerator) GenerateNextState() {
	a := g.state[3]
	b := ((g.state[0] & TinyMT32Mask) ^ g.state[1]) ^ g.state[2]

	a ^= a << TinyMT32SH0
	b ^= (b >> TinyMT32SH0) ^ a

	g.state[0] = g.state[1]
	g.state[1] = g.state[2]
	g.state[2] = a ^ (b << TinyMT32SH1)
	g.state[3] = b

	g.state[1] ^= -(b & 1) & 0x8F7011EE
	g.state[2] ^= -(b & 1) & 0xFC78FF1F
}

func (g *NumberGenerator) Temper() uint32 {
	b := g.state[0] + (g.state[2] >> 8)
	a := g.state[3] ^ b
	return a ^ (-(b & 1) & 0x3793FDFF)
}

func (g *NumberGenerator) GenerateUInt() uint32 {
	g.GenerateNextState()
	return g.Temper()
}

func (g *NumberGenerator) GenerateSingle(exclusiveMaximumValue uint32) uint32 {
	if g.msvcDraw {
		return g.generateSingleMSVC(exclusiveMaximumValue)
	}
	return g.GenerateUInt() % exclusiveMaximumValue
}

// generateSingleMSVC is MSVC STL's _Rng_from_urng::operator() for a 32-bit engine
// (stl/inc/xutility): no draw at all for n <= 1, otherwise redraw while v lands in the
// tail past the last whole multiple of n.
func (g *NumberGenerator) generateSingleMSVC(n uint32) uint32 {
	if n <= 1 {
		return 0
	}
	for {
		v := g.GenerateUInt()
		if v/n < 0xFFFFFFFF/n || 0xFFFFFFFF%n == n-1 {
			return v % n
		}
	}
}

func (g *NumberGenerator) GenerateSigned(minValue int32, maxValue int32) int32 {
	return int32(g.Generate(uint32(minValue), uint32(maxValue)))
}

func (g *NumberGenerator) Generate(minValue uint32, maxValue uint32) uint32 {
	a := minValue + 0x80000000
	b := maxValue + 0x80000000

	roll := g.GenerateSingle((b - a) + 1)

	return (roll + a) + 0x80000000
}

func ManipulateAlpha(value uint32) uint32 {
	return (value ^ (value >> 27)) * TinyMT32Alpha
}

func ManipulateBravo(value uint32) uint32 {
	return (value ^ (value >> 27)) * TinyMT32Bravo
}
