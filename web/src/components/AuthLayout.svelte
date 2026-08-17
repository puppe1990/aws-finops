<script>
  import FlashBanner from './FlashBanner.svelte'
  import LocaleToggle from './LocaleToggle.svelte'
  import { t } from '../i18n.js'
  export let site = {}
  export let flash = {}
  export let title = ''
  export let labels = {}
  export let locale = 'en'
</script>

<div class="relative flex min-h-screen items-center justify-center bg-ink-950 px-4 py-12 text-paper-50">
  <div
    class="pointer-events-none absolute inset-0 opacity-[0.08]"
    style="background-image: radial-gradient(rgb(var(--grain) / 0.9) 0.7px, transparent 0.7px); background-size: 8px 8px;"
  ></div>
  <div class="relative w-full max-w-md">
    <div class="mb-4 flex justify-center">
      <LocaleToggle {locale} {labels} />
    </div>
    <p class="mb-2 text-center font-mono text-[10px] uppercase tracking-[0.35em] text-copper-400">
      {t(labels, 'auth.tagline')}
    </p>
    <h1 class="mb-8 text-center font-display text-5xl text-paper-50">{site.appName || 'Cifra'}</h1>
    <div class="border border-paper-200/20 bg-ink-900 p-7 shadow-[8px_8px_0_0_#c4622d]">
      {#if title}
        <h2 class="mb-5 font-display text-2xl text-paper-50">{title}</h2>
      {/if}
      <div class="mb-4"><FlashBanner {flash} /></div>
      <slot />
    </div>
  </div>
</div>
