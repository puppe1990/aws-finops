<script>
  import { useForm } from '@inertiajs/svelte'
  import AppLayout from '../components/AppLayout.svelte'

  export let accounts = []
  export let errors = {}
  export let site = {}
  export let flash = {}
  export let tenant = null
  export let tenants = []
  export let userEmail = ''
  export let policy = ''

  let authMode = 'default_chain'
  let form = useForm({
    aws_account_id: '',
    alias: '',
    region: 'us-east-1',
    auth_mode: 'default_chain',
    access_key_id: '',
    secret_access_key: '',
  })

  function submit() {
    form.auth_mode = authMode
    form.post('/accounts')
  }
</script>

<svelte:head>
  <title>Contas AWS · Cifra</title>
</svelte:head>

<AppLayout {site} {flash} {tenant} {tenants} {userEmail}>
  <h2 class="font-display text-4xl">Contas AWS</h2>
  <p class="mt-2 max-w-2xl text-sm text-paper-200">
    Cada workspace pode ter várias contas. Use a cadeia local (`~/.aws`) ou cole access keys
    cifradas. Nada é commitado: o account ID de seed vem só de `CIFRA_SEED_AWS_ACCOUNT_ID`.
  </p>

  <ul class="mt-8 divide-y divide-ink-700 border border-paper-200/15">
    {#each accounts as acc}
      <li class="flex flex-wrap items-center justify-between gap-2 bg-ink-900 px-4 py-3">
        <div>
          <p class="font-mono text-lg text-copper-400">{acc.awsAccountId}</p>
          <p class="text-sm text-paper-200">{acc.alias} · {acc.region} · {acc.authMode}</p>
        </div>
        {#if acc.primary}
          <span class="font-mono text-[10px] uppercase tracking-widest text-copper-400"
            >principal</span
          >
        {/if}
      </li>
    {:else}
      <li class="bg-ink-900 px-4 py-6 text-sm">Nenhuma conta neste workspace.</li>
    {/each}
  </ul>

  <form on:submit|preventDefault={submit} class="mt-10 max-w-xl space-y-3 border border-paper-200/15 bg-ink-900 p-5">
    <h3 class="font-display text-2xl">Vincular outra conta</h3>
    <input
      bind:value={form.aws_account_id}
      placeholder="Account ID (12 dígitos)"
      class="w-full border border-paper-200/20 bg-ink-800 p-2.5 text-sm"
    />
    {#if errors.aws_account_id}<p class="text-xs text-copper-400">{errors.aws_account_id}</p>{/if}
    <input
      bind:value={form.alias}
      placeholder="Apelido (ex. staging)"
      class="w-full border border-paper-200/20 bg-ink-800 p-2.5 text-sm"
    />
    {#if errors.alias}<p class="text-xs text-copper-400">{errors.alias}</p>{/if}
    <input
      bind:value={form.region}
      placeholder="Região"
      class="w-full border border-paper-200/20 bg-ink-800 p-2.5 text-sm"
    />
    <select
      bind:value={authMode}
      class="w-full border border-paper-200/20 bg-ink-800 p-2.5 text-sm"
    >
      <option value="default_chain">Cadeia padrão (~/.aws)</option>
      <option value="access_keys">Access key desta conta</option>
    </select>
    {#if authMode === 'access_keys'}
      <input
        bind:value={form.access_key_id}
        placeholder="AWS_ACCESS_KEY_ID"
        class="w-full border border-paper-200/20 bg-ink-800 p-2.5 font-mono text-sm"
      />
      <input
        type="password"
        bind:value={form.secret_access_key}
        placeholder="AWS_SECRET_ACCESS_KEY"
        class="w-full border border-paper-200/20 bg-ink-800 p-2.5 font-mono text-sm"
      />
    {/if}
    <button type="submit" class="bg-copper-500 px-4 py-2 text-sm text-ink-950">Vincular</button>
  </form>

  {#if policy}
    <pre class="mt-8 overflow-x-auto border border-ink-700 bg-ink-900 p-4 font-mono text-[11px] text-paper-200">{policy}</pre>
  {/if}
</AppLayout>
