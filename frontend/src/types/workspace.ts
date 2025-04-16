import type { LabelMap } from '@/types/global';

export const PROGRESS_MAP: LabelMap<number> = {
  1: '未认领',
  2: '已认领',
  3: '已关闭',
};

export const PROGRESS_COLOR_MAP: LabelMap<number> = {
  1: '#f53f3f',
  2: '#f7ba1e',
  3: '#00b42a',
};

export const ALERT_ACTION_MAP: LabelMap<string> = {
  claim: '告警认领',
  closed: '关闭告警',
  opened: '重新打开',
  cancel_claim: '取消认领',
};

export const ALERT_ACTION_COLOR_MAP: LabelMap<string> = {
  claim: '#165dff',
  closed: '#00b42a',
  opened: '#165dff',
  cancel_claim: '#ff5722',
};

export const REFRESH_TIME_MAP: LabelMap<number> = {
  0: '关闭',
  10: '10秒',
  20: '20秒',
  30: '30秒',
  60: '1分钟',
  120: '2分钟',
  300: '5分钟',
  600: '10分钟',
};

export const TIME_RANGE_MAP: LabelMap<string> = {
  '30m': '最近30分钟',
  '1h': '最近1小时',
  '2h': '最近2小时',
  '8h': '最近8小时',
  '1d': '最近1天',
  '2d': '最近2天',
  '7d': '最近7天',
  '14d': '最近14天',
  '30d': '最近30天',
  '3month': '最近3个月',
  '6month': '最近半年',
};

export const STATISTIC_TIME_RANGE_MAP: LabelMap<string> = {
  '1d': '最近1天',
  '7d': '最近7天',
  '14d': '最近14天',
  '30d': '最近30天',
};
