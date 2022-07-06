//go:build tools

package main

import (
	"compress/gzip"
	"encoding/json"
	"os"

	"timeless-jewels/data"
)

//go:generate go run tools.go

func main() {
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

	b, err := json.Marshal(blob)
	if err != nil {
		panic(err)
	}

	out, err := os.OpenFile("./data/"+name+".gz", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
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
