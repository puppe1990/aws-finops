<script>
  import { inertia } from '@inertiajs/svelte'
  export let title = 'Cifra'
  export let site = {}
  export let labels = {}
</script>

<svelte:head>
  <title>Cifra · FinOps AWS</title>
</svelte:head>

<div class="relative min-h-screen overflow-hidden bg-ink-950 text-paper-50">
  <div
    class="pointer-events-none absolute inset-0 opacity-[0.08]"
    style="background-image: radial-gradient(rgb(var(--grain) / 0.9) 0.8px, transparent 0.8px); background-size: 9px 9px;"
  ></div>
  <header class="relative mx-auto flex max-w-6xl items-center justify-between px-6 py-6">
    <p class="font-display text-2xl">{site.appName || 'Cifra'}</p>
    <nav class="flex gap-6 font-mono text-xs uppercase tracking-widest">
      <a href="/login" use:inertia class="text-paper-200 hover:text-copper-400"
        >{labels.login || 'Entrar'}</a
      >
      <a href="/dashboard" use:inertia class="text-paper-200 hover:text-copper-400"
        >{labels.dashboard || 'Livro-caixa'}</a
      >
    </nav>
  </header>

  <main class="relative mx-auto max-w-6xl px-6 pb-24 pt-10">
    <p class="font-mono text-[11px] uppercase tracking-[0.35em] text-copper-400">
      {#if labels.account}
        Multi-tenant · conta {labels.account} já vinculada
      {:else}
        Multi-tenant · vincule uma conta AWS no workspace
      {/if}
    </p>
    <h1
      class="mt-4 max-w-3xl font-display text-6xl leading-[0.92] md:text-8xl"
      data-testid="inertia-ready"
    >
      {labels.heading || 'O livro-caixa da AWS.'}
    </h1>
    <p class="mt-8 max-w-xl text-lg leading-relaxed text-paper-200">
      {labels.subtitle ||
        'Cifra lê o inventário da sua conta, estima o burn mensal e isola cada cliente num workspace. Cost Explorer entra quando a IAM deixar.'}
    </p>
    <div class="mt-10 flex flex-wrap gap-4">
      <a
        href="/login"
        use:inertia
        class="bg-copper-500 px-6 py-3 font-mono text-sm uppercase tracking-widest text-ink-950"
      >
        Abrir o livro
      </a>
      <a
        href="/signup"
        use:inertia
        class="border border-paper-200/30 px-6 py-3 font-mono text-sm uppercase tracking-widest text-paper-50"
      >
        Novo workspace
      </a>
    </div>
    <p class="mt-16 font-mono text-xs text-paper-200">{labels.stack || 'Cais · Go · Svelte · SQLite'}</p>
    <p class="sr-only">{title}</p>
  </main>
</div>
