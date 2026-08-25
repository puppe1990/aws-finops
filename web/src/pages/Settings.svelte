<script>
  import { router } from '@inertiajs/svelte'
  import AppLayout from '../components/AppLayout.svelte'
  import { t, tf } from '../i18n.js'

  export let policy = ''
  export let cloudShell = ''
  export let seededAccount = ''
  export let site = {}
  export let flash = {}
  export let tenant = null
  export let tenants = []
  export let userEmail = ''
  export let labels = {}
  export let locale = 'en'

  let copied = false

  async function copyCloudShell() {
    if (!cloudShell || !navigator.clipboard) return
    try {
      await navigator.clipboard.writeText(cloudShell)
      copied = true
      setTimeout(() => {
        copied = false
      }, 2000)
    } catch (_) {
      /* clipboard may be blocked */
    }
  }
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
  {#if cloudShell}
    <div class="mt-10 flex items-end justify-between gap-3">
      <h3 class="font-display text-2xl">{t(labels, 'set.cloudshell')}</h3>
      <button
        type="button"
        class="border border-paper-200/20 px-3 py-1.5 font-mono text-[10px] uppercase tracking-widest text-copper-400"
        on:click={copyCloudShell}
      >
        {copied ? t(labels, 'set.copied') : t(labels, 'set.copy')}
      </button>
    </div>
    <p class="mt-2 max-w-2xl text-sm leading-relaxed text-paper-200">{t(labels, 'set.cloudshell_lead')}</p>
    <pre class="mt-3 overflow-x-auto whitespace-pre-wrap break-all border border-ink-700 bg-ink-900 p-4 font-mono text-[11px] leading-relaxed text-paper-200">{cloudShell}</pre>
  {/if}
  <h3 class="mt-10 font-display text-2xl">{t(labels, 'set.policy')}</h3>
  <pre class="mt-3 overflow-x-auto border border-ink-700 bg-ink-900 p-4 font-mono text-[11px] leading-relaxed text-paper-200">{policy}</pre>
</AppLayout>
