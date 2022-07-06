<script lang='ts'>
  import '../app.scss';
  import '../wasm_exec.js';
  import { assets } from '$app/paths';
  import { browser } from '$app/env';

  let wasmLoading = true;

  // eslint-disable-next-line no-undef
  const go = new Go();

  if (browser) {
    WebAssembly.instantiateStreaming(fetch(assets + '/calculator.wasm'), go.importObject).then((result) => {
      go.run(result.instance);
      wasmLoading = false;
    });
  }
</script>

{#if wasmLoading}
  <div class='flex flex-row justify-center h-screen'>
    <div class='flex flex-col'>
      <div class='py-10 flex flex-col justify-between'>
        <div>
          <h1 class='text-white mb-10 text-center'>
            Timeless Calculator
          </h1>

          <h2 class='text-center'>
            Loading...
          </h2>
        </div>
      </div>
    </div>
  </div>
{:else}
  <slot />
{/if}