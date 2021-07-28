import request from '@/utils/request';

export async function queryConfig(url, params) {
  let data = {};
  await request({
    url: url,
    method: 'get',
    params,
  }).then(response => {
    data = response.data;
  })
  return data;
}
