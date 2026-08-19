<script>
  export let rows = []
  export let emptyLabel = ''
</script>

<ul class="mt-4 divide-y divide-ink-700 border border-paper-200/15">
  {#each rows as row}
    <li class="flex flex-col gap-2 bg-ink-900 px-4 py-3 sm:flex-row sm:items-center">
      <span class="sm:w-52 sm:shrink-0">{row.name}</span>
      <div class="h-1.5 w-full bg-ink-800 sm:flex-1">
        <!-- width is a runtime percent; Tailwind cannot express it -->
        <div
          class="h-full bg-copper-400"
          data-testid="spend-bar"
          style="width: {row.pct}%"
        ></div>
      </div>
      <span class="font-mono text-copper-400 sm:w-24 sm:text-right">{row.usd}</span>
    </li>
  {:else}
    <li class="bg-ink-900 px-4 py-6 text-sm text-paper-200">{emptyLabel}</li>
  {/each}
</ul>
