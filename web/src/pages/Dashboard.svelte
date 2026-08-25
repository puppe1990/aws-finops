<script>
  import { inertia, router } from '@inertiajs/svelte'
  import AppLayout from '../components/AppLayout.svelte'
  import ServiceSpendList from '../components/ServiceSpendList.svelte'
  import { t, tf } from '../i18n.js'
  import { burnPercent, findingHref, primaryAccountId, shortSync, spendRows, topFindings } from '../ledgerView.js'

  export let summary = {}
  export let services = []
  export let resources = []
  export let findings = []
  export let budgets = []
  export let accounts = []
  export let lastSync = null
  export let site = {}
  export let flash = {}
  export let tenant = null
  export let tenants = []
  export let userEmail = ''
  export let env = ''
  export let totalContacts = 0
  export let labels = {}
  export let locale = 'en'
  export let month = ''
  export let monthLabel = ''
  export let prevMonth = ''
  export let nextMonth = ''
  export let isCurrent = true

  function syncNow() {
    router.post('/sync')
  }

  $: rows = spendRows(services)
  $: shownFindings = topFindings(findings)
  $: accountId = primaryAccountId(accounts)
  $: sourceLabel = summary.source === 'ce' ? t(labels, 'dash.source_ce') : t(labels, 'dash.source_estimate')
</script>

<svelte:head>
  <title>{t(labels, 'dashboard.title')} · Cifra</title>
</svelte:head>

<AppLayout {site} {flash} {tenant} {tenants} {userEmail} {labels} {locale}>
  <div class="flex flex-wrap items-end justify-between gap-4">
    <div>
      <div
        class="inline-flex items-center gap-4 border border-ink-700 px-3 py-1.5 font-mono text-[11px] uppercase tracking-[0.3em] text-copper-400"
      >
        {#if prevMonth}
          <a href="/dashboard?month={prevMonth}" use:inertia class="hover:text-paper-50" aria-label={prevMonth}>‹</a>
        {:else}
          <span class="opacity-25">‹</span>
        {/if}
        <span data-month={month}>{monthLabel || t(labels, 'dash.month')}</span>
        {#if nextMonth}
          <a href="/dashboard?month={nextMonth}" use:inertia class="hover:text-paper-50" aria-label={nextMonth}>›</a>
        {:else}
          <span class="opacity-25">›</span>
        {/if}
      </div>
      <h2 class="mt-4 font-display text-4xl">{t(labels, 'dashboard.title')}</h2>
    </div>
    <button
      type="button"
      class="border border-copper-500 px-4 py-2 font-mono text-xs uppercase tracking-widest text-copper-400 hover:bg-copper-500 hover:text-ink-950"
      on:click={syncNow}
    >
      {t(labels, 'dash.sync')}
    </button>
  </div>

  <section class="mt-10">
    <p class="font-mono text-[10px] uppercase tracking-[0.3em] text-paper-200">{t(labels, 'dash.runrate')}</p>
    <p class="mt-2 font-display text-6xl leading-none tracking-tight text-copper-400 md:text-7xl">
      {summary.monthlyUSD || 'US$ 0,00'}
    </p>
    <div class="mt-4 flex flex-wrap gap-2">
      {#if isCurrent && summary.forecastUSD}
        <p
          class="inline-flex border border-copper-500 px-3 py-1 font-mono text-xs uppercase tracking-widest text-copper-400"
          data-forecast
        >
          {tf(labels, 'dash.forecast', summary.forecastLabel, summary.forecastUSD)}
        </p>
      {/if}
      {#each budgets as b}
        <p class="inline-flex border border-ink-700 px-3 py-1 font-mono text-xs uppercase tracking-widest text-paper-50">
          {tf(labels, 'dash.budget_line', b.name, burnPercent(b.burnBps), b.spent, b.amount)}
        </p>
      {/each}
    </div>
    <p class="mt-3 font-mono text-[11px] uppercase tracking-widest text-paper-200">
      {sourceLabel}
      {#if isCurrent && summary.source !== 'ce' && summary.mtdUSD}
        <span> · {t(labels, 'dash.mtd')} {summary.mtdUSD}</span>
      {/if}
      {#if accountId}
        <span> · {accountId}</span>
      {/if}
      {#if lastSync && lastSync.at}
        <span> · {t(labels, 'dash.last_sync')} {shortSync(lastSync.at)}</span>
      {/if}
      {#if summary.ceDenied}
        <span>
          ·
          <a href="/settings" use:inertia class="text-copper-400 underline">{t(labels, 'nav.settings')}</a>
        </span>
      {/if}
    </p>
  </section>

  <section class="mt-16">
    <h3 class="font-display text-2xl">{t(labels, 'dash.by_service')}</h3>
    <ServiceSpendList {rows} emptyLabel={t(labels, 'dash.nothing_synced')} />
  </section>

  {#if shownFindings.length}
    <section class="mt-16">
      <h3 class="font-display text-2xl">{t(labels, 'dash.findings')}</h3>
      <div class="mt-6 space-y-3">
        {#each shownFindings as f}
          <a
            href={findingHref(f.kind)}
            use:inertia
            class="block border border-paper-200/15 bg-ink-900 p-4 hover:border-copper-500"
          >
            <p class="font-mono text-[10px] uppercase tracking-widest text-copper-400">{f.severity}</p>
            <p class="mt-2 font-display text-xl leading-tight">{f.title}</p>
            <p class="mt-1 text-sm text-paper-200">{f.detail}</p>
          </a>
        {/each}
      </div>
    </section>
  {/if}

  <p class="sr-only">contacts:{totalContacts} env:{env} resources:{(resources || []).length}</p>
</AppLayout>
