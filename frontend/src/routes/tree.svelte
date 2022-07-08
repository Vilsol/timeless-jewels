<script lang="ts">
  import SkillTree from '../lib/components/SkillTree.svelte';
  import Select from 'svelte-select';
  import { page } from '$app/stores';
  import { goto } from '$app/navigation';
  import type { Node } from '../lib/types';
  import { constructQuery, getAffectedNodes, skillTree, translateStat } from '../lib/skill_tree';
  import { syncWrap } from '../lib/worker';
  import { proxy } from 'comlink';
  import type { SearchWithSeed, ReverseSearchConfig, StatConfig } from '../lib/skill_tree';
  import SearchResults from '../lib/components/SearchResults.svelte';

  const searchParams = $page.url.searchParams;

  const jewels = Object.keys(window['TimelessJewels']).map((k) => ({
    value: parseInt(k),
    label: window['TimelessJewels'][k]
  }));

  let selectedJewel = searchParams.has('jewel') ? jewels.find((j) => j.value == searchParams.get('jewel')) : undefined;

  $: conquerors = selectedJewel
    ? window['TimelessJewelConquerors'][selectedJewel.value].map((k) => ({
        value: k,
        label: k
      }))
    : [];

  let selectedConqueror = searchParams.has('conqueror')
    ? {
        value: searchParams.get('conqueror'),
        label: searchParams.get('conqueror')
      }
    : undefined;

  let seed: number = searchParams.has('seed') ? parseInt(searchParams.get('seed')) : 0;

  let circledNode: number | undefined = searchParams.has('location')
    ? parseInt(searchParams.get('location'))
    : undefined;

  let disabled = new Set();
  const clickNode = (node: Node) => {
    if (node.isJewelSocket) {
      circledNode = node.skill;
      updateUrl();
    } else if (!node.isMastery) {
      if (disabled.has(node.skill)) {
        disabled.delete(node.skill);
      } else {
        disabled.add(node.skill);
      }
      // Re-assign to update svelte
      disabled = disabled;
    }
  };

  const updateUrl = () => {
    const url = new URL(window.location.origin + window.location.pathname);
    selectedJewel && url.searchParams.append('jewel', selectedJewel.value.toString());
    selectedConqueror && url.searchParams.append('conqueror', selectedConqueror.value);
    seed && url.searchParams.append('seed', seed.toString());
    circledNode && url.searchParams.append('location', circledNode.toString());
    mode && url.searchParams.append('mode', mode);

    Object.keys(selectedStats).forEach((s) => {
      url.searchParams.append('stat', s.toString());
    });

    goto(url.toString());
  };

  let mode = searchParams.has('mode') ? searchParams.get('mode') : '';
  const setMode = (newMode: string) => {
    mode = newMode;
    updateUrl();
  };

  let allPossibleStats: { [key: string]: { [key: string]: number } } = JSON.parse(window['PossibleStats']);

  let selectedStats: Record<number, StatConfig> = {};
  if (searchParams.has('stat')) {
    searchParams.getAll('stat').map((s) => {
      const nStat = parseInt(s);
      selectedStats[nStat] = {
        weight: 1,
        min: 0,
        id: nStat
      };
    });
  }

  $: availableStats = !selectedJewel ? [] : Object.keys(allPossibleStats[selectedJewel.value]);
  $: statItems = availableStats
    .map((statId) => {
      const id = parseInt(statId);
      return {
        label: translateStat(id),
        value: id
      };
    })
    .filter((s) => !(s.value in selectedStats));

  let statSelector: Select;
  const selectStat = (stat: CustomEvent) => {
    selectedStats[stat.detail.value] = {
      weight: 1,
      min: 0,
      id: stat.detail.value
    };
    selectedStats = selectedStats;
    statSelector['handleClear']();
    updateUrl();
  };

  const removeStat = (id: number) => {
    delete selectedStats[id];
    // Re-assign to update svelte
    selectedStats = selectedStats;
    updateUrl();
  };

  const changeJewel = () => {
    selectedStats = {};
    updateUrl();
  };

  let searching = false;
  let currentSeed = 0;
  let searchGrouped: { [key: number]: SearchWithSeed[] };
  const search = () => {
    if (!circledNode) {
      return;
    }

    searching = true;
    searchGrouped = undefined;

    const affectedNodes = getAffectedNodes(skillTree.nodes[circledNode])
      .filter((n) => {
        return !n.isJewelSocket && !n.isMastery;
      })
      .filter((n) => {
        return !disabled.has(n.skill);
      })
      .map((n) => window['TreeToPassive'][n.skill]);

    const query: ReverseSearchConfig = {
      jewel: selectedJewel.value,
      conqueror: selectedConqueror.value,
      nodes: affectedNodes,
      stats: Object.keys(selectedStats).map((stat) => selectedStats[stat])
    };

    syncWrap
      .search(
        query,
        proxy((s) => (currentSeed = s))
      )
      .then((result) => {
        searchGrouped = result;
        searching = false;
      });
  };

  let highlighted: number[] = [];
  const highlight = (newSeed: number, passives: SearchWithSeed['skills']) => {
    seed = newSeed;
    highlighted = passives.map((s) => s.passive);
    updateUrl();
  };

  const selectAll = () => {
    disabled.clear();
    // Re-assign to update svelte
    disabled = disabled;
  };

  const deselectAll = () => {
    getAffectedNodes(skillTree.nodes[circledNode])
      .filter((n) => {
        return !n.isJewelSocket && !n.isMastery;
      })
      .forEach((n) => disabled.add(n.skill));
    // Re-assign to update svelte
    disabled = disabled;
  };

  const trade = () => {
    const url = new URL('https://www.pathofexile.com/trade/search/Sentinel');
    url.searchParams.set(
      'q',
      JSON.stringify(constructQuery(selectedJewel.value, selectedConqueror.value, searchGrouped))
    );
    window.open(url, '_blank');
  };
</script>

<SkillTree
  {clickNode}
  {circledNode}
  selectedJewel={selectedJewel?.value}
  selectedConqueror={selectedConqueror?.value}
  {highlighted}
  {seed}
  disabled={[...disabled]}
>
  <div
    class="w-1/2 xl:w-1/3 2xl:w-1/4 3xl:w-1/5 absolute top-0 left-0 bg-black/80 backdrop-blur-sm themed rounded-br-lg max-h-screen"
  >
    <div class="p-4 max-h-screen flex flex-col">
      <h3 class="mb-2">Timeless Jewel</h3>
      <Select items={jewels} bind:value={selectedJewel} on:select={changeJewel} />

      {#if selectedJewel}
        <div class="mt-4">
          <h3 class="mb-2">Conqueror</h3>
          <Select items={conquerors} bind:value={selectedConqueror} on:select={updateUrl} />
        </div>

        {#if selectedConqueror && window['TimelessJewelConquerors'][selectedJewel.value].indexOf(selectedConqueror.value) >= 0}
          <div class="mt-4 w-full flex flex-row">
            <button class="selection-button" class:selected={mode === 'seed'} on:click={() => setMode('seed')}>
              Enter Seed
            </button>
            <button class="selection-button" class:selected={mode === 'stats'} on:click={() => setMode('stats')}>
              Select Stats
            </button>
          </div>

          {#if mode === 'seed'}
            <div class="mt-4">
              <h3 class="mb-2">Seed</h3>
              <input
                type="number"
                bind:value={seed}
                on:blur={updateUrl}
                min={window['TimelessJewelSeedRanges'][selectedJewel.value].min}
                max={window['TimelessJewelSeedRanges'][selectedJewel.value].max}
              />
              {#if seed < window['TimelessJewelSeedRanges'][selectedJewel.value].min || seed > window['TimelessJewelSeedRanges'][selectedJewel.value].max}
                <div class="mt-2">
                  Seed must be between {window['TimelessJewelSeedRanges'][selectedJewel.value].min}
                  and {window['TimelessJewelSeedRanges'][selectedJewel.value].max}
                </div>
              {/if}
            </div>
          {:else if mode === 'stats'}
            <div class="mt-4">
              <h3 class="mb-2">Add Stat</h3>
              <Select items={statItems} on:select={selectStat} bind:this={statSelector} />
            </div>
            {#if Object.keys(selectedStats).length > 0}
              <div class="mt-4 flex flex-col overflow-auto min-h-[100px]">
                {#each Object.keys(selectedStats) as s}
                  <div class="mb-2 flex flex-row items-start flex-col stat-config">
                    <div>
                      <button
                        class="p-2 px-4 bg-red-500/40 rounded mr-2"
                        on:click={() => removeStat(selectedStats[s].id)}
                      >
                        -
                      </button>
                      <span>{translateStat(selectedStats[s].id)}</span>
                    </div>
                    <div class="mt-2 flex flex-row">
                      <div class="mr-4 flex flex-row items-center">
                        <div class="mr-2">Min:</div>
                        <input type="number" min="0" bind:value={selectedStats[s].min} />
                      </div>
                      <div class="flex flex-row items-center">
                        <div class="mr-2">Weight:</div>
                        <input type="number" min="0" bind:value={selectedStats[s].weight} />
                      </div>
                    </div>
                  </div>
                {/each}
              </div>
              <div class="flex flex-row">
                <button
                  class="p-2 px-3 bg-yellow-500/40 rounded mt-2 disabled:bg-yellow-900/40 mr-2"
                  on:click={selectAll}
                  disabled={searching || disabled.size == 0}
                >
                  Select All
                </button>
                <button
                  class="p-2 px-3 bg-yellow-500/40 rounded mt-2 disabled:bg-yellow-900/40 mr-2"
                  on:click={deselectAll}
                  disabled={searching}
                >
                  Deselect All
                </button>
                <button
                  class="p-2 px-3 bg-blue-500/40 rounded mt-2 disabled:bg-blue-900/40 mr-2"
                  on:click={trade}
                  disabled={!searchGrouped}
                >
                  Trade
                </button>
                <button
                  class="p-2 px-3 bg-green-500/40 rounded mt-2 disabled:bg-green-900/40 flex-grow"
                  on:click={() => search()}
                  disabled={searching}
                >
                  {#if searching}
                    {currentSeed} / {window['TimelessJewelSeedRanges'][selectedJewel.value].max}
                  {:else}
                    Search
                  {/if}
                </button>
              </div>

              {#if searchGrouped}
                <SearchResults {searchGrouped} {highlight} />
              {/if}
            {/if}
          {/if}

          {#if !circledNode}
            <h2 class="mt-4">Click on a jewel socket</h2>
          {/if}
        {/if}
      {/if}
    </div>
  </div>
</SkillTree>

<style>
  .selection-button {
    @apply bg-neutral-500/20 p-2 px-4 flex-grow;
  }

  .selection-button:first-child {
    @apply rounded-l border-r-2 border-black;
  }

  .selection-button:last-child {
    @apply rounded-r;
  }

  .selected {
    @apply bg-neutral-100/20;
  }

  .stat-config:not(:first-child) {
    @apply border-neutral-100/40 border-t pt-2;
  }
</style>
