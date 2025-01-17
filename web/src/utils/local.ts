import { decrypt, encrypt } from './crypto'

// 7day
const DEFAULT_CACHE_TIME = 7 * 24 * 60 * 60 * 1000

// 设置本地存储
export function setLocal(key: string, value: unknown, expire: number | null = DEFAULT_CACHE_TIME) {
  const storageData = {
    value,
    expire: expire !== null ? new Date().getTime() + expire : null,
  }
  window.localStorage.setItem(key, encrypt(storageData))
}

// 获取本地存储
export function getLocal<T>(key: string) {
  const json = window.localStorage.getItem(key)

  if (json) {
    const storageData = decrypt(json)

    if (storageData) {
      const { value, expire } = storageData
      // no expire or not expired
      if (expire === null || expire >= Date.now())
        return value as T
    }

    // invalid data or expired
    window.localStorage.removeItem(key)
    return null
  }
  return null
}

// 获取本地存储过期时间
export function getLocalExpire(key: string): number | null {
  const json = window.localStorage.getItem(key)

  if (json) {
    const storageData = decrypt(json)

    if (storageData) {
      const { expire } = storageData
      return expire
    }
    return null
  }
  return null
}
