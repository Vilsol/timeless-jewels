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

// Path of Building numbers a replacement by alternate skill index + #legionAdditions.
const pobTimelessJewelAdditions = 337

type abyssComponent struct {
	Type  uint32
	ID    uint32
	Rolls []int32
}

// TestAbyssMatchesPathOfBuilding grades Calculate against Path of Building's own Abyss LUT
// (testdata/abyss, regenerate with dump_fixture.lua against the pinned PoB) at seeds 100, 4242
// and 8000: every Zorath block, every keystone (no block, so no modification), and every node
// the four socket-keyed jewels touch. Component type, id and every roll must be exact.
//
// Rolls are compared signed. PoB's getAbyssJewelComponentRoll flips a negative roll only for
// display, where the stat's minimum is non-negative; the stored roll is what the RNG produced.
func TestAbyssMatchesPathOfBuilding(t *testing.T) {
	rows := gradeAbyssFixture(t, "testdata/abyss/abyss_lut.txt.gz")
	if len(rows) != 15 {
		t.Fatalf("expected 5 jewel types x 3 seeds in the fixture, got %d groups", len(rows))
	}
}

// TestAbyssTailDrawsMatchPathOfBuilding grades every record in PoB's full Abyss LUTs where a
// plain modulo draw disagrees with PoB: each one lands a bounded draw in the biased tail, so
// only a rejection-sampled draw reproduces them. Keys from grading all 52,852,691 records,
// expected values from dump_fixture.lua --keys.
func TestAbyssTailDrawsMatchPathOfBuilding(t *testing.T) {
	rows := gradeAbyssFixture(t, "testdata/abyss/abyss_tail_draws.txt.gz")
	total := 0
	for _, n := range rows {
		total += n
	}
	if total != 1153 {
		t.Fatalf("expected 1153 tail records, got %d", total)
	}
}

type abyssGroup struct {
	jewel data.JewelType
	seed  uint32
}

func gradeAbyssFixture(t *testing.T, path string) map[abyssGroup]int {
	t.Helper()

	byGraphID := make(map[uint32]*data.PassiveSkill, len(data.PassiveSkills))
	for _, p := range data.PassiveSkills {
		if _, ok := byGraphID[p.PassiveSkillGraphID]; !ok {
			byGraphID[p.PassiveSkillGraphID] = p
		}
	}

	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	zr, err := gzip.NewReader(f)
	if err != nil {
		t.Fatal(err)
	}

	rows := make(map[abyssGroup]int)
	failures := make(map[abyssGroup]int)

	sc := bufio.NewScanner(zr)
	for sc.Scan() {
		fields := strings.Fields(sc.Text())
		if fields[0] == "Z" {
			fields = append([]string{"Z", strconv.Itoa(int(data.AbyssZorath))}, fields[1:]...)
		}
		jewel, _ := strconv.Atoi(fields[1])
		seed, _ := strconv.Atoi(fields[2])
		node, _ := strconv.Atoi(fields[3])
		want := parseAbyssComponents(t, fields[4])

		g := abyssGroup{data.JewelType(jewel), uint32(seed)}
		rows[g]++

		passive := byGraphID[uint32(node)]
		if passive == nil {
			t.Fatalf("node %d is not a passive skill", node)
		}
		got := abyssComponents(calculator.Calculate(passive.Index, g.seed, g.jewel, data.Abyss))
		if fmt.Sprint(got) != fmt.Sprint(want) {
			failures[g]++
			if failures[g] <= 5 {
				t.Errorf("%s seed %d node %d (%s): got %v, PoB %v", g.jewel, g.seed, node, passive.Name, got, want)
			}
		}
	}
	if err := sc.Err(); err != nil {
		t.Fatal(err)
	}

	total := 0
	for g, n := range rows {
		if failures[g] > 0 {
			total += failures[g]
			t.Errorf("%s seed %d: %d of %d nodes differ from PoB", g.jewel, g.seed, failures[g], n)
		}
	}
	if total > 0 {
		t.Errorf("%s: %d records differ from PoB", path, total)
	}
	return rows
}

func parseAbyssComponents(t *testing.T, s string) []abyssComponent {
	out := []abyssComponent{}
	if s == "-" {
		return out
	}
	for _, part := range strings.Split(s, ";") {
		f := strings.SplitN(part, ":", 3)
		typ, err1 := strconv.Atoi(f[0])
		id, err2 := strconv.Atoi(f[1])
		if err1 != nil || err2 != nil || len(f) != 3 {
			t.Fatalf("bad component %q", part)
		}
		c := abyssComponent{Type: uint32(typ), ID: uint32(id), Rolls: []int32{}}
		if f[2] != "" {
			for _, r := range strings.Split(f[2], ",") {
				v, err := strconv.Atoi(r)
				if err != nil {
					t.Fatalf("bad roll %q", part)
				}
				c.Rolls = append(c.Rolls, int32(v))
			}
		}
		out = append(out, c)
	}
	return out
}

func abyssComponents(info data.AlternatePassiveSkillInformation) []abyssComponent {
	out := []abyssComponent{}
	if skill := info.AlternatePassiveSkill; skill != nil {
		n := min(len(skill.StatsKeys), 4)
		out = append(out, abyssComponent{Type: 1, ID: skill.Index + pobTimelessJewelAdditions, Rolls: append([]int32{}, info.StatRolls[:n]...)})
	}
	for _, a := range info.AlternatePassiveAdditionInformations {
		n := min(len(a.AlternatePassiveAddition.StatsKeys), 2)
		out = append(out, abyssComponent{Type: 2, ID: a.AlternatePassiveAddition.Index, Rolls: append([]int32{}, a.StatRolls[:n]...)})
	}
	return out
}
