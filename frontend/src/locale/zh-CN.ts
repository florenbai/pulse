import localeMessageBox from '@/components/message-box/locale/zh-CN';
import localeLogin from '@/views/login/locale/zh-CN';

import localeWorkplace from '@/views/dashboard/workplace/locale/zh-CN';
import localeTemplate from '@/views/template/locale/zh-CN';

import localeSettings from './zh-CN/settings';

export default {
  'menu.dashboard': '仪表盘',
  'menu.server.dashboard': '仪表盘-服务端',
  'menu.server.workplace': '工作台-服务端',
  'menu.server.monitor': '实时监控-服务端',
  'menu.list': '列表页',
  'menu.result': '结果页',
  'menu.exception': '异常页',
  'menu.form': '表单页',
  'menu.profile': '详情页',
  'menu.visualization': '数据可视化',
  'menu.user': '个人中心',
  'menu.arcoWebsite': 'Arco Design',
  'menu.faq': '常见问题',
  'menu.alert': '告警中心',
  'menu.dashboard.tv': '电视大盘',
  'menu.alert.workspace': '工作区',
  'menu.alert.event': '告警列表',
  'menu.alert.dashboard': '告警总览',
  'menu.alert.workspace.detail': '工作区详情',
  'menu.template': '告警模板',
  'menu.template.list': '模板管理',
  'menu.notify': '告警通知管理',
  'menu.channel': '通知渠道管理',
  'menu.source': '告警源',
  'menu.source.integration': '告警集成',
  'menu.system': '系统设置',
  'menu.system.team': '团队管理',
  'menu.system.user': '用户管理',
  'menu.system.level': '告警等级管理',
  'menu.system.menu': '菜单管理',
  'menu.system.strategy': '告警策略管理',
  'menu.schedule': '值班管理',
  'menu.schedule.detail': '值班详情',
  'menu.schedule.calendar': '值班日历',
  'menu.dataAnalysis': '数据分析',
  'menu.alert.rule': '告警规则管理',
  'menu.alert.rule.dataSource': '数据源',
  'navbar.docs': '文档中心',
  'navbar.action.locale': '切换为中文',

  ...localeSettings,
  ...localeMessageBox,
  ...localeLogin,
  ...localeWorkplace,
  ...localeTemplate,
};
