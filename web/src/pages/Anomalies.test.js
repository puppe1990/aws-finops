import { render, screen } from '@testing-library/svelte'
import { describe, expect, it, vi } from 'vitest'
import Anomalies from './Anomalies.svelte'

vi.mock('@inertiajs/svelte', () => ({
  router: { post: vi.fn() },
  page: { url: '/anomalies' },
  inertia: () => ({}),
  useForm: (fields) => ({ ...fields, post: vi.fn(), processing: false, errors: {} }),
}))

const labels = {
  'ano.title': 'Anomalies',
  'ano.empty': 'No anomalies in this window.',
  'ano.ce': 'Cost Explorer',
  'ano.spike': 'Spend spike',
  'ano.impact': 'Impact',
  'nav.ledger': 'Ledger',
  'nav.compare': 'Compare',
  'nav.anomalies': 'Anomalies',
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

describe('Anomalies page', () => {
  it('lists anomalies and links to the ledger month', () => {
    render(Anomalies, {
      props: {
        labels,
        locale: 'en',
        anomalies: [
          { kind: 'ce', service: 'Amazon S3', query: '2026-08', usd: 'US$ 8,00', start: '2026-08-10' },
          { kind: 'spike', service: 'Amazon Lightsail', query: '2026-08', usd: 'US$ 15,00' },
        ],
        site: { appName: 'Cifra' },
      },
    })
    expect(screen.getAllByText('Anomalies').length).toBeGreaterThan(0)
    expect(screen.getByText('Amazon S3')).toBeInTheDocument()
    expect(screen.getByText('Spend spike')).toBeInTheDocument()
    expect(document.querySelector('a[href="/dashboard?month=2026-08"]')).toBeTruthy()
  })
})
