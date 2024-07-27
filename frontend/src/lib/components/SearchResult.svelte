<script lang="ts">
  import { openQueryTrade } from '$lib/utils/trade_utils';
  import { constructSingleResultQuery, type SearchWithSeed } from '../skill_tree';
  import { skillTree, translateStat } from '../skill_tree';

  export let highlight: (newSeed: number, passives: number[]) => void;
  export let set: SearchWithSeed;
  export let jewel: number;
  export let conqueror: string;

  const handleOnClick = () =>
    highlight(
      set.seed,
      set.skills.map((s) => s.passive)
    );
</script>

<div
  class="my-2 border-white/50 border p-2 flex flex-col cursor-pointer"
  on:click={handleOnClick}
  on:keydown={handleOnClick}
  role="button"
  tabindex="0">
  <div class="flex flex-row justify-between">
    <!-- Padding -->
    <button class="px-3 invisible">Trade</button>
    <div class="font-bold text-orange-500 text-center">
      Seed {set.seed} (weight {set.weight})
    </div>
    <button
      class="px-3 bg-blue-500/40 rounded"
      on:click={() => openQueryTrade(constructSingleResultQuery(jewel, conqueror, set))}>Trade</button>
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
