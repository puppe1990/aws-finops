<script>
  import { router } from '@inertiajs/svelte'
  import AppLayout from '../components/AppLayout.svelte'
  import { t, tf } from '../i18n.js'

  export let policy = ''
  export let seededAccount = ''
  export let site = {}
  export let flash = {}
  export let tenant = null
  export let tenants = []
  export let userEmail = ''
  export let labels = {}
  export let locale = 'en'
</script>

<svelte:head>
  <title>{t(labels, 'set.title')} · Cifra</title>
</svelte:head>

<AppLayout {site} {flash} {tenant} {tenants} {userEmail} {labels} {locale}>
  <h2 class="font-display text-4xl">{t(labels, 'set.title')}</h2>
  <p class="mt-3 max-w-2xl text-sm leading-relaxed text-paper-200">
    {t(labels, 'set.lead')}
    {#if seededAccount}
      {tf(labels, 'set.seed', seededAccount)}
    {/if}
  </p>
  <button
    type="button"
    class="mt-6 border border-copper-500 px-4 py-2 font-mono text-xs uppercase tracking-widest text-copper-400"
    on:click={() => router.post('/sync')}
  >
    {t(labels, 'set.sync')}
  </button>
  <h3 class="mt-10 font-display text-2xl">{t(labels, 'set.policy')}</h3>
  <pre class="mt-3 overflow-x-auto border border-ink-700 bg-ink-900 p-4 font-mono text-[11px] leading-relaxed text-paper-200">{policy}</pre>
</AppLayout>
