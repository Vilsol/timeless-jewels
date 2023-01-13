package main

import (
	"testing"

	"github.com/MarvinJWendt/testza"

	"github.com/Vilsol/timeless-jewels/calculator"
	"github.com/Vilsol/timeless-jewels/data"
)

var passiveIDs = []uint32{2211, 662, 2095, 600, 944, 75, 2556, 1094, 2589, 2505, 458, 2185, 2030, 625, 854, 1175, 2204, 1174, 13, 2184, 1066, 2192, 702, 1212, 456, 1068, 2547, 457, 1164, 698, 852, 1210, 516, 1245, 60, 447, 2096, 853, 451, 61, 2093, 1211, 134, 2212, 579, 440, 474, 2183, 2408, 2206, 518, 628, 624, 2340, 701, 2031, 2491, 2519, 9, 2205, 11}

func TestReverseGloriousVanity(t *testing.T) {
	statIDs := []uint32{5815}
	result := calculator.ReverseSearch(passiveIDs, statIDs, data.GloriousVanity, data.Xibaqua, nil)
	testza.AssertLen(t, result, 1785)
	testza.AssertLen(t, result[1001], 1)
	testza.AssertEqual(t, uint32(2), result[1001][13][statIDs[0]])
}

func TestReverseElegantHubris(t *testing.T) {
	statIDs := []uint32{25}
	result := calculator.ReverseSearch(passiveIDs, statIDs, data.ElegantHubris, data.Cadiro, nil)
	testza.AssertLen(t, result, 2171)
	testza.AssertLen(t, result[57820], 2)
	testza.AssertEqual(t, uint32(80), result[57820][1068][statIDs[0]])
}

func BenchmarkGloriousVanity(b *testing.B) {
	b.ReportAllocs()
	b.Run("cached", func(b *testing.B) {
		calculator.ReverseSearch(passiveIDs, []uint32{5815}, data.GloriousVanity, data.Xibaqua, nil)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			calculator.ReverseSearch(passiveIDs, []uint32{5815}, data.GloriousVanity, data.Xibaqua, nil)
		}
	})

	b.Run("uncached", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
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
		for i := 0; i < b.N; i++ {
			calculator.ReverseSearch(passiveIDs, []uint32{25}, data.ElegantHubris, data.Cadiro, nil)
		}
	})

	b.Run("uncached", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			calculator.ClearCache()
			calculator.ReverseSearch(passiveIDs, []uint32{25}, data.ElegantHubris, data.Cadiro, nil)
		}
	})
}
