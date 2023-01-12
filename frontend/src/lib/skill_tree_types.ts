export interface Asset {
  '0.1246': string;
  '0.2109': string;
  '0.2972': string;
  '0.3835': string;
}

export interface FlavourTextRect {
  x: number;
  y: number;
  width: number;
  height: number;
}

export interface Ascendancy {
  id: string;
  name: string;
  flavourText?: string;
  flavourTextColour?: string;
  flavourTextRect?: FlavourTextRect;
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

export interface Constants {
  classes: Classes;
  characterAttributes: CharacterAttributes;
  PSSCentreInnerRadius: number;
  skillsPerOrbit: number[];
  orbitRadii: number[];
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

export interface Node {
  skill?: number;
  name?: string;
  icon?: string;
  isNotable?: boolean;
  recipe?: string[];
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

export interface Points {
  totalPoints: number;
  ascendancyPoints: number;
}

export interface Coord {
  x: number;
  y: number;
  w: number;
  h: number;
}

export interface Sprite {
  filename: string;
  coords: { [key: string]: Coord };
}

export interface Sprites {
  normalActive: { [key: string]: Sprite };
  notableActive: { [key: string]: Sprite };
  keystoneActive: { [key: string]: Sprite };
  normalInactive: { [key: string]: Sprite };
  notableInactive: { [key: string]: Sprite };
  keystoneInactive: { [key: string]: Sprite };
  mastery: { [key: string]: Sprite };
  masteryConnected: { [key: string]: Sprite };
  masteryActiveSelected: { [key: string]: Sprite };
  masteryInactive: { [key: string]: Sprite };
  masteryActiveEffect: { [key: string]: Sprite };
  groupBackground: { [key: string]: Sprite };
  frame: { [key: string]: Sprite };
}

export interface RenderParams {
  context: CanvasRenderingContext2D;
  width: number;
  height: number;
}

export type RenderFunc = (params: RenderParams) => void;

export interface Condition {
  min?: number;
  max?: number;
  negated?: boolean;
}

export interface TranslationData {
  condition: Condition[];
  format: string[];
  index_handlers: Array<string[]>;
  string: string;
}

export interface Translation {
  English: TranslationData[];
  ids: string[];
  hidden?: boolean;
}

export interface Class {
  name: string;
  base_str: number;
  base_dex: number;
  base_int: number;
  ascendancies: Ascendancy[];
}

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
  sprites: Sprites;
  imageZoomLevels: number[];
  points: Points;
}
