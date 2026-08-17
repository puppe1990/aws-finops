<script>
  import { useForm } from '@inertiajs/svelte'
  import AppLayout from '../components/AppLayout.svelte'
  import { t } from '../i18n.js'

  export let accounts = []
  export let errors = {}
  export let site = {}
  export let flash = {}
  export let tenant = null
  export let tenants = []
  export let userEmail = ''
  export let policy = ''
  export let labels = {}
  export let locale = 'en'

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
  <title>{t(labels, 'acc.title')} · Cifra</title>
</svelte:head>

<AppLayout {site} {flash} {tenant} {tenants} {userEmail} {labels} {locale}>
  <h2 class="font-display text-4xl">{t(labels, 'acc.title')}</h2>
  <p class="mt-2 max-w-2xl text-sm text-paper-200">
    {t(labels, 'acc.lead')}
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
            >{t(labels, 'acc.primary')}</span
          >
        {/if}
      </li>
    {:else}
      <li class="bg-ink-900 px-4 py-6 text-sm">{t(labels, 'acc.empty')}</li>
    {/each}
  </ul>

  <form on:submit|preventDefault={submit} class="mt-10 max-w-xl space-y-3 border border-paper-200/15 bg-ink-900 p-5">
    <h3 class="font-display text-2xl">{t(labels, 'acc.link_other')}</h3>
    <input
      bind:value={form.aws_account_id}
      placeholder={t(labels, 'acc.ph_id')}
      class="w-full border border-paper-200/20 bg-ink-800 p-2.5 text-sm"
    />
    {#if errors.aws_account_id}<p class="text-xs text-copper-400">{errors.aws_account_id}</p>{/if}
    <input
      bind:value={form.alias}
      placeholder={t(labels, 'acc.ph_alias')}
      class="w-full border border-paper-200/20 bg-ink-800 p-2.5 text-sm"
    />
    {#if errors.alias}<p class="text-xs text-copper-400">{errors.alias}</p>{/if}
    <input
      bind:value={form.region}
      placeholder={t(labels, 'acc.ph_region')}
      class="w-full border border-paper-200/20 bg-ink-800 p-2.5 text-sm"
    />
    <select
      bind:value={authMode}
      class="w-full border border-paper-200/20 bg-ink-800 p-2.5 text-sm"
    >
      <option value="default_chain">{t(labels, 'acc.mode_chain')}</option>
      <option value="access_keys">{t(labels, 'acc.mode_keys')}</option>
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
    <button type="submit" class="bg-copper-500 px-4 py-2 text-sm text-ink-950">{t(labels, 'acc.submit')}</button>
  </form>

  {#if policy}
    <pre class="mt-8 overflow-x-auto border border-ink-700 bg-ink-900 p-4 font-mono text-[11px] text-paper-200">{policy}</pre>
  {/if}
</AppLayout>
