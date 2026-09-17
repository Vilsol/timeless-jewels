package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/Vilsol/timeless-jewels/calculator"
	"github.com/Vilsol/timeless-jewels/data"
)

// In-game recordings of which notable Reclaimed Malevolence rewrote, from
// https://github.com/Parazeya/abyss-jewels observations/ascendancy.json (MIT, poeqol.com).
var abyssAscendancyObservations = []struct {
	seed       uint32
	ascendancy string
	target     string
}{
	{1969, "Deadeye", "Gathering Winds"},
	{1969, "Assassin", "Shadowed Blood"},
	{4501, "Deadeye", "Avidity"},
	{778, "Deadeye", "Gathering Winds"},
	{4906, "Deadeye", "Far Shot"},
	{4906, "Pathfinder", "Nature's Reprisal"},
	{4906, "Elementalist", "Bringer of Ruin"},
	{7760, "Deadeye", "Focal Point"},
	{7247, "Deadeye", "Far Shot"},
	{7247, "Pathfinder", "Nature's Reprisal"},
	{7247, "Assassin", "Unstable Infusion"},
	{7247, "Luminary", "Noble Blood"},
	{7856, "Luminary", "Noble Blood"},
	{7856, "Deadeye", "Ricochet"},
	{7856, "Pathfinder", "Nature's Adrenaline"},
	{7856, "Assassin", "Toxic Delivery"},
	{7856, "Elementalist", "Shaper of Winter"},
	{3418, "Elementalist", "Mastermind of Discord"},
	{3418, "Deadeye", "Avidity"},
	{3418, "Luminary", "Noble Blood"},
	{3418, "Assassin", "Shadowed Blood"},
}

func TestAbyssAscendancyMatchesObservations(t *testing.T) {
	hits := func(shift uint32) int {
		n := 0
		for _, o := range abyssAscendancyObservations {
			if id, ok := calculator.AbyssAscendancyNotable(o.seed+shift, o.ascendancy); ok && id == ascendancyNotableID(t, o.ascendancy, o.target) {
				n++
			}
		}
		return n
	}

	for _, o := range abyssAscendancyObservations {
		want := ascendancyNotableID(t, o.ascendancy, o.target)
		if got, ok := calculator.AbyssAscendancyNotable(o.seed, o.ascendancy); !ok || got != want {
			t.Errorf("seed %d %s: got %d (%v), observed %s (%d)", o.seed, o.ascendancy, got, ok, o.target, want)
		}
	}

	// The same walk on a wrong seed; these are the counts Parazeya's verify.mjs reports.
	for shift, want := range map[uint32]int{137: 10, 911: 7, 4242: 5} {
		if got := hits(shift); got != want {
			t.Errorf("wrong seed +%d: %d of %d match, want %d", shift, got, len(abyssAscendancyObservations), want)
		}
	}
}

// Every Ascendant notable costs five points, so no seed conquers one; Reliquarian has no
// ordinary start and is never entered.
func TestAbyssAscendancyAscendantLosesNothing(t *testing.T) {
	seeds := data.TimelessJewelSeedRanges[data.AbyssZorath]
	for _, ascendancy := range []string{"Ascendant", "Reliquarian"} {
		for seed := seeds.Min; seed <= seeds.Max; seed++ {
			if id, ok := calculator.AbyssAscendancyNotable(seed, ascendancy); ok {
				t.Fatalf("%s seed %d: conquered %d", ascendancy, seed, id)
			}
		}
	}
}

// TestAbyssAscendancyMatchesPathOfBuilding grades the walk against the ASCS section of the
// pinned Path of Building's Zorath LUT (testdata/abyss/dump_ascendancy.lua): every seed for
// Ascendant, Deadeye and Guardian, every 97th seed for the rest.
func TestAbyssAscendancyMatchesPathOfBuilding(t *testing.T) {
	rows, conquered, failures := 0, 0, 0
	for _, fields := range readAbyssFixtureLines(t, "testdata/abyss/abyss_ascendancy.txt.gz") {
		seed, _ := strconv.Atoi(fields[1])
		want := fields[2]
		got := "-"
		if id, ok := calculator.AbyssAscendancyNotable(uint32(seed), fields[0]); ok {
			got = strconv.Itoa(int(id))
		}
		rows++
		if want != "-" {
			conquered++
		}
		if got != want {
			failures++
			if failures <= 10 {
				t.Errorf("%s seed %d: got %s, PoB %s", fields[0], seed, got, want)
			}
		}
	}
	if failures > 0 {
		t.Errorf("%d of %d rows differ from PoB", failures, rows)
	}
	if rows != 25197 || conquered != 17213 {
		t.Fatalf("fixture: %d rows, %d conquering", rows, conquered)
	}
}

// TestAbyssAscendancyEndToEnd is the consumer contract: the walk names the node, Calculate on
// that node's passive skill gives its replacement. Graded against PoB's full modification for
// every conquering row of abyss_ascendancy.txt.gz at a seed abyss_lut.txt.gz carries (100 and
// 8000 for all 19, 4242 for the two sampled at every seed).
func TestAbyssAscendancyEndToEnd(t *testing.T) {
	lut := make(map[string]string)
	for _, fields := range readAbyssFixtureLines(t, "testdata/abyss/abyss_lut.txt.gz") {
		if fields[0] == "Z" {
			lut[fields[1]+" "+fields[2]] = fields[3]
		}
	}

	byGraphID := make(map[uint32]*data.PassiveSkill, len(data.PassiveSkills))
	for _, p := range data.PassiveSkills {
		if _, ok := byGraphID[p.PassiveSkillGraphID]; !ok {
			byGraphID[p.PassiveSkillGraphID] = p
		}
	}

	checked := 0
	for _, fields := range readAbyssFixtureLines(t, "testdata/abyss/abyss_ascendancy.txt.gz") {
		if fields[1] != "100" && fields[1] != "4242" && fields[1] != "8000" {
			continue
		}
		seed, _ := strconv.Atoi(fields[1])
		id, ok := calculator.AbyssAscendancyNotable(uint32(seed), fields[0])
		if !ok {
			continue
		}
		want, found := lut[fmt.Sprintf("%d %d", seed, id)]
		if !found || want == "-" {
			t.Fatalf("%s seed %d: PoB has no modification for node %d", fields[0], seed, id)
		}
		info := calculator.Calculate(byGraphID[id].Index, uint32(seed), data.AbyssZorath, data.Abyss)
		if got := fmt.Sprint(abyssComponents(info)); got != fmt.Sprint(parseAbyssComponents(t, want)) {
			t.Errorf("%s seed %d node %d: got %s, PoB %s", fields[0], seed, id, got, want)
		}
		checked++
	}
	if checked != 40 {
		t.Fatalf("checked %d pairs, want 40", checked)
	}
}

func ascendancyNotableID(t *testing.T, ascendancy, name string) uint32 {
	t.Helper()
	for _, n := range data.SkillTreeData.Nodes {
		if n.Skill != nil && n.AscendancyName != nil && *n.AscendancyName == ascendancy && n.Name != nil && *n.Name == name {
			return uint32(*n.Skill)
		}
	}
	t.Fatalf("no %s notable %q in the tree", ascendancy, name)
	return 0
}

func readAbyssFixtureLines(t *testing.T, path string) [][]string {
	t.Helper()
	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	zr, err := gzip.NewReader(f)
	if err != nil {
		t.Fatal(err)
	}
	var out [][]string
	sc := bufio.NewScanner(zr)
	for sc.Scan() {
		out = append(out, strings.Fields(sc.Text()))
	}
	if err := sc.Err(); err != nil {
		t.Fatal(err)
	}
	return out
}
