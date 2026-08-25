import { render, screen } from '@testing-library/svelte'
import { describe, expect, it, vi } from 'vitest'
import Compare from './Compare.svelte'

vi.mock('@inertiajs/svelte', () => ({
  router: { post: vi.fn() },
  page: { url: '/compare' },
  inertia: () => ({}),
  useForm: (fields) => ({ ...fields, post: vi.fn(), processing: false, errors: {} }),
}))

const labels = {
  'cmp.title': 'Month compare',
  'cmp.by_month': 'By month',
  'cmp.services': 'This month vs last',
  'cmp.by_service': 'By service',
  'cmp.empty': 'No months yet.',
  'cmp.mom': 'vs previous month',
  'nav.ledger': 'Ledger',
  'nav.compare': 'Compare',
  'nav.resources': 'Resources',
  'nav.accounts': 'AWS accounts',
  'nav.budgets': 'Budgets',
  'nav.tenants': 'Workspaces',
  'nav.settings': 'IAM & sync',
  'auth.logout': 'Sign out',
  'theme.light': 'Light mode',
  'theme.dark': 'Dark mode',
  'lang.en': 'EN',
  'lang.pt': 'PT',
  workspace: 'Workspace',
}

describe('Compare page', () => {
  it('lists months with spend and links to the ledger', () => {
    render(Compare, {
      props: {
        labels,
        locale: 'en',
        months: [
          { query: '2026-08', label: 'Aug 2026', usd: 'US$ 19,83', cents: 1983, deltaBps: -3850, current: true },
          { query: '2026-07', label: 'Jul 2026', usd: 'US$ 32,20', cents: 3220 },
        ],
        services: [
          {
            name: 'Amazon Lightsail',
            currentUSD: 'US$ 19,83',
            previousUSD: 'US$ 32,20',
            deltaBps: -3848,
            months: [
              { query: '2026-07', cents: 3220, usd: 'US$ 32,20', pct: 100 },
              { query: '2026-08', cents: 1983, usd: 'US$ 19,83', pct: 62 },
            ],
          },
        ],
        site: { appName: 'Cifra' },
      },
    })
    expect(screen.getByText('Month compare')).toBeInTheDocument()
    expect(screen.getAllByText('Aug 2026').length).toBeGreaterThan(0)
    expect(screen.getAllByText('US$ 19,83').length).toBeGreaterThan(0)
    expect(document.querySelector('a[href="/dashboard?month=2026-07"]')).toBeTruthy()
    expect(screen.getByText('Amazon Lightsail')).toBeInTheDocument()
    expect(document.querySelectorAll('[data-testid="service-month-bar"]').length).toBe(2)
  })
})
