<script lang='ts'>
  import Select from 'svelte-select';

  const jewels = Object.keys(TimelessJewels).map(k => ({
    value: parseInt(k),
    label: TimelessJewels[k]
  }));

  let selectedJewel = undefined;

  $: conquerors = selectedJewel ? TimelessJewelConquerors[selectedJewel.value].map(k => ({
    value: k,
    label: k
  })) : [];

  let selectedConqueror = undefined;

  const passiveSkills = Object.keys(PassiveSkills).map(k => ({
    value: PassiveSkills[k],
    label: k
  }));

  let selectedPassiveSkill = undefined;

  let seed = 0;

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
      console.log(selectedPassiveSkill.value, seed, selectedJewel.value, selectedConqueror.value);
      result = Calculate(selectedPassiveSkill.value, seed, selectedJewel.value, selectedConqueror.value);
    }
  }
</script>

<div class='themed'>
  <h3 class='mb-2'>Timeless Jewel</h3>
  <Select items={jewels} bind:value={selectedJewel} on:select={() => selectedConqueror = undefined}></Select>

  {#if selectedJewel}
    <div class='mt-4'>
      <h3 class='mb-2'>Conqueror</h3>
      <Select items={conquerors} bind:value={selectedConqueror}></Select>
    </div>

    {#if selectedConqueror && TimelessJewelConquerors[selectedJewel.value].indexOf(selectedConqueror.value) >= 0}
      <div class='mt-4'>
        <h3 class='mb-2'>Passive Skill</h3>
        <Select items={passiveSkills} bind:value={selectedPassiveSkill}></Select>
      </div>

      {#if selectedPassiveSkill}
        <div class='mt-4'>
          <h3 class='mb-2'>Seed</h3>
          <input type='number' bind:value={seed} class='seed' />
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

<style>
    .themed {
        --background: #393939;
        --listBackground: #393939;
        --itemHoverBG: #c2410c;
        --itemIsActiveBG: #696969;
        --borderFocusColor: #c2410c;
        --inputColor: #ffffff;
        --border: #696969;
    }

    .seed {
        background: #393939;
        height: 42px;
        padding: 16px;
        width: 100%;
        border-radius: 3px;
    }
</style>
