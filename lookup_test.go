package main

import (
	"testing"

	"github.com/Vilsol/timeless-jewels/data"
)

// Tests name rows by their game Id, never by row index: GGG inserts rows into these tables every
// league, which shifts the indices but not the Ids.

func lookupByID[T any](tb testing.TB, table []T, id string, idOf func(T) string) T {
	tb.Helper()
	for _, row := range table {
		if idOf(row) == id {
			return row
		}
	}
	tb.Fatalf("no row with Id %q", id)
	var zero T
	return zero
}

func passiveIndex(tb testing.TB, id string) uint32 {
	tb.Helper()
	return lookupByID(tb, data.PassiveSkills, id, func(p *data.PassiveSkill) string { return p.ID }).Index
}

func passiveIndices(tb testing.TB, ids []string) []uint32 {
	tb.Helper()
	out := make([]uint32, len(ids))
	for i, id := range ids {
		out[i] = passiveIndex(tb, id)
	}
	return out
}

func statIndex(tb testing.TB, id string) uint32 {
	tb.Helper()
	return lookupByID(tb, data.Stats, id, func(s *data.Stat) string { return s.ID }).Index
}

func alternatePassiveSkill(tb testing.TB, id string) *data.AlternatePassiveSkill {
	tb.Helper()
	return lookupByID(tb, data.AlternatePassiveSkills, id, func(s *data.AlternatePassiveSkill) string { return s.ID })
}

func alternatePassiveAddition(tb testing.TB, id string) *data.AlternatePassiveAddition {
	tb.Helper()
	return lookupByID(tb, data.AlternatePassiveAdditions, id, func(a *data.AlternatePassiveAddition) string { return a.ID })
}
