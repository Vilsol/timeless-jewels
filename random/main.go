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
}

func NewRNG() *NumberGenerator {
	return &NumberGenerator{state: [4]uint32{}}
}

func (g *NumberGenerator) Reset(passiveSkill *data.PassiveSkill, timelessJewel data.TimelessJewel) {
	g.state[0] = InitialStateConstant0
	g.state[1] = InitialStateConstant1
	g.state[2] = InitialStateConstant2
	g.state[3] = InitialStateConstant3

	g.Initialize([]uint32{
		passiveSkill.PassiveSkillGraphID,
		timelessJewel.GetSeed(),
	})
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

	for i := 0; i < 5; i++ {
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

	for i := 0; i < 4; i++ {
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

	for i := 0; i < 8; i++ {
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

	if (b & 1) != 0 {
		return a ^ 0x3793FDFF
	}

	return a
}

func (g *NumberGenerator) GenerateUInt() uint32 {
	g.GenerateNextState()
	return g.Temper()
}

func (g *NumberGenerator) GenerateSingle(exclusiveMaximumValue uint32) uint32 {
	return g.GenerateUInt() % exclusiveMaximumValue
}

func (g *NumberGenerator) Generate(min uint32, max uint32) uint32 {
	a := min + 0x80000000
	b := max + 0x80000000

	if min >= 0x80000000 {
		a = min + 0x80000000
	}

	if max >= 0x80000000 {
		b = max + 0x80000000
	}

	roll := g.GenerateSingle((b - a) + 1)

	return (roll + a) + 0x80000000
}

func ManipulateAlpha(value uint32) uint32 {
	return (value ^ (value >> 27)) * TinyMT32Alpha
}

func ManipulateBravo(value uint32) uint32 {
	return (value ^ (value >> 27)) * TinyMT32Bravo
}
