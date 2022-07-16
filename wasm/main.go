//go:build wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/Vilsol/crystalline"

	"github.com/Vilsol/timeless-jewels/calculator"
	"github.com/Vilsol/timeless-jewels/data"
)

func main() {
	js.Global().Set("Calculate", crystalline.MapOrPanic(calculator.Calculate))

	js.Global().Set("ReverseSearch", crystalline.MapOrPanic(calculator.ReverseSearch))

	js.Global().Set("GetStatByIndex", crystalline.MapOrPanic(data.GetStatByIndex))

	js.Global().Set("GetAlternatePassiveSkillByIndex", crystalline.MapOrPanic(data.GetAlternatePassiveSkillByIndex))

	js.Global().Set("GetAlternatePassiveAdditionByIndex", crystalline.MapOrPanic(data.GetAlternatePassiveAdditionByIndex))

	js.Global().Set("GetPassiveByIndex", crystalline.MapOrPanic(data.GetPassiveSkillByIndex))

	js.Global().Set("TimelessJewels", crystalline.MapOrPanic(map[data.JewelType]string{
		data.GloriousVanity:  data.GloriousVanity.String(),
		data.LethalPride:     data.LethalPride.String(),
		data.BrutalRestraint: data.BrutalRestraint.String(),
		data.MilitantFaith:   data.MilitantFaith.String(),
		data.ElegantHubris:   data.ElegantHubris.String(),
	}))

	js.Global().Set("TimelessJewelConquerors", crystalline.MapOrPanic(data.TimelessJewelConquerors))

	js.Global().Set("TimelessJewelSeedRanges", crystalline.MapOrPanic(data.TimelessJewelSeedRanges))

	js.Global().Set("PassiveSkills", crystalline.MapOrPanic(data.GetApplicablePassives()))

	js.Global().Set("SkillTree", string(data.SkillTreeJSON))

	treeToPassive := make(map[uint32]*data.PassiveSkill)
	for _, skill := range data.PassiveSkills {
		treeToPassive[skill.PassiveSkillGraphID] = skill
	}

	js.Global().Set("TreeToPassive", crystalline.MapOrPanic(treeToPassive))

	js.Global().Set("PassiveTranslations", string(data.PassiveSkillTranslationsJSON))

	js.Global().Set("PossibleStats", string(data.PossibleStatsJSON))

	fmt.Println("Calculator Initialized")
	select {}
}
