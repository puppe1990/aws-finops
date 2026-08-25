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
  'dash.forecast': 'Forecast %s · %s',
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
    expect(screen.getByText('Amazon Lightsail')).toBeInTheDocument()
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

  it('renders prev month link and omits next when current', () => {
    render(Dashboard, {
      props: baseProps({
        monthLabel: 'Aug 2026',
        prevMonth: '2026-07',
        nextMonth: '',
        isCurrent: true,
      }),
    })
    const prev = document.querySelector('a[href="/dashboard?month=2026-07"]')
    expect(prev).toBeTruthy()
    expect(document.querySelector('a[href="/dashboard?month=2026-09"]')).toBeNull()
    expect(screen.queryByText(/Month-to-date/)).not.toBeInTheDocument()
  })

  it('renders next month link when viewing the past', () => {
    render(Dashboard, {
      props: baseProps({
        monthLabel: 'Jul 2026',
        prevMonth: '2026-06',
        nextMonth: '2026-08',
        isCurrent: false,
      }),
    })
    expect(document.querySelector('a[href="/dashboard?month=2026-08"]')).toBeTruthy()
  })

  it('shows next-month forecast on the current month', () => {
    render(Dashboard, {
      props: baseProps({
        isCurrent: true,
        monthLabel: 'Aug 2026',
        prevMonth: '2026-07',
        nextMonth: '',
        summary: {
          monthlyUSD: 'US$ 32,15',
          source: 'ce',
          ceDenied: false,
          mtdUSD: 'US$ 32,15',
          forecastUSD: 'US$ 38,50',
          forecastLabel: 'Sep 2026',
        },
      }),
    })
    expect(screen.getByText('Forecast Sep 2026 · US$ 38,50')).toBeInTheDocument()
  })

  it('hides forecast on a past month', () => {
    render(Dashboard, {
      props: baseProps({
        isCurrent: false,
        monthLabel: 'Jul 2026',
        prevMonth: '2026-06',
        nextMonth: '2026-08',
        summary: {
          monthlyUSD: 'US$ 19,47',
          source: 'ce',
          ceDenied: false,
          forecastUSD: 'US$ 38,50',
          forecastLabel: 'Sep 2026',
        },
      }),
    })
    expect(screen.queryByText(/Forecast/)).not.toBeInTheDocument()
  })

  it('hides month-to-date on a past month even if source is estimate', () => {
    render(Dashboard, {
      props: baseProps({
        isCurrent: false,
        summary: { monthlyUSD: 'US$ 0,00', source: 'estimate', ceDenied: true, mtdUSD: 'US$ 7,35' },
        monthLabel: 'Jul 2026',
        prevMonth: '2026-06',
        nextMonth: '2026-08',
      }),
    })
    expect(screen.queryByText(/Month-to-date/)).not.toBeInTheDocument()
  })
})
