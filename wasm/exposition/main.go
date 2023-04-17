package exposition

import (
	"github.com/Vilsol/crystalline"

	"github.com/Vilsol/timeless-jewels/calculator"
	"github.com/Vilsol/timeless-jewels/data"
)

func Expose() *crystalline.Exposer {
	e := crystalline.NewExposer("timeless-jewels")

	e.ExposeFuncOrPanic(calculator.Calculate)
	e.ExposeFuncOrPanic(calculator.ReverseSearch)
	e.ExposeFuncOrPanic(data.GetStatByIndex)
	e.ExposeFuncOrPanic(data.GetAlternatePassiveSkillByIndex)
	e.ExposeFuncOrPanic(data.GetAlternatePassiveAdditionByIndex)
	e.ExposeFuncOrPanic(data.GetPassiveSkillByIndex)

	e.ExposeOrPanic(map[data.JewelType]string{
		data.GloriousVanity:  data.GloriousVanity.String(),
		data.LethalPride:     data.LethalPride.String(),
		data.BrutalRestraint: data.BrutalRestraint.String(),
		data.MilitantFaith:   data.MilitantFaith.String(),
		data.ElegantHubris:   data.ElegantHubris.String(),
	}, "data", "TimelessJewels")

	e.ExposeOrPanic(data.TimelessJewelConquerors, "data", "TimelessJewelConquerors")
	e.ExposeOrPanic(data.TimelessJewelSeedRanges, "data", "TimelessJewelSeedRanges")
	e.ExposeOrPanic(data.GetApplicablePassives(), "data", "PassiveSkills")
	e.ExposeOrPanic(string(data.SkillTreeJSON), "data", "SkillTree")

	treeToPassive := make(map[uint32]*data.PassiveSkill)
	for _, skill := range data.PassiveSkills {
		treeToPassive[skill.PassiveSkillGraphID] = skill
	}

	e.ExposeOrPanic(treeToPassive, "data", "TreeToPassive")
	e.ExposeOrPanic(string(data.StatTranslationsJSON), "data", "PassiveTranslations")
	e.ExposeOrPanic(string(data.PossibleStatsJSON), "data", "PossibleStats")

	return e
}
