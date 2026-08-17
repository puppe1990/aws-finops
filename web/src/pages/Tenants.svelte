<script>
  import { useForm, router } from '@inertiajs/svelte'
  import AppLayout from '../components/AppLayout.svelte'

  export let tenants = []
  export let tenant = null
  export let site = {}
  export let flash = {}
  export let userEmail = ''

  let form = useForm({ name: '', slug: '' })
  function submit() {
    form.post('/tenants')
  }
  function activate(id) {
    router.post('/tenants/switch', { tenant_id: id })
  }
</script>

<svelte:head>
  <title>Workspaces · Cifra</title>
</svelte:head>

<AppLayout {site} {flash} {tenant} {tenants} {userEmail}>
  <h2 class="font-display text-4xl">Workspaces</h2>
  <p class="mt-2 max-w-2xl text-sm text-paper-200">
    Multi-tenant: cada workspace tem suas contas, orçamentos e inventário. Nada vaza entre clientes.
  </p>

  <ul class="mt-8 divide-y divide-ink-700 border border-paper-200/15">
    {#each tenants as t}
      <li class="flex items-center justify-between bg-ink-900 px-4 py-3">
        <div>
          <p class="text-lg">{t.name}</p>
          <p class="font-mono text-xs text-paper-200">{t.slug}</p>
        </div>
        {#if tenant && t.id === tenant.id}
          <span class="font-mono text-[10px] uppercase tracking-widest text-copper-400">ativo</span>
        {:else}
          <button
            type="button"
            class="border border-paper-200/20 px-3 py-1 text-xs"
            on:click={() => activate(t.id)}>Ativar</button
          >
        {/if}
      </li>
    {/each}
  </ul>

  <form on:submit|preventDefault={submit} class="mt-10 max-w-md space-y-3 border border-paper-200/15 bg-ink-900 p-5">
    <h3 class="font-display text-2xl">Novo cliente</h3>
    <input
      bind:value={form.name}
      placeholder="Nome do workspace"
      class="w-full border border-paper-200/20 bg-ink-800 p-2.5 text-sm"
    />
    <input
      bind:value={form.slug}
      placeholder="slug (opcional)"
      class="w-full border border-paper-200/20 bg-ink-800 p-2.5 text-sm"
    />
    <button type="submit" class="bg-copper-500 px-4 py-2 text-sm text-ink-950">Criar workspace</button>
  </form>
</AppLayout>
