import { getToken } from '@/utils/token'
import { toast } from '@/composables/useToast'

// 响应 HTTP 请求
export async function handleResult(resp: Response) {
  if (resp.status !== 200) {
    let reason = await resp.text()
    // TODO:
    if (/json/i.test(resp.headers.get('Content-Type') || '')) {
      const data = JSON.parse(reason)
      reason = data.error || reason
    }
    if (!reason)
      reason = resp.statusText

    // global error handler
    toast.error(reason)
    console.error(reason)

    return Promise.reject(reason)
  }
  return await resp.json()
}

// 发送请求函数
export async function sendReq(method: string, url: RequestInfo, data?: any, headerInit?: HeadersInit) {
  const headers = new Headers({
    'Content-Type': 'application/json',
    ...headerInit,
  })

  const token = getToken()
  token && headers.set('Authorization', `Bearer ${token}`)

  const resp = await fetch(url, {
    method,
    body: JSON.stringify(data || {}),
    headers,
  })
  return await handleResult(resp)
}

// 封装 HTTP 请求类
class Request {
  async delete(url: RequestInfo, data?: any) {
    return await sendReq('DELETE', url, data)
  }

  async put(url: RequestInfo, data?: any) {
    return await sendReq('PUT', url, data)
  }

  async patch(url: RequestInfo, data?: any) {
    return await sendReq('PATCH', url, data)
  }

  async post(url: RequestInfo, data?: any) {
    return await sendReq('POST', url, data)
  }

  async get(url: RequestInfo) {
    const headers = new Headers({})
    const token = getToken()
    token && headers.set('Authorization', `Bearer ${token}`)

    const resp = await fetch(url, { headers })

    return await handleResult(resp)
  }
}

// 导出请求类
const request = new Request()
export default request
