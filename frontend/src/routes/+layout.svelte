<script lang="ts">
  import '../i18n';

  import { browser } from '$app/environment';
  import { assets } from '$app/paths';
  import { _, locale } from 'svelte-i18n';
  import '../app.scss';
  import { loadSkillTree } from '../lib/skill_tree';
  import { initializeCrystalline } from '../lib/types';
  import { syncWrap } from '../lib/worker';
  import '../wasm_exec.js';

  let wasmLoading = true;

  // eslint-disable-next-line no-undef
  const go = new Go();

  if (browser) {
    const currentLocale = localStorage.getItem('locale') || navigator.language || 'en';
    fetch(assets + '/calculator.wasm')
      .then((data) => data.arrayBuffer())
      .then((data) => {
        WebAssembly.instantiate(data, go.importObject).then((result) => {
          go.run(result.instance);
          wasmLoading = false;
          initializeCrystalline();

          loadSkillTree(currentLocale);
        });

        syncWrap.boot(data);
      });

    document.title = $_('app_name');

    locale.set(currentLocale);
  }
</script>

{#if wasmLoading}
  <div class="flex flex-row justify-center h-screen">
    <div class="flex flex-col">
      <div class="py-10 flex flex-col justify-between">
        <div>
          <h1 class="text-white mb-10 text-center">{$_('app_name')}</h1>

          <h2 class="text-center">{$_('loading')}</h2>
        </div>
      </div>
    </div>
  </div>
{:else}
  <slot />
{/if}
