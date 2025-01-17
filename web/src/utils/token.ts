import { getLocal, getLocalExpire, setLocal } from './local'

const TOKEN_KEY = 'access_token'
const DURATION = 24 * 60 * 60 * 1000 // 24 hour

// 从本地存储中获取 token
export function getToken() {
  return getLocal(TOKEN_KEY)
}

// 将 token 存储到本地存储中, 并设置过期时间
export function setToken(token: string) {
  setLocal(TOKEN_KEY, token, DURATION)
}

// 从本地存储中删除 token
export function removeToken() {
  window.localStorage.removeItem(TOKEN_KEY)
}

// 刷新 token, 避免过期
export async function refreshAccessToken() {
  const expire = getLocalExpire(TOKEN_KEY)

  if (!expire || expire - new Date().getTime() > 1000 * 60 * 30)
    return

  // TODO: refresh token api
  try {
    setToken('new token')
  }
  // 无感刷新，有异常也不提示
  catch {}
}
