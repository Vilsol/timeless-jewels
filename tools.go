//go:build tools

package main

import (
	"compress/gzip"
	"encoding/json"
	"os"

	"github.com/Vilsol/timeless-jewels/calculator"
	"github.com/Vilsol/timeless-jewels/data"
)

// Uses separate steps so finder step has the new data loaded by data package

//go:generate go run tools.go re-zip
//go:generate go run tools.go find

func main() {
	if len(os.Args) < 2 {
		return
	}

	switch os.Args[1] {
	case "re-zip":
		reZip()
	case "find":
		findAll()
	}
}

func reZip() {
	reMarshalZip[[]*data.AlternatePassiveAddition]("alternate_passive_additions.json")
	reMarshalZip[[]*data.AlternatePassiveSkill]("alternate_passive_skills.json")
	reMarshalZip[[]*data.AlternateTreeVersion]("alternate_tree_versions.json")
	reMarshalZip[[]*data.PassiveSkill]("passive_skills.json")
	reMarshalZip[[]*data.Stat]("stats.json")
	reMarshalZip[map[string]interface{}]("SkillTree.json")
	reMarshalZip[[]interface{}]("passive_skill.min.json")
}

func reMarshalZip[T any](name string) {
	in, err := os.ReadFile("./source_data/" + name)
	if err != nil {
		panic(err)
	}

	var blob = new(T)
	if err := json.Unmarshal(in, &blob); err != nil {
		panic(err)
	}

	writeZipped("./data/"+name+".gz", blob)
}

func writeZipped(path string, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	out, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}

	writer := gzip.NewWriter(out)
	if _, err := writer.Write(b); err != nil {
		panic(err)
	}

	if err := writer.Close(); err != nil {
		panic(err)
	}
}

func findAll() {
	foundStats := make(map[data.JewelType]map[uint32]int)
	applicable := data.GetApplicablePassives()

	for jewelType := range data.TimelessJewelConquerors {
		foundStats[jewelType] = make(map[uint32]int)

		var firstConqueror data.Conqueror
		for conqueror := range data.TimelessJewelConquerors[jewelType] {
			firstConqueror = conqueror
			break
		}

		println(jewelType.String())
		for seed := data.TimelessJewelSeedRanges[jewelType].Min; seed <= data.TimelessJewelSeedRanges[jewelType].Max; seed++ {
			if seed%500 == 0 {
				println(seed)
			}

			for _, skill := range applicable {
				if skill.IsKeystone {
					continue
				}

				results := calculator.Calculate(skill.Index, seed, jewelType, firstConqueror)
				if results.AlternatePassiveSkill != nil {
					for _, key := range results.AlternatePassiveSkill.StatsKeys {
						foundStats[jewelType][key]++
					}
				}

				for _, info := range results.AlternatePassiveAdditionInformations {
					if info.AlternatePassiveAddition != nil {
						for _, key := range info.AlternatePassiveAddition.StatsKeys {
							foundStats[jewelType][key]++
						}
					}
				}
			}
		}
	}

	writeZipped("./data/possible_stats.json.gz", foundStats)
}
