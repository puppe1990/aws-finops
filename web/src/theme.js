export const THEME_KEY = 'cifra-theme'
export const THEME_LIGHT = 'light'
export const THEME_DARK = 'dark'

export function readTheme() {
  try {
    return localStorage.getItem(THEME_KEY) === THEME_LIGHT ? THEME_LIGHT : THEME_DARK
  } catch {
    return THEME_DARK
  }
}

export function applyTheme(theme) {
  const light = theme === THEME_LIGHT
  document.documentElement.classList.toggle('light', light)
  try {
    localStorage.setItem(THEME_KEY, light ? THEME_LIGHT : THEME_DARK)
  } catch {
    /* private mode */
  }
  const meta = document.querySelector('meta[name="theme-color"]')
  if (meta) {
    meta.setAttribute('content', light ? '#f7f1e2' : '#121816')
  }
}

export function toggleTheme(current) {
  const next = current === THEME_LIGHT ? THEME_DARK : THEME_LIGHT
  applyTheme(next)
  return next
}
