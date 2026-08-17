<script>
  import { router } from '@inertiajs/svelte'
  import AppLayout from '../components/AppLayout.svelte'

  export let policy = ''
  export let seededAccount = ''
  export let site = {}
  export let flash = {}
  export let tenant = null
  export let tenants = []
  export let userEmail = ''
</script>

<svelte:head>
  <title>IAM · Cifra</title>
</svelte:head>

<AppLayout {site} {flash} {tenant} {tenants} {userEmail}>
  <h2 class="font-display text-4xl">IAM & sync</h2>
  <p class="mt-3 max-w-2xl text-sm leading-relaxed text-paper-200">
    Sem <span class="font-mono">ce:GetCostAndUsage</span> o Cifra estima o gasto pelo inventário
    (Lightsail + S3). Cole a policy abaixo numa role de leitura de FinOps.
    {#if seededAccount}
      Seed local: <span class="font-mono text-copper-400">{seededAccount}</span>
      (só via env, não vai no repositório).
    {/if}
  </p>
  <button
    type="button"
    class="mt-6 border border-copper-500 px-4 py-2 font-mono text-xs uppercase tracking-widest text-copper-400"
    on:click={() => router.post('/sync')}
  >
    Sincronizar agora
  </button>
  <h3 class="mt-10 font-display text-2xl">Policy mínima</h3>
  <pre class="mt-3 overflow-x-auto border border-ink-700 bg-ink-900 p-4 font-mono text-[11px] leading-relaxed text-paper-200">{policy}</pre>
</AppLayout>
