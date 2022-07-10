package random

import "github.com/Vilsol/timeless-jewels/data"

const (
	InitialStateConstant0 = 0x40336050
	InitialStateConstant1 = 0xCFA3723C
	InitialStateConstant2 = 0x3CAC5F6F
	InitialStateConstant3 = 0x3793FDFF
)

type NumberGenerator struct {
	state []uint32
}

func NewRNG(passiveSkill *data.PassiveSkill, timelessJewel *data.TimelessJewel) *NumberGenerator {
	rng := &NumberGenerator{
		state: []uint32{
			0,
			InitialStateConstant0,
			InitialStateConstant1,
			InitialStateConstant2,
			InitialStateConstant3,
		},
	}

	rng.Initialize([]uint32{
		passiveSkill.PassiveSkillGraphID,
		timelessJewel.GetSeed(),
	})

	return rng
}

func (g *NumberGenerator) Initialize(seeds []uint32) {
	index := uint32(1)

	for _, seed := range seeds {
		roundState := ManipulateAlpha(
			g.state[(index%4)+1] ^
				g.state[((index+1)%4)+1] ^
				g.state[(((index+4)-1)%4)+1])

		g.state[((index+1)%4)+1] += roundState

		roundState += seed + index

		g.state[(((index+1)+1)%4)+1] += roundState
		g.state[(index%4)+1] = roundState

		index = (index + 1) % 4
	}

	for i := 0; i < 5; i++ {
		roundState := ManipulateAlpha(
			g.state[(index%4)+1] ^
				g.state[((index+1)%4)+1] ^
				g.state[(((index+4)-1)%4)+1])

		g.state[((index+1)%4)+1] += roundState

		roundState += index

		g.state[(((index+1)+1)%4)+1] += roundState
		g.state[(index%4)+1] = roundState

		index = (index + 1) % 4
	}

	for i := 0; i < 4; i++ {
		roundState := ManipulateBravo(
			g.state[(index%4)+1] +
				g.state[((index+1)%4)+1] +
				g.state[(((index+4)-1)%4)+1])

		g.state[((index+1)%4)+1] ^= roundState

		roundState -= index

		g.state[(((index+1)+1)%4)+1] ^= roundState
		g.state[(index%4)+1] = roundState

		index = (index + 1) % 4
	}

	for i := 0; i < 8; i++ {
		g.GenerateNextState()
	}
}

func (g *NumberGenerator) GenerateNextState() {
	a := uint32(0)
	b := uint32(0)

	a = g.state[4]
	b = ((g.state[1] & 0x7FFFFFFF) ^ g.state[2]) ^ g.state[3]

	a ^= a << 1
	b ^= (b >> 1) ^ a

	g.state[1] = g.state[2]
	g.state[2] = g.state[3]
	g.state[3] = a ^ (b << 10)
	g.state[4] = b

	g.state[2] ^= (uint32)(-((int)(b & 1)) & 0x8F7011EE)
	g.state[3] ^= (uint32)(-((int)(b & 1)) & 0xFC78FF1F)

	g.state[0]++
}

func (g *NumberGenerator) Temper() uint32 {
	a := g.state[4]

	b := g.state[1] + (g.state[3] >> 8)

	a ^= b

	if (b & 1) != 0 {
		a ^= 0x3793FDFF
	}

	return a
}

func (g *NumberGenerator) GenerateUInt() uint32 {
	g.GenerateNextState()
	return g.Temper()
}

func (g *NumberGenerator) GenerateSingle(exclusiveMaximumValue uint32) uint32 {
	maximumValue := exclusiveMaximumValue - 1
	roundState := uint32(0)
	value := uint32(0)

	for {
		for {
			value = g.GenerateUInt() | (2 * (value << 31))

			roundState = 0xFFFFFFFF | (2 * (roundState << 31))

			if !(roundState < maximumValue) {
				break
			}
		}

		if !(((value / exclusiveMaximumValue) >= roundState) && ((roundState % exclusiveMaximumValue) != maximumValue)) {
			break
		}
	}

	return value % exclusiveMaximumValue
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
	return (value ^ (value >> 27)) * 0x19660D
}

func ManipulateBravo(value uint32) uint32 {
	return (value ^ (value >> 27)) * 0x5D588B65
}
