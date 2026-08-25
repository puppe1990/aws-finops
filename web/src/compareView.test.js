import { describe, expect, it } from 'vitest'
import { chartMonths, deltaPercent, deltaTone, monthBars, monthShort } from './compareView.js'

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

  it('lays chart months oldest-left and drops leading zeros', () => {
    const rows = chartMonths([
      { query: '2026-08', label: 'Aug 2026', cents: 1983 },
      { query: '2026-07', label: 'Jul 2026', cents: 3966 },
      { query: '2026-06', label: 'Jun 2026', cents: 0 },
    ])
    expect(rows.map((r) => r.query)).toEqual(['2026-07', '2026-08'])
    expect(rows[0].pct).toBe(100)
  })

  it('shortens month labels and tones deltas', () => {
    expect(monthShort('Aug 2026')).toBe('Aug')
    expect(monthShort('2026-08')).toBe('8')
    expect(deltaTone(-3850)).toBe('down')
    expect(deltaTone(1200)).toBe('up')
    expect(deltaTone(null)).toBe('flat')
  })
})

