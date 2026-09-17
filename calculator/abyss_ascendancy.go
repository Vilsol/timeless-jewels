// The walk below is ported from abyssAscendancyNotable and scripts/build-abyss-ascendancy.mjs in
// https://github.com/Parazeya/abyss-jewels, used under the MIT License:
//
// Copyright (c) 2026 poeqol.com
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package calculator

import (
	"reflect"
	"slices"
	"strconv"
	"sync"

	"github.com/Vilsol/timeless-jewels/data"
	"github.com/Vilsol/timeless-jewels/random"
)

// Notables costing this many ascendancy points or more are walked through, never taken, which
// is why Ascendant (every notable costs five) loses nothing.
const abyssAscendancyMaxCost = 4

type abyssAscendancyNode struct {
	ascendancy  string // "" for a class start
	links       []uint32
	collectable bool
}

type abyssAscendancyGraph struct {
	rootEdges  []uint32
	classStart map[string]uint32
	nodes      map[uint32]*abyssAscendancyNode
}

var (
	abyssAscendancyMu      sync.Mutex
	abyssAscendancyTree    uintptr
	abyssAscendancyCurrent *abyssAscendancyGraph
)

// AbyssAscendancyNotable returns the graph id of the ascendancy notable Reclaimed Malevolence
// (data.AbyssZorath) conquers for a seed, or false when it conquers none. ascendancy is the
// tree's internal name ("Raider", not "Warden"). The socket plays no part. Reads
// data.SkillTreeData; the replacement itself is Calculate on that notable's passive skill.
func AbyssAscendancyNotable(seed uint32, ascendancy string) (uint32, bool) {
	graph := currentAbyssAscendancyGraph()
	if graph == nil || ascendancy == "" {
		return 0, false
	}
	bridge, hasBridge := graph.classStart[ascendancy]

	var rng random.NumberGenerator
	rng.ResetSingle(seed)

	const root = ^uint32(0)
	visited := make(map[uint32]bool)
	frontier := []uint32{root}
	for guard := 0; guard < 500 && len(frontier) > 0; guard++ {
		i := rng.GenerateSingle(uint32(len(frontier)))
		g := frontier[i]
		frontier = slices.Delete(frontier, int(i), int(i)+1)

		edges := graph.rootEdges
		if g != root {
			edges = graph.nodes[g].links
		}
		for _, t := range edges {
			if visited[t] {
				continue
			}
			visited[t] = true
			node := graph.nodes[t]
			if node == nil {
				continue
			}
			if hasBridge && t == bridge {
				frontier = append(frontier, t)
				continue
			}
			if node.ascendancy != ascendancy {
				continue
			}
			if node.collectable {
				return t, true
			}
			frontier = append(frontier, t)
		}
	}
	return 0, false
}

// The graph is rebuilt only when data.SkillTreeData.Nodes is replaced.
func currentAbyssAscendancyGraph() *abyssAscendancyGraph {
	abyssAscendancyMu.Lock()
	defer abyssAscendancyMu.Unlock()

	nodes := data.SkillTreeData.Nodes
	if len(nodes) == 0 {
		return nil
	}
	key := reflect.ValueOf(nodes).Pointer()
	if abyssAscendancyCurrent == nil || abyssAscendancyTree != key {
		abyssAscendancyCurrent = buildAbyssAscendancyGraph(nodes)
		abyssAscendancyTree = key
	}
	return abyssAscendancyCurrent
}

func buildAbyssAscendancyGraph(treeNodes map[string]data.Node) *abyssAscendancyGraph {
	bySkill := make(map[uint32]data.Node, len(treeNodes))
	for _, n := range treeNodes {
		if n.Skill != nil {
			bySkill[uint32(*n.Skill)] = n
		}
	}

	// An ordinary ascendancy start hangs off its class start; the alternate ones have no in-edges.
	// The tree's root lists only the class starts and the alternates, so the ordinary starts are
	// appended, in id order.
	var ordinaryStarts []uint32
	for id, n := range bySkill {
		if isTrue(n.IsAscendancyStart) && len(n.In) > 0 {
			ordinaryStarts = append(ordinaryStarts, id)
		}
	}
	slices.Sort(ordinaryStarts)

	graph := &abyssAscendancyGraph{
		rootEdges:  append(parseGraphIDs(treeNodes["root"].Out), ordinaryStarts...),
		classStart: make(map[string]uint32),
		nodes:      make(map[uint32]*abyssAscendancyNode),
	}

	keep := make(map[uint32]bool)
	for _, id := range ordinaryStarts {
		n := bySkill[id]
		bridge := parseGraphIDs(n.In[:1])[0]
		graph.classStart[*n.AscendancyName] = bridge
		keep[bridge] = true
	}
	for id, n := range bySkill {
		if n.AscendancyName != nil {
			keep[id] = true
		}
	}

	// Ascendancy points to allocate a node, counted from its own start.
	cost := make(map[uint32]int)
	for _, start := range ordinaryStarts {
		name := *bySkill[start].AscendancyName
		cost[start] = 0
		queue := []uint32{start}
		for head := 0; head < len(queue); head++ {
			for _, o := range links(bySkill[queue[head]]) {
				t, ok := bySkill[o]
				if _, seen := cost[o]; seen || !ok || t.AscendancyName == nil || *t.AscendancyName != name {
					continue
				}
				cost[o] = cost[queue[head]] + 1
				queue = append(queue, o)
			}
		}
	}

	for id := range keep {
		n := bySkill[id]
		node := &abyssAscendancyNode{links: links(n)}
		if n.AscendancyName != nil {
			node.ascendancy = *n.AscendancyName
		}
		c, costed := cost[id]
		node.collectable = isTrue(n.IsNotable) && !isTrue(n.IsMultipleChoice) && !isTrue(n.IsMultipleChoiceOption) &&
			!isTrue(n.IsAscendancyStart) && costed && c < abyssAscendancyMaxCost
		graph.nodes[id] = node
	}
	return graph
}

// links is in-edges then out-edges, in the tree's own order: the draw indexes the frontier.
func links(n data.Node) []uint32 {
	return append(parseGraphIDs(n.In), parseGraphIDs(n.Out)...)
}

func parseGraphIDs(ids []string) []uint32 {
	out := make([]uint32, 0, len(ids))
	for _, s := range ids {
		v, err := strconv.ParseUint(s, 10, 32)
		if err != nil {
			panic("timeless-jewels: bad skill tree edge " + strconv.Quote(s))
		}
		out = append(out, uint32(v))
	}
	return out
}

func isTrue(b *bool) bool {
	return b != nil && *b
}
