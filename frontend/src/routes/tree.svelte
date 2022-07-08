<script lang="ts">
  import SkillTree from '../lib/components/SkillTree.svelte';
  import Select from 'svelte-select';
  import { page } from '$app/stores';
  import { browser } from '$app/env';
  import { goto } from '$app/navigation';
  import type { Node } from '../lib/types';
  import { formatStats, getAffectedNodes, inverseTranslations, passiveToTree, skillTree } from '../lib/skill_tree';
  import { syncWrap } from '../lib/worker';
  import { proxy } from 'comlink';
  import { getStat } from '../lib/skill_tree.js';

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

  const clickNode = (node: Node) => {
    if (node.isJewelSocket) {
      circledNode = node.skill;
      updateUrl();
    }
  };

  const updateUrl = () => {
    if (browser) {
      const url = new URL(window.location.origin + window.location.pathname);
      selectedJewel && url.searchParams.append('jewel', selectedJewel.value.toString());
      selectedConqueror && url.searchParams.append('conqueror', selectedConqueror.value);
      seed && url.searchParams.append('seed', seed.toString());
      circledNode && url.searchParams.append('location', circledNode.toString());
      mode && url.searchParams.append('mode', mode);

      selectedStats.forEach((s) => {
        url.searchParams.append('stat', s.toString());
      });

      goto(url.toString());
    }
  };

  let mode = searchParams.has('mode') ? searchParams.get('mode') : '';
  const setMode = (newMode: string) => {
    mode = newMode;
    updateUrl();
  };

  let allPossibleStats: { [key: string]: { [key: string]: number } };
  if (browser) {
    allPossibleStats = JSON.parse(window['PossibleStats']);
  }

  let selectedStats = new Set<number>();
  if (searchParams.has('stat')) {
    searchParams.getAll('stat').forEach((s) => selectedStats.add(parseInt(s)));
  }

  const translateStat = (id: number, roll?: number | undefined): string => {
    const stat = getStat(id);
    const translation = inverseTranslations[stat.id];
    if (roll) {
      return formatStats(translation, roll);
    }

    let translationText = stat.text || stat.id;
    if (translation && translation.English && translation.English.length) {
      translationText = translation.English[0].string;
      translation.English[0].format.forEach((f, i) => {
        translationText = translationText.replace(`{${i}}`, f);
      });
    }
    return translationText;
  };

  $: availableStats = !selectedJewel ? [] : Object.keys(allPossibleStats[selectedJewel.value]);
  $: statItems = availableStats
    .map((statId) => {
      const id = parseInt(statId);
      return {
        label: translateStat(id),
        value: id
      };
    })
    .filter((s) => !selectedStats.has(s.value));

  let statSelector: Select;
  const selectStat = (stat: CustomEvent) => {
    selectedStats = selectedStats.add(stat.detail.value);
    statSelector['handleClear']();
    updateUrl();
  };

  const removeStat = (id: number) => {
    selectedStats.delete(id);
    // Re-assign to update svelte
    selectedStats = selectedStats;
    updateUrl();
  };

  const changeJewel = () => {
    selectedStats.clear();
    updateUrl();
  };

  type SearchResults = { [key: string]: { [key: string]: { [key: string]: number } } };
  type SearchWithSeed = {
    seed: number;
    skills: {
      passive: number;
      stats: { [key: string]: number };
    }[];
  };

  let searching = false;
  let currentSeed = 0;
  let searchResult: SearchResults;
  let searchGrouped: { [key: number]: SearchWithSeed[] };
  let expandedGroup = '';
  const search = () => {
    if (!circledNode) {
      return;
    }

    searching = true;
    searchResult = undefined;

    const affectedNodes = getAffectedNodes(skillTree.nodes[circledNode])
      .filter((n) => {
        return !n.isJewelSocket && !n.isMastery;
      })
      .map((n) => window['TreeToPassive'][n.skill]);

    const args = [affectedNodes, [...selectedStats], selectedJewel.value, selectedConqueror.value];

    syncWrap
      .search(
        args,
        proxy((s) => (currentSeed = s))
      )
      .then((result) => {
        searchResult = result;
        searching = false;

        searchGrouped = {};
        Object.keys(searchResult).forEach((seed) => {
          const skills = Object.keys(searchResult[seed]).map((skillID) => {
            return {
              passive: passiveToTree[skillID],
              stats: searchResult[seed][skillID]
            };
          });

          const len = Object.keys(searchResult[seed]).length;
          searchGrouped[len] = [
            ...(searchGrouped[len] || []),
            {
              skills: skills,
              seed: parseInt(seed)
            }
          ];
        });

        expandedGroup = Object.keys(searchGrouped).sort().reverse()[0];
      });
  };

  let highlighted: number[] = [];
  const highlight = (newSeed: number, passives: SearchWithSeed['skills']) => {
    seed = newSeed;
    highlighted = passives.map((s) => s.passive);
    updateUrl();
  };
</script>

<SkillTree
  {clickNode}
  {circledNode}
  selectedJewel={selectedJewel?.value}
  selectedConqueror={selectedConqueror?.value}
  {highlighted}
  {seed}
>
  <div
    class="w-1/2 lg:w-1/3 xl:w-1/4 2xl:w-1/5 absolute top-0 left-0 bg-black/80 backdrop-blur-sm themed rounded-br-lg max-h-screen"
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
                class="seed"
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
            {#if selectedStats.size > 0}
              <div class="mt-4 flex flex-col overflow-auto min-h-[50px]">
                {#each [...selectedStats] as s}
                  <div class="mb-2 flex flex-row items-center">
                    <button class="p-2 px-4 bg-red-500/40 rounded mr-2" on:click={() => removeStat(s)}> - </button>
                    <span>{translateStat(s)}</span>
                  </div>
                {/each}
              </div>
              <button
                class="p-2 px-4 bg-green-500/40 rounded mt-2 disabled:bg-green-700/40"
                on:click={() => search()}
                disabled={searching}
              >
                {#if searching}
                  {currentSeed} / {window['TimelessJewelSeedRanges'][selectedJewel.value].max}
                {:else}
                  Search
                {/if}
              </button>

              {#if searchResult}
                <div class="mt-4 flex flex-col overflow-auto">
                  {#each Object.keys(searchGrouped).sort().reverse() as k}
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
                                Seed {set.seed}
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
    @apply rounded-l-lg border-r-2 border-black;
  }

  .selection-button:last-child {
    @apply rounded-r-lg;
  }

  .selected {
    @apply bg-neutral-100/20;
  }
</style>
