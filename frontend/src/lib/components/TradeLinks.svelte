<script lang="ts">
  import { openQueryTrade, type Query } from '$lib/utils/trade_utils';

  // queries to display trade link buttons for
  export let queries: Query[];
</script>

<div class="flex flex-wrap gap-2.5 my-2">
  {#each queries as query, index}
    <button class="p-1 px-3 bg-blue-500/40 rounded" on:click={() => openQueryTrade(query)}>
      Trade Batch {index + 1}
    </button>
  {/each}
</div>
