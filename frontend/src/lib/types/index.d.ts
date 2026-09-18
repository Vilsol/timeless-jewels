/* eslint-disable */

/// <reference lib="es2018" />
/// <reference lib="dom" />
/// <reference lib="esnext.disposable" />
export declare namespace calculator {
  interface CalculateStats {
    readonly Calls: number;
    readonly TotalNanos: number;
    readonly MaxNanos: number;
  }
  function Calculate(passiveID: number, seed: number, timelessJewelType: data.JewelType, conqueror: data.Conqueror): data.AlternatePassiveSkillInformation;
  function GetCalculateStats(): calculator.CalculateStats;
  function ResetCalculateStats(): void;
  function ReverseSearch(passiveIDs: Array<number> | undefined, statIDs: Array<number> | undefined, timelessJewelType: data.JewelType, conqueror: data.Conqueror, updates: (seed: number) => void): Promise<(Record<number, Record<number, Record<number, number> | undefined> | undefined> | undefined)>;
  function SetCalculateTracking(enabled: boolean): void;
}
export declare namespace data {
  interface AlternatePassiveAddition {
    readonly Index: number;
    readonly ID: string;
    readonly AlternateTreeVersionsKey: number;
    readonly SpawnWeight: number;
    readonly StatsKeys?: Array<number>;
    readonly Stat1Min: number;
    readonly Stat1Max: number;
    readonly Stat2Min: number;
    readonly Stat2Max: number;
    readonly PassiveType?: Array<data.PassiveSkillType>;
  }
  interface AlternatePassiveAdditionInformation {
    readonly AlternatePassiveAddition?: data.AlternatePassiveAddition;
    readonly StatRolls?: Array<number>;
  }
  interface AlternatePassiveSkill {
    readonly Index: number;
    readonly ID: string;
    readonly AlternateTreeVersionsKey: number;
    readonly Name: string;
    readonly PassiveType?: Array<data.PassiveSkillType>;
    readonly StatsKeys?: Array<number>;
    readonly Stat1Min: number;
    readonly Stat1Max: number;
    readonly Stat2Min: number;
    readonly Stat2Max: number;
    readonly Stat3Min: number;
    readonly Stat3Max: number;
    readonly Stat4Min: number;
    readonly Stat4Max: number;
    readonly SpawnWeight: number;
    readonly ConquerorIndex: number;
    readonly RandomMin: number;
    readonly RandomMax: number;
    readonly ConquerorVersion: number;
  }
  interface AlternatePassiveSkillInformation {
    readonly AlternatePassiveSkill?: data.AlternatePassiveSkill;
    readonly StatRolls?: Array<number>;
    readonly AlternatePassiveAdditionInformations?: Array<data.AlternatePassiveAdditionInformation>;
  }
  type Conqueror = "Xibaqua" | "Zerphi" | "Ahuana" | "Doryani" | "Kaom" | "Rakiata" | "Kiloava" | "Akoya" | "Deshret" | "Balbala" | "Asenath" | "Nasima" | "Venarius" | "Maxarius" | "Dominus" | "Avarius" | "Cadiro" | "Victario" | "Chitus" | "Caspiro" | "Abyss" | "Vorana" | "Uhtred" | "Medved";
  const Conqueror: {
    readonly Xibaqua: "Xibaqua";
    readonly Zerphi: "Zerphi";
    readonly Ahuana: "Ahuana";
    readonly Doryani: "Doryani";
    readonly Kaom: "Kaom";
    readonly Rakiata: "Rakiata";
    readonly Kiloava: "Kiloava";
    readonly Akoya: "Akoya";
    readonly Deshret: "Deshret";
    readonly Balbala: "Balbala";
    readonly Asenath: "Asenath";
    readonly Nasima: "Nasima";
    readonly Venarius: "Venarius";
    readonly Maxarius: "Maxarius";
    readonly Dominus: "Dominus";
    readonly Avarius: "Avarius";
    readonly Cadiro: "Cadiro";
    readonly Victario: "Victario";
    readonly Chitus: "Chitus";
    readonly Caspiro: "Caspiro";
    readonly Abyss: "Abyss";
    readonly Vorana: "Vorana";
    readonly Uhtred: "Uhtred";
    readonly Medved: "Medved";
  };
  type JewelType = 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11;
  const JewelType: {
    readonly GloriousVanity: 1;
    readonly LethalPride: 2;
    readonly BrutalRestraint: 3;
    readonly MilitantFaith: 4;
    readonly ElegantHubris: 5;
    readonly HeroicTragedy: 6;
    readonly AbyssTecrod: 7;
    readonly AbyssUlaman: 8;
    readonly AbyssKurgal: 9;
    readonly AbyssAmanamu: 10;
    readonly AbyssZorath: 11;
  };
  interface PassiveSkill {
    readonly Index: number;
    readonly ID: string;
    readonly StatIndices?: Array<number>;
    readonly PassiveSkillGraphID: number;
    readonly Name: string;
    readonly IsKeystone: boolean;
    readonly IsNotable: boolean;
    readonly IsJewelSocket: boolean;
    readonly AscendancyKey?: number;
    readonly DescendancyKey?: number;
  }
  type PassiveSkillType = 0 | 1 | 2 | 3 | 4 | 5;
  const PassiveSkillType: {
    readonly None: 0;
    readonly SmallAttribute: 1;
    readonly SmallNormal: 2;
    readonly Notable: 3;
    readonly KeyStone: 4;
    readonly JewelSocket: 5;
  };
  interface Range {
    readonly Min: number;
    readonly Max: number;
    readonly Special: boolean;
  }
  interface Stat {
    readonly Index: number;
    readonly ID: string;
    readonly Text: string;
    readonly Category?: number;
  }
  interface TimelessJewelConqueror {
    readonly Index: number;
    readonly Version: number;
  }
  function GetAlternatePassiveAdditionByIndex(index: number): (data.AlternatePassiveAddition | undefined);
  function GetAlternatePassiveSkillByIndex(index: number): (data.AlternatePassiveSkill | undefined);
  function GetPassiveSkillByIndex(index: number): (data.PassiveSkill | undefined);
  function GetStatByIndex(index: number): (data.Stat | undefined);
  const PassiveSkillAuraStatTranslationsJSON: string;
  const PassiveSkillStatTranslationsJSON: string;
  const PassiveSkills: Array<data.PassiveSkill | undefined> | undefined;
  const PossibleStats: string;
  const SkillTree: string;
  const StatTranslationsJSON: string;
  const TimelessJewelConquerors: Record<data.JewelType, Record<data.Conqueror, data.TimelessJewelConqueror | undefined> | undefined> | undefined;
  const TimelessJewelSeedRanges: Record<data.JewelType, data.Range> | undefined;
  const TimelessJewels: Record<data.JewelType, string> | undefined;
  const TreeToPassive: Record<number, data.PassiveSkill | undefined> | undefined;
}
export function boot(wasm: string | URL | BufferSource): Promise<{ calculator: typeof calculator; data: typeof data }>;
export const initializeCrystalline: () => void;