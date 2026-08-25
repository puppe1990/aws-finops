export function monthBars(months) {
  const list = months || []
  const max = list.reduce((m, row) => Math.max(m, Number(row.cents) || 0), 0)
  return list.map((row) => ({
    ...row,
    pct: max === 0 ? 0 : Math.round((Number(row.cents) / max) * 100),
  }))
}

export function deltaPercent(bps) {
  if (bps === null || bps === undefined || bps === '') {
    return ''
  }
  const pct = Math.round(Number(bps) / 100)
  const sign = pct > 0 ? '+' : ''
  return `${sign}${pct}%`
}

export function deltaTone(bps) {
  if (bps === null || bps === undefined || bps === '') return 'flat'
  const n = Number(bps)
  if (n > 0) return 'up'
  if (n < 0) return 'down'
  return 'flat'
}

export function monthShort(label) {
  const s = String(label || '')
  const iso = s.match(/^(\d{4})-(\d{2})/)
  if (iso) return String(Number(iso[2]))
  return s.split(/\s+/)[0] || ''
}

export function chartMonths(months) {
  const bars = monthBars(months)
  const chrono = [...bars].reverse()
  const first = chrono.findIndex((row) => Number(row.cents) > 0)
  return first < 0 ? [] : chrono.slice(first)
}
