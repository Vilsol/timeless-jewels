<script lang="ts">
  import type { SearchWithSeed } from '../skill_tree';
  import { skillTree, translateStat } from '../skill_tree';

  export let searchGrouped: { [key: number]: SearchWithSeed[] };
  export let highlight: (newSeed: number, passives: SearchWithSeed['skills']) => void;

  let expandedGroup = '';
</script>

<div class="mt-4 flex flex-col overflow-auto">
  {#each Object.keys(searchGrouped)
    .map((x) => parseInt(x))
    .sort((a, b) => a - b)
    .reverse() as k}
    <div class="mb-2">
      <button
        class="text-lg w-full p-2 px-4 bg-neutral-500/30 rounded flex flex-row justify-between"
        on:click={() => (expandedGroup = expandedGroup === k ? '' : k)}
      >
        <span>
          {k} Matches [{searchGrouped[k].length}]
        </span>
        <span>
          {expandedGroup === k ? '^' : 'V'}
        </span>
      </button>

      {#if expandedGroup === k}
        <div class="flex flex-col">
          {#each searchGrouped[k] as set}
            <div
              class="my-2 border-white/50 border p-2 flex flex-col cursor-pointer"
              on:click={() => highlight(set.seed, set.skills)}
            >
              <div class="font-bold text-orange-500 text-center">
                Seed {set.seed} (weight {set.weight})
              </div>
              {#each set.skills as skill}
                <div class="mt-2">
                  <span>
                    {skillTree.nodes[skill.passive].name} ({skill.passive})
                  </span>
                  <ul class="list-disc pl-6 font-bold">
                    {#each Object.keys(skill.stats) as stat}
                      <li>{translateStat(stat, skill.stats[stat])}</li>
                    {/each}
                  </ul>
                </div>
              {/each}
            </div>
          {/each}
        </div>
      {/if}
    </div>
  {/each}
</div>
