<script>
  import { inertia } from '@inertiajs/svelte'
  import LocaleToggle from '../components/LocaleToggle.svelte'
  import { t } from '../i18n.js'
  export let title = 'Cifra'
  export let site = {}
  export let labels = {}
  export let locale = 'en'
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
    <nav class="flex items-center gap-6 font-mono text-xs uppercase tracking-widest">
      <LocaleToggle {locale} {labels} />
      <a href="/login" use:inertia class="text-paper-200 hover:text-copper-400"
        >{t(labels, 'auth.login_submit')}</a
      >
      <a href="/dashboard" use:inertia class="text-paper-200 hover:text-copper-400"
        >{t(labels, 'dashboard.title')}</a
      >
    </nav>
  </header>

  <main class="relative mx-auto max-w-6xl px-6 pb-24 pt-10">
    <p class="font-mono text-[11px] uppercase tracking-[0.35em] text-copper-400">
      {labels.eyebrow || t(labels, 'home.eyebrow_empty')}
    </p>
    <h1
      class="mt-4 max-w-3xl font-display text-6xl leading-[0.92] md:text-8xl"
      data-testid="inertia-ready"
    >
      {labels.heading || t(labels, 'home.rails_heading')}
    </h1>
    <p class="mt-8 max-w-xl text-lg leading-relaxed text-paper-200">
      {labels.subtitle || t(labels, 'home.rails_subtitle')}
    </p>
    <div class="mt-10 flex flex-wrap gap-4">
      <a
        href="/login"
        use:inertia
        class="bg-copper-500 px-6 py-3 font-mono text-sm uppercase tracking-widest text-ink-950"
      >
        {t(labels, 'home.open_ledger')}
      </a>
      <a
        href="/signup"
        use:inertia
        class="border border-paper-200/30 px-6 py-3 font-mono text-sm uppercase tracking-widest text-paper-50"
      >
        {t(labels, 'home.new_workspace')}
      </a>
    </div>
    <p class="mt-16 font-mono text-xs text-paper-200">{t(labels, 'home.stack')}</p>
    <p class="sr-only">{title}</p>
  </main>
</div>
