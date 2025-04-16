import axios from 'axios';
import { S } from 'mockjs';
import qs from 'query-string';

export interface AlertRuleSourceRecode {
  id: number;
  sourceName: string;
  sourceType: string;
  mark: string;
  address: string;
  sign: string;
}

export interface AlertRuleGroupRecode {
  id: number;
  groupName: string;
  ruleSource: number;
  file: string;
  children: [];
}

export interface AlertRuleRecode {
  id: number;
  name: string;
  health: string;
  sourceType: string;
  labels: Map<string, string>;
  annotations: Map<string, string>;
  duration: number;
  query: string;
  gid: string;
  updatedAt: string;
}

export interface AlertRuleLabels {
  tag: string;
  value: string;
}

export interface PageParams extends Partial<AlertRuleSourceRecode> {
  page: number;
  pageSize: number;
}

export interface AlertRuleSourceListRes {
  list: AlertRuleSourceRecode[];
  total: number;
}

export interface Agent {
  id: number;
  clientId: string;
  heartbeatTime: string;
}

export function submitAlertRuleSource(data: any) {
  return axios.post(`/api/v1/alert-rule/source`, data);
}

export function submitAlertRuleGroup(data: any) {
  return axios.post(`/api/v1/rule-group`, data);
}

export function updateAlertRuleSource(id: number, data: any) {
  return axios.put(`/api/v1/alert-rule/source/${id}`, data);
}

export function queryAlertRuleGroup() {
  return axios.get(`/api/v1/rule-group/all`);
}

export function querySourceLabels(id: number) {
  return axios.get(`/api/v1/rule-group/label/${id}`);
}

export function deleteAlertRuleGroup(id: number) {
  return axios.delete(`/api/v1/rule-group/${id}`);
}

export function reloadAlertRuleSource(id: number) {
  return axios.post(`/api/v1/alert-rule/source/reload/${id}`);
}

export function deleteAlertRule(id: number) {
  return axios.delete(`/api/v1/alert-rule/${id}`);
}

export function queryAlertRuleList(id: number) {
  return axios.get(`/api/v1/alert-rule/list/gid/${id}`);
}

export function submitAlertRule(data: any) {
  return axios.post(`/api/v1/alert-rule`, data);
}

export function updateAlertRule(id: number, data: any) {
  return axios.put(`/api/v1/alert-rule/${id}`, data);
}

export function queryAlertRuleSource(params: PageParams) {
  return axios.get<AlertRuleSourceListRes>('/api/v1/alert-rule/source/list', {
    params,
    paramsSerializer: (obj) => {
      return qs.stringify(obj);
    },
  });
}

export function queryAllAlertRuleSource() {
  return axios.get(`/api/v1/alert-rule/source/all`);
}

export function queryAllAlertRuleSourceAgent(id: number) {
  return axios.get(`/api/v1/alert-rule/source/agent/${id}`);
}

export function deleteAlertRuleSource(id: number) {
  return axios.delete(`/api/v1/alert-rule/source/${id}`);
}
