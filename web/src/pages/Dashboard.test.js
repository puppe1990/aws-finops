import { render, screen } from '@testing-library/svelte'
import { describe, expect, it, vi } from 'vitest'
import Dashboard from './Dashboard.svelte'

vi.mock('@inertiajs/svelte', () => ({
  router: { post: vi.fn() },
  page: { url: '/dashboard' },
  inertia: () => ({}),
  useForm: (fields) => ({ ...fields, post: vi.fn(), processing: false, errors: {} }),
}))

const labels = {
  'dashboard.title': 'Ledger',
  'dash.month': 'Current month',
  'dash.sync': 'Sync AWS',
  'dash.runrate': 'Monthly run-rate',
  'dash.source_ce': 'Source: Cost Explorer',
  'dash.source_estimate': 'Source: inventory estimate',
  'dash.mtd': 'Month-to-date',
  'dash.by_service': 'By service',
  'dash.nothing_synced': 'Nothing synced yet.',
  'dash.findings': 'Findings',
  'dash.budget_line': '%s · burn %s · %s / %s',
  'dash.last_sync': 'Last sync',
  'nav.settings': 'IAM & sync',
  'res.title': 'Inventory',
  'nav.ledger': 'Ledger',
  'nav.resources': 'Resources',
  'nav.accounts': 'AWS accounts',
  'nav.budgets': 'Budgets',
  'nav.tenants': 'Workspaces',
  workspace: 'Workspace',
  'auth.logout': 'Sign out',
  'theme.light': 'Light mode',
  'theme.dark': 'Dark mode',
  'lang.en': 'EN',
  'lang.pt': 'PT',
}

function baseProps(extra = {}) {
  return {
    labels,
    locale: 'en',
    summary: { monthlyUSD: 'US$ 32,15', source: 'ce', ceDenied: false, mtdUSD: 'US$ 32,15' },
    services: [
      { name: 'Amazon Lightsail', cents: 1947, usd: 'US$ 19,47' },
      { name: 'AWS Glue', cents: 0, usd: 'US$ 0,00' },
    ],
    findings: [],
    budgets: [],
    accounts: [{ awsAccountId: '840298254452', alias: 'principal' }],
    lastSync: { at: '2026-08-19T22:33:21Z', status: 'ok', source: 'ce' },
    resources: [{ name: 'should-not-show' }],
    site: { appName: 'Cifra' },
    ...extra,
  }
}

describe('Dashboard ledger', () => {
  it('shows the run-rate and hides the costliest table', () => {
    render(Dashboard, { props: baseProps() })
    expect(screen.getByText('US$ 32,15')).toBeInTheDocument()
    expect(screen.queryByText('Costliest resources')).not.toBeInTheDocument()
    expect(screen.queryByText('should-not-show')).not.toBeInTheDocument()
    expect(screen.queryByText('AWS Glue')).not.toBeInTheDocument()
    expect(screen.queryByText('Findings')).not.toBeInTheDocument()
  })

  it('does not show month-to-date when source is Cost Explorer', () => {
    render(Dashboard, { props: baseProps() })
    expect(screen.queryByText(/Month-to-date/)).not.toBeInTheDocument()
  })
})
