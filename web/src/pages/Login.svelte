<script>
  import { useForm } from '@inertiajs/svelte'
  import AuthLayout from '../components/AuthLayout.svelte'
  import PasswordInput from '../components/PasswordInput.svelte'
  export let errors = {}
  export let site = {}
  export let flash = {}
  export let labels = {}
  export let locale = 'en'
  let form = useForm({ email: 'demo@example.com', password: 'password' })
  function submit() {
    form.post('/login')
  }
</script>

<svelte:head>
  <title>{labels['auth.login_title'] || 'Sign in'} · Cifra</title>
</svelte:head>

<AuthLayout {site} {flash} {labels} {locale} title={labels['auth.login_book'] || 'Sign in'}>
  <form on:submit|preventDefault={submit} class="space-y-3">
    <div>
      <input
        type="email"
        bind:value={form.email}
        placeholder="Email"
        class="block w-full border border-paper-200/20 bg-ink-800 p-2.5 text-sm text-paper-50 outline-none focus:border-copper-500"
      />
      {#if errors.email}<p class="mt-1 text-xs text-copper-400">{errors.email}</p>{/if}
    </div>
    <div>
      <PasswordInput
        bind:value={form.password}
        name="password"
        autocomplete="current-password"
        className="block w-full border border-paper-200/20 bg-ink-800 p-2.5 text-sm text-paper-50"
      />
    </div>
    <button
      type="submit"
      class="w-full bg-copper-500 px-4 py-2.5 text-sm font-medium tracking-wide text-ink-950 hover:bg-copper-400"
    >
      {labels['auth.login_submit'] || 'Sign in'}
    </button>
    <p class="pt-1 text-center text-xs text-paper-200">
      {labels['auth.demo_hint'] || 'Demo: demo@example.com / password'}
      <br />
      <a href="/signup" class="text-copper-400 underline">{labels['auth.signup_prompt'] || 'Create account'}</a>
      ·
      <a href="/forgot-password" class="text-copper-400 underline">{labels['auth.forgot_password'] || 'Forgot password?'}</a>
    </p>
  </form>
</AuthLayout>
