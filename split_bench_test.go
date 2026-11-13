package main

import (
	"testing"

	"github.com/Vilsol/timeless-jewels/calculator"
	"github.com/Vilsol/timeless-jewels/data"
)

// These benchmarks exist to A/B the data-package split (embedded init -> injected
// Initialize). They deliberately cover the surfaces the split rebuilds and nothing else:
//
//	idToPassiveSkill                 GetPassiveSkillByIndex
//	idToAlternateTreeVersion         GetAlternateTreeVersionIndex
//	idToAlternatePassiveSkill        GetAlternatePassiveSkillByIndex
//	idToAlternatePassiveAddition     GetAlternatePassiveAdditionByIndex
//	reverseAlternatePassiveSkills    GetApplicableAlternatePassiveSkills
//	reverseAlternatePassiveAdditions GetApplicableAlternatePassiveAdditions
//	AlternatePassiveSkills (slice)   GetAlternatePassiveSkillKeyStone (linear scan)
//
// ⚠️ BenchmarkAll and the reverse-search benchmarks are NOT usable as an A/B: BenchmarkAll
// sweeps every seed of every jewel (minutes per iteration) and both iterate
// TimelessJewelConquerors, a map, so the conqueror sampled changes per run. Everything here
// is an explicit ordered pair with fixed seeds so repeated runs measure the same work.

// benchPairs is one conqueror per jewel type with a seed inside that type's range.
// ElegantHubris is Special (seed is stored /20), so its seed must be a multiple of 20.
var benchPairs = []struct {
	name      string
	jewel     data.JewelType
	conqueror data.Conqueror
	seed      uint32
}{
	{"GloriousVanity", data.GloriousVanity, data.Xibaqua, 2000},
	{"LethalPride", data.LethalPride, data.Kaom, 12000},
	{"BrutalRestraint", data.BrutalRestraint, data.Deshret, 2000},
	{"MilitantFaith", data.MilitantFaith, data.Venarius, 4000},
	{"ElegantHubris", data.ElegantHubris, data.Cadiro, 45060},
	{"HeroicTragedy", data.HeroicTragedy, data.Vorana, 2000},
}

// TestBenchCorpusIsNonDegenerate is the control for the benchmarks below, and it is not
// optional. Calculate returns an EMPTY struct immediately when the passive is not valid for
// alteration or the conqueror does not resolve — so a benchmark over the wrong ids measures
// six early returns per pair and reports a beautifully stable ns/op that would be flat across
// ANY change to the indexes, including one that broke them. This REFUSES rather than passing
// vacuously if the corpus stops producing real work.
func TestBenchCorpusIsNonDegenerate(t *testing.T) {
	passiveIDs := passiveIndices(t, reversePassives)
	for _, p := range benchPairs {
		replaced, augmented := 0, 0
		for _, id := range passiveIDs {
			result := calculator.Calculate(id, p.seed, p.jewel, p.conqueror)
			if result.AlternatePassiveSkill != nil {
				replaced++
			}
			if len(result.AlternatePassiveAdditionInformations) > 0 {
				augmented++
			}
		}
		if replaced+augmented == 0 {
			t.Errorf("REFUSING: %s/%s seed %d produced no replacements and no augments over %d passives — the benchmark would be timing early returns",
				p.jewel, p.conqueror, p.seed, len(passiveIDs))
			continue
		}
		t.Logf("%-16s %-9s seed %6d: %2d replaced, %2d augmented of %d", p.jewel, p.conqueror, p.seed, replaced, augmented, len(passiveIDs))
	}
}

// BenchmarkCalculate is the hot path go-pob will call once per in-radius node per build.
// reversePassives comes from reverse_test.go — 61 ids already proven valid for alteration.
func BenchmarkCalculate(b *testing.B) {
	passiveIDs := passiveIndices(b, reversePassives)
	b.ReportAllocs()

	for _, p := range benchPairs {
		b.Run(p.name, func(b *testing.B) {
			for range b.N {
				for _, id := range passiveIDs {
					calculator.Calculate(id, p.seed, p.jewel, p.conqueror)
				}
			}
		})
	}
}

// BenchmarkLookups isolates the index reads from the calculation around them. If the split
// were to leave an index stale or swap a map for a scan, this moves before BenchmarkCalculate
// does — Calculate spends most of its time in the RNG, which would mask a lookup regression.
func BenchmarkLookups(b *testing.B) {
	passiveIDs := passiveIndices(b, reversePassives)
	b.ReportAllocs()

	b.Run("PassiveSkillByIndex", func(b *testing.B) {
		for range b.N {
			for _, id := range passiveIDs {
				sinkPassive = data.GetPassiveSkillByIndex(id)
			}
		}
	})

	b.Run("AlternateTreeVersionIndex", func(b *testing.B) {
		for range b.N {
			for _, p := range benchPairs {
				sinkTreeVersion = data.GetAlternateTreeVersionIndex(uint32(p.jewel))
			}
		}
	})

	b.Run("AlternatePassiveSkillByIndex", func(b *testing.B) {
		for range b.N {
			for i := range uint32(128) {
				sinkAltSkill = data.GetAlternatePassiveSkillByIndex(i)
			}
		}
	})

	b.Run("AlternatePassiveAdditionByIndex", func(b *testing.B) {
		for range b.N {
			for i := range uint32(128) {
				sinkAltAddition = data.GetAlternatePassiveAdditionByIndex(i)
			}
		}
	})

	// The two reverse indexes, reached the way tree_manager.go reaches them.
	b.Run("ApplicableAlternates", func(b *testing.B) {
		jewels := make([]data.TimelessJewel, 0, len(benchPairs))
		for _, p := range benchPairs {
			jewels = append(jewels, data.TimelessJewel{
				Seed:                   p.seed,
				AlternateTreeVersion:   data.GetAlternateTreeVersionIndex(uint32(p.jewel)),
				TimelessJewelConqueror: data.TimelessJewelConquerors[p.jewel][p.conqueror],
			})
		}
		skills := make([]*data.PassiveSkill, 0, len(passiveIDs))
		for _, id := range passiveIDs {
			skills = append(skills, data.GetPassiveSkillByIndex(id))
		}

		b.ResetTimer()
		for range b.N {
			for _, j := range jewels {
				for _, s := range skills {
					sinkAltSkills = data.GetApplicableAlternatePassiveSkills(s, j)
					sinkAltAdditions = data.GetApplicableAlternatePassiveAdditions(s, j)
				}
			}
		}
	})

	// The one lookup that is a LINEAR SCAN of the exported slice rather than a map read,
	// so it is the one a re-ordered injected slice would change the cost of.
	b.Run("KeyStone", func(b *testing.B) {
		jewels := make([]data.TimelessJewel, 0, len(benchPairs))
		for _, p := range benchPairs {
			jewels = append(jewels, data.TimelessJewel{
				Seed:                   p.seed,
				AlternateTreeVersion:   data.GetAlternateTreeVersionIndex(uint32(p.jewel)),
				TimelessJewelConqueror: data.TimelessJewelConquerors[p.jewel][p.conqueror],
			})
		}

		b.ResetTimer()
		for range b.N {
			for _, j := range jewels {
				sinkAltSkill = data.GetAlternatePassiveSkillKeyStone(j)
			}
		}
	})
}

// Package-level sinks so the compiler cannot eliminate the lookups above.
var (
	sinkPassive      *data.PassiveSkill
	sinkTreeVersion  *data.AlternateTreeVersion
	sinkAltSkill     *data.AlternatePassiveSkill
	sinkAltAddition  *data.AlternatePassiveAddition
	sinkAltSkills    []*data.AlternatePassiveSkill
	sinkAltAdditions []*data.AlternatePassiveAddition
)
