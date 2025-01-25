import request from '@/utils/request'

export function createUser(data) {
  return request({
    url: '/api/v1/user',
    method: 'post',
    data
  })
}

export function getUser(id) {
  return request({
    url: '/api/v1/user/' + id,
    method: 'get'
  })
}

export function getList(data) {
  return request({
    url: '/api/v1/user',
    method: 'get',
    data
  })
}

export function updateUser(data) {
  return request({
    url: '/api/v1/user',
    method: 'put',
    data
  })
}

export function deleteUser(data) {
  return request({
    url: '/api/v1/user/',
    method: 'delete',
    data
  })
}

export function login(data) {
  return request({
    url: '/login',
    method: 'post',
    data
  })
}

export function getInfo() {
  return request({
    url: '/user/info',
    method: 'get'
  })
}

export function refreshToken() {
  return request({
    url: '/api/auth/refresh',
    method: 'get'
  })
}

export function logout() {
  return request({
    url: '/user/logout',
    method: 'post'
  })
}
