<script>
  import { inertia } from '@inertiajs/svelte'
  import AppLayout from '../components/AppLayout.svelte'
  import { t } from '../i18n.js'
  import { deltaPercent, monthBars } from '../compareView.js'

  export let months = []
  export let services = []
  export let ceDenied = false
  export let site = {}
  export let flash = {}
  export let tenant = null
  export let tenants = []
  export let userEmail = ''
  export let labels = {}
  export let locale = 'en'

  $: rows = monthBars(months)
  $: current = (months || []).find((m) => m.current) || (months || [])[0]
</script>

<svelte:head>
  <title>{t(labels, 'cmp.title')} · Cifra</title>
</svelte:head>

<AppLayout {site} {flash} {tenant} {tenants} {userEmail} {labels} {locale}>
  <h2 class="font-display text-4xl">{t(labels, 'cmp.title')}</h2>

  {#if current}
    <section class="mt-10">
      <p class="font-mono text-[10px] uppercase tracking-[0.3em] text-paper-200">{current.label}</p>
      <p class="mt-2 font-display text-6xl leading-none tracking-tight text-copper-400 md:text-7xl">
        {current.usd}
      </p>
      {#if deltaPercent(current.deltaBps)}
        <p class="mt-3 font-mono text-xs text-paper-200">
          {deltaPercent(current.deltaBps)}
          {#if current.deltaUSD}
            <span> · {current.deltaUSD}</span>
          {/if}
          <span> · {t(labels, 'cmp.mom')}</span>
        </p>
      {/if}
    </section>
  {/if}

  {#if ceDenied}
    <p class="mt-6 text-sm text-paper-200">
      <a href="/settings" use:inertia class="text-copper-400 underline">{t(labels, 'nav.settings')}</a>
    </p>
  {/if}

  <section class="mt-12">
    <h3 class="font-display text-2xl">{t(labels, 'cmp.by_month')}</h3>
    <ul class="mt-4 divide-y divide-ink-700 border border-paper-200/15">
      {#each rows as row}
        <li class="bg-ink-900">
          <a href="/dashboard?month={row.query}" use:inertia class="block px-4 py-3 hover:bg-ink-800">
            <div class="flex flex-col gap-2 sm:flex-row sm:items-center">
              <span class="sm:w-36 sm:shrink-0">{row.label}</span>
              <div class="h-1.5 w-full bg-ink-800 sm:flex-1">
                <div class="h-full bg-copper-400" data-testid="compare-bar" style="width: {row.pct}%"></div>
              </div>
              <span class="font-mono text-copper-400 sm:w-24 sm:text-right">{row.usd}</span>
              <span class="font-mono text-xs text-paper-200 sm:w-16 sm:text-right">
                {deltaPercent(row.deltaBps) || '—'}
              </span>
            </div>
          </a>
        </li>
      {:else}
        <li class="bg-ink-900 px-4 py-6 text-sm text-paper-200">{t(labels, 'cmp.empty')}</li>
      {/each}
    </ul>
  </section>

  {#if services.length}
    <section class="mt-12">
      <h3 class="font-display text-2xl">{t(labels, 'cmp.services')}</h3>
      <ul class="mt-4 divide-y divide-ink-700 border border-paper-200/15">
        {#each services as svc}
          <li class="flex flex-col gap-1 bg-ink-900 px-4 py-3 sm:flex-row sm:items-baseline sm:justify-between">
            <span>{svc.name}</span>
            <span class="font-mono text-xs text-paper-200">
              {svc.previousUSD} → <span class="text-copper-400">{svc.currentUSD}</span>
              {#if deltaPercent(svc.deltaBps)}
                <span> · {deltaPercent(svc.deltaBps)}</span>
              {/if}
            </span>
          </li>
        {/each}
      </ul>
    </section>
  {/if}
</AppLayout>
