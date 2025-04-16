import axios from 'axios';
import qs from 'query-string';

export interface BaseInfoRequest {
  workspaceName: string;
  workspaceDesc: string;
  teams: number[];
  strategy: number | undefined;
}

export interface StrategyItem {
  members: number[];
  teams: number[];
  schedules: number[];
  alertChannels: string[][];
  alertLoop: {
    enable: boolean;
    interval: number;
    count: number;
  };
  progression: {
    enable: boolean;
    condition: number;
    duration: number;
  };
}

export interface TimeSlot {
  enable: boolean;
  type?: string;
  weeks?: number[];
  times?: string[];
  calendar?: number;
}

export interface StrategyFilter {
  tag: string;
  operation: string;
  value: string[];
}

export interface AlertStrategyRequest {
  delay: number;
  strategyName: string;
  strategyDesc: string;
  templateId: number;
  timeSlot: TimeSlot;
  filters: StrategyFilter[];
  strategySet: StrategyItem[];
}

export interface AlertSourceRequest {
  alertSource: number[];
}

export interface ConditionItem {
  entry: number;
  type: number;
  value: string;
}

export interface Gather {
  conditionItem: ConditionItem[];
}
export interface DownFrequencyRequest {
  alertAggregation: {
    enable: boolean;
    dimension?: number[];
    window?: number;
    stormWatch?: number;
  };
  alertQuiet: {
    enable: boolean;
    quietType: string;
    quietStartTime?: string;
    quietEndTime?: string;
    condition: {
      enable: boolean;
      gather: Gather[];
    };
  };
}

export interface WorkspaceRecode {
  id: number;
  workspaceName: string;
  workspaceDesc: string;
  enable: boolean;
  pending: number;
  processing: number;
  sourceCount: number;
  readySourceCount: number;
}

export type WorkspaceRequest = BaseInfoRequest & AlertSourceRequest;

export interface WorkspaceParams extends Partial<WorkspaceRecode> {
  page: number;
  pageSize: number;
}

export interface WorkspaceListRes {
  list: WorkspaceRecode[];
  total: number;
}

export interface LevelMappingReq {
  id: number;
  alertSourceLevel: string;
}

export function submitWorkspaceForm(data: BaseInfoRequest) {
  return axios.post('/api/v1/workspace', data);
}

export function editWorkspaceForm(id: number, data: any) {
  return axios.put(`/api/v1/workspace/${id}`, data);
}

export function queryWorkspaceBase(id: number) {
  return axios.get(`/api/v1/workspace/${id}`);
}

export function deleteWorkspaceItem(id: number) {
  return axios.delete(`/api/v1/workspace/${id}`);
}

export function changeWorkspaceStatus(id: number) {
  return axios.put(`/api/v1/workspace/status/${id}`);
}

export function relatedWorkspaceAlertSourceForm(id: number, data: any) {
  return axios.post(`/api/v1/workspace/alert-source/${id}`, data);
}

export function editWorkspaceAlertSourceName(id: number, data: any) {
  return axios.put(`/api/v1/workspace/alert-source/${id}`, data);
}

export function refreshWorkspaceAlertSourceToken(id: number) {
  return axios.put(`/api/v1/workspace/alert-source/refresh-token/${id}`);
}

export function updateWorkspaceAlertSourceLevelMapping(data: LevelMappingReq) {
  return axios.put(`/api/v1/workspace/alert-source/level-mapping`, data);
}

export function queryWorkspaceList(params: WorkspaceParams) {
  return axios.get<WorkspaceListRes>('/api/v1/workspace/list', {
    params,
    paramsSerializer: (obj) => {
      return qs.stringify(obj);
    },
  });
}

export function queryWorkspaceAll() {
  return axios.get('/api/v1/workspace/all');
}
