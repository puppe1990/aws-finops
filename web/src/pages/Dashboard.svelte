<script>
  import { router } from '@inertiajs/svelte'
  import AppLayout from '../components/AppLayout.svelte'
  import { t, tf } from '../i18n.js'

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

  function syncNow() {
    router.post('/sync')
  }

  $: burnLabel = (bps) => `${Math.round((Number(bps) || 0) / 100)}%`
</script>

<svelte:head>
  <title>{t(labels, 'dashboard.title')} · Cifra</title>
</svelte:head>

<AppLayout {site} {flash} {tenant} {tenants} {userEmail} {labels} {locale}>
  <div class="flex flex-wrap items-end justify-between gap-4">
    <div>
      <p class="font-mono text-[11px] uppercase tracking-[0.3em] text-copper-400">{t(labels, 'dash.month')}</p>
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

  <section class="mt-8 grid gap-4 md:grid-cols-3">
    <article class="border border-paper-200/15 bg-ink-900 p-5">
      <p class="font-mono text-[10px] uppercase tracking-widest text-paper-200">{t(labels, 'dash.runrate')}</p>
      <p class="mt-2 font-mono text-3xl text-copper-400">{summary.monthlyUSD || 'US$ 0,00'}</p>
      <p class="mt-2 text-xs text-paper-200">
        {summary.source === 'ce' ? t(labels, 'dash.source_ce') : t(labels, 'dash.source_estimate')}
      </p>
    </article>
    <article class="border border-paper-200/15 bg-ink-900 p-5">
      <p class="font-mono text-[10px] uppercase tracking-widest text-paper-200">{t(labels, 'dash.mtd')}</p>
      <p class="mt-2 font-mono text-3xl">{summary.mtdUSD || 'US$ 0,00'}</p>
      <p class="mt-2 text-xs text-paper-200">{tf(labels, 'dash.resources_visible', summary.resourceCount || 0)}</p>
    </article>
    <article class="border border-paper-200/15 bg-ink-900 p-5">
      <p class="font-mono text-[10px] uppercase tracking-widest text-paper-200">{t(labels, 'dash.accounts')}</p>
      <p class="mt-2 font-mono text-3xl">{summary.accountCount || 0}</p>
      <p class="mt-2 font-mono text-xs text-paper-200">
        {#each accounts as acc}
          {acc.awsAccountId} · {acc.alias}<br />
        {/each}
      </p>
    </article>
  </section>

  {#if summary.ceDenied}
    <aside class="mt-6 border border-copper-500/60 bg-copper-600/10 p-4 text-sm text-paper-50">
      {t(labels, 'dash.ce_banner')}
      <a href="/settings" class="text-copper-400 underline">{t(labels, 'nav.settings')}</a>.
    </aside>
  {/if}

  <section class="mt-10 grid gap-8 lg:grid-cols-5">
    <div class="lg:col-span-3">
      <h3 class="font-display text-2xl">{t(labels, 'dash.by_service')}</h3>
      <ul class="mt-4 divide-y divide-ink-700 border border-paper-200/15">
        {#each services as row}
          <li class="flex items-center justify-between bg-ink-900 px-4 py-3">
            <span>{row.name}</span>
            <span class="font-mono text-copper-400">{row.usd}</span>
          </li>
        {:else}
          <li class="bg-ink-900 px-4 py-6 text-sm text-paper-200">
            {t(labels, 'dash.nothing_synced')}
          </li>
        {/each}
      </ul>
    </div>
    <div class="lg:col-span-2">
      <h3 class="font-display text-2xl">{t(labels, 'dash.findings')}</h3>
      <ul class="mt-4 space-y-3">
        {#each findings as f}
          <li class="border border-paper-200/15 bg-ink-900 p-3">
            <p class="font-mono text-[10px] uppercase tracking-widest text-copper-400">
              {f.severity}
            </p>
            <p class="mt-1 text-sm">{f.title}</p>
            <p class="mt-1 text-xs text-paper-200">{f.detail}</p>
          </li>
        {:else}
          <li class="text-sm text-paper-200">{t(labels, 'dash.no_findings')}</li>
        {/each}
      </ul>
    </div>
  </section>

  <section class="mt-10">
    <h3 class="font-display text-2xl">{t(labels, 'dash.costliest')}</h3>
    <div class="mt-4 overflow-x-auto border border-paper-200/15">
      <table class="w-full text-left text-sm">
        <thead class="bg-ink-800 font-mono text-[10px] uppercase tracking-widest text-paper-200">
          <tr>
            <th class="px-3 py-2">{t(labels, 'dash.col_resource')}</th>
            <th class="px-3 py-2">{t(labels, 'dash.col_type')}</th>
            <th class="px-3 py-2">{t(labels, 'dash.col_region')}</th>
            <th class="px-3 py-2">{t(labels, 'dash.col_state')}</th>
            <th class="px-3 py-2 text-right">{t(labels, 'dash.col_month')}</th>
          </tr>
        </thead>
        <tbody>
          {#each resources.slice(0, 8) as r}
            <tr class="border-t border-ink-700 bg-ink-900">
              <td class="px-3 py-2 font-mono">{r.name}</td>
              <td class="px-3 py-2">{r.label}</td>
              <td class="px-3 py-2">{r.region}</td>
              <td class="px-3 py-2">{r.state || '—'}</td>
              <td class="px-3 py-2 text-right font-mono text-copper-400">{r.usd}</td>
            </tr>
          {:else}
            <tr><td class="px-3 py-6 text-paper-200" colspan="5">{t(labels, 'dash.no_inventory')}</td></tr>
          {/each}
        </tbody>
      </table>
    </div>
  </section>

  {#if budgets.length}
    <section class="mt-10">
      <h3 class="font-display text-2xl">{t(labels, 'dash.budgets')}</h3>
      <ul class="mt-4 grid gap-3 md:grid-cols-2">
        {#each budgets as b}
          <li class="border border-paper-200/15 bg-ink-900 p-4">
            <p class="text-sm">{b.name}</p>
            <p class="mt-1 font-mono text-xl">{b.spent} / {b.amount}</p>
            <p class="mt-1 text-xs text-paper-200">{tf(labels, 'dash.burn', burnLabel(b.burnBps))}</p>
          </li>
        {/each}
      </ul>
    </section>
  {/if}

  {#if lastSync}
    <p class="mt-8 font-mono text-[11px] text-paper-200">
      {t(labels, 'dash.last_sync')}: {lastSync.status} · {lastSync.source || '—'} · {lastSync.at || ''}
    </p>
  {/if}
  <p class="sr-only">contacts:{totalContacts} env:{env}</p>
</AppLayout>
