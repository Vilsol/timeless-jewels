package main

import (
	"testing"

	"github.com/Vilsol/timeless-jewels/data"
)

// TestDoubleResetPopulation sizes the redundancy the Reset memo removes, because the memo
// measured ~0 on BenchmarkCalculate and "the mechanism is real" is not the same claim as "the
// mechanism runs".
//
// The duplicate Reset only happens on the notable branch of IsPassiveSkillReplaced, and only
// when the tree version's NotableReplacementSpawnWeight is strictly between 0 and 100 — at 0
// or >=100 the function returns before it ever seeds. So the memo's ceiling is the share of
// (passive, jewel) pairs that are notables under a tree version in that band.
func TestDoubleResetPopulation(t *testing.T) {
	passiveIDs := passiveIndices(t, reversePassives)
	bands := map[string][]string{}
	for _, v := range data.AlternateTreeVersions {
		band := "0 (never seeds)"
		switch {
		case v.NotableReplacementSpawnWeight >= 100:
			band = ">=100 (never seeds)"
		case v.NotableReplacementSpawnWeight > 0:
			band = "1..99 (SEEDS TWICE)"
		}
		bands[band] = append(bands[band], v.ID)
	}
	for band, ids := range bands {
		t.Logf("NotableReplacementSpawnWeight %-20s %2d tree versions: %v", band, len(ids), ids)
	}

	// ⚠️ THE MEMO IN random.Reset EXISTS ONLY FOR THIS BAND. If a data update empties it, the
	// memo is dead weight carrying 24 bytes and a comparison on every seed, and nothing else
	// would ever say so — the benchmarks would simply stop showing a difference, which reads
	// identically to "the optimisation was never worth much".
	if len(bands["1..99 (SEEDS TWICE)"]) == 0 {
		t.Error("no tree version seeds twice any more — random.(*NumberGenerator).Reset's memo is now dead code and should be removed")
	}

	notables := 0
	for _, id := range passiveIDs {
		if skill := data.GetPassiveSkillByIndex(id); skill != nil && skill.IsNotable {
			notables++
		}
	}
	t.Logf("benchmark corpus: %d of %d passives are notable", notables, len(passiveIDs))

	// The corpus-wide share, so the answer does not depend on my 61 hand-picked ids.
	allNotable, allValid := 0, 0
	for _, skill := range data.PassiveSkills {
		if !data.IsPassiveSkillValidForAlteration(skill) {
			continue
		}
		allValid++
		if skill.IsNotable {
			allNotable++
		}
	}
	t.Logf("whole tree: %d of %d alterable passives are notable (%.1f%%)",
		allNotable, allValid, 100*float64(allNotable)/float64(allValid))
}
