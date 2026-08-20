<script>
  import { inertia, router } from '@inertiajs/svelte'
  import AppLayout from '../components/AppLayout.svelte'
  import ServiceSpendList from '../components/ServiceSpendList.svelte'
  import { t, tf } from '../i18n.js'
  import { burnPercent, findingHref, primaryAccountId, spendRows, topFindings } from '../ledgerView.js'

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
      <div class="flex items-center gap-3 font-mono text-[11px] uppercase tracking-[0.3em] text-copper-400">
        {#if prevMonth}
          <a href="/dashboard?month={prevMonth}" use:inertia class="text-copper-400 hover:text-paper-50" aria-label={prevMonth}>‹</a>
        {/if}
        <span>{monthLabel || t(labels, 'dash.month')}</span>
        {#if nextMonth}
          <a href="/dashboard?month={nextMonth}" use:inertia class="text-copper-400 hover:text-paper-50" aria-label={nextMonth}>›</a>
        {/if}
      </div>
      <h2 class="mt-1 font-display text-4xl">{t(labels, 'dashboard.title')}</h2>
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
    <p class="mt-3 font-mono text-xs text-paper-200">
      {sourceLabel}
      {#if isCurrent && summary.source !== 'ce' && summary.mtdUSD}
        <span> · {t(labels, 'dash.mtd')} {summary.mtdUSD}</span>
      {/if}
      {#if accountId}
        <span> · {accountId}</span>
      {/if}
      {#if lastSync && lastSync.at}
        <span> · {t(labels, 'dash.last_sync')} {lastSync.at}</span>
      {/if}
      {#if summary.ceDenied}
        <span>
          ·
          <a href="/settings" use:inertia class="text-copper-400 underline">{t(labels, 'nav.settings')}</a>
        </span>
      {/if}
    </p>
    {#each budgets as b}
      <p class="mt-2 font-mono text-xs text-paper-200">
        {tf(labels, 'dash.budget_line', b.name, burnPercent(b.burnBps), b.spent, b.amount)}
      </p>
    {/each}
  </section>

  <section class="mt-12">
    <h3 class="font-display text-2xl">{t(labels, 'dash.by_service')}</h3>
    <ServiceSpendList {rows} emptyLabel={t(labels, 'dash.nothing_synced')} />
  </section>

  {#if shownFindings.length}
    <section class="mt-10 space-y-3">
      <h3 class="font-display text-2xl">{t(labels, 'dash.findings')}</h3>
      {#each shownFindings as f}
        <a
          href={findingHref(f.kind)}
          use:inertia
          class="block border border-paper-200/15 bg-ink-900 p-3 hover:border-copper-500"
        >
          <p class="font-mono text-[10px] uppercase tracking-widest text-copper-400">{f.severity}</p>
          <p class="mt-1 text-sm">{f.title}</p>
          <p class="mt-1 text-xs text-paper-200">{f.detail}</p>
        </a>
      {/each}
    </section>
  {/if}

  <p class="sr-only">contacts:{totalContacts} env:{env} resources:{(resources || []).length}</p>
</AppLayout>
