import axios from 'axios';
import qs from 'query-string';

export interface AlertEventRecode {
  id: number;
  alertTitle: string;
  sourceId: number;
  cluster: string;
  level: number;
  description: string;
  firstTriggerTime: string;
  triggerTime: string;
  triggerValue: string;
  recoverTime: string;
  annotations: string;
  ruleConfig: string;
  isRecovered: boolean;
  notifyCurNumber: number;
  progress: number;
  tags: Map<string, string>;
  levelName: string;
  icon: string;
  nickname: string;
  sourceName: string;
  color: string;
  sourceIp: string;
}
export interface WorkspaceAlertEventParams extends Partial<AlertEventRecode> {
  alertTitle?: string;
  progress?: number;
  alertLevel?: string;
  timeRange?: string;
  page: number;
  pageSize: number;
}

export interface WorkspaceAlertEventListRes {
  list: AlertEventRecode[];
  total: number;
}

export function queryWorkspaceAlertEventList(
  id: number,
  params: WorkspaceAlertEventParams
) {
  return axios.get<WorkspaceAlertEventListRes>(
    `/api/v1/alert/event/workspace-list/${id}`,
    {
      params,
      paramsSerializer: (obj) => {
        return qs.stringify(obj);
      },
    }
  );
}

export function queryAllAlertEventList(params: WorkspaceAlertEventParams) {
  return axios.get<WorkspaceAlertEventListRes>(`/api/v1/alert/event/list`, {
    params,
    paramsSerializer: (obj) => {
      return qs.stringify(obj);
    },
  });
}

export function claimAlertEvent(data: any) {
  return axios.post('/api/v1/alert/event/claim', { events: data });
}

export function closeAlertEvent(data: any) {
  return axios.post('/api/v1/alert/event/closed', data);
}
