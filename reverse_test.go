package main

import (
	"testing"

	"github.com/MarvinJWendt/testza"
	"github.com/Vilsol/timeless-jewels/calculator"
	"github.com/Vilsol/timeless-jewels/data"
)

var passiveIDs = []uint32{2211, 662, 2095, 600, 944, 75, 2556, 1094, 2589, 2505, 458, 2185, 2030, 625, 854, 1175, 2204, 1174, 13, 2184, 1066, 2192, 702, 1212, 456, 1068, 2547, 457, 1164, 698, 852, 1210, 516, 1245, 60, 447, 2096, 853, 451, 61, 2093, 1211, 134, 2212, 579, 440, 474, 2183, 2408, 2206, 518, 628, 624, 2340, 701, 2031, 2491, 2519, 9, 2205, 11}

func TestReverseGloriousVanity(t *testing.T) {
	statIDs := []uint32{25}
	result := calculator.ReverseSearch(passiveIDs, statIDs, data.GloriousVanity, data.Xibaqua, nil)
	testza.AssertLen(t, result, 7294)
	testza.AssertLen(t, result[1001], 2)
	testza.AssertEqual(t, int32(7), result[1001][1175][statIDs[0]])
}

func TestReverseElegantHubris(t *testing.T) {
	statIDs := []uint32{25}
	result := calculator.ReverseSearch(passiveIDs, statIDs, data.ElegantHubris, data.Cadiro, nil)
	testza.AssertLen(t, result, 2375)
	testza.AssertLen(t, result[45060], 1)
	testza.AssertEqual(t, int32(80), result[45060][2556][statIDs[0]])
}

func BenchmarkGloriousVanity(b *testing.B) {
	b.ReportAllocs()
	b.Run("cached", func(b *testing.B) {
		calculator.ReverseSearch(passiveIDs, []uint32{5815}, data.GloriousVanity, data.Xibaqua, nil)
		b.ResetTimer()
		for range b.N {
			calculator.ReverseSearch(passiveIDs, []uint32{5815}, data.GloriousVanity, data.Xibaqua, nil)
		}
	})

	b.Run("uncached", func(b *testing.B) {
		for range b.N {
			calculator.ClearCache()
			calculator.ReverseSearch(passiveIDs, []uint32{5815}, data.GloriousVanity, data.Xibaqua, nil)
		}
	})
}

func BenchmarkElegantHubris(b *testing.B) {
	b.ReportAllocs()
	b.Run("cached", func(b *testing.B) {
		calculator.ReverseSearch(passiveIDs, []uint32{25}, data.ElegantHubris, data.Cadiro, nil)
		b.ResetTimer()
		for range b.N {
			calculator.ReverseSearch(passiveIDs, []uint32{25}, data.ElegantHubris, data.Cadiro, nil)
		}
	})

	b.Run("uncached", func(b *testing.B) {
		for range b.N {
			calculator.ClearCache()
			calculator.ReverseSearch(passiveIDs, []uint32{25}, data.ElegantHubris, data.Cadiro, nil)
		}
	})
}
