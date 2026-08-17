<script>
  import { useForm } from '@inertiajs/svelte'
  import AuthLayout from '../components/AuthLayout.svelte'
  export let errors = {}
  export let site = {}
  export let flash = {}
  let form = useForm({ email: '' })
  function submit() { form.post('/forgot-password') }
</script>

<svelte:head>
  <title>Forgot password · aws-finops</title>
</svelte:head>

<AuthLayout {site} {flash} title="Forgot password">
  <form on:submit|preventDefault={submit} class="space-y-3">
    <div>
      <input bind:value={form.email} type="email" placeholder="Email" class="block w-full rounded-lg border border-stone-300 p-2.5 text-sm" />
      {#if errors.email}<p class="mt-1 text-xs text-red-600">{errors.email}</p>{/if}
    </div>
    <button type="submit" class="w-full rounded-lg bg-stone-800 px-4 py-2.5 text-sm font-medium text-white hover:bg-stone-900">Send reset link</button>
    <p class="pt-1 text-center text-xs text-stone-500">
      <a href="/login" class="text-stone-700 underline hover:text-stone-900">Back to login</a>
    </p>
  </form>
</AuthLayout>
