package data

import (
	"bytes"
	"compress/gzip"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strings"
)

//go:embed alternate_passive_additions.json.gz
var alternatePassiveAdditionsGz []byte
var AlternatePassiveAdditions []*AlternatePassiveAddition

var idToAlternatePassiveAddition = make(map[uint32]*AlternatePassiveAddition)
var reverseAlternatePassiveAdditions = make(map[PassiveSkillType]map[uint32][]*AlternatePassiveAddition)

//go:embed alternate_passive_skills.json.gz
var alternatePassiveSkillsGz []byte
var AlternatePassiveSkills []*AlternatePassiveSkill

var idToAlternatePassiveSkill = make(map[uint32]*AlternatePassiveSkill)
var reverseAlternatePassiveSkills = make(map[PassiveSkillType]map[uint32][]*AlternatePassiveSkill)

//go:embed alternate_tree_versions.json.gz
var alternateTreeVersionsGz []byte
var AlternateTreeVersions []*AlternateTreeVersion

var idToAlternateTreeVersion = make(map[uint32]*AlternateTreeVersion)

//go:embed passive_skills.json.gz
var passiveSkillsGz []byte
var PassiveSkills []*PassiveSkill

var idToPassiveSkill = make(map[uint32]*PassiveSkill)

//go:embed stats.json.gz
var statsGz []byte
var Stats []*Stat

var idToStat = make(map[uint32]*Stat)

//go:embed SkillTree.json.gz
var skillTreeGz []byte
var SkillTreeJSON []byte
var SkillTreeData SkillTree

//go:embed trs/*
var trs embed.FS

//go:embed possible_stats.json.gz
var possibleStatsGz []byte
var PossibleStatsJSON []byte

func init() {
	AlternatePassiveAdditions = unzipJSONTo[[]*AlternatePassiveAddition](alternatePassiveAdditionsGz)

	for _, alt := range AlternatePassiveAdditions {
		idToAlternatePassiveAddition[alt.Index] = alt

		for _, skillType := range alt.PassiveType {
			if _, ok := reverseAlternatePassiveAdditions[skillType]; !ok {
				reverseAlternatePassiveAdditions[skillType] = make(map[uint32][]*AlternatePassiveAddition)
			}

			reverseAlternatePassiveAdditions[skillType][alt.AlternateTreeVersionsKey] = append(reverseAlternatePassiveAdditions[skillType][alt.AlternateTreeVersionsKey], alt)
		}
	}

	AlternatePassiveSkills = unzipJSONTo[[]*AlternatePassiveSkill](alternatePassiveSkillsGz)

	for _, alt := range AlternatePassiveSkills {
		idToAlternatePassiveSkill[alt.Index] = alt

		for _, skillType := range alt.PassiveType {
			if _, ok := reverseAlternatePassiveSkills[skillType]; !ok {
				reverseAlternatePassiveSkills[skillType] = make(map[uint32][]*AlternatePassiveSkill)
			}

			reverseAlternatePassiveSkills[skillType][alt.AlternateTreeVersionsKey] = append(reverseAlternatePassiveSkills[skillType][alt.AlternateTreeVersionsKey], alt)
		}
	}

	AlternateTreeVersions = unzipJSONTo[[]*AlternateTreeVersion](alternateTreeVersionsGz)

	for _, alt := range AlternateTreeVersions {
		idToAlternateTreeVersion[alt.Index] = alt
	}

	PassiveSkills = unzipJSONTo[[]*PassiveSkill](passiveSkillsGz)

	for _, skill := range PassiveSkills {
		idToPassiveSkill[skill.Index] = skill
	}

	Stats = unzipJSONTo[[]*Stat](statsGz)

	for _, stat := range Stats {
		idToStat[stat.Index] = stat
	}

	SkillTreeData = unzipJSONTo[SkillTree](skillTreeGz)

	var err error
	SkillTreeJSON, err = json.Marshal(SkillTreeData)
	if err != nil {
		panic(err)
	}

	PossibleStatsJSON = unzipTo(possibleStatsGz)
}

func getLocale(locale string) string {
	matches := regexp.MustCompile(`(?mi)en|zh`).FindStringSubmatch(locale)
	if len(matches) > 0 {
		return strings.ToLower(matches[0])
	}

	return "en"
}

func StatTranslationsJSON(locale string) string {
	var gz, err = trs.ReadFile(fmt.Sprintf("trs/%s/stat_descriptions.json.gz", getLocale(locale)))
	if err != nil {
		panic(err)
	}

	return string(unzipTo(gz))
}

func PassiveSkillStatTranslationsJSON(locale string) string {
	var gz, err = trs.ReadFile(fmt.Sprintf("trs/%s/passive_skill_stat_descriptions.json.gz", getLocale(locale)))
	if err != nil {
		panic(err)
	}

	return string(unzipTo(gz))
}

func PassiveSkillAuraStatTranslationsJSON(locale string) string {
	var gz, err = trs.ReadFile(fmt.Sprintf("trs/%s/passive_skill_aura_stat_descriptions.json.gz", getLocale(locale)))
	if err != nil {
		panic(err)
	}

	return string(unzipTo(gz))
}

func unzipJSONTo[T any](data []byte) T {
	var out = new(T)
	if err := json.Unmarshal(unzipTo(data), &out); err != nil {
		panic(err)
	}

	return *out
}

func unzipTo(data []byte) []byte {
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}

	all, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	return all
}
