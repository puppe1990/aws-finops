export function t(labels, key) {
  if (labels && labels[key]) {
    return labels[key]
  }
  return key
}

export function tf(labels, key, ...args) {
  let out = t(labels, key)
  for (const arg of args) {
    out = out.replace('%s', String(arg))
  }
  return out
}
