import { DEFAULT_LAYOUT } from '../base';
import { AppRouteRecordRaw } from '../types';

const ALERT_RULE: AppRouteRecordRaw = {
  path: '/alert-rule',
  name: 'alert-rule',
  component: DEFAULT_LAYOUT,
  meta: {
    locale: 'menu.alert.rule',
    requiresAuth: true,
    order: 6,
  },
  children: [
    {
      path: 'source',
      name: 'alert-rule-source',
      component: () => import('@/views/rules/source/index.vue'),
      meta: {
        locale: 'menu.alert.source',
        requiresAuth: true,
        roles: ['*'],
      },
    },
    {
      path: 'list',
      name: 'alert-rule-list',
      component: () => import('@/views/rules/list/index.vue'),
      meta: {
        locale: 'menu.alert.rule',
        requiresAuth: true,
        roles: ['*'],
      },
    },
  ],
};

export default ALERT_RULE;
