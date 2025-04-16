import axios from 'axios';
import qs from 'query-string';

export interface AlertMttaRecord {
  today: number;
  yesterday: number;
  diff: number;
  direction: string;
}

export interface AlertLevelDistributionRecord {
  count: number;
  levelName: string;
  color: string;
}

export interface TopAlertRecord {
  alertTitle: string;
  num: number;
}

export interface TopAlertParams {
  timeRange?: string;
}

export interface AlertCountRecord {
  date: string;
  num: number;
}

export function queryAlertMtta(id: number) {
  return axios.get<AlertMttaRecord>(`/api/v1/alert/statistic/mtta/${id}`);
}

export function queryAlertMttr(id: number) {
  return axios.get<AlertMttaRecord>(`/api/v1/alert/statistic/mttr/${id}`);
}

export function queryAlertCount(id: number) {
  return axios.get<AlertMttaRecord>(`/api/v1/alert/statistic/count/${id}`);
}

export function queryAlertLevel(id: number) {
  return axios.get<AlertLevelDistributionRecord[]>(
    `/api/v1/alert/statistic/level/${id}`
  );
}

export function queryAllAlertMtta() {
  return axios.get<AlertMttaRecord>(`/api/v1/alert/statistic/mtta`);
}

export function queryAllAlertMttr() {
  return axios.get<AlertMttaRecord>(`/api/v1/alert/statistic/mttr`);
}

export function queryAllAlertCount() {
  return axios.get<AlertMttaRecord>(`/api/v1/alert/statistic/count`);
}

export function queryAllAlertLevel() {
  return axios.get<AlertLevelDistributionRecord[]>(
    `/api/v1/alert/statistic/level`
  );
}

export function queryTopAlert(params: TopAlertParams) {
  return axios.get<TopAlertRecord[]>(`/api/v1/alert/statistic/top10`, {
    params,
    paramsSerializer: (obj) => {
      return qs.stringify(obj);
    },
  });
}

export function queryAlertLevelStatistic(params: TopAlertParams) {
  return axios.get<AlertLevelDistributionRecord[]>(
    `/api/v1/alert/statistic/level-alert`,
    {
      params,
      paramsSerializer: (obj) => {
        return qs.stringify(obj);
      },
    }
  );
}

export function queryAlertCountStatistic(params: TopAlertParams) {
  return axios.get<AlertCountRecord[]>(`/api/v1/alert/statistic/alert-count`, {
    params,
    paramsSerializer: (obj) => {
      return qs.stringify(obj);
    },
  });
}
