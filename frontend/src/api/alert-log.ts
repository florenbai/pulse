import axios from 'axios';

export interface AlertLogRecord {
  action: string;
  nickname: string;
  createdAt: string;
}

export interface StrategyLogRecord {
  channels: string[];
  isNotify: boolean;
  errMessage: string;
  createdAt: string;
  strategyType: number;
  nicknames: string[];
  strategyTitle: string;
  notifyType: number;
}

export function queryAlertLog(id: number) {
  return axios.get<AlertLogRecord[]>(`/api/v1/alert/log/${id}`);
}

export function queryStrategyLog(id: number) {
  return axios.get<StrategyLogRecord[]>(`/api/v1/strategy/log/${id}`);
}
