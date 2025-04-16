import axios from 'axios';
import qs from 'query-string';
import { MenuPermission } from './menu';

export interface TeamRecord {
  id: number;
  teamName: string;
  teamDesc: string;
  updatedAt: string;
  updatedName: string;
  status: boolean;
}

export interface TeamUserRecord {
  userId: number;
  teamName: string;
  nickname: string;
}

export interface TeamInfoRecord {
  teamName: string;
  teamDesc: string;
  updatedAt: string;
  updatedName: string;
  teamMembers: number[];
  status: boolean;
}

export interface TeamRequest {
  teamName: string;
  teamDesc: string;
  teamMembers: number[];
  status: boolean;
}

export interface TeamParams extends Partial<TeamRecord> {
  page: number;
  pageSize: number;
}

export interface TeamListRes {
  list: TeamRecord[];
  total: number;
}

export function queryTeamInfo(id: number) {
  return axios.get<TeamInfoRecord>(`/api/v1/team/${id}`);
}

export function submitTeamForm(data: TeamRequest) {
  return axios.post('/api/v1/team', data);
}

export function submitEditTeamForm(id: number, data: TeamRequest) {
  return axios.put(`/api/v1/team/${id}`, data);
}

export function changeTeamStatus(id: number) {
  return axios.put(`/api/v1/team/status/${id}`);
}

export function queryAllTeam() {
  return axios.get(`/api/v1/team/all`);
}

export function queryAllTeamUser(team: number[]) {
  const t = team.join(',');
  return axios.get(`/api/v1/team/user/all/${t}`);
}

export function queryTeamList(params: TeamParams) {
  return axios.get<TeamListRes>('/api/v1/team/list', {
    params,
    paramsSerializer: (obj) => {
      return qs.stringify(obj);
    },
  });
}

export function deleteTeam(id: number) {
  return axios.delete(`/api/v1/team/${id}`);
}

export function teamPermission(id: number, data: MenuPermission) {
  return axios.put(`/api/v1/team/permission/${id}`, data);
}

export function queryAuthorizationInfo(id: number) {
  return axios.get(`/api/v1/team/authorization/${id}`);
}
