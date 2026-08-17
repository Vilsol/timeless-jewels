import { expose } from 'comlink';
import '../wasm_exec.js';
import { loadSkillTree, passiveToTree } from './skill_tree';
import type { SearchWithSeed, ReverseSearchConfig, SearchResults } from './skill_tree';
import { calculator, initializeCrystalline } from './types';

const obj = {
  boot(wasm: ArrayBuffer) {
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-ignore
    const go = new Go();
    WebAssembly.instantiate(wasm, go.importObject).then((result) => {
      go.run(result.instance);

      initializeCrystalline();

      loadSkillTree();
    });
  },
  async search(args: ReverseSearchConfig, callback: (seed: number) => Promise<void>): Promise<SearchResults> {
    const searchResult = await calculator.ReverseSearch(
      [...args.nodes, ...args.anointNodes],
      args.stats.map((s) => s.id),
      args.jewel,
      args.conqueror,
      callback
    );

    const anointSet = new Set(args.anointNodes);

    const searchGrouped: { [key: number]: SearchWithSeed[] } = {};
    Object.keys(searchResult).forEach((seedStr) => {
      const seed = parseInt(seedStr);

      const statCounts: Record<number, number> = {};
      const anointCandidates: { score: number; skillID: number }[] = [];

      const skills = Object.keys(searchResult[seed])
        .filter((skillIDStr) => !anointSet.has(parseInt(skillIDStr)))
        .map((skillIDStr) => {
          const skillID = parseInt(skillIDStr);
          Object.keys(searchResult[seed][skillID]).forEach((st) => {
            const n = parseInt(st);
            statCounts[n] = (statCounts[n] || 0) + 1;
          });

          return {
            passive: passiveToTree[skillID],
            stats: searchResult[seed][skillID]
          };
        });

      const runningCounts = { ...statCounts };

      if (args.anoints > 0) {
        const seedResult = (searchResult ?? {})[seed] ?? {};
        const anointSkillIDs = Object.keys(seedResult)
          .filter((skillIDStr) => anointSet.has(parseInt(skillIDStr)))
          .map((skillIDStr) => parseInt(skillIDStr));

        // Marginal weight of taking this node next, given what is already counted.
        const scoreCandidate = (skillID: number) => {
          let nodeScore = 0;
          const nodeStats = seedResult[skillID] ?? {};
          for (const stat of args.stats) {
            const count = nodeStats[stat.id] !== undefined ? 1 : 0;
            const remainingCap = stat.max > 0 ? Math.max(0, stat.max - (runningCounts[stat.id] || 0)) : count;
            const effectiveCount = stat.max > 0 ? Math.min(count, remainingCap) : count;
            nodeScore += effectiveCount * stat.weight;
          }
          return nodeScore;
        };

        const remaining = new Set(anointSkillIDs);
        for (let i = 0; i < args.anoints && remaining.size > 0; i++) {
          let bestID = -1;
          let bestScore = -1;
          for (const skillID of remaining) {
            const score = scoreCandidate(skillID);
            if (score > bestScore) {
              bestScore = score;
              bestID = skillID;
            }
          }
          if (bestID === -1) {
            break;
          }

          remaining.delete(bestID);
          anointCandidates.push({ score: bestScore, skillID: bestID });
          Object.keys(seedResult[bestID] ?? {}).forEach((st) => {
            const n = parseInt(st);
            runningCounts[n] = (runningCounts[n] || 0) + 1;
          });
          skills.push({
            passive: passiveToTree[bestID],
            stats: (seedResult[bestID] ?? {}) as { [key: string]: number }
          });
        }
      }

      let weight = 0;
      for (const stat of args.stats) {
        const count = statCounts[stat.id] || 0;
        const effectiveCount = stat.max > 0 ? Math.min(count, stat.max) : count;
        weight += effectiveCount * stat.weight;
      }
      for (let i = 0; i < Math.min(args.anoints, anointCandidates.length); i++) {
        weight += anointCandidates[i].score;
      }

      const len = skills.length;
      searchGrouped[len] = [
        ...(searchGrouped[len] || []),
        {
          skills: skills,
          seed,
          weight,
          statCounts: runningCounts
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
        searchGrouped[nLen] = searchGrouped[nLen].sort((a, b) => b.weight - a.weight);
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
} as const;

expose(obj);

export type WorkerType = typeof obj;
