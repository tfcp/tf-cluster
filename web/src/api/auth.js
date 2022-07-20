import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/auth/user-list',
    method: 'get',
    params,
    baseURL: process.env.VUE_APP_URL
  })
}

export function getCount(params) {
  return request({
    url: '/auth/user-count',
    method: 'get',
    params,
    baseURL: process.env.VUE_APP_URL
  })
}

export function getDetail(id) {
  return request({
    url: '/auth/user-detail?id=' + id,
    method: 'get',
    baseURL: process.env.VUE_APP_URL
  })
}

export function Delete(id) {
  return request({
    url: '/auth/user-delete?id=' + id,
    method: 'post',
    baseURL: process.env.VUE_APP_URL
  })
}

export function enable(id) {
  const params = {
    id: id,
    status: 1
  }
  return request({
    url: '/auth/user-change',
    method: 'post',
    params,
    baseURL: process.env.VUE_APP_URL
  })
}

export function disable(id) {
  const params = {
    id: id,
    status: 2
  }
  return request({
    url: '/auth/user-change',
    method: 'post',
    params,
    baseURL: process.env.VUE_APP_URL
  })
}

export function save(params) {
  return request({
    url: '/auth/user-save',
    method: 'post',
    params,
    baseURL: process.env.VUE_APP_URL
  })
}

export function changePwd(params) {
  return request({
    url: '/auth/pwd-edit',
    method: 'post',
    params,
    baseURL: process.env.VUE_APP_URL
  })
}
