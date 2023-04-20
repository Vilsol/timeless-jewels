#!/usr/bin/env bash

if [ "$#" -ne 2 ]; then
    echo "usage: $0 <tree_version> <game_version>"
    echo ""
    echo "example: $0 3.21.0 3.21"
    exit 1
fi

set -ex

curl -L "https://raw.githubusercontent.com/grindinggear/skilltree-export/$1/data.json" | gzip > ./data/SkillTree.json.gz

curl -L "https://go-pob-data.pages.dev/data/$2/raw/AlternatePassiveAdditions.json.gz" > ./data/alternate_passive_additions.json.gz
curl -L "https://go-pob-data.pages.dev/data/$2/raw/AlternatePassiveSkills.json.gz" > ./data/alternate_passive_skills.json.gz
curl -L "https://go-pob-data.pages.dev/data/$2/raw/AlternateTreeVersions.json.gz" > ./data/alternate_tree_versions.json.gz
curl -L "https://go-pob-data.pages.dev/data/$2/raw/PassiveSkills.json.gz" > ./data/passive_skills.json.gz
curl -L "https://go-pob-data.pages.dev/data/$2/raw/Stats.json.gz" > ./data/stats.json.gz

curl -L "https://go-pob-data.pages.dev/data/$2/stat_translations/en/stat_descriptions.json.gz" > ./data/stat_descriptions.json.gz
curl -L "https://go-pob-data.pages.dev/data/$2/stat_translations/en/passive_skill_stat_descriptions.json.gz" > ./data/passive_skill_stat_descriptions.json.gz
curl -L "https://go-pob-data.pages.dev/data/$2/stat_translations/en/passive_skill_aura_stat_descriptions.json.gz" > ./data/passive_skill_aura_stat_descriptions.json.gz

go generate -tags tools -x ./...