<script>
  import { useForm } from '@inertiajs/svelte'
  import AuthLayout from '../components/AuthLayout.svelte'
  import PasswordInput from '../components/PasswordInput.svelte'
  export let errors = {}
  export let site = {}
  export let flash = {}
  export let token = ''
  export let labels = {}
  export let locale = 'en'
  let form = useForm({ token, password: '', password_confirmation: '' })
  function submit() { form.post('/reset-password') }
</script>

<svelte:head>
  <title>{labels['auth.reset_password_title'] || 'Reset password'} · Cifra</title>
</svelte:head>

<AuthLayout {site} {flash} {labels} {locale} title={labels['auth.reset_password_title'] || 'Reset password'}>
  {#if errors.token}
    <p class="mb-3 text-sm text-red-600">{errors.token}</p>
  {/if}
  <form on:submit|preventDefault={submit} class="space-y-3">
    <input type="hidden" bind:value={form.token} />
    <div>
      <PasswordInput bind:value={form.password} name="password" autocomplete="new-password" placeholder="New password" className="block w-full rounded-lg border border-stone-300 p-2.5 text-sm" />
      {#if errors.password}<p class="mt-1 text-xs text-red-600">{errors.password}</p>{/if}
    </div>
    <div>
      <PasswordInput bind:value={form.password_confirmation} name="password_confirmation" autocomplete="new-password" placeholder="Confirm password" className="block w-full rounded-lg border border-stone-300 p-2.5 text-sm" />
      {#if errors.password_confirmation}<p class="mt-1 text-xs text-red-600">{errors.password_confirmation}</p>{/if}
    </div>
    <button type="submit" class="w-full rounded-lg bg-stone-800 px-4 py-2.5 text-sm font-medium text-white hover:bg-stone-900">{labels['auth.reset_password_submit'] || 'Update password'}</button>
  </form>
</AuthLayout>
