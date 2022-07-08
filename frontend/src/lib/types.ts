// noinspection JSUnusedGlobalSymbols

export interface SkillTreeData {
  tree: string;
  classes: Class[];
  groups: { [key: string]: Group };
  nodes: { [key: string]: Node };
  extraImages: { [key: string]: ExtraImage };
  jewelSlots: number[];
  min_x: number;
  min_y: number;
  max_x: number;
  max_y: number;
  assets: { [key: string]: Asset };
  constants: Constants;
  skillSprites: SkillSprites;
  imageZoomLevels: number[];
  points: Points;
}

export interface Asset {
  '0.1246': string;
  '0.2109': string;
  '0.2972': string;
  '0.3835': string;
}

export interface JewelCircle {
  '1': string;
}

export interface Class {
  name: string;
  base_str: number;
  base_dex: number;
  base_int: number;
  ascendancies: Ascendancy[];
}

export interface Ascendancy {
  id: string;
  name: string;
  flavourText?: string;
  flavourTextColour?: string;
  flavourTextRect?: FlavourTextRect;
}

export interface FlavourTextRect {
  x: number;
  y: number;
  width: number;
  height: number;
}

export interface Constants {
  classes: Classes;
  characterAttributes: CharacterAttributes;
  PSSCentreInnerRadius: number;
  skillsPerOrbit: number[];
  orbitRadii: number[];
}

export interface CharacterAttributes {
  Strength: number;
  Dexterity: number;
  Intelligence: number;
}

export interface Classes {
  StrDexIntClass: number;
  StrClass: number;
  DexClass: number;
  IntClass: number;
  StrDexClass: number;
  StrIntClass: number;
  DexIntClass: number;
}

export interface ExtraImage {
  x: number;
  y: number;
  image: string;
}

export interface Group {
  x: number;
  y: number;
  orbits: number[];
  nodes: string[];
  isProxy?: boolean;
}

export interface Node {
  skill?: number;
  name?: string;
  icon?: string;
  isNotable?: boolean;
  recipe?: Recipe[];
  stats?: string[];
  group?: number;
  orbit?: number;
  orbitIndex?: number;
  out?: string[];
  in?: string[];
  reminderText?: string[];
  isMastery?: boolean;
  inactiveIcon?: string;
  activeIcon?: string;
  activeEffectImage?: string;
  masteryEffects?: MasteryEffect[];
  grantedStrength?: number;
  ascendancyName?: string;
  grantedDexterity?: number;
  isAscendancyStart?: boolean;
  isMultipleChoice?: boolean;
  grantedIntelligence?: number;
  isJewelSocket?: boolean;
  expansionJewel?: ExpansionJewel;
  grantedPassivePoints?: number;
  isKeystone?: boolean;
  flavourText?: string[];
  isProxy?: boolean;
  isMultipleChoiceOption?: boolean;
  isBlighted?: boolean;
  classStartIndex?: number;
}

export interface ExpansionJewel {
  size: number;
  index: number;
  proxy: string;
  parent?: string;
}

export interface MasteryEffect {
  effect: number;
  stats: string[];
  reminderText?: string[];
}

export enum Recipe {
  AmberOil = 'AmberOil',
  AzureOil = 'AzureOil',
  BlackOil = 'BlackOil',
  ClearOil = 'ClearOil',
  CrimsonOil = 'CrimsonOil',
  GoldenOil = 'GoldenOil',
  IndigoOil = 'IndigoOil',
  OpalescentOil = 'OpalescentOil',
  SepiaOil = 'SepiaOil',
  SilverOil = 'SilverOil',
  TealOil = 'TealOil',
  VerdantOil = 'VerdantOil',
  VioletOil = 'VioletOil'
}

export interface Points {
  totalPoints: number;
  ascendancyPoints: number;
}

export interface SkillSprites {
  normalActive: Sprite[];
  notableActive: Sprite[];
  keystoneActive: Sprite[];
  normalInactive: Sprite[];
  notableInactive: Sprite[];
  keystoneInactive: Sprite[];
  mastery: Sprite[];
  masteryConnected: Sprite[];
  masteryActiveSelected: Sprite[];
  masteryInactive: Sprite[];
  masteryActiveEffect: Sprite[];
}

export interface Sprite {
  filename: string;
  coords: { [key: string]: Coord };
}

export interface Coord {
  x: number;
  y: number;
  w: number;
  h: number;
}

export interface RenderParams {
  context: CanvasRenderingContext2D;
  width: number;
  height: number;
}

export type RenderFunc = (params: RenderParams) => void;

export interface Translation {
  English: TranslationData[];
  ids: string[];
  hidden?: boolean;
}

export interface TranslationData {
  condition: Condition[];
  format: Format[];
  index_handlers: Array<IndexHandler[]>;
  string: string;
}

export interface Condition {
  min?: number;
  max?: number;
  negated?: boolean;
}

export enum Format {
  Empty = '#',
  Format = '+#',
  Ignore = 'ignore'
}

export enum IndexHandler {
  AfflictionRewardType = 'affliction_reward_type',
  CanonicalStat = 'canonical_stat',
  DecisecondsToSeconds = 'deciseconds_to_seconds',
  DisplayIndexableSupport = 'display_indexable_support',
  DivideByFifteen0DP = 'divide_by_fifteen_0dp',
  DivideByFifty = 'divide_by_fifty',
  DivideByFive = 'divide_by_five',
  DivideByFour = 'divide_by_four',
  DivideByOneHundred = 'divide_by_one_hundred',
  DivideByOneHundred2DP = 'divide_by_one_hundred_2dp',
  DivideByOneHundred2DPIfRequired = 'divide_by_one_hundred_2dp_if_required',
  DivideByOneHundredAndNegate = 'divide_by_one_hundred_and_negate',
  DivideByOneThousand = 'divide_by_one_thousand',
  DivideBySix = 'divide_by_six',
  DivideByTen0DP = 'divide_by_ten_0dp',
  DivideByTen1DP = 'divide_by_ten_1dp',
  DivideByTen1DPIfRequired = 'divide_by_ten_1dp_if_required',
  DivideByThree = 'divide_by_three',
  DivideByTwelve = 'divide_by_twelve',
  DivideByTwentyThenDouble0DP = 'divide_by_twenty_then_double_0dp',
  DivideByTwo0DP = 'divide_by_two_0dp',
  Double = 'double',
  MetamorphosisRewardDescription = 'metamorphosis_reward_description',
  MillisecondsToSeconds = 'milliseconds_to_seconds',
  MillisecondsToSeconds0DP = 'milliseconds_to_seconds_0dp',
  MillisecondsToSeconds1DP = 'milliseconds_to_seconds_1dp',
  MillisecondsToSeconds2DP = 'milliseconds_to_seconds_2dp',
  MillisecondsToSeconds2DPIfRequired = 'milliseconds_to_seconds_2dp_if_required',
  ModValueToItemClass = 'mod_value_to_item_class',
  MultiplicativeDamageModifier = 'multiplicative_damage_modifier',
  MultiplyByFour = 'multiply_by_four',
  Negate = 'negate',
  NegateAndDouble = 'negate_and_double',
  OldLeechPercent = 'old_leech_percent',
  OldLeechPermyriad = 'old_leech_permyriad',
  PassiveHash = 'passive_hash',
  PerMinuteToPerSecond = 'per_minute_to_per_second',
  PerMinuteToPerSecond0DP = 'per_minute_to_per_second_0dp',
  PerMinuteToPerSecond1DP = 'per_minute_to_per_second_1dp',
  PerMinuteToPerSecond2DP = 'per_minute_to_per_second_2dp',
  PerMinuteToPerSecond2DPIfRequired = 'per_minute_to_per_second_2dp_if_required',
  The30_OfValue = '30%_of_value',
  The60_OfValue = '60%_of_value',
  TimesOnePointFive = 'times_one_point_five',
  TimesTwenty = 'times_twenty',
  TreeExpansionJewelPassive = 'tree_expansion_jewel_passive'
}
