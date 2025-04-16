import axios from 'axios';
import { UserState } from '@/store/modules/user/types';
import qs from 'query-string';

export interface UserRecord {
  id: number;
  username: string;
  nickname: string;
  password: string;
  email: string;
  phone: string;
  userid: string;
  status: boolean;
}

export interface LoginData {
  username: string;
  password: string;
}

export interface LoginRes {
  token: string;
}

export interface UserFormRequest {
  username: string;
  nickname: string;
  password: string;
  email: string;
  phone: string;
  userid: string;
  status: boolean;
}

export interface UserParams extends Partial<UserRecord> {
  nickName: string;
  page: number;
  pageSize: number;
}

export interface UserListRes {
  list: UserRecord[];
  total: number;
}

export function login(data: LoginData) {
  return axios.post<LoginRes>('/api/v1/user/login', data);
}

export function logout() {
  return axios.post<LoginRes>('/api/v1/user/logout');
}

export function getUserInfo() {
  return axios.post<UserState>('/api/user/info');
}

export function queryUserInfo(id: number) {
  return axios.get<UserRecord>(`/api/v1/user/${id}`);
}

export function getMenuList() {
  return axios.get('/api/v1/user/menu');
}

export function getMenuNameList() {
  return axios.get('/api/v1/menu/name-list');
}

export function submitUserForm(data: UserFormRequest) {
  return axios.post('/api/v1/user', data);
}

export function submitEditUserForm(id: number, data: UserFormRequest) {
  return axios.put(`/api/v1/user/${id}`, data);
}

export function deleteUser(id: number) {
  return axios.delete(`/api/v1/user/${id}`);
}

export function changeUserStatus(id: number) {
  return axios.put(`/api/v1/mon/user/status/${id}`);
}

export function queryAllUser() {
  return axios.get(`/api/v1/user/all`);
}

export function queryUserList(params: UserParams) {
  return axios.get<UserListRes>('/api/v1/user/list', {
    params,
    paramsSerializer: (obj) => {
      return qs.stringify(obj);
    },
  });
}
