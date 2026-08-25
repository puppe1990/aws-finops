<script>
  import { inertia } from '@inertiajs/svelte'
  import AppLayout from '../components/AppLayout.svelte'
  import { t } from '../i18n.js'

  export let anomalies = []
  export let ceDenied = false
  export let site = {}
  export let flash = {}
  export let tenant = null
  export let tenants = []
  export let userEmail = ''
  export let labels = {}
  export let locale = 'en'

  function kindLabel(kind) {
    return kind === 'ce' ? t(labels, 'ano.ce') : t(labels, 'ano.spike')
  }
</script>

<svelte:head>
  <title>{t(labels, 'ano.title')} · Cifra</title>
</svelte:head>

<AppLayout {site} {flash} {tenant} {tenants} {userEmail} {labels} {locale}>
  <p class="font-mono text-[11px] uppercase tracking-[0.3em] text-copper-400">{t(labels, 'nav.anomalies')}</p>
  <h2 class="mt-1 font-display text-4xl">{t(labels, 'ano.title')}</h2>
  <p class="mt-8 font-display text-6xl leading-none tracking-tight text-copper-400">{(anomalies || []).length}</p>

  {#if ceDenied}
    <p class="mt-6 text-sm text-paper-200">
      <a href="/settings" use:inertia class="text-copper-400 underline">{t(labels, 'nav.settings')}</a>
    </p>
  {/if}

  <ul class="mt-12 space-y-3">
    {#each anomalies as a}
      <li>
        <a
          href={a.query ? `/dashboard?month=${a.query}` : '/compare'}
          use:inertia
          class="block border border-paper-200/15 bg-ink-900 p-4 hover:border-copper-500"
        >
          <p class="font-mono text-[10px] uppercase tracking-widest text-copper-400">{kindLabel(a.kind)}</p>
          <div class="mt-2 flex items-baseline justify-between gap-4">
            <p class="font-display text-xl leading-tight">{a.service || a.query}</p>
            <p class="shrink-0 font-mono text-lg text-copper-400">{a.usd}</p>
          </div>
          {#if a.start}
            <p class="mt-2 font-mono text-[11px] uppercase tracking-widest text-paper-200">{a.start}</p>
          {/if}
        </a>
      </li>
    {:else}
      <li class="border border-paper-200/15 bg-ink-900 px-4 py-6 text-sm text-paper-200">{t(labels, 'ano.empty')}</li>
    {/each}
  </ul>
</AppLayout>
