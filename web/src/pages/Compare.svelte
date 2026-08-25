<script>
  import { inertia } from '@inertiajs/svelte'
  import AppLayout from '../components/AppLayout.svelte'
  import CompareMonthChart from '../components/CompareMonthChart.svelte'
  import ServiceHistoryList from '../components/ServiceHistoryList.svelte'
  import { t } from '../i18n.js'
  import { deltaPercent, deltaTone } from '../compareView.js'

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

  $: current = (months || []).find((m) => m.current) || (months || [])[0]
  $: tone = current ? deltaTone(current.deltaBps) : 'flat'
</script>

<svelte:head>
  <title>{t(labels, 'cmp.title')} · Cifra</title>
</svelte:head>

<AppLayout {site} {flash} {tenant} {tenants} {userEmail} {labels} {locale}>
  <p class="font-mono text-[11px] uppercase tracking-[0.3em] text-copper-400">{t(labels, 'nav.compare')}</p>
  <h2 class="mt-1 font-display text-4xl">{t(labels, 'cmp.title')}</h2>

  {#if current}
    <section class="mt-10">
      <p class="font-mono text-[10px] uppercase tracking-[0.3em] text-paper-200">{current.label}</p>
      <p class="mt-2 font-display text-6xl leading-none tracking-tight text-copper-400 md:text-7xl">
        {current.usd}
      </p>
      {#if deltaPercent(current.deltaBps)}
        <p
          class="mt-4 inline-flex border px-3 py-1 font-mono text-xs uppercase tracking-widest {tone === 'up'
            ? 'border-copper-500 text-copper-400'
            : 'border-ink-700 text-paper-50'}"
        >
          {deltaPercent(current.deltaBps)}
          {#if current.deltaUSD}
            <span> · {current.deltaUSD}</span>
          {/if}
        </p>
        <p class="mt-2 font-mono text-[11px] uppercase tracking-widest text-paper-200">{t(labels, 'cmp.mom')}</p>
      {/if}
    </section>
  {/if}

  {#if ceDenied}
    <p class="mt-6 text-sm text-paper-200">
      <a href="/settings" use:inertia class="text-copper-400 underline">{t(labels, 'nav.settings')}</a>
    </p>
  {/if}

  <section class="mt-14">
    <h3 class="font-display text-2xl">{t(labels, 'cmp.by_month')}</h3>
    <CompareMonthChart {months} currentQuery={current && current.query} />
  </section>

  {#if services.length}
    <section class="mt-16">
      <h3 class="font-display text-2xl">{t(labels, 'cmp.by_service')}</h3>
      <ServiceHistoryList {services} />
    </section>
  {:else if !current}
    <p class="mt-10 text-sm text-paper-200">{t(labels, 'cmp.empty')}</p>
  {/if}
</AppLayout>
