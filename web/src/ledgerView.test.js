import { describe, expect, it } from 'vitest'
import {
  spendRows,
  topFindings,
  findingHref,
  primaryAccountId,
  burnPercent,
  shortSync,
  usageLabel,
} from './ledgerView.js'

describe('spendRows', () => {
  it('drops zero-cent services and sizes bars against the largest remaining', () => {
    const rows = spendRows([
      { name: 'Lightsail', cents: 1947, usd: 'US$ 19,47' },
      { name: 'Glue', cents: 0, usd: 'US$ 0,00' },
      { name: 'Amplify', cents: 721, usd: 'US$ 7,21' },
    ])
    expect(rows.map((r) => r.name)).toEqual(['Lightsail', 'Amplify'])
    expect(rows[0].pct).toBe(100)
    expect(rows[1].pct).toBe(37)
  })

  it('returns empty when nothing has spend', () => {
    expect(spendRows([{ name: 'Glue', cents: 0 }])).toEqual([])
  })

  it('keeps usage-type details on visible rows', () => {
    const rows = spendRows([
      {
        name: 'Amazon Lightsail',
        cents: 1947,
        usd: 'US$ 19,47',
        details: [{ name: 'BoxUsage:small_3_0', cents: 1900, usd: 'US$ 19,00' }],
      },
    ])
    expect(rows[0].details).toEqual([{ name: 'BoxUsage:small_3_0', cents: 1900, usd: 'US$ 19,00' }])
  })
})

describe('topFindings', () => {
  it('keeps the first three', () => {
    const got = topFindings([{ kind: 'a' }, { kind: 'b' }, { kind: 'c' }, { kind: 'd' }])
    expect(got.map((f) => f.kind)).toEqual(['a', 'b', 'c'])
  })
})

describe('findingHref', () => {
  it('sends ce_denied to IAM settings', () => {
    expect(findingHref('ce_denied')).toBe('/settings')
  })
  it('sends other findings to inventory', () => {
    expect(findingHref('unknown_s3_size')).toBe('/resources')
  })
})

describe('primaryAccountId', () => {
  it('uses the first account (primary is already first)', () => {
    expect(primaryAccountId([{ awsAccountId: '840298254452' }])).toBe('840298254452')
  })
  it('is empty without accounts', () => {
    expect(primaryAccountId([])).toBe('')
  })
})

describe('burnPercent', () => {
  it('renders basis points as a percent', () => {
    expect(burnPercent(6400)).toBe('64%')
  })
})

describe('shortSync', () => {
  it('keeps the calendar day of an RFC3339 stamp', () => {
    expect(shortSync('2026-08-19T22:33:21Z')).toBe('2026-08-19')
  })
})

describe('usageLabel', () => {
  it('strips the AWS region prefix from a usage type', () => {
    expect(usageLabel('USE1-BundleUsage:2GB')).toBe('BundleUsage:2GB')
  })
})
