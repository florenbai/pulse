import axios from 'axios';
import qs from 'query-string';

export interface ScheduleForm {
  scheduleName: string;
  scheduleDesc: string;
  teamId: number | undefined;
  scheduleTimeRange: string[];
  schedulePeriod: string[];
  users: number[];
  // scheduleUserCount: number;
}

export interface ScheduleRecord {
  id: number;
  scheduleName: string;
  scheduleDesc: string;
  teamId: number;
  teamName: string;
  periodType: string;
  period: number;
  startRange: string;
  endRange: string;
  createdAt: string;
  updatedAt: string;
}

export interface SchedulePlanItem {
  day: string;
  isSchedule: boolean;
}

export interface SchedulePlanRecord {
  id: number;
  scheduleId: number;
  month: string;
  plan: SchedulePlanItem[];
  uid: string;
  nickname: string;
}

export interface ScheduleParams extends Partial<ScheduleRecord> {
  page: number;
  pageSize: number;
}

export interface ScheduleListRes {
  list: ScheduleRecord[];
  total: number;
}

export interface SchedulePlanParams {
  month: string;
}

export interface SchedulePlanRequest {
  index: number[];
}

export function submitScheduleForm(data: ScheduleForm) {
  return axios.post('/api/v1/schedule', data);
}

export function querySchedulePlanList(id: number, params: SchedulePlanParams) {
  return axios.get<SchedulePlanRecord[]>(`/api/v1/schedule-plan/list/${id}`, {
    params,
    paramsSerializer: (obj) => {
      return qs.stringify(obj);
    },
  });
}

export function queryScheduleList(params: ScheduleParams) {
  return axios.get<ScheduleListRes>('/api/v1/schedule/list', {
    params,
    paramsSerializer: (obj) => {
      return qs.stringify(obj);
    },
  });
}

export function queryAllSchedule() {
  return axios.get(`/api/v1/schedule/all`);
}

export function deleteSchedule(id: number) {
  return axios.delete(`/api/v1/schedule/${id}`);
}

export function setSchedulePlan(id: number, data: SchedulePlanRequest) {
  return axios.put(`/api/v1/schedule-plan/${id}`, data);
}

export function clearSchedulePlan(id: number, data: SchedulePlanRequest) {
  return axios.put(`/api/v1/schedule-plan/clear/${id}`, data);
}
