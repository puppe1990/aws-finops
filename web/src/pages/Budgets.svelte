<script>
  import { useForm } from '@inertiajs/svelte'
  import AppLayout from '../components/AppLayout.svelte'
  import { t, tf } from '../i18n.js'

  export let budgets = []
  export let spentUSD = ''
  export let site = {}
  export let flash = {}
  export let tenant = null
  export let tenants = []
  export let userEmail = ''
  export let labels = {}
  export let locale = 'en'

  let form = useForm({ name: '', amount_usd: '' })
  function submit() {
    form.post('/budgets')
  }
</script>

<svelte:head>
  <title>{t(labels, 'bud.title')} · Cifra</title>
</svelte:head>

<AppLayout {site} {flash} {tenant} {tenants} {userEmail} {labels} {locale}>
  <h2 class="font-display text-4xl">{t(labels, 'bud.title')}</h2>
  <p class="mt-2 text-sm text-paper-200">{tf(labels, 'bud.spent', spentUSD || '—')}</p>

  <ul class="mt-8 space-y-3">
    {#each budgets as b}
      <li class="border border-paper-200/15 bg-ink-900 p-4">
        <p class="text-sm">{b.name}</p>
        <p class="font-mono text-2xl text-copper-400">{b.spent} / {b.amount}</p>
        <p class="text-xs text-paper-200">{tf(labels, 'bud.period_burn', b.period, `${Math.round((b.burnBps || 0) / 100)}%`)}</p>
      </li>
    {:else}
      <li class="text-sm text-paper-200">{t(labels, 'bud.empty')}</li>
    {/each}
  </ul>

  <form on:submit|preventDefault={submit} class="mt-10 max-w-md space-y-3 border border-paper-200/15 bg-ink-900 p-5">
    <h3 class="font-display text-2xl">{t(labels, 'bud.new')}</h3>
    <input
      bind:value={form.name}
      placeholder={t(labels, 'bud.ph_name')}
      class="w-full border border-paper-200/20 bg-ink-800 p-2.5 text-sm"
    />
    <input
      bind:value={form.amount_usd}
      placeholder={t(labels, 'bud.ph_amount')}
      class="w-full border border-paper-200/20 bg-ink-800 p-2.5 text-sm"
    />
    <button type="submit" class="bg-copper-500 px-4 py-2 text-sm text-ink-950">{t(labels, 'bud.create')}</button>
  </form>
</AppLayout>
