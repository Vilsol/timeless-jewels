<script lang='ts'>
  import SkillTree from '../lib/components/SkillTree.svelte';
  import Select from 'svelte-select';
  import { page } from '$app/stores';
  import { browser } from '$app/env';
  import { goto } from '$app/navigation';
  import type { Node } from '../lib/types';

  const searchParams = $page.url.searchParams;

  const jewels = Object.keys(TimelessJewels).map(k => ({
    value: parseInt(k),
    label: TimelessJewels[k]
  }));

  let selectedJewel = searchParams.has('jewel') ? jewels.find(j => j.value == searchParams.get('jewel')) : undefined;

  $: conquerors = selectedJewel ? TimelessJewelConquerors[selectedJewel.value].map(k => ({
    value: k,
    label: k
  })) : [];

  let selectedConqueror = searchParams.has('conqueror') ? {
    value: searchParams.get('conqueror'),
    label: searchParams.get('conqueror')
  } : undefined;

  let seed: number = searchParams.has('seed') ? parseInt(searchParams.get('seed')) : 0;

  let circledNode: number | undefined = searchParams.has('location') ? parseInt(searchParams.get('location')) : undefined;
  const clickNode = (node: Node) => {
    if (node.isJewelSocket) {
      circledNode = node.skill;
      updateUrl();
    }
  };

  const updateUrl = () => {
    if (browser) {
      const params = {};
      selectedJewel && (params['jewel'] = selectedJewel.value);
      selectedConqueror && (params['conqueror'] = selectedConqueror.value);
      seed && (params['seed'] = seed);
      circledNode && (params['location'] = circledNode);

      const resultQuery = Object.keys(params)
        .map((key) => key + '=' + encodeURIComponent(params[key]))
        .join('&');

      goto($page.url.pathname + '?' + resultQuery);
    }
  };
</script>

<SkillTree {clickNode} {circledNode} selectedJewel={selectedJewel?.value} selectedConqueror={selectedConqueror?.value}
           {seed}>
  <div
    class='w-screen md:w-1/2 lg:w-1/3 xl:w-1/4 2xl:w-1/5 absolute top-0 left-0 bg-black/80 backdrop-blur-sm themed p-4 rounded-br-lg'>

    <h3 class='mb-2'>Timeless Jewel</h3>
    <Select items={jewels} bind:value={selectedJewel} on:select={updateUrl}></Select>

    {#if selectedJewel}
      <div class='mt-4'>
        <h3 class='mb-2'>Conqueror</h3>
        <Select items={conquerors} bind:value={selectedConqueror} on:select={updateUrl}></Select>
      </div>

      {#if selectedConqueror && TimelessJewelConquerors[selectedJewel.value].indexOf(selectedConqueror.value) >= 0}
        <div class='mt-4'>
          <h3 class='mb-2'>Seed</h3>
          <input type='number' bind:value={seed} class='seed' on:blur={updateUrl} />
        </div>

        {#if !circledNode}
          <h2 class='mt-4'>
            Click on a jewel socket
          </h2>
        {/if}
      {/if}
    {/if}

  </div>
</SkillTree>