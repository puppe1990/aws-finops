<script>
  import { useForm } from '@inertiajs/svelte'
  import AuthLayout from '../components/AuthLayout.svelte'
  import PasswordInput from '../components/PasswordInput.svelte'
  export let errors = {}
  export let site = {}
  export let flash = {}
  let form = useForm({ email: 'demo@example.com', password: 'password' })
  function submit() {
    form.post('/login')
  }
</script>

<svelte:head>
  <title>Entrar · Cifra</title>
</svelte:head>

<AuthLayout {site} {flash} title="Entrar no livro">
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
      Entrar
    </button>
    <p class="pt-1 text-center text-xs text-paper-200">
      Demo: demo@example.com / password
      <br />
      <a href="/signup" class="text-copper-400 underline">Criar conta</a>
      ·
      <a href="/forgot-password" class="text-copper-400 underline">Esqueci a senha</a>
    </p>
  </form>
</AuthLayout>
