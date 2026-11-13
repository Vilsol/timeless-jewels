#!/usr/bin/env bash

if [ "$#" -ne 2 ]; then
    echo "usage: $0 <tree_version> <game_version>"
    echo ""
    echo "example: $0 3.21.0 3.21"
    exit 1
fi

set -ex

curl -L "https://raw.githubusercontent.com/grindinggear/skilltree-export/$1/data.json" | gzip > ./data/embedded/SkillTree.json.gz

curl -L "https://go-pob-data.pages.dev/data/$2/raw/AlternatePassiveAdditions.json.gz" > ./data/embedded/alternate_passive_additions.json.gz
curl -L "https://go-pob-data.pages.dev/data/$2/raw/AlternatePassiveSkills.json.gz" > ./data/embedded/alternate_passive_skills.json.gz
curl -L "https://go-pob-data.pages.dev/data/$2/raw/AlternateTreeVersions.json.gz" > ./data/embedded/alternate_tree_versions.json.gz
curl -L "https://go-pob-data.pages.dev/data/$2/raw/PassiveSkills.json.gz" > ./data/embedded/passive_skills.json.gz
curl -L "https://go-pob-data.pages.dev/data/$2/raw/Stats.json.gz" > ./data/embedded/stats.json.gz

go run tools.go translations "$2"

go generate -tags tools -x ./...