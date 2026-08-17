import { afterEach, describe, expect, it } from 'vitest'
import { applyTheme, readTheme, THEME_KEY, toggleTheme } from './theme.js'

afterEach(() => {
  document.documentElement.classList.remove('light')
  localStorage.removeItem(THEME_KEY)
})

describe('readTheme', () => {
  it('defaults to dark when nothing is stored', () => {
    expect(readTheme()).toBe('dark')
  })

  it('reads light from localStorage', () => {
    localStorage.setItem(THEME_KEY, 'light')
    expect(readTheme()).toBe('light')
  })
})

describe('applyTheme', () => {
  it('adds light class and persists', () => {
    applyTheme('light')
    expect(document.documentElement.classList.contains('light')).toBe(true)
    expect(localStorage.getItem(THEME_KEY)).toBe('light')
  })

  it('removes light class for dark', () => {
    document.documentElement.classList.add('light')
    applyTheme('dark')
    expect(document.documentElement.classList.contains('light')).toBe(false)
    expect(localStorage.getItem(THEME_KEY)).toBe('dark')
  })
})

describe('toggleTheme', () => {
  it('switches dark to light', () => {
    expect(toggleTheme('dark')).toBe('light')
  })

  it('switches light to dark', () => {
    expect(toggleTheme('light')).toBe('dark')
  })
})
