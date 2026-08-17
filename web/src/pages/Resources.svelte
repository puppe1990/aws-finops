<script>
  import AppLayout from '../components/AppLayout.svelte'
  export let resources = []
  export let site = {}
  export let flash = {}
  export let tenant = null
  export let tenants = []
  export let userEmail = ''
</script>

<svelte:head>
  <title>Recursos · Cifra</title>
</svelte:head>

<AppLayout {site} {flash} {tenant} {tenants} {userEmail}>
  <h2 class="font-display text-4xl">Inventário</h2>
  <p class="mt-2 max-w-2xl text-sm text-paper-200">
    Tudo que o Cifra consegue ver com a IAM atual, isolado neste workspace.
  </p>
  <div class="mt-8 overflow-x-auto border border-paper-200/15">
    <table class="w-full text-left text-sm">
      <thead class="bg-ink-800 font-mono text-[10px] uppercase tracking-widest text-paper-200">
        <tr>
          <th class="px-3 py-2">Nome</th>
          <th class="px-3 py-2">Tipo</th>
          <th class="px-3 py-2">Região</th>
          <th class="px-3 py-2">Estado</th>
          <th class="px-3 py-2 text-right">Estimativa /mês</th>
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
          <tr><td class="px-3 py-6" colspan="5">Nenhum recurso. Sincronize a conta.</td></tr>
        {/each}
      </tbody>
    </table>
  </div>
</AppLayout>
