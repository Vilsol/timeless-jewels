package random

import (
	"testing"

	"github.com/Vilsol/timeless-jewels/data"
)

// A memo hit on (graph id, seed) must still take the draw mode from the jewel being reset to.
func TestResetMemoTakesDrawModeFromJewel(t *testing.T) {
	const n = 0xC0000000 // anything >= n is in the tail MSVC rejects

	passive := &data.PassiveSkill{PassiveSkillGraphID: 1}
	abyss := data.TimelessJewel{AlternateTreeVersion: &data.AlternateTreeVersion{Index: uint32(data.AbyssTecrod)}}
	legion := data.TimelessJewel{AlternateTreeVersion: &data.AlternateTreeVersion{Index: uint32(data.GloriousVanity)}}

	for seed := uint32(100); ; seed++ {
		if seed > 10000 {
			t.Fatal("no seed draws into the tail")
		}
		abyss.Seed, legion.Seed = seed, seed

		var probe NumberGenerator
		probe.Reset(passive, legion)
		if probe.GenerateUInt() < n {
			continue
		}

		draw := func(jewel data.TimelessJewel) uint32 {
			var g NumberGenerator
			g.Reset(passive, jewel)
			return g.GenerateSingle(n)
		}
		wantLegion, wantAbyss := draw(legion), draw(abyss)
		if wantLegion == wantAbyss {
			t.Fatalf("seed %d: modes agree on a tail draw", seed)
		}

		var g NumberGenerator
		g.Reset(passive, abyss)
		g.Reset(passive, legion)
		if got := g.GenerateSingle(n); got != wantLegion {
			t.Errorf("abyss then legion: got %d, want %d", got, wantLegion)
		}
		g.Reset(passive, abyss)
		if got := g.GenerateSingle(n); got != wantAbyss {
			t.Errorf("legion then abyss: got %d, want %d", got, wantAbyss)
		}
		return
	}
}
