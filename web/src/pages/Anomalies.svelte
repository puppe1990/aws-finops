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
  <h2 class="font-display text-4xl">{t(labels, 'ano.title')}</h2>
  <p class="mt-2 font-mono text-xs text-paper-200">{(anomalies || []).length}</p>

  {#if ceDenied}
    <p class="mt-6 text-sm text-paper-200">
      <a href="/settings" use:inertia class="text-copper-400 underline">{t(labels, 'nav.settings')}</a>
    </p>
  {/if}

  <ul class="mt-10 divide-y divide-ink-700 border border-paper-200/15">
    {#each anomalies as a}
      <li class="bg-ink-900">
        <a
          href={a.query ? `/dashboard?month=${a.query}` : '/compare'}
          use:inertia
          class="block px-4 py-3 hover:bg-ink-800"
        >
          <p class="font-mono text-[10px] uppercase tracking-widest text-copper-400">{kindLabel(a.kind)}</p>
          <p class="mt-1 text-sm">{a.service || a.query}</p>
          <p class="mt-1 font-mono text-xs text-paper-200">
            {t(labels, 'ano.impact')}
            {a.usd}
            {#if a.start}
              <span> · {a.start}</span>
            {/if}
          </p>
        </a>
      </li>
    {:else}
      <li class="bg-ink-900 px-4 py-6 text-sm text-paper-200">{t(labels, 'ano.empty')}</li>
    {/each}
  </ul>
</AppLayout>
