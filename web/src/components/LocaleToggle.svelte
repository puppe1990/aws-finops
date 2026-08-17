<script>
  import { useForm } from '@inertiajs/svelte'
  import { onMount } from 'svelte'
  import { t } from '../i18n.js'

  export let locale = 'en'
  export let labels = {}

  let form = useForm({ locale: locale || 'en' })

  function setLocale(next) {
    if (next === locale) return
    form.locale = next
    form.post('/locale')
  }

  onMount(() => {
    document.documentElement.lang = locale === 'pt' ? 'pt-BR' : 'en'
  })
</script>

<div class="inline-flex overflow-hidden border border-ink-700" data-testid="locale-toggle">
  <button
    type="button"
    class="px-2 py-1 font-mono text-[10px] uppercase tracking-widest {locale === 'en'
      ? 'bg-copper-500 text-ink-950'
      : 'text-paper-200 hover:text-paper-50'}"
    aria-pressed={locale === 'en'}
    on:click={() => setLocale('en')}
  >
    {t(labels, 'lang.en')}
  </button>
  <button
    type="button"
    class="px-2 py-1 font-mono text-[10px] uppercase tracking-widest {locale === 'pt'
      ? 'bg-copper-500 text-ink-950'
      : 'text-paper-200 hover:text-paper-50'}"
    aria-pressed={locale === 'pt'}
    on:click={() => setLocale('pt')}
  >
    {t(labels, 'lang.pt')}
  </button>
</div>
