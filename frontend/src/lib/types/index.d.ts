/* eslint-disable */
export declare namespace calculator {
  function Calculate(passiveID: number, seed: number, timelessJewelType: number, conqueror: string): data.AlternatePassiveSkillInformation;
  function ReverseSearch(passiveIDs?: Array<number>, statIDs?: Array<number>, timelessJewelType: number, conqueror: string, updates: (arg1: number) => Promise<void>): Promise<(Record<number, Record<number, Record<number, number> | undefined> | undefined> | undefined)>;
}
export declare namespace data {
  interface AlternatePassiveAddition {
    Index: number;
    ID: string;
    AlternateTreeVersionsKey: number;
    SpawnWeight: number;
    StatsKeys?: Array<number>;
    Stat1Min: number;
    Stat1Max: number;
    Stat2Min: number;
    Stat2Max: number;
    PassiveType?: Array<number>;
    GetStatMinMax(min: boolean, index: number): number;
  }
  interface AlternatePassiveAdditionInformation {
    AlternatePassiveAddition?: data.AlternatePassiveAddition;
    StatRolls?: Record<number, number>;
  }
  interface AlternatePassiveSkill {
    Index: number;
    ID: string;
    AlternateTreeVersionsKey: number;
    Name: string;
    PassiveType?: Array<number>;
    StatsKeys?: Array<number>;
    Stat1Min: number;
    Stat1Max: number;
    Stat2Min: number;
    Stat2Max: number;
    Stat3Min: number;
    Stat3Max: number;
    Stat4Min: number;
    Stat4Max: number;
    SpawnWeight: number;
    ConquerorIndex: number;
    RandomMin: number;
    RandomMax: number;
    ConquerorVersion: number;
    GetStatMinMax(min: boolean, index: number): number;
  }
  interface AlternatePassiveSkillInformation {
    AlternatePassiveSkill?: data.AlternatePassiveSkill;
    StatRolls?: Record<number, number>;
    AlternatePassiveAdditionInformations?: Array<data.AlternatePassiveAdditionInformation>;
  }
  interface PassiveSkill {
    Index: number;
    ID: string;
    StatIndices?: Array<number>;
    PassiveSkillGraphID: number;
    Name: string;
    IsKeystone: boolean;
    IsNotable: boolean;
    IsJewelSocket: boolean;
  }
  interface Range {
    Min: number;
    Max: number;
    Special: boolean;
  }
  interface Stat {
    Index: number;
    ID: string;
    Text: string;
    Category?: number;
  }
  interface TimelessJewelConqueror {
    Index: number;
    Version: number;
  }
  function GetAlternatePassiveAdditionByIndex(index: number): (data.AlternatePassiveAddition | undefined);
  function GetAlternatePassiveSkillByIndex(index: number): (data.AlternatePassiveSkill | undefined);
  function GetPassiveSkillByIndex(index: number): (data.PassiveSkill | undefined);
  function GetStatByIndex(index: number): (data.Stat | undefined);
  const PassiveSkills: Array<data.PassiveSkill | undefined> | undefined;
  const PassiveTranslations: string;
  const PossibleStats: string;
  const SkillTree: string;
  const TimelessJewelConquerors: Record<number, Record<string, data.TimelessJewelConqueror | undefined> | undefined> | undefined;
  const TimelessJewelSeedRanges: Record<number, data.Range> | undefined;
  const TimelessJewels: Record<number, string> | undefined;
  const TreeToPassive: Record<number, data.PassiveSkill | undefined> | undefined;
}
export const initializeCrystalline: () => void;