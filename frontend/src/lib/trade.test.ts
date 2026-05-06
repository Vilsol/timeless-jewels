import { describe, expect, it } from 'vitest';
import { constructQuery, tradeStatNames, tradeUrl } from './trade';

const seeds = (...n: number[]) => n.map((seed) => ({ seed }));
const stripPrefix = (id: string) => id.replace('explicit.pseudo_timeless_jewel_', '');

describe('constructQuery — multi-seed (≤180 results)', () => {
  // Repros issues #25, #26, #31, #51 (and #15, #33).
  // The user picks one conqueror; the four PoE-trade "count" boxes should each
  // target exactly one conqueror, with only the chosen one enabled.

  it('returns four DISTINCT stat-group objects (not aliased)', () => {
    const q = constructQuery(5, 'Caspiro', seeds(12780, 13000, 13500));
    const groups = q.query.stats;
    expect(groups).toHaveLength(4);
    // Cross-product of references must all differ
    for (let i = 0; i < groups.length; i++) {
      for (let j = i + 1; j < groups.length; j++) {
        expect(groups[i]).not.toBe(groups[j]);
      }
    }
  });

  it('each stat-group contains filters for exactly ONE conqueror', () => {
    const q = constructQuery(5, 'Caspiro', seeds(12780, 13000, 13500));
    const conquerorsPerGroup = q.query.stats.map(
      (g) => new Set(g.filters.map((f: { id: string }) => stripPrefix(f.id)))
    );
    for (const set of conquerorsPerGroup) {
      expect(set.size).toBe(1);
    }
    // And the union across all groups should cover all four conquerors
    const all = new Set(conquerorsPerGroup.flatMap((s) => [...s]));
    expect(all).toEqual(new Set(['cadiro', 'victario', 'chitus', 'caspiro']));
  });

  it('each stat-group has exactly result.length filters (no cross-conqueror accumulation)', () => {
    const seedList = seeds(12780, 13000, 13500);
    const q = constructQuery(5, 'Caspiro', seedList);
    for (const g of q.query.stats) {
      expect(g.filters).toHaveLength(seedList.length);
    }
  });

  it('only the SELECTED conqueror group is enabled (others disabled)', () => {
    const q = constructQuery(5, 'Caspiro', seeds(12780, 13000, 13500));
    const enabled = q.query.stats.filter((g) => !g.disabled);
    expect(enabled).toHaveLength(1);
    const enabledConqs = new Set(enabled[0].filters.map((f: { id: string }) => stripPrefix(f.id)));
    expect(enabledConqs).toEqual(new Set(['caspiro']));
  });

  // Repros the specific report in #25: "I picked Caspiro, got Cadiro/Victario instead".
  it('selecting any non-last conqueror still yields exactly one enabled group for that conqueror', () => {
    for (const target of ['Cadiro', 'Victario', 'Chitus']) {
      const q = constructQuery(5, target, seeds(100, 200));
      const enabled = q.query.stats.filter((g) => !g.disabled);
      expect(enabled, `selected=${target}`).toHaveLength(1);
      const ids = new Set(enabled[0].filters.map((f: { id: string }) => stripPrefix(f.id)));
      expect(ids, `selected=${target}`).toEqual(new Set([target.toLowerCase()]));
    }
  });

  // PR #55 (haggys22) tried to fix this by sorting the conqueror loop so the
  // selected conqueror iterates first, rescuing its filters from the slice(0,45)
  // truncation in the original buggy code. This test locks the stronger guarantee:
  // selected-conqueror coverage is independent of iteration position, at any
  // result-count regime, including the slice-relevant 50-seed and 200-seed cases
  // where PR #55's workaround would have left the boxes disabled-by-default or
  // truncated to 45.
  it('selected-conqueror coverage is independent of iteration order (PR #55)', () => {
    const jewelFiveOrder = Object.keys(tradeStatNames[5]); // Cadiro, Victario, Chitus, Caspiro
    for (const target of jewelFiveOrder) {
      for (const count of [3, 30, 50, 200]) {
        const list = Array.from({ length: count }, (_, i) => ({ seed: i + 1 }));
        const q = constructQuery(5, target, list);

        // Whatever branch we land in, the selected conqueror must be the only
        // conqueror enabled by default (or, in deep-dive mode, the only
        // conqueror represented at all).
        const enabled = q.query.stats.filter((g) => !g.disabled);
        expect(enabled.length, `${target}@${count}: enabled groups`).toBeGreaterThanOrEqual(1);
        const enabledIds = new Set(enabled.flatMap((g) => g.filters.map((f: { id: string }) => stripPrefix(f.id))));
        expect(enabledIds, `${target}@${count}: enabled-conqueror ids`).toEqual(new Set([target.toLowerCase()]));

        // And the selected conqueror's seeds are never silently truncated
        // beyond PoE's hard cap (max_query_length = 180).
        const seedSet = new Set(
          q.query.stats.flatMap((g) =>
            g.filters
              .filter((f: { id: string }) => stripPrefix(f.id) === target.toLowerCase())
              .map((f: { value: { min: number } }) => f.value.min)
          )
        );
        expect(seedSet.size, `${target}@${count}: selected seeds in query`).toBe(Math.min(count, 180));
      }
    }
  });

  // The 4-group layout is intentional — the 3 disabled groups must remain
  // functional comparators so the user can toggle them on in PoE trade and
  // immediately see how the same seeds look for a different conqueror.
  // This locks the full compare-mode contract in one assertion.
  it('compare mode: 4 groups (1 enabled selected + 3 disabled functional comparators)', () => {
    const seedList = seeds(100, 200, 300);
    const q = constructQuery(5, 'Victario', seedList);
    expect(q.query.stats).toHaveLength(4);

    const enabled = q.query.stats.filter((g) => !g.disabled);
    const disabled = q.query.stats.filter((g) => g.disabled);
    expect(enabled).toHaveLength(1);
    expect(disabled).toHaveLength(3);

    const conquerorOf = (g: { filters: { id: string }[] }) => stripPrefix(g.filters[0].id);
    expect(conquerorOf(enabled[0])).toBe('victario');
    expect(new Set(disabled.map(conquerorOf))).toEqual(new Set(['cadiro', 'chitus', 'caspiro']));

    // Each disabled group must mirror the enabled one — same seed values,
    // same filter count, just a different stat ID. That's what makes them
    // useful comparators when toggled on.
    const seedValues = (g: { filters: { value: { min: number; max: number } }[] }) =>
      g.filters.map((f) => [f.value.min, f.value.max]);
    for (const d of disabled) {
      expect(d.filters).toHaveLength(seedList.length);
      expect(seedValues(d)).toEqual(seedValues(enabled[0]));
      // and only one conqueror's stat ID per group
      const ids = new Set(d.filters.map((f) => stripPrefix(f.id)));
      expect(ids.size).toBe(1);
      expect(ids.has('victario')).toBe(false);
    }
  });
});

describe('constructQuery — single-seed', () => {
  it('returns one stat group with one filter per conqueror; only selected one enabled', () => {
    const q = constructQuery(5, 'Caspiro', seeds(12780));
    expect(q.query.stats).toHaveLength(1);
    const g = q.query.stats[0];
    expect(g.filters).toHaveLength(4);
    const enabledFilters = g.filters.filter((f: { disabled: boolean }) => !f.disabled);
    expect(enabledFilters).toHaveLength(1);
    expect(stripPrefix(enabledFilters[0].id)).toBe('caspiro');
  });
});

describe('constructQuery — overflow (>180 results)', () => {
  it('packs only the selected conqueror into 4 groups, first enabled', () => {
    const many = Array.from({ length: 200 }, (_, i) => ({ seed: i + 1 }));
    const q = constructQuery(5, 'Chitus', many);
    expect(q.query.stats).toHaveLength(4);
    expect(q.query.stats[0].disabled).toBe(false);
    for (let i = 1; i < 4; i++) {
      expect(q.query.stats[i].disabled).toBe(true);
    }
    const total = q.query.stats.reduce((acc: number, g) => acc + g.filters.length, 0);
    expect(total).toBe(180); // 45 * 4
    const allIds = new Set(q.query.stats.flatMap((g) => g.filters.map((f: { id: string }) => stripPrefix(f.id))));
    expect(allIds).toEqual(new Set(['chitus']));
  });
});

describe('Heroic Tragedy support (issue #56)', () => {
  it('tradeStatNames has an entry for jewel type 6 with Vorana, Uhtred, Medved', () => {
    expect(tradeStatNames[6]).toBeDefined();
    expect(Object.keys(tradeStatNames[6]).sort()).toEqual(['Medved', 'Uhtred', 'Vorana']);
  });

  it('constructQuery does not throw for Heroic Tragedy', () => {
    expect(() => constructQuery(6, 'Vorana', seeds(150, 200))).not.toThrow();
  });
});

describe('tradeUrl — platform/realm mapping (issue #52)', () => {
  it('PC has no realm path segment', () => {
    expect(tradeUrl('PC', 'Standard')).toBe('https://www.pathofexile.com/trade/search/Standard');
  });

  it('Xbox maps to /xbox/', () => {
    expect(tradeUrl('Xbox', 'Standard')).toBe('https://www.pathofexile.com/trade/search/xbox/Standard');
  });

  // Repros #52: GGG renamed playstation realm to "sony".
  it('Playstation maps to /sony/ (not /playstation/)', () => {
    expect(tradeUrl('Playstation', 'Standard')).toBe('https://www.pathofexile.com/trade/search/sony/Standard');
  });

  it('falls back to PC + Standard when platform/league are missing', () => {
    expect(tradeUrl('', '')).toBe('https://www.pathofexile.com/trade/search/Standard');
    // @ts-expect-error: testing runtime fallback behavior
    expect(tradeUrl(undefined, undefined)).toBe('https://www.pathofexile.com/trade/search/Standard');
  });
});

// ---------------------------------------------------------------------------
// RED tests below: each captures a real gap still present in trade.ts.
// They are intentionally failing — leave them red until the underlying issue
// is addressed, then turn them green.
// ---------------------------------------------------------------------------

describe('RED — constructQuery rejects degenerate input', () => {
  // Currently: returns a query with 4 empty stat groups. PoE rejects it
  // silently and the user gets the "no results" error from issue #46.
  it('throws (or returns no stats) for an empty result array', () => {
    expect(() => constructQuery(5, 'Caspiro', [])).toThrow();
  });

  // Currently: Object.keys(tradeStatNames[99] ?? {}) === [], so the function
  // happily returns a stats:[] query — a confusing silent no-op for callers.
  it('throws for an unknown jewel type', () => {
    expect(() => constructQuery(99, 'Whoever', seeds(1))).toThrow(/jewel/i);
  });

  // Currently: produces 4 groups, all disabled. PoE returns nothing and the
  // user can't tell why. Likely root cause behind some "wrong conqueror"
  // reports where stale localStorage held a renamed conqueror name.
  it('throws for an unknown conqueror name', () => {
    expect(() => constructQuery(5, 'NotAConqueror', seeds(1))).toThrow(/conqueror/i);
  });
});

describe('RED — selected-conqueror seed coverage (issues #7, #14)', () => {
  // The multi-result branch slices each conqueror's filters to 45, so a user
  // with 100 matching seeds for their selected conqueror only sees the first
  // 45 in PoE trade. The remaining 55 are silently dropped. This is the
  // root cause behind "more than 28 jewels not appearing" — our cap is 45,
  // but the bug is that we drop *selected*-conqueror seeds in the first place.
  it('does not silently drop selected-conqueror seeds when 46 ≤ count ≤ 180', () => {
    const hundred = Array.from({ length: 100 }, (_, i) => ({ seed: i + 1 }));
    const q = constructQuery(5, 'Caspiro', hundred);
    // Seeds may live in disabled groups (toggleable in PoE trade UI) but must
    // be *present* in the query — silent truncation is the bug.
    const seedSet = new Set(
      q.query.stats.flatMap((g) => g.filters.map((f: { value: { min: number } }) => f.value.min))
    );
    expect(seedSet.size).toBe(100);
    // And every filter still targets the selected conqueror only.
    const ids = new Set(q.query.stats.flatMap((g) => g.filters.map((f: { id: string }) => stripPrefix(f.id))));
    expect(ids).toEqual(new Set(['caspiro']));
  });
});

describe('RED — tradeUrl encodes league names safely', () => {
  // PoE has shipped league names with spaces ("Solo Self-Found", "Phrecia Hardcore").
  // We currently inject the raw string into the URL path, so the resulting URL
  // contains literal spaces — most browsers tolerate this on window.open, but
  // anything that round-trips through URL parsing or copy-paste breaks.
  it('percent-encodes spaces in league names', () => {
    const url = tradeUrl('PC', 'Phrecia Solo Self-Found');
    expect(url).not.toContain(' ');
    expect(url).toContain('Phrecia%20Solo%20Self-Found');
  });

  // Same concern for league names containing other reserved characters.
  it('percent-encodes ampersands and slashes in league names', () => {
    const url = tradeUrl('PC', 'A&B/C');
    expect(url).toContain('A%26B%2FC');
  });
});
