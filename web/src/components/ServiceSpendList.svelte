<script>
  import { usageLabel } from '../ledgerView.js'

  export let rows = []
  export let emptyLabel = ''
</script>

<ul class="mt-6 space-y-8">
  {#each rows as row}
    <li class="border-b border-paper-200/10 pb-8 last:border-b-0 last:pb-0">
      <div class="flex items-baseline justify-between gap-4">
        <h4 class="font-display text-xl leading-tight">{row.name}</h4>
        <p class="shrink-0 font-mono text-lg text-copper-400">{row.usd}</p>
      </div>
      <div class="mt-4 w-full bg-ink-800" style="height: 0.45rem">
        <div class="h-full bg-copper-400" data-testid="spend-bar" style="width: {row.pct}%"></div>
      </div>
      {#if row.details && row.details.length}
        <ul class="mt-4 space-y-1.5">
          {#each row.details as detail}
            <li class="flex justify-between gap-4 font-mono text-[11px] text-paper-200">
              <span class="truncate">{usageLabel(detail.name)}</span>
              <span class="shrink-0 text-copper-400">{detail.usd}</span>
            </li>
          {/each}
        </ul>
      {/if}
    </li>
  {:else}
    <li class="text-sm text-paper-200">{emptyLabel}</li>
  {/each}
</ul>
