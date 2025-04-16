import axios from 'axios';
import qs from 'query-string';

export interface AlertChannelRequest {
  channelName: string;
  channelType: string;
  channelSign: string;
  channelGroup: string;
}

export interface AlertChannelRecode {
  id: number;
  channelName: string;
  channelType: string;
  channelSign: string;
  channelGroup: string;
}
export interface AlertTemplateChannel {
  id: number;
  channelName: string;
  finished: boolean;
  content: string;
}

export interface AlertTemplateChannelRequest {
  content: string;
}
export interface AlertTemplateRequest {
  templateName: string;
  templateDesc: string;
  channels: number[];
}
export interface AlertTemplateRecode {
  id: number;
  templateName: string;
  templateDesc: string;
  channels: AlertTemplateChannel[];
  createdAt: string;
  updatedAt: string;
}

export interface WechatBotRecode {
  channelType: string;
  channelName: string;
  id: number;
}

export interface AlertTemplateParams extends Partial<AlertTemplateRecode> {
  page: number;
  pageSize: number;
}

export interface AlertTemplateListRes {
  list: AlertTemplateRecode[];
  total: number;
}

export interface ChannelGroup {
  data: Map<string, WechatBotRecode[]>;
}

export interface AlertChannelListRes {
  list: AlertChannelRecode[];
  total: number;
}

export function queryAlertTemplateList(params: AlertTemplateParams) {
  return axios.get<AlertTemplateListRes>('/api/v1/alert/template/list', {
    params,
    paramsSerializer: (obj) => {
      return qs.stringify(obj);
    },
  });
}

export function queryAlertChannelList(params: AlertTemplateParams) {
  return axios.get<AlertChannelListRes>('/api/v1/alert/channel/list', {
    params,
    paramsSerializer: (obj) => {
      return qs.stringify(obj);
    },
  });
}

export function submitTemplateForm(data: AlertTemplateRequest) {
  return axios.post('/api/v1/alert/template', data);
}

export function submitChannelForm(data: AlertChannelRequest) {
  return axios.post('/api/v1/alert/channel', data);
}

export function editTemplateForm(id: number, data: AlertTemplateRequest) {
  return axios.put(`/api/v1/alert/template/${id}`, data);
}

export function editChannelForm(id: number, data: AlertChannelRequest) {
  return axios.put(`/api/v1/alert/channel/${id}`, data);
}

export function editTemplateChannelForm(
  id: number,
  data: AlertTemplateChannelRequest
) {
  return axios.put(`/api/v1/alert/template/channel/${id}`, data);
}

export function deleteTemplate(id: number) {
  return axios.delete(`/api/v1/alert/template/${id}`);
}

export function deleteChannel(id: number) {
  return axios.delete(`/api/v1/alert/channel/${id}`);
}

export function queryAlertTemplateDetail(id: number) {
  return axios.get<AlertTemplateRecode>(`/api/v1/alert/template/detail/${id}`);
}

export function queryChannelsByTemplateId(id: number) {
  return axios.get<number[]>(`/api/v1/alert/template/channels/${id}`);
}

export function queryChannelGroup() {
  return axios.get(`/api/v1/alert/channel/group`);
}

export function queryTemplateEnableChannel(id: number) {
  return axios.get(`/api/v1/alert/template/channel-group/${id}`);
}
