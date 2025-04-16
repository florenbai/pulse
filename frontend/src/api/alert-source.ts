import axios from 'axios';

export interface AlertSourceRecode {
  id: number;
  sourceName: string;
  icon: string;
  category: string;
  createdAt: string;
  updatedAt: string;
  status: boolean;
}

export interface WorkspaceAlertSourceRecode {
  id: number;
  sourceName: string;
  description: string;
  verified: boolean;
  configuration: string;
  icon: string;
  token: string;
  levelField: string;
  eventCount: number;
}

export function queryAllAlertSource() {
  return axios.get(`/api/v1/alert/source/all`);
}

export function queryWorkspaceAlertSource(id: number | undefined) {
  return axios.get<WorkspaceAlertSourceRecode[]>(
    `/api/v1/workspace/alert-source/${id}`
  );
}

export function deleteAlertSource(id: number) {
  return axios.delete(`/api/v1/workspace/alert-source/${id}`);
}
