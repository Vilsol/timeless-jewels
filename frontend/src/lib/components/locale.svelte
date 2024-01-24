<script lang="ts">
  const selectLocale = (newLocale: string) => {
    localStorage.setItem('locale', newLocale);
    location.reload();
  };
</script>

<div class="gap-1 text-sm text-center">
  <a href="#" class="text-gray-300 hover:text-orange-400" on:click={() => selectLocale('en')}>English</a>
  <span class="text-gray-800">/</span>
  <a href="#" class="text-gray-300 hover:text-orange-400" on:click={() => selectLocale('zh')}>中文</a>
</div>
