import axios from 'axios';
import qs from 'query-string';
import { StrategyItem } from './alert-strategy';

export interface SystemStrategyRequest {
  delay: number;
  strategyName: string;
  strategyDesc: string;
  templateId: number;
  strategySet: StrategyItem[];
}

export interface SystemStrategyRecord {
  id: number;
  delay: number;
  strategyName: string;
  strategyDesc: string;
  templateId: number;
  nickname: string;
  strategySet: StrategyItem[];
}

export interface SystemStrategyListRes {
  list: SystemStrategyRecord[];
  total: number;
}

export interface SystemStrategyParams extends Partial<SystemStrategyRecord> {
  page: number;
  pageSize: number;
}

export function submitSystemStrategyForm(data: SystemStrategyRequest) {
  return axios.post('/api/v1/system/strategy', data);
}

export function editSystemStrategyForm(
  id: number,
  data: SystemStrategyRequest
) {
  return axios.put(`/api/v1/system/strategy/${id}`, data);
}

export function querySystemStrategyList(params: SystemStrategyParams) {
  return axios.get<SystemStrategyListRes>('/api/v1/system/strategy/list', {
    params,
    paramsSerializer: (obj) => {
      return qs.stringify(obj);
    },
  });
}

export function deleteSystemStrategy(id: number) {
  return axios.delete(`/api/v1/system/strategy/${id}`);
}

export function queryAllSystemStrategy() {
  return axios.get('/api/v1/system/strategy/all', { withCredentials: true });
}

export function querySystemStrategy(id: number) {
  return axios.get(`/api/v1/system/strategy/${id}`);
}
