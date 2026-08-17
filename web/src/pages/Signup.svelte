<script>
  import { useForm } from '@inertiajs/svelte'
  import AuthLayout from '../components/AuthLayout.svelte'
  import PasswordInput from '../components/PasswordInput.svelte'
  export let errors = {}
  export let site = {}
  export let flash = {}
  export let labels = {}
  export let locale = 'en'
  let form = useForm({ email: '', password: '', password_confirmation: '' })
  function submit() { form.post('/signup') }
</script>

<svelte:head>
  <title>{labels['auth.signup_title'] || 'Sign up'} · Cifra</title>
</svelte:head>

<AuthLayout {site} {flash} {labels} {locale} title={labels['home.new_workspace'] || 'New workspace'}>
  <form on:submit|preventDefault={submit} class="space-y-3">
    <div>
      <input bind:value={form.email} type="email" placeholder="Email" class="block w-full border border-paper-200/20 bg-ink-800 p-2.5 text-sm text-paper-50" />
      {#if errors.email}<p class="mt-1 text-xs text-copper-400">{errors.email}</p>{/if}
    </div>
    <div>
      <PasswordInput bind:value={form.password} name="password" autocomplete="new-password" className="block w-full border border-paper-200/20 bg-ink-800 p-2.5 text-sm text-paper-50" />
      {#if errors.password}<p class="mt-1 text-xs text-copper-400">{errors.password}</p>{/if}
    </div>
    <div>
      <PasswordInput bind:value={form.password_confirmation} name="password_confirmation" autocomplete="new-password" className="block w-full border border-paper-200/20 bg-ink-800 p-2.5 text-sm text-paper-50" />
      {#if errors.password_confirmation}<p class="mt-1 text-xs text-copper-400">{errors.password_confirmation}</p>{/if}
    </div>
    <button type="submit" class="w-full bg-copper-500 px-4 py-2.5 text-sm font-medium text-ink-950">{labels['auth.signup_submit'] || 'Sign up'}</button>
    <p class="pt-1 text-center text-xs text-paper-200">
      <a href="/login" class="text-copper-400 underline">{labels['auth.login_prompt'] || 'Already have an account?'}</a>
    </p>
  </form>
</AuthLayout>
