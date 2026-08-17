import { describe, expect, it } from 'vitest'
import { t, tf } from './i18n.js'

describe('t', () => {
  it('returns the label when present', () => {
    expect(t({ 'nav.ledger': 'Ledger' }, 'nav.ledger')).toBe('Ledger')
  })

  it('returns the key when missing', () => {
    expect(t({}, 'nav.ledger')).toBe('nav.ledger')
  })

  it('substitutes %s placeholders', () => {
    expect(tf({ 'dash.burn': 'Burn %s' }, 'dash.burn', '24%')).toBe('Burn 24%')
  })
})
