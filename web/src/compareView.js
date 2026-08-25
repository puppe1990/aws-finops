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
