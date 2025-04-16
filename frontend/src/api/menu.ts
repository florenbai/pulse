import axios from 'axios';

export interface MenuRecode {
  id: number;
  name: string;
  pid: number;
  type: number;
  icon: string;
  component: string;
  routePath: string;
  routeName: string;
  redirect: string;
  sort: number;
  requiresAuth: boolean;
  permission: string;
  locale: string;
  hideMenu: boolean;
  hideChildMenu: boolean;
  isLink: boolean;
  cache: boolean;
  affix: boolean;
  enabled: boolean;
  activeMenu: boolean;
}

export interface MenuForm {
  name: string;
  pid: number;
  type: number;
  icon: string;
  component: string;
  routePath: string;
  routeName: string;
  redirect: string;
  sort: number;
  requiresAuth: boolean;
  permission: string;
  locale: string;
  hideMenu: boolean;
  hideChildMenu: boolean;
  isLink: boolean;
  cache: boolean;
  affix: boolean;
  enabled: boolean;
  activeMenu: boolean;
}

export interface MenuPermission {
  menus: number[];
  workspace: number;
}

export function getMenuList() {
  return axios.get('/api/v1/menu/list');
}

export function submitMenuForm(data: any) {
  return axios.post(`/api/v1/menu`, data);
}

export function editMenuForm(id: number, data: any) {
  return axios.put(`/api/v1/menu/${id}`, data);
}

export function deleteMenu(id: number) {
  return axios.delete(`/api/v1/menu/${id}`);
}

export function changeMenuSort(data: number[]) {
  return axios.put(`/api/v1/menu/sort`, data);
}

export function changeMenuStatus(id: number) {
  return axios.put(`/api/v1/menu/status/${id}`);
}
