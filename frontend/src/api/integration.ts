import axios from 'axios';
import { StrategyFilterValues } from './alert-strategy';

export interface SystemIntegrationRecode {
  id: number;
  name: string;
  description: string;
  token: string;
  levelField: string;
  icon: string;
  eventCount: number;
  disabled: boolean;
}
export interface IntegrationRouterFilter {
  tag: '';
  operation: 'IN';
  value: [];
}

export interface IntegrationRouter {
  id: number;
  workspaces: number[];
  workspaceNames: string[];
  filters: StrategyFilterValues[];
  next: number;
  sort: number;
}

export interface TagExtract {
  oldTag?: string;
  expression?: string;
  newTag?: string;
  value?: string;
  deleteTag?: string[];
}

export interface TagRewriteItem {
  id: number;
  filters: IntegrationRouterFilter[];
  rewriteType: string;
  rewriteItem: TagExtract;
}

export function querySystemIntegration() {
  return axios.get(`/api/v1/integration/list`);
}

export function queryIntegrationDetail(id: number) {
  return axios.get(`/api/v1/integration/detail/${id}`);
}

export function deleteSystemIntegration(id: number) {
  return axios.delete(`/api/v1/integration/${id}`);
}

export function submitSystemIntegration(data: any) {
  return axios.post(`/api/v1/integration`, data);
}

export function changeIntegrationStatusById(id: number) {
  return axios.put(`/api/v1/integration/status/${id}`, id);
}

export function updateSystemIntegration(id: number, data: any) {
  return axios.put(`/api/v1/integration/${id}`, data);
}

export function submitIntegrationRouterForm(id: number, data: any) {
  return axios.post(`/api/v1/integration/router/${id}`, data);
}

export function updateIntegrationRouterForm(id: number, data: any) {
  return axios.put(`/api/v1/integration/router/${id}`, data);
}

export function queryIntegrationAllRouters(id: number) {
  return axios.get<IntegrationRouter[]>(`/api/v1/integration/router/all/${id}`);
}

export function deleteIntegrationRouterRecord(id: number) {
  return axios.delete(`/api/v1/integration/router/${id}`);
}

export function changeRouterSort(data: number[]) {
  return axios.put(`/api/v1/integration/router/sort`, data);
}

export function submitTagRewriteForm(id: number, data: any) {
  return axios.post(`/api/v1/integration/tag-rewrite/${id}`, data);
}

export function updateTagRewriteForm(id: number, data: any) {
  return axios.put(`/api/v1/integration/tag-rewrite/${id}`, data);
}

export function queryAllTagRewrite(id: number) {
  return axios.get<TagRewriteItem[]>(
    `/api/v1/integration/tag-rewrite/all/${id}`
  );
}

export function deleteTagRewriteRecord(id: number) {
  return axios.delete(`/api/v1/integration/tag-rewrite/${id}`);
}
