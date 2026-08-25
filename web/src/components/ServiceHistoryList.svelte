<script>
  import { deltaPercent, deltaTone, monthShort } from '../compareView.js'

  export let services = []

  function chipClass(bps) {
    const tone = deltaTone(bps)
    if (tone === 'down') return 'text-paper-50'
    if (tone === 'up') return 'text-copper-400'
    return 'text-paper-200'
  }
</script>

<ul class="mt-6 space-y-6">
  {#each services as svc}
    <li class="border-b border-paper-200/10 pb-6 last:border-b-0 last:pb-0">
      <div class="flex items-baseline justify-between gap-4">
        <h4 class="font-display text-xl leading-tight">{svc.name}</h4>
        <div class="shrink-0 text-right">
          <p class="font-mono text-lg text-copper-400">{svc.currentUSD}</p>
          <p class="font-mono text-[11px] uppercase tracking-widest {chipClass(svc.deltaBps)}">
            {deltaPercent(svc.deltaBps) || '—'}
          </p>
        </div>
      </div>
      {#if svc.months && svc.months.length}
        <div class="mt-4 flex items-end gap-1">
          {#each svc.months as m}
            <div class="flex min-w-0 flex-1 flex-col items-center gap-1">
              <div
                class="flex w-full items-end bg-ink-800"
                style="height: 3.25rem"
                title="{m.query} {m.usd}"
                data-testid="service-month-bar"
              >
                <div class="w-full bg-copper-400" style="height: {Math.max(m.pct, 4)}%"></div>
              </div>
              <span class="hidden font-mono text-[9px] uppercase tracking-wider text-paper-200 sm:block">
                {monthShort(m.query)}
              </span>
            </div>
          {/each}
        </div>
      {/if}
    </li>
  {/each}
</ul>
