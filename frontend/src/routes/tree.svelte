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

  $: affectedNodes = getAffectedNodes(skillTree.nodes[circledNode]).filter((n) => !n.isJewelSocket && !n.isMastery);

  $: seedResults =
    !seed ||
    !selectedJewel ||
    !selectedConqueror ||
    window['TimelessJewelConquerors'][selectedJewel.value].indexOf(selectedConqueror.value) < 0
      ? []
      : affectedNodes.map((n) => ({
          node: n.skill,
          result: Calculate(TreeToPassive[n.skill], seed, selectedJewel.value, selectedConqueror.value)
        }));

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

  let minTotalWeight = 0;
  let searching = false;
  let currentSeed = 0;
  let searchResults: SearchResults;
  const search = () => {
    if (!circledNode) {
      return;
    }

    searching = true;
    searchResults = undefined;

    const query: ReverseSearchConfig = {
      jewel: selectedJewel.value,
      conqueror: selectedConqueror.value,
      nodes: affectedNodes
        .filter((n) => {
          return !disabled.has(n.skill);
        })
        .map((n) => window['TreeToPassive'][n.skill]),
      stats: Object.keys(selectedStats).map((stat) => selectedStats[stat]),
      minTotalWeight
    };

    syncWrap
      .search(
        query,
        proxy((s) => (currentSeed = s))
      )
      .then((result) => {
        searchResults = result;
        searching = false;
        collapsed = true;
      });
  };

  let highlighted: number[] = [];
  const highlight = (newSeed: number, passives: number[]) => {
    seed = newSeed;
    highlighted = passives;
    updateUrl();
  };

  const selectAll = () => {
    disabled.clear();
    // Re-assign to update svelte
    disabled = disabled;
  };

  const selectAllNotables = () => {
    affectedNodes.forEach((n) => {
      if (n.isNotable) {
        disabled.delete(n.skill);
      }
    });
    // Re-assign to update svelte
    disabled = disabled;
  };

  const selectAllPassives = () => {
    affectedNodes.forEach((n) => {
      if (!n.isNotable) {
        disabled.delete(n.skill);
      }
    });
    // Re-assign to update svelte
    disabled = disabled;
  };

  const deselectAll = () => {
    affectedNodes
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
      JSON.stringify(constructQuery(selectedJewel.value, selectedConqueror.value, searchResults.raw))
    );
    window.open(url, '_blank');
  };

  let groupResults = true;
  let collapsed = false;

  type CombinedResult = {
    stat: string;
    passives: number[];
  };

  const combineResults = (results: unknown[]): CombinedResult[] => {
    const mappedStats: { [key: number]: number[] } = {};
    results.forEach((r) => {
      if ('alternatePassiveSkill' in r.result) {
        const alt = GetAlternatePassiveSkillByIndex(r.result['alternatePassiveSkill']);
        if ('statsKeys' in alt) {
          alt['statsKeys'].forEach((key) => {
            mappedStats[key] = [...(mappedStats[key] || []), r.node];
          });
        }
      }

      if ('alternatePassiveAdditionInformations' in r.result) {
        r.result['alternatePassiveAdditionInformations'].forEach((info) => {
          const addition = GetAlternatePassiveAdditionByIndex(info['alternatePassiveSkillAddition']);
          if ('statsKeys' in addition) {
            addition['statsKeys'].forEach((key) => {
              mappedStats[key] = [...(mappedStats[key] || []), r.node];
            });
          }
        });
      }
    });

    return Object.keys(mappedStats)
      .map((statID) => ({
        stat: translateStat(parseInt(statID)),
        passives: mappedStats[statID]
      }))
      .sort((a, b) => b.passives.length - a.passives.length);
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
      <div class="flex flex-row justify-between mb-2">
        <h3 class="flex-grow">
          {collapsed ? 'Results' : 'Timeless Jewel'}
        </h3>
        {#if searchResults}
          <div class="flex flex-row">
            {#if collapsed}
              <button
                class="p-1 px-3 bg-blue-500/40 rounded disabled:bg-blue-900/40 mr-2"
                on:click={trade}
                disabled={!searchResults}
              >
                Trade
              </button>
              <button
                class="p-1 px-3 bg-blue-500/40 rounded disabled:bg-blue-900/40 mr-2"
                class:grouped={groupResults}
                on:click={() => (groupResults = !groupResults)}
                disabled={!searchResults}
              >
                Grouped
              </button>
            {/if}
            <button class="bg-neutral-100/20 px-4 p-1 rounded" on:click={() => (collapsed = !collapsed)}>
              {collapsed ? 'Config' : 'Results'}
            </button>
          </div>
        {/if}
      </div>

      {#if !collapsed}
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

              {#if seed >= window['TimelessJewelSeedRanges'][selectedJewel.value].min && seed <= window['TimelessJewelSeedRanges'][selectedJewel.value].max}
                <ul class="mt-4 overflow-auto">
                  {#each combineResults(seedResults) as r}
                    <li class="cursor-pointer" on:click={() => highlight(seed, r.passives)}>
                      <span class="font-bold">({r.passives.length})</span>
                      {r.stat}
                    </li>
                  {/each}
                </ul>
              {/if}
            {:else if mode === 'stats'}
              <div class="mt-4">
                <h3 class="mb-2">Add Stat</h3>
                <Select items={statItems} on:select={selectStat} bind:this={statSelector} />
              </div>
              {#if Object.keys(selectedStats).length > 0}
                <div class="mt-4 flex flex-col overflow-auto min-h-[100px]">
                  {#each Object.keys(selectedStats) as s}
                    <div class="mb-4 flex flex-row items-start flex-col border-neutral-100/40 border-b pb-4">
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
                <div class="flex flex-col mt-2">
                  <div class="flex flex-row items-center">
                    <div class="mr-2 min-w-fit">Min Total Weight:</div>
                    <input type="number" min="0" bind:value={minTotalWeight} />
                  </div>
                </div>
                <div class="flex flex-col mt-4">
                  <div class="flex flex-row">
                    <button
                      class="p-2 px-2 bg-yellow-500/40 rounded disabled:bg-yellow-900/40 mr-2"
                      on:click={selectAll}
                      disabled={searching || disabled.size == 0}
                    >
                      Select All
                    </button>
                    <button
                      class="p-2 px-2 bg-yellow-500/40 rounded disabled:bg-yellow-900/40 mr-2"
                      on:click={selectAllNotables}
                      disabled={searching || disabled.size == 0}
                    >
                      Notables
                    </button>
                    <button
                      class="p-2 px-2 bg-yellow-500/40 rounded disabled:bg-yellow-900/40 mr-2"
                      on:click={selectAllPassives}
                      disabled={searching || disabled.size == 0}
                    >
                      Passives
                    </button>
                    <button
                      class="p-2 px-2 bg-yellow-500/40 rounded disabled:bg-yellow-900/40 flex-grow"
                      on:click={deselectAll}
                      disabled={searching || disabled.size >= affectedNodes.length}
                    >
                      Deselect
                    </button>
                  </div>
                  <div class="flex flex-row mt-2">
                    <button
                      class="p-2 px-3 bg-green-500/40 rounded disabled:bg-green-900/40 flex-grow"
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
                </div>
              {/if}
            {/if}

            {#if !circledNode}
              <h2 class="mt-4">Click on a jewel socket</h2>
            {/if}
          {/if}
        {/if}
      {/if}

      {#if searchResults && collapsed}
        <SearchResults {searchResults} {groupResults} {highlight} />
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

  .grouped {
    @apply bg-pink-500/40 disabled:bg-pink-900/40;
  }
</style>
