/* eslint-disable */
export let calculator;
export let data;

export const initializeCrystalline = () => {
  calculator = {
    Calculate: globalThis["go"]["timeless-jewels"]["calculator"]["Calculate"],
    ReverseSearch: globalThis["go"]["timeless-jewels"]["calculator"]["ReverseSearch"],
  }
  data = {
    GetAlternatePassiveAdditionByIndex: globalThis["go"]["timeless-jewels"]["data"]["GetAlternatePassiveAdditionByIndex"],
    GetAlternatePassiveSkillByIndex: globalThis["go"]["timeless-jewels"]["data"]["GetAlternatePassiveSkillByIndex"],
    GetPassiveSkillByIndex: globalThis["go"]["timeless-jewels"]["data"]["GetPassiveSkillByIndex"],
    GetStatByIndex: globalThis["go"]["timeless-jewels"]["data"]["GetStatByIndex"],
    PassiveSkillAuraStatTranslationsJSON: globalThis["go"]["timeless-jewels"]["data"]["PassiveSkillAuraStatTranslationsJSON"],
    PassiveSkillStatTranslationsJSON: globalThis["go"]["timeless-jewels"]["data"]["PassiveSkillStatTranslationsJSON"],
    PassiveSkills: globalThis["go"]["timeless-jewels"]["data"]["PassiveSkills"],
    PossibleStats: globalThis["go"]["timeless-jewels"]["data"]["PossibleStats"],
    SkillTree: globalThis["go"]["timeless-jewels"]["data"]["SkillTree"],
    StatTranslationsJSON: globalThis["go"]["timeless-jewels"]["data"]["StatTranslationsJSON"],
    TimelessJewelConquerors: globalThis["go"]["timeless-jewels"]["data"]["TimelessJewelConquerors"],
    TimelessJewelSeedRanges: globalThis["go"]["timeless-jewels"]["data"]["TimelessJewelSeedRanges"],
    TimelessJewels: globalThis["go"]["timeless-jewels"]["data"]["TimelessJewels"],
    TreeToPassive: globalThis["go"]["timeless-jewels"]["data"]["TreeToPassive"],
  }
}