<script>
  export let flash = {}
</script>

{#if flash.notice}
  <p
    class="rounded-sm border border-sea-700 bg-sea-800/80 px-3 py-2 text-sm text-paper-50"
    data-testid="flash-notice"
  >
    {flash.notice}
  </p>
{/if}
{#if flash.success}
  <p
    class="rounded-sm border border-sea-700 bg-sea-800/80 px-3 py-2 text-sm text-paper-50"
    data-testid="flash-success"
  >
    {flash.success}
  </p>
{/if}
{#if flash.alert}
  <p
    class="rounded-sm border border-copper-600 bg-copper-600/20 px-3 py-2 text-sm text-copper-400"
    data-testid="flash-alert"
  >
    {flash.alert}
  </p>
{/if}
