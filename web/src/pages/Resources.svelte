<script>
  import AppLayout from '../components/AppLayout.svelte'
  import { t } from '../i18n.js'
  export let resources = []
  export let site = {}
  export let flash = {}
  export let tenant = null
  export let tenants = []
  export let userEmail = ''
  export let labels = {}
  export let locale = 'en'
</script>

<svelte:head>
  <title>{t(labels, 'res.title')} · Cifra</title>
</svelte:head>

<AppLayout {site} {flash} {tenant} {tenants} {userEmail} {labels} {locale}>
  <h2 class="font-display text-4xl">{t(labels, 'res.title')}</h2>
  <p class="mt-2 max-w-2xl text-sm text-paper-200">
    {t(labels, 'res.lead')}
  </p>
  <div class="mt-8 overflow-x-auto border border-paper-200/15">
    <table class="w-full text-left text-sm">
      <thead class="bg-ink-800 font-mono text-[10px] uppercase tracking-widest text-paper-200">
        <tr>
          <th class="px-3 py-2">{t(labels, 'dash.col_resource')}</th>
          <th class="px-3 py-2">{t(labels, 'dash.col_type')}</th>
          <th class="px-3 py-2">{t(labels, 'dash.col_region')}</th>
          <th class="px-3 py-2">{t(labels, 'dash.col_state')}</th>
          <th class="px-3 py-2 text-right">{t(labels, 'res.estimate_month')}</th>
        </tr>
      </thead>
      <tbody>
        {#each resources as r}
          <tr class="border-t border-ink-700 bg-ink-900">
            <td class="px-3 py-2 font-mono">{r.name}</td>
            <td class="px-3 py-2">{r.label}</td>
            <td class="px-3 py-2">{r.region}</td>
            <td class="px-3 py-2">{r.state || '—'}</td>
            <td class="px-3 py-2 text-right font-mono text-copper-400">{r.usd}</td>
          </tr>
        {:else}
          <tr><td class="px-3 py-6" colspan="5">{t(labels, 'res.empty')}</td></tr>
        {/each}
      </tbody>
    </table>
  </div>
</AppLayout>
