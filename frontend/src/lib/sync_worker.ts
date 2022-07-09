import { expose } from 'comlink';
import '../wasm_exec.js';
import { loadSkillTree, passiveToTree } from './skill_tree';
import type { SearchWithSeed, ReverseSearchConfig, SearchResults } from './skill_tree';

const obj = {
  boot(data: ArrayBuffer) {
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-ignore
    const go = new Go();
    WebAssembly.instantiate(data, go.importObject).then((result) => {
      go.run(result.instance);

      // eslint-disable-next-line @typescript-eslint/ban-ts-comment
      // @ts-ignore
      loadSkillTree(SkillTree, PassiveTranslations, TreeToPassive);
    });
  },
  search(args: ReverseSearchConfig, callback: (seed: number) => void): SearchResults {
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-ignore
    const searchResult = ReverseSearch(
      args.nodes,
      args.stats.map((s) => s.id),
      args.jewel,
      args.conqueror,
      callback
    );

    const searchGrouped: { [key: number]: SearchWithSeed[] } = {};
    Object.keys(searchResult).forEach((seed) => {
      let weight = 0;

      const statCounts: Record<number, number> = {};
      const skills = Object.keys(searchResult[seed]).map((skillID) => {
        Object.keys(searchResult[seed][skillID]).forEach((st) => {
          const n = parseInt(st);
          statCounts[n] = (statCounts[n] || 0) + 1;
          weight += args.stats.find((s) => s.id == n)?.weight || 0;
        });

        return {
          passive: passiveToTree[parseInt(skillID)],
          stats: searchResult[seed][skillID]
        };
      });

      const len = Object.keys(searchResult[seed]).length;
      searchGrouped[len] = [
        ...(searchGrouped[len] || []),
        {
          skills: skills,
          seed: parseInt(seed),
          weight,
          statCounts
        }
      ];
    });

    Object.keys(searchGrouped).forEach((len) => {
      const nLen = parseInt(len);
      searchGrouped[nLen] = searchGrouped[nLen].filter((g) => {
        if (g.weight < args.minTotalWeight) {
          return false;
        }

        for (const stat of args.stats) {
          if ((g.statCounts[stat.id] === undefined && stat.min > 0) || g.statCounts[stat.id] < stat.min) {
            return false;
          }
        }

        return true;
      });

      if (Object.keys(searchGrouped[nLen]).length == 0) {
        delete searchGrouped[nLen];
      } else {
        searchGrouped[nLen] = searchGrouped[nLen].sort((a, b) => {
          return b.weight - a.weight;
        });
      }
    });

    return {
      grouped: searchGrouped,
      raw: Object.keys(searchGrouped)
        .map((x) => searchGrouped[parseInt(x)])
        .flat()
        .sort((a, b) => b.weight - a.weight)
    };
  }
};

expose(obj);
