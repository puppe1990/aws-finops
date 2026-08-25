<script>
  import { inertia, page, router } from '@inertiajs/svelte'
  import FlashBanner from './FlashBanner.svelte'
  import LocaleToggle from './LocaleToggle.svelte'
  import { t } from '../i18n.js'
  import { readTheme, toggleTheme } from '../theme.js'

  export let site = {}
  export let flash = {}
  export let labels = {}
  export let locale = 'en'
  export let tenant = null
  export let tenants = []
  export let userEmail = ''

  let theme = typeof document === 'undefined' ? 'dark' : readTheme()

  function flipTheme() {
    theme = toggleTheme(theme)
  }

  $: nav = [
    { href: '/dashboard', label: t(labels, 'nav.ledger') },
    { href: '/compare', label: t(labels, 'nav.compare') },
    { href: '/anomalies', label: t(labels, 'nav.anomalies') },
    { href: '/resources', label: t(labels, 'nav.resources') },
    { href: '/accounts', label: t(labels, 'nav.accounts') },
    { href: '/budgets', label: t(labels, 'nav.budgets') },
    { href: '/tenants', label: t(labels, 'nav.tenants') },
    { href: '/settings', label: t(labels, 'nav.settings') },
  ]

  function isActive(href) {
    const url = page.url || ''
    const path = url.split('?')[0]
    if (href === '/dashboard') return path === '/dashboard'
    return path === href || path.startsWith(href + '/')
  }

  function logout() {
    router.post('/logout')
  }

  function switchTenant(event) {
    const tenant_id = event.target.value
    if (!tenant_id) return
    router.post('/tenants/switch', { tenant_id })
  }
</script>

<div class="min-h-screen bg-ink-950 text-paper-50">
  <div
    class="pointer-events-none fixed inset-0 opacity-[0.07]"
    style="background-image: radial-gradient(rgb(var(--grain) / 0.9) 0.6px, transparent 0.6px); background-size: 7px 7px;"
  ></div>

  <div class="relative flex min-h-screen flex-col md:flex-row">
    <aside
      class="flex w-full shrink-0 flex-col border-b border-ink-700 bg-ink-900 md:w-64 md:border-b-0 md:border-r"
    >
      <div class="flex items-end justify-between border-b border-ink-700 px-5 py-4 md:block md:px-5 md:py-6">
        <a href="/dashboard" use:inertia class="block">
          <p class="font-mono text-[10px] uppercase tracking-[0.28em] text-copper-400">
            FinOps AWS
          </p>
          <h1 class="mt-1 font-display text-3xl leading-none text-paper-50">
            {site.appName || 'Cifra'}
          </h1>
        </a>
        {#if tenant}
          <p class="mt-3 hidden font-mono text-xs text-paper-200 md:block">{tenant.name}</p>
        {/if}
        <div class="flex items-center gap-3 md:hidden">
          <LocaleToggle {locale} {labels} />
          <button
            type="button"
            class="font-mono text-[11px] uppercase tracking-widest text-paper-200"
            data-testid="theme-toggle-mobile"
            aria-pressed={theme === 'light'}
            on:click={flipTheme}
          >
            {theme === 'light' ? t(labels, 'theme.dark') : t(labels, 'theme.light')}
          </button>
          <button
            type="button"
            class="font-mono text-[11px] uppercase tracking-widest text-paper-200"
            on:click={logout}
          >
            {t(labels, 'auth.logout')}
          </button>
        </div>
      </div>

      <nav class="flex gap-1 overflow-x-auto px-3 py-3 text-sm md:flex-1 md:flex-col md:space-y-0.5 md:overflow-visible md:py-5">
        {#each nav as item}
          <a
            href={item.href}
            use:inertia
            class="whitespace-nowrap rounded-sm px-3 py-2 tracking-wide {isActive(item.href)
              ? 'bg-copper-500 text-ink-950'
              : 'text-paper-200 hover:bg-ink-800 hover:text-paper-50'}"
            aria-current={isActive(item.href) ? 'page' : undefined}
          >
            {item.label}
          </a>
        {/each}
      </nav>

      <div class="hidden space-y-3 border-t border-ink-700 p-4 md:block">
        {#if tenants?.length > 1}
          <label class="block font-mono text-[10px] uppercase tracking-widest text-paper-200">
            {t(labels, 'workspace')}
            <select
              class="mt-1 w-full rounded-sm border border-ink-700 bg-ink-800 px-2 py-1.5 text-xs text-paper-50"
              on:change={switchTenant}
            >
              {#each tenants as ws}
                <option value={ws.id} selected={tenant && ws.id === tenant.id}>{ws.name}</option>
              {/each}
            </select>
          </label>
        {/if}
        {#if userEmail}
          <p class="truncate font-mono text-[11px] text-paper-200">{userEmail}</p>
        {/if}
        <LocaleToggle {locale} {labels} />
        <button
          type="button"
          class="flex w-full items-center justify-between rounded-sm border border-ink-700 px-3 py-2 text-left text-sm text-paper-200 hover:border-copper-500 hover:text-copper-400"
          data-testid="theme-toggle"
          aria-pressed={theme === 'light'}
          on:click={flipTheme}
        >
          {theme === 'light' ? t(labels, 'theme.dark') : t(labels, 'theme.light')}
        </button>
        <button
          type="button"
          class="w-full rounded-sm border border-ink-700 px-3 py-2 text-left text-sm text-paper-200 hover:border-copper-500 hover:text-copper-400"
          data-testid="logout-button"
          on:click={logout}
        >
          {t(labels, 'auth.logout')}
        </button>
      </div>
    </aside>

    <main class="min-w-0 flex-1 px-6 py-8 md:px-10">
      <div class="mb-6 space-y-2">
        <FlashBanner {flash} />
      </div>
      <slot />
    </main>
  </div>
</div>
