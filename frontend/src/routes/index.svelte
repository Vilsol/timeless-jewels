<script lang='ts'>
  import Select from 'svelte-select';
  import { browser } from '$app/env';
  import { goto } from '$app/navigation';
  import { page } from '$app/stores';
  import type { SkillTreeData } from '../lib/types';
  import { base } from '$app/paths';

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

  const skillTree: SkillTreeData = JSON.parse(window['SkillTree']);

  const passiveSkills = Object.keys(PassiveSkills).map(k => ({
    value: PassiveSkills[k],
    label: k
  })).filter(entry => {
    return true;
  });

  let selectedPassiveSkill = searchParams.has('passive_skill') ? passiveSkills.find(j => j.value == searchParams.get('passive_skill')) : undefined;

  let seed = searchParams.has('seed') ? searchParams.get('seed') : 0;

  interface Result {
    alternatePassiveAdditionInformations: {
      alternatePassiveSkillAddition: number;
      statRolls: Record<number, number>;
    }[] | undefined;
    statRolls: Record<number, number> | undefined;
    alternatePassiveSkill: number | undefined;
  }

  let result: undefined | Result;

  $: {
    if (selectedPassiveSkill && seed && selectedJewel && selectedConqueror) {
      result = Calculate(selectedPassiveSkill.value, typeof seed === 'string' ? parseInt(seed) : seed, selectedJewel.value, selectedConqueror.value);
    }
  }

  const updateUrl = () => {
    if (browser) {
      const params = {};
      selectedJewel && (params['jewel'] = selectedJewel.value);
      selectedConqueror && (params['conqueror'] = selectedConqueror.value);
      selectedPassiveSkill && (params['passive_skill'] = selectedPassiveSkill.value);
      seed && (params['seed'] = seed);

      const resultQuery = Object.keys(params)
        .map((key) => key + '=' + encodeURIComponent(params[key]))
        .join('&');

      goto($page.url.pathname + '?' + resultQuery);
    }
  };
</script>

<div class='py-10 flex flex-row justify-center w-screen h-screen'>
  <div class='flex flex-col justify-between w-1/3'>
    <div>
      <h1 class='text-white mb-10 text-center'>
        Timeless Calculator
      </h1>

      <a href='{base}/tree'>
        <h2 class='text-white mb-10 text-center underline text-orange-500'>
          Skill Tree View
        </h2>
      </a>

      <div class='themed'>
        <h3 class='mb-2'>Timeless Jewel</h3>
        <Select items={jewels} bind:value={selectedJewel} on:select={updateUrl}></Select>

        {#if selectedJewel}
          <div class='mt-4'>
            <h3 class='mb-2'>Conqueror</h3>
            <Select items={conquerors} bind:value={selectedConqueror} on:select={updateUrl}></Select>
          </div>

          {#if selectedConqueror && TimelessJewelConquerors[selectedJewel.value].indexOf(selectedConqueror.value) >= 0}
            <div class='mt-4'>
              <h3 class='mb-2'>Passive Skill</h3>
              <Select items={passiveSkills} bind:value={selectedPassiveSkill} on:select={updateUrl}></Select>
            </div>

            {#if selectedPassiveSkill}
              <div class='mt-4'>
                <h3 class='mb-2'>Seed</h3>
                <input type='number' bind:value={seed} class='seed' on:blur={updateUrl} />
              </div>

              {#if result}
                {#if 'alternatePassiveSkill' in result}
                  {@const alt = GetAlternatePassiveSkillByIndex(result['alternatePassiveSkill'])}
                  <div class='mt-4'>
                    <h3>Alternate Passive Skill</h3>
                    <span>{alt.name} ({alt.id}) ({result['alternatePassiveSkill']})</span>
                  </div>

                  {#if 'statRolls' in result && Object.keys(result['statRolls']).length > 0}
                    <ol class='mt-4 list-decimal pl-8'>
                      {#each Object.keys(result['statRolls']) as roll, i}
                        {@const stat = GetStatByIndex(alt.statsKeys[i])}
                        <li>{stat.text || '<no name>'} ({stat.id}) - {result['statRolls'][roll]}</li>
                      {/each}
                    </ol>
                  {/if}
                {/if}

                {#if 'alternatePassiveAdditionInformations' in result && result.alternatePassiveAdditionInformations.length > 0}
                  <div class='mt-4'>
                    <h3>Additions</h3>
                    <ul class='list-disc pl-8'>
                      {#each result.alternatePassiveAdditionInformations as info}
                        {@const addition = GetAlternatePassiveAdditionByIndex(info.alternatePassiveSkillAddition)}
                        <li class='mt-4'>
                          <span>{addition.id} ({addition.index})</span>

                          {#if 'statRolls' in info && Object.keys(info['statRolls']).length > 0}
                            <ol class='list-decimal pl-8'>
                              {#each Object.keys(info['statRolls']) as roll, i}
                                {@const stat = GetStatByIndex(addition.statsKeys[i])}
                                <li>{stat.text || '<no name>'} ({stat.id}) - {info['statRolls'][roll]}</li>
                              {/each}
                            </ol>
                          {/if}
                        </li>
                      {/each}
                    </ul>
                  </div>
                {/if}
              {/if}
            {/if}
          {/if}
        {/if}
      </div>
    </div>

    <div class='text-right text-orange-500'>
      <a href='https://github.com/Vilsol/timeless-jewels' target='_blank' rel='noopener'>Source (Github)</a>
    </div>
  </div>
</div>
