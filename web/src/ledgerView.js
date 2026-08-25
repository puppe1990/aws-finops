export function spendRows(services) {
  const visible = (services || []).filter((s) => Number(s.cents) > 0)
  const max = visible.reduce((m, s) => Math.max(m, Number(s.cents) || 0), 0)
  return visible.map((s) => ({
    ...s,
    pct: max === 0 ? 0 : Math.round((Number(s.cents) / max) * 100),
  }))
}

export function topFindings(findings, limit = 3) {
  return (findings || []).slice(0, limit)
}

export function findingHref(kind) {
  return kind === 'ce_denied' ? '/settings' : '/resources'
}

export function primaryAccountId(accounts) {
  if (!accounts || !accounts[0] || !accounts[0].awsAccountId) return ''
  return accounts[0].awsAccountId
}

export function burnPercent(bps) {
  return `${Math.round((Number(bps) || 0) / 100)}%`
}

export function shortSync(at) {
  const s = String(at || '')
  return s.length >= 10 ? s.slice(0, 10) : s
}

export function usageLabel(name) {
  return String(name || '').replace(/^[A-Z]{2,4}\d-/, '')
}
