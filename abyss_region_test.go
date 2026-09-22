package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/Vilsol/timeless-jewels/calculator"
	"github.com/Vilsol/timeless-jewels/data"
)

const abyssRegionFixture = "testdata/abyss/abyss_regions.txt.gz"

// TestAbyssRegionMatchesPathOfBuilding grades AbyssConqueredNodes against a sample of the ordered
// node lists in PoB's ABYS LUTs: every socket at every 395th seed from 100 to 8000 and at 4242, plus
// the first record of every region size, which covers the socket drawn twice (59) and the six
// sockets cut short by cluster-jewel scaffolding (44 to 48). Node set and order must be exact, for
// all four jewels. Regenerate with TestAbyssRegionMatchesPathOfBuildingExhaustive.
func TestAbyssRegionMatchesPathOfBuilding(t *testing.T) {
	rows := readAbyssRegionFixture(t)
	if len(rows) != 477 {
		t.Fatalf("expected 477 records in %s, got %d", abyssRegionFixture, len(rows))
	}
	for _, r := range rows {
		for jewel := data.AbyssTecrod; jewel <= data.AbyssAmanamu; jewel++ {
			got, ok := calculator.AbyssConqueredNodes(r.socket, r.seed, jewel)
			if !ok || !slices.Equal(got, r.nodes) {
				t.Errorf("%s socket %d seed %d: got %v (ok %v), PoB %v", jewel, r.socket, r.seed, got, ok, r.nodes)
			}
		}
	}
}

// TestAbyssRegionWrongSeedControl shows the grader can tell regions apart: the same walk one seed
// off must disagree with PoB almost everywhere, or a match above means nothing.
func TestAbyssRegionWrongSeedControl(t *testing.T) {
	rows := readAbyssRegionFixture(t)
	matched := 0
	for _, r := range rows {
		got, _ := calculator.AbyssConqueredNodes(r.socket, r.seed%7901+100, data.AbyssTecrod)
		if slices.Equal(got, r.nodes) {
			matched++
		}
	}
	if matched > 0 {
		t.Fatalf("a wrong seed reproduced %d of %d records; the fixture cannot discriminate", matched, len(rows))
	}
}

// TestAbyssRegionRejectsWhatHasNoRecord pins the false cases: a socket PoB has no block for, a
// seed out of range, and a jewel that is not socket-keyed.
func TestAbyssRegionRejectsWhatHasNoRecord(t *testing.T) {
	for _, c := range []struct {
		name   string
		socket uint32
		seed   uint32
		jewel  data.JewelType
	}{
		{"cluster sub-socket", 2311, 4242, data.AbyssTecrod},
		{"not a socket", 27140, 4242, data.AbyssTecrod},
		{"seed below range", 26725, 99, data.AbyssTecrod},
		{"seed above range", 26725, 8001, data.AbyssTecrod},
		{"legion jewel", 26725, 4242, data.GloriousVanity},
		{"Zorath", 26725, 4242, data.AbyssZorath},
	} {
		if nodes, ok := calculator.AbyssConqueredNodes(c.socket, c.seed, c.jewel); ok {
			t.Errorf("%s: got %v, want no region", c.name, nodes)
		}
	}
}

// TestAbyssRegionMatchesPathOfBuildingExhaustive grades every record in the four ABYS LUTs of the
// pinned PoB checkout named by TIMELESS_JEWELS_POB (165,921 per jewel), node set and order exact.
// Set TIMELESS_JEWELS_WRITE_FIXTURE=1 to rewrite abyss_regions.txt.gz from the same records.
func TestAbyssRegionMatchesPathOfBuildingExhaustive(t *testing.T) {
	pob := os.Getenv("TIMELESS_JEWELS_POB")
	if pob == "" {
		t.Skip("exhaustive grade NOT run: set TIMELESS_JEWELS_POB to a Path of Building checkout")
	}

	var reference *abyssLUT
	for jewel := data.AbyssTecrod; jewel <= data.AbyssAmanamu; jewel++ {
		lut := loadAbyssLUT(t, pob, jewel)
		if len(lut.sockets) != 21 || lut.seedMin != 100 || lut.seedMax != 8000 || lut.seedInc != 1 || lut.size != 60 {
			t.Fatalf("%s: unexpected LUT shape %d sockets, seeds %d..%d step %d, size %d",
				jewel, len(lut.sockets), lut.seedMin, lut.seedMax, lut.seedInc, lut.size)
		}

		graded, bad := 0, 0
		for _, socket := range lut.sockets {
			for i, want := range lut.records[socket] {
				seed := lut.seedMin + uint32(i)*lut.seedInc
				got, ok := calculator.AbyssConqueredNodes(socket, seed, jewel)
				graded++
				if !ok || !slices.Equal(got, want) {
					bad++
					if bad <= 10 {
						t.Errorf("%s socket %d seed %d: got %v (ok %v), PoB %v", jewel, socket, seed, got, ok, want)
					}
				}
			}
		}
		t.Logf("%s: %d of %d records exact", jewel, graded-bad, graded)
		if reference == nil {
			reference = lut
		}
	}

	if os.Getenv("TIMELESS_JEWELS_WRITE_FIXTURE") != "" {
		writeAbyssRegionFixture(t, reference)
	}
}

type abyssRegionRow struct {
	socket, seed uint32
	nodes        []uint32
}

func readAbyssRegionFixture(t *testing.T) []abyssRegionRow {
	t.Helper()
	f, err := os.Open(abyssRegionFixture)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		t.Fatal(err)
	}

	var rows []abyssRegionRow
	sc := bufio.NewScanner(gz)
	sc.Buffer(make([]byte, 1<<16), 1<<20)
	for sc.Scan() {
		fields := strings.Fields(sc.Text())
		if len(fields) != 3 {
			t.Fatalf("bad fixture line %q", sc.Text())
		}
		row := abyssRegionRow{socket: mustUint32(t, fields[0]), seed: mustUint32(t, fields[1])}
		for _, id := range strings.Split(fields[2], ",") {
			row.nodes = append(row.nodes, mustUint32(t, id))
		}
		rows = append(rows, row)
	}
	if err := sc.Err(); err != nil {
		t.Fatal(err)
	}
	return rows
}

func writeAbyssRegionFixture(t *testing.T, lut *abyssLUT) {
	t.Helper()
	pick := func(i int) bool { return i%395 == 0 || uint32(i)+lut.seedMin == 4242 }

	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	seenSize := map[int]bool{}
	for _, socket := range lut.sockets {
		for i, nodes := range lut.records[socket] {
			if !pick(i) && seenSize[len(nodes)] {
				continue
			}
			seenSize[len(nodes)] = true
			ids := make([]string, len(nodes))
			for k, n := range nodes {
				ids[k] = strconv.Itoa(int(n))
			}
			fmt.Fprintf(gz, "%d %d %s\n", socket, lut.seedMin+uint32(i)*lut.seedInc, strings.Join(ids, ","))
		}
	}
	if err := gz.Close(); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(abyssRegionFixture, buf.Bytes(), 0o644); err != nil {
		t.Fatal(err)
	}
	t.Logf("wrote %s", abyssRegionFixture)
}

// abyssLUT is PoB's ABYS format, read the way DataAbyssJewelLookUpTableHelper.lua reads it, but
// keeping each record's node order and dropping the modifications.
type abyssLUT struct {
	seedMin, seedMax, seedInc uint32
	size                      int
	sockets                   []uint32
	records                   map[uint32][][]uint32
}

func loadAbyssLUT(t *testing.T, pob string, jewel data.JewelType) *abyssLUT {
	t.Helper()
	name := strings.ReplaceAll(jewel.String(), " ", "")
	parts, err := filepath.Glob(filepath.Join(pob, "src/Data/TimelessJewelData", name+".zip.part*"))
	if err != nil || len(parts) == 0 {
		t.Fatalf("%s: no LUT parts under %s (err %v)", jewel, pob, err)
	}
	slices.SortFunc(parts, func(a, b string) int {
		return mustAtoi(t, a[strings.LastIndex(a, "part")+4:]) - mustAtoi(t, b[strings.LastIndex(b, "part")+4:])
	})
	var compressed bytes.Buffer
	for _, p := range parts {
		b, err := os.ReadFile(p)
		if err != nil {
			t.Fatal(err)
		}
		compressed.Write(b)
	}
	zr, err := zlib.NewReader(&compressed)
	if err != nil {
		t.Fatal(err)
	}
	blob, err := io.ReadAll(zr)
	if err != nil {
		t.Fatal(err)
	}

	o := 0
	u8 := func() int { o++; return int(blob[o-1]) }
	u16 := func() uint32 { o += 2; return uint32(binary.LittleEndian.Uint16(blob[o-2:])) }
	if string(blob[:4]) != "ABYS" || blob[4] != 1 || data.JewelType(blob[5]) != jewel {
		t.Fatalf("%s: bad header %q v%d type %d", jewel, blob[:4], blob[4], blob[5])
	}
	o = 6
	lut := &abyssLUT{seedMin: u16(), seedMax: u16(), seedInc: u16(), records: map[uint32][][]uint32{}}
	sockets := u8()
	lut.size = u8()
	for range sockets {
		lut.sockets = append(lut.sockets, u16())
	}
	seeds := int((lut.seedMax-lut.seedMin)/lut.seedInc) + 1
	for _, socket := range lut.sockets {
		perSeed := make([][]uint32, seeds)
		for i := range perSeed {
			n := u8()
			nodes := make([]uint32, n)
			for k := range nodes {
				nodes[k] = u16()
				for range u8() { // skip the modification
					o += 3 + int(blob[o+2])*2
				}
			}
			perSeed[i] = nodes
		}
		lut.records[socket] = perSeed
	}
	if o != len(blob) {
		t.Fatalf("%s: parsed %d of %d bytes", jewel, o, len(blob))
	}
	return lut
}

func mustUint32(t *testing.T, s string) uint32 {
	t.Helper()
	v, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		t.Fatal(err)
	}
	return uint32(v)
}

func mustAtoi(t *testing.T, s string) int {
	t.Helper()
	v, err := strconv.Atoi(s)
	if err != nil {
		t.Fatal(err)
	}
	return v
}
