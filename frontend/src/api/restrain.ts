import axios from 'axios';

export interface RestrainRequest {
  restrainType: string;
  fields?: string[];
  cumulativeTime: number;
}

export interface RestrainRecord {
  id: number;
  restrainType: string;
  fields?: string[];
  cumulativeTime: number;
}

export function submitRestrainForm(id: number, data: RestrainRequest) {
  return axios.post(`/api/v1/alert/restrain/${id}`, data);
}

export function editRestrainForm(id: number, data: RestrainRequest) {
  return axios.put(`/api/v1/alert/restrain/${id}`, data);
}

export function queryAllAlertRestrain(id: number) {
  return axios.get<RestrainRecord[]>(`/api/v1/alert/restrain/all/${id}`);
}

export function deleteRestrainRecord(id: number) {
  return axios.delete(`/api/v1/alert/restrain/${id}`);
}
