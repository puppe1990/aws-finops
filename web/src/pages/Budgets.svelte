<script>
  import { useForm } from '@inertiajs/svelte'
  import AppLayout from '../components/AppLayout.svelte'

  export let budgets = []
  export let spentUSD = ''
  export let site = {}
  export let flash = {}
  export let tenant = null
  export let tenants = []
  export let userEmail = ''

  let form = useForm({ name: '', amount_usd: '' })
  function submit() {
    form.post('/budgets')
  }
</script>

<svelte:head>
  <title>Orçamentos · Cifra</title>
</svelte:head>

<AppLayout {site} {flash} {tenant} {tenants} {userEmail}>
  <h2 class="font-display text-4xl">Orçamentos</h2>
  <p class="mt-2 text-sm text-paper-200">Gasto atual do workspace: {spentUSD || '—'}</p>

  <ul class="mt-8 space-y-3">
    {#each budgets as b}
      <li class="border border-paper-200/15 bg-ink-900 p-4">
        <p class="text-sm">{b.name}</p>
        <p class="font-mono text-2xl text-copper-400">{b.spent} / {b.amount}</p>
        <p class="text-xs text-paper-200">{b.period} · queima {Math.round((b.burnBps || 0) / 100)}%</p>
      </li>
    {:else}
      <li class="text-sm text-paper-200">Nenhum orçamento ainda.</li>
    {/each}
  </ul>

  <form on:submit|preventDefault={submit} class="mt-10 max-w-md space-y-3 border border-paper-200/15 bg-ink-900 p-5">
    <h3 class="font-display text-2xl">Novo teto</h3>
    <input
      bind:value={form.name}
      placeholder="Nome"
      class="w-full border border-paper-200/20 bg-ink-800 p-2.5 text-sm"
    />
    <input
      bind:value={form.amount_usd}
      placeholder="Valor em USD (ex. 50)"
      class="w-full border border-paper-200/20 bg-ink-800 p-2.5 text-sm"
    />
    <button type="submit" class="bg-copper-500 px-4 py-2 text-sm text-ink-950">Criar</button>
  </form>
</AppLayout>
