<script lang="ts">
    import { _ } from 'svelte-i18n';
    import type { SearchWithSeed } from '../skill_tree';
    import { openTrade, skillTree, translateStat } from '../skill_tree';

    export let highlight: (newSeed: number, passives: number[]) => void;
    export let set: SearchWithSeed;
    export let jewel: number;
    export let conqueror: string;
    export let platform: string;
    export let league: string;
</script>

<div
    class="my-2 border-gray-800 border p-4 flex flex-col"
    on:click={() =>
        highlight(
            set.seed,
            set.skills.map((s) => s.passive)
        )}>
    <div class="flex flex-row justify-between">
        <!-- Padding -->
        <div class="font-bold text-orange-500">
            Seed {set.seed} (weight {set.weight})
        </div>
        <div class="flex gap-1 items-center">
            <span class="text-gray-500 mr-2">{$_('Trade')}</span>
            <button
                class="px-3 text-sm bg-blue-500/40 hover:bg-blue-500/70 rounded cursor-pointer"
                on:click={() => openTrade(jewel, conqueror, [set])}>Global</button>
            <button
                class="px-3 text-sm bg-red-500/40 hover:bg-red-500/70 rounded cursor-pointer"
                on:click={() => openTrade(jewel, conqueror, [set], 'tencent')}>CN</button>
        </div>
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
