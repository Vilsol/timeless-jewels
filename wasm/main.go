//go:build wasm

package main

import (
	"fmt"
	"strconv"
	"syscall/js"

	"timeless-jewels/calculator"
	"timeless-jewels/data"
)

func main() {
	js.Global().Set("Calculate", js.FuncOf(func(_ js.Value, args []js.Value) any {
		passiveID := uint32(args[0].Int())
		seed := uint32(args[1].Int())
		jewelType := data.JewelType(args[2].Int())
		conqueror := data.Conqueror(args[3].String())

		result := calculator.Calculate(passiveID, seed, jewelType, conqueror)

		out := make(map[string]interface{})

		statRolls := make(map[string]interface{})
		for k, v := range result.StatRolls {
			statRolls[strconv.Itoa(int(k))] = int(v)
		}
		out["statRolls"] = statRolls

		if result.AlternatePassiveSkill != nil {
			out["alternatePassiveSkill"] = result.AlternatePassiveSkill.Index
		}

		informations := make([]interface{}, len(result.AlternatePassiveAdditionInformations))
		for i, information := range result.AlternatePassiveAdditionInformations {
			alt := make(map[string]interface{})
			alt["alternatePassiveSkillAddition"] = information.AlternatePassiveAddition.Index

			altStatRolls := make(map[string]interface{})
			for k, v := range information.StatRolls {
				altStatRolls[strconv.Itoa(int(k))] = int(v)
			}
			alt["statRolls"] = altStatRolls

			informations[i] = alt
		}
		out["alternatePassiveAdditionInformations"] = informations

		return js.ValueOf(out)
	}))

	js.Global().Set("GetStatByIndex", js.FuncOf(func(_ js.Value, args []js.Value) any {
		stat := data.GetStatByIndex(uint32(args[0].Int()))

		return js.ValueOf(map[string]interface{}{
			"index": stat.Index,
			"id":    stat.ID,
			"text":  stat.Text,
		})
	}))

	js.Global().Set("GetAlternatePassiveSkillByIndex", js.FuncOf(func(_ js.Value, args []js.Value) any {
		stat := data.GetAlternatePassiveSkillByIndex(uint32(args[0].Int()))

		statsKeys := make([]interface{}, len(stat.StatsKeys))
		for i, key := range stat.StatsKeys {
			statsKeys[i] = key
		}

		return js.ValueOf(map[string]interface{}{
			"index":     stat.Index,
			"id":        stat.ID,
			"name":      stat.Name,
			"statsKeys": statsKeys,
		})
	}))

	js.Global().Set("GetAlternatePassiveAdditionByIndex", js.FuncOf(func(_ js.Value, args []js.Value) any {
		stat := data.GetAlternatePassiveAdditionByIndex(uint32(args[0].Int()))

		statsKeys := make([]interface{}, len(stat.StatsKeys))
		for i, key := range stat.StatsKeys {
			statsKeys[i] = key
		}

		return js.ValueOf(map[string]interface{}{
			"index":     stat.Index,
			"id":        stat.ID,
			"statsKeys": statsKeys,
		})
	}))

	remappedTimelessJewels := map[string]interface{}{
		strconv.Itoa(int(data.GloriousVanity)):  data.GloriousVanity.String(),
		strconv.Itoa(int(data.LethalPride)):     data.LethalPride.String(),
		strconv.Itoa(int(data.BrutalRestraint)): data.BrutalRestraint.String(),
		strconv.Itoa(int(data.MilitantFaith)):   data.MilitantFaith.String(),
		strconv.Itoa(int(data.ElegantHubris)):   data.ElegantHubris.String(),
	}

	js.Global().Set("TimelessJewels", js.ValueOf(remappedTimelessJewels))

	remappedTimelessJewelConquerors := make(map[string]interface{})
	for jewelType, sub := range data.TimelessJewelConquerors {
		conquerors := make([]interface{}, 0)
		for conqueror := range sub {
			conquerors = append(conquerors, string(conqueror))
		}
		remappedTimelessJewelConquerors[strconv.Itoa(int(jewelType))] = conquerors
	}

	js.Global().Set("TimelessJewelConquerors", js.ValueOf(remappedTimelessJewelConquerors))

	remappedTimelessJewelSeedRanges := make(map[string]interface{})
	for jewelType, sub := range data.TimelessJewelSeedRanges {
		remappedTimelessJewelSeedRanges[strconv.Itoa(int(jewelType))] = map[string]interface{}{
			"min": int(sub.Min),
			"max": int(sub.Max),
		}
	}

	js.Global().Set("TimelessJewelSeedRanges", js.ValueOf(remappedTimelessJewelSeedRanges))

	remappedPassives := make(map[string]interface{})
	for _, skill := range data.PassiveSkills {
		if skill.Name != "" {
			remappedPassives[skill.Name+" ("+skill.ID+")"] = skill.Index
		}
	}

	js.Global().Set("PassiveSkills", js.ValueOf(remappedPassives))

	fmt.Println("Calculator Initialized")
	select {}
}
