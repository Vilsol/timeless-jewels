<script lang="ts">
  import SkillTree from '../../lib/components/SkillTree.svelte';
  import Select from 'svelte-select';
  import { page } from '$app/stores';
  import { goto } from '$app/navigation';
  import type { Node } from '../../lib/skill_tree_types';
  import { getAffectedNodes, skillTree, translateStat, openTrade } from '../../lib/skill_tree';
  import { syncWrap } from '../../lib/worker';
  import { proxy } from 'comlink';
  import type { ReverseSearchConfig, StatConfig } from '../../lib/skill_tree';
  import SearchResults from '../../lib/components/SearchResults.svelte';
  import { statValues } from '../../lib/values';
  import { data, calculator } from '../../lib/types';

  const searchParams = $page.url.searchParams;

  const jewels = Object.keys(data.TimelessJewels).map((k) => ({
    value: parseInt(k),
    label: data.TimelessJewels[k]
  }));

  let selectedJewel = searchParams.has('jewel') ? jewels.find((j) => j.value == searchParams.get('jewel')) : undefined;

  $: conquerors = selectedJewel
    ? Object.keys(data.TimelessJewelConquerors[selectedJewel.value]).map((k) => ({
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

  $: affectedNodes = circledNode
    ? getAffectedNodes(skillTree.nodes[circledNode]).filter((n) => !n.isJewelSocket && !n.isMastery)
    : [];

  $: seedResults =
    !seed ||
    !selectedJewel ||
    !selectedConqueror ||
    Object.keys(data.TimelessJewelConquerors[selectedJewel.value]).indexOf(selectedConqueror.value) < 0
      ? []
      : affectedNodes
          .filter((n) => !!data.TreeToPassive[n.skill])
          .map((n) => ({
            node: n.skill,
            result: calculator.Calculate(
              data.TreeToPassive[n.skill].Index,
              seed,
              selectedJewel.value,
              selectedConqueror.value
            )
          }));

  let selectedStats: Record<number, StatConfig> = {};
  if (searchParams.has('stat')) {
    searchParams.getAll('stat').forEach((s) => {
      const nStat = parseInt(s);
      selectedStats[nStat] = {
        weight: 1,
        min: 0,
        id: nStat
      };
    });
  }

  let mode = searchParams.has('mode') ? searchParams.get('mode') : '';

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

  const setMode = (newMode: string) => {
    mode = newMode;
    updateUrl();
  };

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

  const allPossibleStats: { [key: string]: { [key: string]: number } } = JSON.parse(data.PossibleStats);

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
    statSelector.handleClear();
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

  let results = false;
  let minTotalWeight = 0;
  let searching = false;
  let currentSeed = 0;
  let searchResults: SearchResults;
  let searchJewel = 1;
  let searchConqueror = '';
  const search = () => {
    if (!circledNode) {
      return;
    }

    searchJewel = selectedJewel.value;
    searchConqueror = selectedConqueror.value;
    searching = true;
    searchResults = undefined;

    const query: ReverseSearchConfig = {
      jewel: selectedJewel.value,
      conqueror: selectedConqueror.value,
      nodes: affectedNodes
        .filter((n) => !disabled.has(n.skill))
        .map((n) => data.TreeToPassive[n.skill])
        .filter((n) => !!n)
        .map((n) => n.Index),
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
        results = true;
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
    affectedNodes.filter((n) => !n.isJewelSocket && !n.isMastery).forEach((n) => disabled.add(n.skill));
    // Re-assign to update svelte
    disabled = disabled;
  };

  let groupResults =
    localStorage.getItem('groupResults') === null ? true : localStorage.getItem('groupResults') === 'true';
  $: localStorage.setItem('groupResults', groupResults ? 'true' : 'false');

  type CombinedResult = {
    id: string;
    rawStat: string;
    stat: string;
    passives: number[];
  };

  export const colorKeys = {
    physical: '#c79d93',
    cast: '#b3f8fe',
    fire: '#ff9a77',
    cold: '#93d8ff',
    lightning: '#f8cb76',
    attack: '#da814d',
    life: '#c96e6e',
    chaos: '#d8a7d3',
    unique: '#af6025',
    critical: '#b2a7d6'
  };

  const colorMessage = (message: string): string => {
    Object.keys(colorKeys).forEach((key) => {
      const value = colorKeys[key];
      message = message.replace(
        new RegExp(`(${key}(?:$|\\s))|((?:^|\\s)${key})`, 'gi'),
        `<span style='color: ${value}; font-weight: bold'>$1$2</span>`
      );
    });

    return message;
  };

  const combineResults = (
    rawResults: { result: data.AlternatePassiveSkillInformation; node: number }[],
    withColors: boolean,
    only: 'notables' | 'passives' | 'all'
  ): CombinedResult[] => {
    const mappedStats: { [key: number]: number[] } = {};
    rawResults.forEach((r) => {
      if (skillTree.nodes[r.node].isKeystone) {
        return;
      }

      if (only !== 'all') {
        if (only === 'notables' && !skillTree.nodes[r.node].isNotable) {
          return;
        }

        if (only === 'passives' && skillTree.nodes[r.node].isNotable) {
          return;
        }
      }

      if (r.result.AlternatePassiveSkill && r.result.AlternatePassiveSkill.StatsKeys) {
        r.result.AlternatePassiveSkill.StatsKeys.forEach((key) => {
          mappedStats[key] = [...(mappedStats[key] || []), r.node];
        });
      }

      if (r.result.AlternatePassiveAdditionInformations) {
        r.result.AlternatePassiveAdditionInformations.forEach((info) => {
          if (info.AlternatePassiveAddition.StatsKeys) {
            info.AlternatePassiveAddition.StatsKeys.forEach((key) => {
              mappedStats[key] = [...(mappedStats[key] || []), r.node];
            });
          }
        });
      }
    });

    return Object.keys(mappedStats).map((statID) => {
      const translated = translateStat(parseInt(statID));
      return {
        stat: withColors ? colorMessage(translated) : translated,
        rawStat: translated,
        id: statID,
        passives: mappedStats[statID]
      };
    });
  };

  const sortCombined = (
    combinedResults: CombinedResult[],
    order: 'count' | 'alphabet' | 'rarity' | 'value'
  ): CombinedResult[] => {
    switch (order) {
      case 'alphabet':
        return combinedResults.sort((a, b) =>
          a.rawStat
            .replace(/[#+%]/gi, '')
            .trim()
            .toLowerCase()
            .localeCompare(b.rawStat.replace(/[#+%]/gi, '').trim().toLowerCase())
        );
      case 'count':
        return combinedResults.sort((a, b) => b.passives.length - a.passives.length);
      case 'rarity':
        return combinedResults.sort(
          (a, b) => allPossibleStats[selectedJewel.value][a.id] - allPossibleStats[selectedJewel.value][b.id]
        );
      case 'value':
        return combinedResults.sort((a, b) => {
          const aValue = statValues[a.id] || 0;
          const bValue = statValues[b.id] || 0;
          if (aValue != bValue) {
            return bValue - aValue;
          }
          return allPossibleStats[selectedJewel.value][a.id] - allPossibleStats[selectedJewel.value][b.id];
        });
    }

    return combinedResults;
  };

  const sortResults = [
    {
      label: 'Count',
      value: 'count'
    },
    {
      label: 'Alphabetical',
      value: 'alphabet'
    },
    {
      label: 'Rarity',
      value: 'rarity'
    },
    {
      label: 'Value',
      value: 'value'
    }
  ] as const;

  let sortOrder = sortResults.find((r) => r.value === (localStorage.getItem('sortOrder') || 'count'));
  $: localStorage.setItem('sortOrder', sortOrder.value);

  let colored = localStorage.getItem('colored') === null ? true : localStorage.getItem('colored') === 'true';
  $: localStorage.setItem('colored', colored ? 'true' : 'false');

  let split = localStorage.getItem('split') === null ? true : localStorage.getItem('split') === 'true';
  $: localStorage.setItem('split', split ? 'true' : 'false');

  const onPaste = (event: ClipboardEvent) => {
    if (event.type !== 'paste') {
      return;
    }

    const paste = (event.clipboardData || window.clipboardData).getData('text');
    const lines = paste.split('\n');

    if (lines.length < 14) {
      return;
    }

    const jewel = jewels.find((j) => j.label === lines[2]);
    if (!jewel) {
      return;
    }

    let newSeed: number | undefined;
    let conqueror: string | undefined;
    for (let i = 10; i < lines.length; i++) {
      conqueror = Object.keys(data.TimelessJewelConquerors[jewel.value]).find((k) => lines[i].indexOf(k) >= 0);
      if (conqueror) {
        const matches = /(\d+)/.exec(lines[i]);
        if (matches.length === 0) {
          continue;
        }

        newSeed = parseInt(matches[1]);
        break;
      }
    }

    if (!conqueror || !newSeed) {
      return;
    }

    results = false;
    mode = 'seed';
    seed = newSeed;
    selectedJewel = jewel;
    selectedConqueror = { label: conqueror, value: conqueror };
    updateUrl();
  };

  let collapsed = false;
</script>

<svelte:window on:paste={onPaste} />

<SkillTree
  {clickNode}
  {circledNode}
  selectedJewel={selectedJewel?.value}
  selectedConqueror={selectedConqueror?.value}
  {highlighted}
  {seed}
  highlightJewels={!circledNode}
  disabled={[...disabled]}>
  {#if !collapsed}
    <div
      class="w-screen md:w-10/12 lg:w-2/3 xl:w-1/2 2xl:w-5/12 3xl:w-1/3 4xl:w-1/4 absolute top-0 left-0 bg-black/80 backdrop-blur-sm themed rounded-br-lg max-h-screen">
      <div class="p-4 max-h-screen flex flex-col">
        <div class="flex flex-row justify-between mb-2">
          <div class="flex flex-row items-center">
            <button class="burger-menu mr-3" on:click={() => (collapsed = true)}>
              <div />
              <div />
              <div />
            </button>

            <h3 class="flex-grow">
              {results ? 'Results' : 'Timeless Jewel'}
            </h3>
          </div>
          {#if searchResults}
            <div class="flex flex-row">
              {#if results}
                <button
                  class="p-1 px-3 bg-blue-500/40 rounded disabled:bg-blue-900/40 mr-2"
                  on:click={() => openTrade(searchJewel, searchConqueror, searchResults.raw)}
                  disabled={!searchResults}>
                  Trade
                </button>
                <button
                  class="p-1 px-3 bg-blue-500/40 rounded disabled:bg-blue-900/40 mr-2"
                  class:grouped={groupResults}
                  on:click={() => (groupResults = !groupResults)}
                  disabled={!searchResults}>
                  Grouped
                </button>
              {/if}
              <button class="bg-neutral-100/20 px-4 p-1 rounded" on:click={() => (results = !results)}>
                {results ? 'Config' : 'Results'}
              </button>
            </div>
          {/if}
        </div>

        {#if !results}
          <Select items={jewels} bind:value={selectedJewel} on:change={changeJewel} />

          {#if selectedJewel}
            <div class="mt-4">
              <h3 class="mb-2">Conqueror</h3>
              <Select items={conquerors} bind:value={selectedConqueror} on:change={updateUrl} />
            </div>

            {#if selectedConqueror && Object.keys(data.TimelessJewelConquerors[selectedJewel.value]).indexOf(selectedConqueror.value) >= 0}
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
                    min={data.TimelessJewelSeedRanges[selectedJewel.value].Min}
                    max={data.TimelessJewelSeedRanges[selectedJewel.value].Max} />
                  {#if seed < data.TimelessJewelSeedRanges[selectedJewel.value].Min || seed > data.TimelessJewelSeedRanges[selectedJewel.value].Max}
                    <div class="mt-2">
                      Seed must be between {data.TimelessJewelSeedRanges[selectedJewel.value].Min}
                      and {data.TimelessJewelSeedRanges[selectedJewel.value].Max}
                    </div>
                  {/if}
                </div>

                {#if seed >= data.TimelessJewelSeedRanges[selectedJewel.value].Min && seed <= data.TimelessJewelSeedRanges[selectedJewel.value].Max}
                  <div class="flex flex-row mt-4 items-end">
                    <div class="flex-grow">
                      <h3 class="mb-2">Sort Order</h3>
                      <Select items={sortResults} bind:value={sortOrder} />
                    </div>
                    <div class="ml-2">
                      <button
                        class="bg-neutral-500/20 p-2 px-4 rounded"
                        class:selected={colored}
                        on:click={() => (colored = !colored)}>
                        Colors
                      </button>
                    </div>
                    <div class="ml-2">
                      <button
                        class="bg-neutral-500/20 p-2 px-4 rounded"
                        class:selected={split}
                        on:click={() => (split = !split)}>
                        Split
                      </button>
                    </div>
                  </div>

                  {#if !split}
                    <ul class="mt-4 overflow-auto" class:rainbow={colored}>
                      {#each sortCombined(combineResults(seedResults, colored, 'all'), sortOrder.value) as r}
                        <li class="cursor-pointer" on:click={() => highlight(seed, r.passives)}>
                          <span class="font-bold" class:text-white={(statValues[r.id] || 0) < 3}
                            >({r.passives.length})</span>
                          <span class="text-white">{@html r.stat}</span>
                        </li>
                      {/each}
                    </ul>
                  {:else}
                    <div class="overflow-auto mt-4">
                      <h3>Notables</h3>
                      <ul class="mt-1" class:rainbow={colored}>
                        {#each sortCombined(combineResults(seedResults, colored, 'notables'), sortOrder.value) as r}
                          <li class="cursor-pointer" on:click={() => highlight(seed, r.passives)}>
                            <span class="font-bold" class:text-white={(statValues[r.id] || 0) < 3}
                              >({r.passives.length})</span>
                            <span class="text-white">{@html r.stat}</span>
                          </li>
                        {/each}
                      </ul>

                      <h3 class="mt-2">Smalls</h3>
                      <ul class="mt-1" class:rainbow={colored}>
                        {#each sortCombined(combineResults(seedResults, colored, 'passives'), sortOrder.value) as r}
                          <li class="cursor-pointer" on:click={() => highlight(seed, r.passives)}>
                            <span class="font-bold" class:text-white={(statValues[r.id] || 0) < 3}
                              >({r.passives.length})</span>
                            <span class="text-white">{@html r.stat}</span>
                          </li>
                        {/each}
                      </ul>
                    </div>
                  {/if}
                {/if}
              {:else if mode === 'stats'}
                <div class="mt-4">
                  <h3 class="mb-2">Add Stat</h3>
                  <Select items={statItems} on:change={selectStat} bind:this={statSelector} />
                </div>
                {#if Object.keys(selectedStats).length > 0}
                  <div class="mt-4 flex flex-col overflow-auto min-h-[100px]">
                    {#each Object.keys(selectedStats) as s}
                      <div class="mb-4 flex flex-row items-start flex-col border-neutral-100/40 border-b pb-4">
                        <div>
                          <button
                            class="p-2 px-4 bg-red-500/40 rounded mr-2"
                            on:click={() => removeStat(selectedStats[s].id)}>
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
                        disabled={searching || disabled.size == 0}>
                        Select All
                      </button>
                      <button
                        class="p-2 px-2 bg-yellow-500/40 rounded disabled:bg-yellow-900/40 mr-2"
                        on:click={selectAllNotables}
                        disabled={searching || disabled.size == 0}>
                        Notables
                      </button>
                      <button
                        class="p-2 px-2 bg-yellow-500/40 rounded disabled:bg-yellow-900/40 mr-2"
                        on:click={selectAllPassives}
                        disabled={searching || disabled.size == 0}>
                        Passives
                      </button>
                      <button
                        class="p-2 px-2 bg-yellow-500/40 rounded disabled:bg-yellow-900/40 flex-grow"
                        on:click={deselectAll}
                        disabled={searching || disabled.size >= affectedNodes.length}>
                        Deselect
                      </button>
                    </div>
                    <div class="flex flex-row mt-2">
                      <button
                        class="p-2 px-3 bg-green-500/40 rounded disabled:bg-green-900/40 flex-grow"
                        on:click={() => search()}
                        disabled={searching}>
                        {#if searching}
                          {currentSeed} / {data.TimelessJewelSeedRanges[selectedJewel.value].Max}
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

        {#if searchResults && results}
          <SearchResults {searchResults} {groupResults} {highlight} jewel={searchJewel} conqueror={searchConqueror} />
        {/if}
      </div>
    </div>
  {:else}
    <button
      class="burger-menu absolute top-0 left-0 bg-black/80 backdrop-blur-sm rounded-br-lg p-4 pt-5"
      on:click={() => (collapsed = false)}>
      <div />
      <div />
      <div />
    </button>
  {/if}

  <div class="text-orange-500 absolute bottom-0 right-0 m-2">
    <a href="https://github.com/Vilsol/timeless-jewels" target="_blank" rel="noopener">Source (Github)</a>
  </div>
</SkillTree>

<style lang="postcss">
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

  .rainbow {
    animation: colorRotate 2s linear 0s infinite;
  }

  @keyframes colorRotate {
    from {
      color: hsl(0, 100%, 50%);
    }
    25% {
      color: hsl(90, 100%, 50%);
    }
    50% {
      color: hsl(180, 100%, 50%);
    }
    75% {
      color: hsl(270, 100%, 50%);
    }
    100% {
      color: hsl(359, 100%, 50%);
    }
  }
</style>
