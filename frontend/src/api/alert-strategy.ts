import axios from 'axios';

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

export interface StrategyFilterValues {
  values: StrategyFilter[];
}

export interface AlertStrategyRequest {
  workspaceId: number;
  delay: number;
  systemStrategy: boolean;
  continuous: boolean;
  strategyName: string;
  strategyDesc: string;
  templateId: number;
  timeSlot: TimeSlot;
  filters: StrategyFilterValues[];
  strategySet: StrategyItem[];
}

export interface AlertStrategyRecord {
  id: number;
  status: string;
  workspaceId: number;
  delay: number;
  continuous: boolean;
  systemStrategy: boolean;
  strategyName: string;
  weight: number;
  strategyDesc: string;
  templateId: number;
  timeSlot: TimeSlot;
  filters: StrategyFilterValues[];
  strategySet: StrategyItem[];
  nickname: string;
  updatedAt: string;
}

export interface AlertStrategySort {
  id: number;
  weight: number;
}

export function submitStrategyForm(data: AlertStrategyRequest) {
  return axios.post('/api/v1/alert/strategy', data);
}

export function editStrategyForm(id: number, data: AlertStrategyRequest) {
  return axios.put(`/api/v1/alert/strategy/${id}`, data);
}

export function editStrategySort(data: AlertStrategySort[]) {
  return axios.put(`/api/v1/alert/strategy/sort`, { newWeight: data });
}

export function queryAllAlertStrategy(id: number) {
  return axios.get<AlertStrategyRecord[]>(`/api/v1/alert/strategy/all/${id}`);
}

export function changeStrategyStatus(id: number) {
  return axios.put(`/api/v1/alert/strategy/status/${id}`);
}

export function deleteStrategyRecord(id: number) {
  return axios.delete(`/api/v1/alert/strategy/${id}`);
}
