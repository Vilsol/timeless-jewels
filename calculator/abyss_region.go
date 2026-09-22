// The walk below is ported from abyssRegion and scripts/build-abyss-links.mjs in
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
	"regexp"
	"slices"
	"sync"

	"github.com/Vilsol/timeless-jewels/data"
	"github.com/Vilsol/timeless-jewels/random"
)

// Attempts, not conquered passives: the jewel's passive_tree_abyss_size, which PoB's ABYS header
// also stores.
const abyssRegionSize = 60

const (
	abyssWeightNotable   = 25
	abyssWeightAttribute = 1
	abyssWeightOther     = 5
)

// A small passive whose only stat is +N to one attribute.
var abyssAttributeStat = regexp.MustCompile(`^\+\d+ to (Strength|Dexterity|Intelligence)$`)

type abyssRegionNode struct {
	links       []uint32
	weight      uint32
	conquerable bool
}

type abyssRegionGraph struct {
	nodes   map[uint32]*abyssRegionNode
	sockets map[uint32]bool
}

var (
	abyssRegionMu      sync.Mutex
	abyssRegionTree    uintptr
	abyssRegionCurrent *abyssRegionGraph
)

// AbyssConqueredNodes returns the graph ids of the passives a socket-keyed Abyss jewel
// (data.AbyssTecrod through data.AbyssAmanamu) conquers for a seed, in the order the game draws
// them, which is also PoB's record order. It returns false when the jewel type, seed or socket has
// no region. The region does not depend on which of the four jewels it is.
//
// Reads data.SkillTreeData. Each node's modification is Calculate on that node's passive skill, and
// Calculate takes the passive's Index, not its graph id: map through PassiveSkillGraphID first.
func AbyssConqueredNodes(socketID uint32, seed uint32, jewelType data.JewelType) ([]uint32, bool) {
	if jewelType < data.AbyssTecrod || jewelType > data.AbyssAmanamu {
		return nil, false
	}
	seeds := data.TimelessJewelSeedRanges[jewelType]
	if seed < seeds.Min || seed > seeds.Max {
		return nil, false
	}
	graph := currentAbyssRegionGraph()
	if graph == nil || !graph.sockets[socketID] {
		return nil, false
	}

	var rng random.NumberGenerator
	rng.ResetAbyss(socketID, seed)

	type entry struct {
		id     uint32
		weight uint32
	}

	// The socket is never marked visited: from the second draw it is an ordinary neighbour again,
	// and can come back and be drawn twice.
	visited := make(map[uint32]bool)
	frontier := []entry{{socketID, 1}}
	total := uint32(1)
	conquered := make([]uint32, 0, abyssRegionSize)

	for attempt := 0; attempt < abyssRegionSize && len(frontier) > 0; attempt++ {
		roll := rng.GenerateSingle(total)
		i := 0
		for ; i < len(frontier)-1; i++ {
			if frontier[i].weight > roll {
				break
			}
			roll -= frontier[i].weight
		}
		picked := frontier[i]
		total -= picked.weight
		// In place: the order of the rest is what the next draw indexes.
		frontier = slices.Delete(frontier, i, i+1)

		for _, o := range graph.nodes[picked.id].links {
			if visited[o] {
				continue
			}
			visited[o] = true
			w := graph.nodes[o].weight
			total += w
			frontier = append(frontier, entry{o, w})
		}

		// A socket or the cluster-jewel scaffolding still spends the attempt.
		if graph.nodes[picked.id].conquerable {
			conquered = append(conquered, picked.id)
		}
	}
	return conquered, true
}

// The graph is rebuilt only when data.SkillTreeData.Nodes is replaced.
func currentAbyssRegionGraph() *abyssRegionGraph {
	abyssRegionMu.Lock()
	defer abyssRegionMu.Unlock()

	nodes := data.SkillTreeData.Nodes
	if len(nodes) == 0 {
		return nil
	}
	key := reflect.ValueOf(nodes).Pointer()
	if abyssRegionCurrent == nil || abyssRegionTree != key {
		abyssRegionCurrent = buildAbyssRegionGraph(nodes)
		abyssRegionTree = key
	}
	return abyssRegionCurrent
}

func buildAbyssRegionGraph(treeNodes map[string]data.Node) *abyssRegionGraph {
	bySkill := make(map[uint32]data.Node, len(treeNodes))
	for _, n := range treeNodes {
		if n.Skill != nil {
			bySkill[uint32(*n.Skill)] = n
		}
	}

	// Masteries and class starts are never pushed onto the frontier, so leaving them out of the
	// lists is the same as filtering them during the walk.
	walkable := func(id uint32) bool {
		n, ok := bySkill[id]
		return ok && !isTrue(n.IsMastery) && n.ClassStartIndex == nil
	}

	graph := &abyssRegionGraph{
		nodes:   make(map[uint32]*abyssRegionNode, len(bySkill)),
		sockets: make(map[uint32]bool),
	}
	for id, n := range bySkill {
		if !walkable(id) {
			continue
		}
		node := &abyssRegionNode{
			links:       slices.DeleteFunc(links(n), func(o uint32) bool { return !walkable(o) }),
			weight:      abyssWeightOther,
			conquerable: !isTrue(n.IsJewelSocket) && !abyssScaffolding(n) && n.AscendancyName == nil,
		}
		switch {
		case isTrue(n.IsNotable):
			node.weight = abyssWeightNotable
		case len(n.Stats) == 1 && abyssAttributeStat.MatchString(n.Stats[0]):
			node.weight = abyssWeightAttribute
		}
		graph.nodes[id] = node

		if isTrue(n.IsJewelSocket) && !abyssScaffolding(n) && n.AscendancyName == nil {
			graph.sockets[id] = true
		}
	}
	return graph
}

// abyssScaffolding is the hidden cluster-jewel structure: walkable, never conquerable. An attempt
// that lands on it is spent, which is what cuts the six sockets bordering it short.
func abyssScaffolding(n data.Node) bool {
	return isTrue(n.IsProxy) || (n.ExpansionJewel != nil && n.ExpansionJewel.Parent != nil)
}
