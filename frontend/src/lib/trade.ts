export interface SeedRef {
  seed: number;
}

export const tradeStatNames: { [key: number]: { [key: string]: string } } = {
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
  },
  6: {
    Vorana: 'explicit.pseudo_timeless_jewel_vorana',
    Uhtred: 'explicit.pseudo_timeless_jewel_uhtred',
    Medved: 'explicit.pseudo_timeless_jewel_medved'
  }
};

const platformRealms: { [key: string]: string } = {
  PC: '',
  Xbox: 'xbox',
  Playstation: 'sony'
};

export const constructQuery = <T extends SeedRef>(
  jewel: number,
  conqueror: string,
  result: T[],
  isLegacyTradersMode = false
) => {
  const max_filter_length = 45;
  const max_filters = 4;
  const max_query_length = max_filter_length * max_filters;

  if (result.length === 0) {
    throw new Error('constructQuery: result must not be empty');
  }
  const conquerors = Object.keys(tradeStatNames[jewel] ?? {});
  if (conquerors.length === 0) {
    throw new Error(`constructQuery: unknown jewel type ${jewel}`);
  }
  if (!conquerors.includes(conqueror)) {
    throw new Error(`constructQuery: unknown conqueror "${conqueror}" for jewel ${jewel}`);
  }

  interface TradeFilter {
    id: string;
    value: { min: number; max: number };
    disabled?: boolean;
  }
  interface TradeStatGroup {
    type: 'count';
    value: { min: number };
    filters: TradeFilter[];
    disabled: boolean;
  }
  const final_query: TradeStatGroup[] = [];

  if (result.length === 1) {
    final_query.push({
      type: 'count',
      value: { min: 1 },
      disabled: false,
      filters: conquerors.map((conq) => ({
        id: tradeStatNames[jewel][conq],
        value: { min: result[0].seed, max: result[0].seed },
        disabled: conq != conqueror
      }))
    });
  } else if (result.length <= max_filter_length) {
    // Multi-conqueror compare layout: one stat group per conqueror, only the
    // selected one enabled. The disabled groups stay toggleable in PoE trade
    // so users can quickly compare seed coverage across conquerors.
    for (const conq of conquerors) {
      final_query.push({
        type: 'count',
        value: { min: 1 },
        filters: result.map((r) => ({
          id: tradeStatNames[jewel][conq],
          value: { min: r.seed, max: r.seed }
        })),
        disabled: conq != conqueror
      });
    }
  } else {
    // Selected-conqueror deep-dive: too many matches to compare against
    // other conquerors meaningfully, so chunk the selected conqueror's seeds
    // across all 4 stat groups. First group enabled, rest toggleable.
    const selectedId = tradeStatNames[jewel][conqueror];
    for (let i = 0; i < max_filters; i++) {
      final_query.push({
        type: 'count',
        value: { min: 1 },
        filters: [],
        disabled: i != 0
      });
    }
    for (const [i, r] of result.slice(0, max_query_length).entries()) {
      const index = Math.floor(i / max_filter_length);
      final_query[index].filters.push({
        id: selectedId,
        value: { min: r.seed, max: r.seed }
      });
    }
  }

  return {
    query: {
      status: { option: isLegacyTradersMode ? 'online' : 'any' },
      stats: final_query
    },
    sort: { price: 'asc' }
  };
};

export const tradeUrl = (platform: string, league: string): string => {
  if (!platform || typeof platform !== 'string') {
    platform = 'PC';
  }

  if (!league || typeof league !== 'string') {
    league = 'Standard';
  }

  const realm = platformRealms[platform] ?? platform.toLowerCase();
  return `https://www.pathofexile.com/trade/search${realm ? `/${realm}` : ''}/${encodeURIComponent(league)}`;
};

export const openTrade = <T extends SeedRef>(
  jewel: number,
  conqueror: string,
  results: T[],
  platform: string,
  league: string,
  isLegacyTradersMode = false
) => {
  const url = new URL(tradeUrl(platform, league));
  url.searchParams.set('q', JSON.stringify(constructQuery(jewel, conqueror, results, isLegacyTradersMode)));

  console.log('opening trade', url);

  window.open(url, '_blank');
};
