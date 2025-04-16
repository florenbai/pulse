import axios from 'axios';

export interface AlertLevelRecode {
  id: number;
  levelName: string;
  levelDesc: string;
  color: string;
  isDefault: boolean;
}

export interface AlertLevelRequest {
  levelName: string;
  levelDesc: string;
  color: string;
}

export interface AlertLevelMappingRecode {
  id: number;
  levelId: number;
  integrationId: number;
  alertLevel: string;
  levelName: string;
  color: string;
}

export function queryAllAlertLevel() {
  return axios.get<AlertLevelRecode[]>(`/api/v1/alert/level/all`);
}

export function updateLevelMapping(data: any) {
  return axios.put(`/api/v1/alert/level/mapping`, data);
}

export function editLevelForm(id: number, data: any) {
  return axios.put(`/api/v1/alert/level/${id}`, data);
}

export function submitLevelForm(data: any) {
  return axios.post(`/api/v1/alert/level`, data);
}

export function deleteLevel(id: number) {
  return axios.delete(`/api/v1/alert/level/${id}`);
}

export function queryAllAlertMappingLevel(id: number) {
  return axios.get<AlertLevelMappingRecode[]>(
    `/api/v1/alert/level/all-mapping/${id}`
  );
}
