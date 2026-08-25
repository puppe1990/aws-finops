import { describe, expect, it } from 'vitest'
import { deltaPercent, monthBars } from './compareView.js'

describe('compareView', () => {
  it('scales month bars to the peak month', () => {
    const rows = monthBars([
      { query: '2026-08', cents: 1983, usd: 'US$ 19,83' },
      { query: '2026-07', cents: 3966, usd: 'US$ 39,66' },
    ])
    expect(rows[1].pct).toBe(100)
    expect(rows[0].pct).toBe(50)
  })

  it('formats month-over-month percent with a sign', () => {
    expect(deltaPercent(-3850)).toBe('-38%')
    expect(deltaPercent(1200)).toBe('+12%')
    expect(deltaPercent(null)).toBe('')
  })
})
