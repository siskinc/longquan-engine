import request from '@/utils/request';

export function createNamespace(data) {
  return request({
    url: '/namespace',
    method: 'post',
    data,
  });
}

export function queryNamespace(params) {
  return request({
    url: '/namespace',
    method: 'get',
    params: params,
  });
}

export async function queryNamespaceDataList(params) {
  let data = [];
  await request({
    url: '/namespace',
    method: 'get',
    params: params,
  }).then(response => {
    data = response.data;
  })
  return data;
}

export function deleteNamespace(id) {
  return request({
    url: `/namespace/${id}`,
    method: 'delete',
  });
}

export function updateNamespace(id, data) {
  return request({
    url: `/namespace/${id}`,
    method: 'patch',
    data,
  });
}
