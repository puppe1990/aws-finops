import { render, screen } from '@testing-library/svelte'
import { describe, expect, it } from 'vitest'
import ServiceSpendList from './ServiceSpendList.svelte'

describe('ServiceSpendList', () => {
  it('renders name, amount, and a bar width from pct', () => {
    render(ServiceSpendList, {
      props: {
        rows: [{ name: 'Amazon Lightsail', usd: 'US$ 19,47', pct: 100 }],
        emptyLabel: 'Nothing synced yet.',
      },
    })
    expect(screen.getByText('Amazon Lightsail')).toBeInTheDocument()
    expect(screen.getByText('US$ 19,47')).toBeInTheDocument()
    const bar = document.querySelector('[data-testid="spend-bar"]')
    expect(bar).toHaveStyle({ width: '100%' })
  })

  it('shows empty copy when there are no rows', () => {
    render(ServiceSpendList, {
      props: { rows: [], emptyLabel: 'Nothing synced yet.' },
    })
    expect(screen.getByText('Nothing synced yet.')).toBeInTheDocument()
  })
})
