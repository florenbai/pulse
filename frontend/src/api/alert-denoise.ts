import axios from 'axios';
import { StrategyFilter, StrategyFilterValues } from './alert-strategy';

export interface AlertAggregationRequest {
  aggregationName: string;
  aggregationDesc: string;
  workspaceId: number;
  tagsDimension: StrategyFilterValues[];
  windows: number;
  storm: number;
  titleDimension: boolean;
  levelDimension: boolean;
}

export interface AlertAggregationRecord {
  id: number;
  aggregationName: string;
  aggregationDesc: string;
  workspaceId: number;
  tagsDimension: StrategyFilterValues[];
  windows: number;
  storm: number;
  titleDimension: boolean;
  levelDimension: boolean;
  status: string;
  nickname: string;
  updatedAt: string;
}

export interface SilenceTime {
  rangeTime?: string[];
  weeks?: number[];
  times?: string[];
}

export interface SilenceRequest {
  workspaceId: number;
  silenceName: string;
  silenceDesc: string;
  silenceType: string;
  silenceTime: SilenceTime;
  filters: StrategyFilter[];
}

export interface SilenceRecord {
  id: number;
  workspaceId: number;
  silenceName: string;
  silenceDesc: string;
  silenceType: string;
  silenceTime: SilenceTime;
  filters: StrategyFilter[];
  nickname: string;
  updatedAt: string;
}

export function submitAggregationForm(data: AlertAggregationRequest) {
  return axios.post('/api/v1/alert/aggregation', data);
}

export function queryAllAlertAggregation(id: number) {
  return axios.get<AlertAggregationRecord[]>(
    `/api/v1/alert/aggregation/all/${id}`
  );
}

export function editAggregationForm(id: number, data: AlertAggregationRequest) {
  return axios.put(`/api/v1/alert/aggregation/${id}`, data);
}

export function changeAggregationStatus(id: number) {
  return axios.put(`/api/v1/alert/aggregation/status/${id}`);
}

export function deleteAggregationRecord(id: number) {
  return axios.delete(`/api/v1/alert/aggregation/${id}`);
}

export function submitSilenceForm(data: SilenceRequest) {
  return axios.post('/api/v1/alert/silence', data);
}

export function queryAllAlertSilence(id: number) {
  return axios.get<SilenceRecord[]>(`/api/v1/alert/silence/all/${id}`);
}

export function editSilenceForm(id: number, data: SilenceRequest) {
  return axios.put(`/api/v1/alert/silence/${id}`, data);
}

export function deleteSilenceRecord(id: number) {
  return axios.delete(`/api/v1/alert/silence/${id}`);
}
