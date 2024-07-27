import type { Translation, Node, SkillTreeData, Group, Sprite, TranslationFile } from './skill_tree_types';
import { data } from './types';
import { type Filter, type Query, filterGroupsToQuery, filtersToFilterGroup } from './utils/trade_utils';
import { chunkArray } from './utils/utils';

export let skillTree: SkillTreeData;

export const drawnGroups: Record<number, Group> = {};
export const drawnNodes: Record<number, Node> = {};

export const inverseSprites: Record<string, Sprite> = {};
export const inverseSpritesActive: Record<string, Sprite> = {};

export const inverseTranslations: Record<string, Translation> = {};

export const passiveToTree: Record<number, number> = {};

export const loadSkillTree = () => {
  skillTree = JSON.parse(data.SkillTree);
  console.log('Loaded skill tree', skillTree);

  Object.keys(skillTree.groups).forEach((groupId) => {
    const group = skillTree.groups[groupId];
    group.nodes.forEach((nodeId) => {
      const node = skillTree.nodes[nodeId];

      // Do not care about proxy passives
      if (node.isProxy) {
        return;
      }

      // Do not care about class starting nodes
      if ('classStartIndex' in node) {
        return;
      }

      // Do not care about cluster jewels
      if (node.expansionJewel) {
        if (node.expansionJewel.parent) {
          return;
        }
      }

      // Do not care about blighted nodes
      if (node.isBlighted) {
        return;
      }

      // Do not care about ascendancies
      if (node.ascendancyName) {
        return;
      }

      drawnGroups[parseInt(groupId)] = group;
      drawnNodes[parseInt(nodeId)] = node;
    });
  });

  Object.keys(skillTree.sprites.keystoneInactive['0.3835'].coords).forEach(
    (c) => (inverseSprites[c] = skillTree.sprites.keystoneInactive['0.3835'])
  );
  Object.keys(skillTree.sprites.notableInactive['0.3835'].coords).forEach(
    (c) => (inverseSprites[c] = skillTree.sprites.notableInactive['0.3835'])
  );
  Object.keys(skillTree.sprites.normalInactive['0.3835'].coords).forEach(
    (c) => (inverseSprites[c] = skillTree.sprites.normalInactive['0.3835'])
  );
  Object.keys(skillTree.sprites.masteryInactive['0.3835'].coords).forEach(
    (c) => (inverseSprites[c] = skillTree.sprites.masteryInactive['0.3835'])
  );

  Object.keys(skillTree.sprites.keystoneActive['0.3835'].coords).forEach(
    (c) => (inverseSpritesActive[c] = skillTree.sprites.keystoneActive['0.3835'])
  );
  Object.keys(skillTree.sprites.notableActive['0.3835'].coords).forEach(
    (c) => (inverseSpritesActive[c] = skillTree.sprites.notableActive['0.3835'])
  );
  Object.keys(skillTree.sprites.normalActive['0.3835'].coords).forEach(
    (c) => (inverseSpritesActive[c] = skillTree.sprites.normalActive['0.3835'])
  );
  Object.keys(skillTree.sprites.masteryInactive['0.3835'].coords).forEach(
    (c) => (inverseSpritesActive[c] = skillTree.sprites.masteryInactive['0.3835'])
  );

  Object.keys(skillTree.sprites.groupBackground['0.3835'].coords).forEach(
    (c) => (inverseSprites[c] = skillTree.sprites.groupBackground['0.3835'])
  );
  Object.keys(skillTree.sprites.frame['0.3835'].coords).forEach(
    (c) => (inverseSprites[c] = skillTree.sprites.frame['0.3835'])
  );

  const translationFiles = [
    data.StatTranslationsJSON,
    data.PassiveSkillStatTranslationsJSON,
    data.PassiveSkillAuraStatTranslationsJSON
  ];

  translationFiles.forEach((f) => {
    const translations: TranslationFile = JSON.parse(f);

    translations.descriptors.forEach((t) => {
      t.ids.forEach((id) => {
        if (!(id in inverseTranslations)) {
          inverseTranslations[id] = t;
        }
      });
    });
  });

  Object.keys(data.TreeToPassive).forEach((k) => {
    passiveToTree[data.TreeToPassive[parseInt(k)].Index] = parseInt(k);
  });
};

const indexHandlers: Record<string, number> = {
  negate: -1,
  times_twenty: 1 / 20,
  canonical_stat: 1,
  per_minute_to_per_second: 60,
  milliseconds_to_seconds: 1000,
  display_indexable_support: 1,
  divide_by_one_hundred: 100,
  milliseconds_to_seconds_2dp_if_required: 1000,
  deciseconds_to_seconds: 10,
  old_leech_percent: 1,
  old_leech_permyriad: 10000,
  times_one_point_five: 1 / 1.5,
  '30%_of_value': 100 / 30,
  divide_by_one_thousand: 1000,
  divide_by_twelve: 12,
  divide_by_six: 6,
  per_minute_to_per_second_2dp_if_required: 60,
  '60%_of_value': 100 / 60,
  double: 1 / 2,
  negate_and_double: 1 / -2,
  multiply_by_four: 1 / 4,
  per_minute_to_per_second_0dp: 60,
  milliseconds_to_seconds_0dp: 1000,
  mod_value_to_item_class: 1,
  milliseconds_to_seconds_2dp: 1000,
  multiplicative_damage_modifier: 1,
  divide_by_one_hundred_2dp: 100,
  per_minute_to_per_second_1dp: 60,
  divide_by_one_hundred_2dp_if_required: 100,
  divide_by_ten_1dp_if_required: 10,
  milliseconds_to_seconds_1dp: 1000,
  divide_by_fifty: 50,
  per_minute_to_per_second_2dp: 60,
  divide_by_ten_0dp: 10,
  divide_by_one_hundred_and_negate: -100,
  tree_expansion_jewel_passive: 1,
  passive_hash: 1,
  divide_by_ten_1dp: 10,
  affliction_reward_type: 1,
  divide_by_five: 5,
  metamorphosis_reward_description: 1,
  divide_by_two_0dp: 2,
  divide_by_fifteen_0dp: 15,
  divide_by_three: 3,
  divide_by_twenty_then_double_0dp: 10,
  divide_by_four: 4
};

export type Point = {
  x: number;
  y: number;
};

export const toCanvasCoords = (x: number, y: number, offsetX: number, offsetY: number, scaling: number): Point => ({
  x: (Math.abs(skillTree.min_x) + x + offsetX) / scaling,
  y: (Math.abs(skillTree.min_y) + y + offsetY) / scaling
});

export const rotateAroundPoint = (center: Point, target: Point, angle: number): Point => {
  const radians = (Math.PI / 180) * angle;
  const cos = Math.cos(radians);
  const sin = Math.sin(radians);
  const nx = cos * (target.x - center.x) + sin * (target.y - center.y) + center.x;
  const ny = cos * (target.y - center.y) - sin * (target.x - center.x) + center.y;
  return {
    x: nx,
    y: ny
  };
};

export const orbit16Angles = [0, 30, 45, 60, 90, 120, 135, 150, 180, 210, 225, 240, 270, 300, 315, 330];
export const orbit40Angles = [
  0, 10, 20, 30, 40, 45, 50, 60, 70, 80, 90, 100, 110, 120, 130, 135, 140, 150, 160, 170, 180, 190, 200, 210, 220, 225,
  230, 240, 250, 260, 270, 280, 290, 300, 310, 315, 320, 330, 340, 350
];

export const orbitAngleAt = (orbit: number, index: number): number => {
  const nodesInOrbit = skillTree.constants.skillsPerOrbit[orbit];
  if (nodesInOrbit == 16) {
    return orbit16Angles[orbit16Angles.length - index] || 0;
  } else if (nodesInOrbit == 40) {
    return orbit40Angles[orbit40Angles.length - index] || 0;
  } else {
    return 360 - (360 / nodesInOrbit) * index;
  }
};

export const calculateNodePos = (node: Node, offsetX: number, offsetY: number, scaling: number): Point => {
  if (node.group === undefined || node.orbit === undefined || node.orbitIndex === undefined) {
    return { x: 0, y: 0 };
  }

  const targetGroup = skillTree.groups[node.group];
  const targetAngle = orbitAngleAt(node.orbit, node.orbitIndex);

  const targetGroupPos = toCanvasCoords(targetGroup.x, targetGroup.y, offsetX, offsetY, scaling);
  const targetNodePos = toCanvasCoords(
    targetGroup.x,
    targetGroup.y - skillTree.constants.orbitRadii[node.orbit],
    offsetX,
    offsetY,
    scaling
  );
  return rotateAroundPoint(targetGroupPos, targetNodePos, targetAngle);
};

export const distance = (p1: Point, p2: Point): number =>
  Math.sqrt(Math.pow(p1.x - p2.x, 2) + Math.pow(p1.y - p2.y, 2));

export const formatStats = (translation: Translation, stat: number): string | undefined => {
  let selectedTranslation = -1;

  for (let i = 0; i < translation.list.length; i++) {
    const t = translation.list[i];

    let matches = true;
    if (t.conditions?.length > 0) {
      const first = t.conditions[0];
      if (first.min !== undefined) {
        if (stat < first.min) {
          matches = false;
        }
      }

      if (first.max !== undefined) {
        if (stat > first.max) {
          matches = false;
        }
      }

      if (first.negated) {
        matches = !matches;
      }
    }

    if (matches) {
      selectedTranslation = i;
      break;
    }
  }

  if (selectedTranslation == -1) {
    return undefined;
  }

  const datum = translation.list[selectedTranslation];

  let finalStat = stat;

  if (datum.index_handlers !== undefined) {
    if (Array.isArray(datum.index_handlers)) {
      if (datum.index_handlers?.length > 0) {
        datum.index_handlers[0].forEach((handler) => {
          finalStat = finalStat / (indexHandlers[handler] || 1);
        });
      }
    } else {
      Object.keys(datum.index_handlers).forEach((handler) => {
        finalStat = finalStat / (indexHandlers[handler] || 1);
      });
    }
  }

  return datum.string
    .replace(/\{0(?::(.*?)d(.*?))\}/, '$1' + finalStat.toString() + '$2')
    .replace(`{0}`, parseFloat(finalStat.toFixed(2)).toString());
};

export const baseJewelRadius = 1800;

export const getAffectedNodes = (socket: Node): Node[] => {
  const result: Node[] = [];

  const socketPos = calculateNodePos(socket, 0, 0, 1);
  for (const node of Object.values(drawnNodes)) {
    const nodePos = calculateNodePos(node, 0, 0, 1);

    if (distance(nodePos, socketPos) < baseJewelRadius) {
      result.push(node);
    }
  }

  return result;
};

type Stat = { Index: number; ID: string; Text: string };

const statCache: Record<number, Stat> = {};
export const getStat = (id: number | string): Stat => {
  const nId = typeof id === 'string' ? parseInt(id) : id;
  if (!(nId in statCache)) {
    statCache[nId] = data.GetStatByIndex(nId);
  }
  return statCache[nId];
};

export interface StatConfig {
  min: number;
  id: number;
  weight: number;
}

export interface ReverseSearchConfig {
  jewel: number;
  conqueror: string;
  nodes: number[];
  stats: StatConfig[];
  minTotalWeight: number;
}

export interface SearchWithSeed {
  seed: number;
  weight: number;
  statCounts: Record<number, number>;
  skills: {
    passive: number;
    stats: { [key: string]: number };
  }[];
}

export interface SearchResults {
  grouped: { [key: number]: SearchWithSeed[] };
  raw: SearchWithSeed[];
}

export const translateStat = (id: number, roll?: number | undefined): string => {
  const stat = getStat(id);
  const translation = inverseTranslations[stat.ID];
  if (roll) {
    return formatStats(translation, roll) || stat.ID;
  }

  let translationText = stat.Text || stat.ID;
  if (translation && translation.list && translation.list.length) {
    translationText = translation.list[0].string;
    translationText = translationText.replace(/\{\d(?::(.*?)d(.*?))\}/, '$1#$2').replace(/\{\d\}/, '#');
  }
  return translationText;
};

const tradeStatNames: { [key: number]: { [key: string]: string } } = {
  1: {
    Ahuana: 'explicit.pseudo_timeless_jewel_ahuana',
    Xibaqua: 'explicit.pseudo_timeless_jewel_xibaqua',
    Doryani: 'explicit.pseudo_timeless_jewel_doryani',
    Zerphi: 'explicit.pseudo_timeless_jewel_zerphi'
  },
  2: {
    Kaom: 'explicit.pseudo_timeless_jewel_kaom',
    Rakiata: 'explicit.pseudo_timeless_jewel_rakiata',
    Kiloava: 'explicit.pseudo_timeless_jewel_kiloava',
    Akoya: 'explicit.pseudo_timeless_jewel_akoya'
  },
  3: {
    Deshret: 'explicit.pseudo_timeless_jewel_deshret',
    Balbala: 'explicit.pseudo_timeless_jewel_balbala',
    Asenath: 'explicit.pseudo_timeless_jewel_asenath',
    Nasima: 'explicit.pseudo_timeless_jewel_nasima'
  },
  4: {
    Venarius: 'explicit.pseudo_timeless_jewel_venarius',
    Maxarius: 'explicit.pseudo_timeless_jewel_maxarius',
    Dominus: 'explicit.pseudo_timeless_jewel_dominus',
    Avarius: 'explicit.pseudo_timeless_jewel_avarius'
  },
  5: {
    Cadiro: 'explicit.pseudo_timeless_jewel_cadiro',
    Victario: 'explicit.pseudo_timeless_jewel_victario',
    Chitus: 'explicit.pseudo_timeless_jewel_chitus',
    Caspiro: 'explicit.pseudo_timeless_jewel_caspiro'
  }
};

export const constructSingleResultQuery = (jewel: number, conqueror: string | null, result: SearchWithSeed): Query => {
  const anyConqueror = conqueror === null;

  const filters: Filter[] = Object.keys(tradeStatNames[jewel]).map((conq) => ({
    id: tradeStatNames[jewel][conq],
    value: {
      min: result.seed,
      max: result.seed
    },
    disabled: anyConqueror ? false : conq != conqueror
  }));

  const filterGroup = filtersToFilterGroup(filters, false);
  const query: Query = filterGroupsToQuery([filterGroup]);
  return query;
};

const constructSearchFilter = (jewel: number, conqueror: string | null, result: SearchWithSeed): Filter[] => {
  // null conqueror indicates to search for any conqueror
  const anyConqueror = conqueror === null;
  const conquerors = anyConqueror ? Object.keys(tradeStatNames[jewel]) : [conqueror];

  return conquerors.map((conq) => ({
    id: tradeStatNames[jewel][conq],
    value: {
      min: result.seed,
      max: result.seed
    }
  }));
};

export const constructQueries = (jewel: number, conqueror: string | null, results: SearchWithSeed[]) => {
  const max_filter_length = 50;
  const max_filters = 4;
  const max_query_length = max_filter_length * max_filters;

  // convert all results into filters
  const allFilters = results.flatMap((result) => constructSearchFilter(jewel, conqueror, result));

  // group filters into groups of max_query_length, where each group is further grouped into chunks of max_filter_length
  // this represents multiple trade links, where each trade link has multiple filter groups, and each filter group has multiple filters
  const queryFilterGroups = chunkArray(allFilters, max_query_length).map((chunk) =>
    chunkArray(chunk, max_filter_length)
  );

  // map filters groups to queries
  const tradeQueries = queryFilterGroups.map((queryFilterGroup) => {
    // for each query, map the chunks within it to filter groups
    const tradeFilterGroups = queryFilterGroup.map((filters, index) => filtersToFilterGroup(filters, index !== 0));
    const tradeQuery: Query = filterGroupsToQuery(tradeFilterGroups);
    return tradeQuery;
  });

  return tradeQueries;
};
