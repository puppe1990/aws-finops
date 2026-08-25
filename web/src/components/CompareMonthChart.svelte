<script>
  import { inertia } from '@inertiajs/svelte'
  import { chartMonths, monthShort } from '../compareView.js'

  export let months = []
  export let currentQuery = ''

  $: cols = chartMonths(months)
</script>

{#if cols.length}
  <div class="mt-8 flex items-end gap-2 sm:gap-3">
    {#each cols as col}
      <a
        href="/dashboard?month={col.query}"
        use:inertia
        class="group flex min-w-0 flex-1 flex-col items-center gap-2"
      >
        <span class="font-mono text-[10px] {col.query === currentQuery ? 'text-copper-400' : 'text-paper-200'}">
          {col.query === currentQuery ? col.usd : ''}
        </span>
        <div class="flex w-full items-end bg-ink-800" style="height: 9rem">
          <div
            class="w-full bg-copper-400"
            data-testid="compare-bar"
            style="height: {Math.max(col.pct, 3)}%; opacity: {col.query === currentQuery ? 1 : 0.55}"
          ></div>
        </div>
        <span class="font-mono text-[10px] uppercase tracking-widest {col.query === currentQuery ? 'text-copper-400' : 'text-paper-200'}">
          {monthShort(col.label)}
        </span>
      </a>
    {/each}
  </div>
{/if}
