import request from '@/utils/request';

export function createPropertySet(data) {
  return request({
    url: '/property-set',
    method: 'post',
    data,
  });
}

export async function queryPropertySet(params) {
  let data = [];
  await request({
    url: '/property-set',
    method: 'get',
    params: params,
  }).then((response) => {
    data = response.data;
  })
  return data;
}

export function deletePropertySet(id) {
  return request({
    url: `/property-set/${id}`,
    method: 'delete',
  })
}