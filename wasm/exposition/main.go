package exposition

import (
	"github.com/Vilsol/crystalline/bind"
	"github.com/Vilsol/timeless-jewels/calculator"
	"github.com/Vilsol/timeless-jewels/data"

	// The exposed lookups answer from the shipped tables, so the wasm build installs them.
	// Since the data package stopped embedding its own assets, this import is what carries
	// them — without it every exposed getter returns nil and SkillTreeJSON is empty.
	_ "github.com/Vilsol/timeless-jewels/data/embedded"
)

//crystalline:exports
func Exports(r bind.Registry) {
	// Nothing on the JS side writes to these or calls their methods; they are
	// read once per seed and thrown away, so a live wrapper is all cost.
	r.Type(data.AlternatePassiveSkillInformation{}, bind.Plain())
	r.Type(data.AlternatePassiveSkill{}, bind.Plain())
	r.Type(data.AlternatePassiveAddition{}, bind.Plain())
	r.Type(data.PassiveSkill{}, bind.Plain())
	r.Type(data.Stat{}, bind.Plain())
	r.Type(data.Range{}, bind.Plain())
	r.Type(data.TimelessJewelConqueror{}, bind.Plain())
	r.Type(calculator.CalculateStats{}, bind.Plain())

	r.Func(calculator.Calculate)
	r.Func(calculator.ReverseSearch)
	r.Func(calculator.SetCalculateTracking)
	r.Func(calculator.GetCalculateStats)
	r.Func(calculator.ResetCalculateStats)
	r.Func(data.GetStatByIndex)
	r.Func(data.GetAlternatePassiveSkillByIndex)
	r.Func(data.GetAlternatePassiveAdditionByIndex)
	r.Func(data.GetPassiveSkillByIndex)

	r.Value("TimelessJewels", map[data.JewelType]string{
		data.GloriousVanity:  data.GloriousVanity.String(),
		data.LethalPride:     data.LethalPride.String(),
		data.BrutalRestraint: data.BrutalRestraint.String(),
		data.MilitantFaith:   data.MilitantFaith.String(),
		data.ElegantHubris:   data.ElegantHubris.String(),
		data.HeroicTragedy:   data.HeroicTragedy.String(),
	})

	r.Value("TimelessJewelConquerors", data.TimelessJewelConquerors)
	r.Value("TimelessJewelSeedRanges", data.TimelessJewelSeedRanges)
	r.Value("PassiveSkills", data.GetApplicablePassives())
	r.Value("SkillTree", string(data.SkillTreeJSON), bind.InNamespace("data"))

	treeToPassive := make(map[uint32]*data.PassiveSkill)
	for _, skill := range data.PassiveSkills {
		treeToPassive[skill.PassiveSkillGraphID] = skill
	}

	r.Value("TreeToPassive", treeToPassive)
	r.Value("StatTranslationsJSON", string(data.StatTranslationsJSON), bind.InNamespace("data"))
	r.Value("PassiveSkillStatTranslationsJSON", string(data.PassiveSkillStatTranslationsJSON), bind.InNamespace("data"))
	r.Value("PassiveSkillAuraStatTranslationsJSON", string(data.PassiveSkillAuraStatTranslationsJSON), bind.InNamespace("data"))
	r.Value("PossibleStats", string(data.PossibleStatsJSON), bind.InNamespace("data"))
}
