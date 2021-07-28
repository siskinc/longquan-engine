import request from '@/utils/request';
import { queryConfig } from './config';

export const queryPropertyTypeEnum = async () => { return await queryConfig(`/config/property-type-enum`) };

export const queryPropertySystemClassEnum = async () => { return await queryConfig(`/config/property-system-class-enum`) };

export function createProperty(data) {
  return request({
    url: '/property',
    method: 'post',
    data,
  });
}
