package calculator

import (
	"github.com/Vilsol/timeless-jewels/data"
	"github.com/Vilsol/timeless-jewels/random"
)

type UpdateFunc func(seed uint32)

// calculationCache is keyed by (conqueror, jewelType, realSeed, passiveIdx).
// The per-seed inner map is intentionally small (~60 entries) to keep
// each lookup hot in L1: a single 474k-entry flat map regressed cached-bench
// throughput by ~5× because the bucket array no longer fits in L2.
var calculationCache = make(map[data.Conqueror]map[data.JewelType]map[uint32]map[uint32]data.AlternatePassiveSkillInformation)

func Calculate(passiveID uint32, seed uint32, timelessJewelType data.JewelType, conqueror data.Conqueror) data.AlternatePassiveSkillInformation {
	passiveSkill := data.GetPassiveSkillByIndex(passiveID)

	if !data.IsPassiveSkillValidForAlteration(passiveSkill) {
		return data.AlternatePassiveSkillInformation{}
	}

	alternateTreeVersion := data.GetAlternateTreeVersionIndex(uint32(timelessJewelType))
	if alternateTreeVersion == nil {
		return data.AlternatePassiveSkillInformation{}
	}

	timelessJewelConqueror := data.TimelessJewelConquerors[timelessJewelType][conqueror]
	if timelessJewelConqueror == nil {
		return data.AlternatePassiveSkillInformation{}
	}

	timelessJewel := data.TimelessJewel{
		Seed:                   seed,
		AlternateTreeVersion:   alternateTreeVersion,
		TimelessJewelConqueror: timelessJewelConqueror,
	}

	alternateTreeManager := AlternateTreeManager{
		PassiveSkill:  passiveSkill,
		TimelessJewel: timelessJewel,
	}

	var rng random.NumberGenerator
	if alternateTreeManager.IsPassiveSkillReplaced(&rng) {
		return alternateTreeManager.ReplacePassiveSkill(&rng)
	}

	return data.AlternatePassiveSkillInformation{
		AlternatePassiveAdditionInformations: alternateTreeManager.AugmentPassiveSkill(&rng),
	}
}

func ReverseSearch(passiveIDs []uint32, statIDs []uint32, timelessJewelType data.JewelType, conqueror data.Conqueror, updates UpdateFunc) map[uint32]map[uint32]map[uint32]int32 {
	passiveSkills := make(map[uint32]*data.PassiveSkill)
	for _, id := range passiveIDs {
		skill := data.GetPassiveSkillByIndex(id)
		if data.IsPassiveSkillValidForAlteration(skill) {
			passiveSkills[id] = skill
		}
	}

	alternateTreeVersion := data.GetAlternateTreeVersionIndex(uint32(timelessJewelType))
	if alternateTreeVersion == nil {
		return nil
	}

	timelessJewelConqueror := data.TimelessJewelConquerors[timelessJewelType][conqueror]
	if timelessJewelConqueror == nil {
		return nil
	}

	timelessJewel := data.TimelessJewel{
		AlternateTreeVersion:   alternateTreeVersion,
		TimelessJewelConqueror: timelessJewelConqueror,
	}

	alternateTreeManager := AlternateTreeManager{}

	statMap := make(map[uint32]bool)
	for _, id := range statIDs {
		statMap[id] = true
	}

	conquerorCache, ok := calculationCache[conqueror]
	if !ok {
		conquerorCache = make(map[data.JewelType]map[uint32]map[uint32]data.AlternatePassiveSkillInformation)
		calculationCache[conqueror] = conquerorCache
	}

	jewelCache, ok := conquerorCache[timelessJewelType]
	if !ok {
		jewelCache = make(map[uint32]map[uint32]data.AlternatePassiveSkillInformation)
		conquerorCache[timelessJewelType] = jewelCache
	}

	results := make(map[uint32]map[uint32]map[uint32]int32)

	seedRange := data.TimelessJewelSeedRanges[timelessJewelType]
	seedMin := seedRange.Min
	seedMax := seedRange.Max
	if seedRange.Special {
		seedMin /= 20
		seedMax /= 20
	}

	var rng random.NumberGenerator
	for seed := seedMin; seed <= seedMax; seed++ {
		realSeed := seed
		if seedRange.Special {
			realSeed *= 20
		}

		if seed%10 == 0 && updates != nil {
			updates(realSeed)
		}

		timelessJewel.Seed = realSeed
		alternateTreeManager.TimelessJewel = timelessJewel

		seedCache, ok := jewelCache[realSeed]
		if !ok {
			seedCache = make(map[uint32]data.AlternatePassiveSkillInformation, len(passiveSkills))
			jewelCache[realSeed] = seedCache
		}

		for _, skill := range passiveSkills {
			alternateTreeManager.PassiveSkill = skill

			var result data.AlternatePassiveSkillInformation
			if cacheHit, ok := seedCache[skill.Index]; ok {
				result = cacheHit
			} else {
				if alternateTreeManager.IsPassiveSkillReplaced(&rng) {
					result = alternateTreeManager.ReplacePassiveSkill(&rng)
				} else {
					result = data.AlternatePassiveSkillInformation{
						AlternatePassiveAdditionInformations: alternateTreeManager.AugmentPassiveSkill(&rng),
					}
				}
				seedCache[skill.Index] = result
			}

			// Cache the inner result maps locally per (realSeed, skill) so we don't re-lookup
			// (and re-allocate when missing) on every matching stat key.
			var skillResults map[uint32]int32

			if result.AlternatePassiveSkill != nil {
				for i, key := range result.AlternatePassiveSkill.StatsKeys {
					if _, ok := statMap[key]; ok {
						if skillResults == nil {
							seedResults := results[realSeed]
							if seedResults == nil {
								seedResults = make(map[uint32]map[uint32]int32)
								results[realSeed] = seedResults
							}
							if existing, ok := seedResults[skill.Index]; ok {
								skillResults = existing
							} else {
								skillResults = make(map[uint32]int32)
								seedResults[skill.Index] = skillResults
							}
						}
						skillResults[key] = result.StatRolls[i]
					}
				}
			}

			for _, augment := range result.AlternatePassiveAdditionInformations {
				if augment.AlternatePassiveAddition != nil {
					for i, key := range augment.AlternatePassiveAddition.StatsKeys {
						if _, ok := statMap[key]; ok {
							if skillResults == nil {
								seedResults := results[realSeed]
								if seedResults == nil {
									seedResults = make(map[uint32]map[uint32]int32)
									results[realSeed] = seedResults
								}
								if existing, ok := seedResults[skill.Index]; ok {
									skillResults = existing
								} else {
									skillResults = make(map[uint32]int32)
									seedResults[skill.Index] = skillResults
								}
							}
							skillResults[key] = augment.StatRolls[i]
						}
					}
				}
			}
		}
	}

	return results
}

func ClearCache() {
	calculationCache = make(map[data.Conqueror]map[data.JewelType]map[uint32]map[uint32]data.AlternatePassiveSkillInformation)
}
