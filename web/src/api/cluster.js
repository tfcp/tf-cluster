import request from '@/utils/request'

export function getConfigList(params) {
  return request({
    url: '/cluster/config-list',
    method: 'get',
    params,
    baseURL: process.env.VUE_APP_URL
  })
}

export function getConfigCount(params) {
  return request({
    url: '/cluster/config-count',
    method: 'get',
    params,
    baseURL: process.env.VUE_APP_URL
  })
}

export function saveConfig(params) {
  return request({
    url: '/cluster/config-save',
    method: 'post',
    params,
    baseURL: process.env.VUE_APP_URL
  })
}

export function getDetail(id) {
  return request({
    url: '/cluster/config-info?id=' + id,
    method: 'get',
    baseURL: process.env.VUE_APP_URL
  })
}
