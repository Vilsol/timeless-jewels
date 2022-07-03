package data

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"encoding/json"
	"io"
)

//go:embed alternate_passive_additions.json.gz
var alternatePassiveAdditionsGz []byte
var AlternatePassiveAdditions []*AlternatePassiveAddition

//go:embed alternate_passive_skills.json.gz
var alternatePassiveSkillsGz []byte
var AlternatePassiveSkills []*AlternatePassiveSkill

//go:embed alternate_tree_versions.json.gz
var alternateTreeVersionsGz []byte
var AlternateTreeVersions []*AlternateTreeVersion

//go:embed passive_skills.json.gz
var passiveSkillsGz []byte
var PassiveSkills []*PassiveSkill

//go:embed stats.json.gz
var statsGz []byte
var Stats []*Stat

func init() {
	AlternatePassiveAdditions = unzipJSONTo[AlternatePassiveAddition](alternatePassiveAdditionsGz)
	AlternatePassiveSkills = unzipJSONTo[AlternatePassiveSkill](alternatePassiveSkillsGz)
	AlternateTreeVersions = unzipJSONTo[AlternateTreeVersion](alternateTreeVersionsGz)
	PassiveSkills = unzipJSONTo[PassiveSkill](passiveSkillsGz)
	Stats = unzipJSONTo[Stat](statsGz)
}

func unzipJSONTo[T any](data []byte) []*T {
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	all, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	var out = make([]*T, 0)
	if err := json.Unmarshal(all, &out); err != nil {
		panic(err)
	}

	return out
}
