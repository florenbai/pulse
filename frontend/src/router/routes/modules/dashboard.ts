import { DEFAULT_LAYOUT } from '../base';
import { AppRouteRecordRaw } from '../types';

const DASHBOARD: AppRouteRecordRaw = {
  path: '/dashboard',
  name: 'dashboard',
  component: DEFAULT_LAYOUT,
  meta: {
    locale: 'menu.dashboard',
    requiresAuth: true,
    icon: 'icon-dashboard',
    order: 0,
  },
  children: [
    {
      path: 'workplace',
      name: 'Workplace',
      component: () => import('@/views/alert/event/index.vue'),
      meta: {
        locale: 'menu.alert.dashboard',
        requiresAuth: true,
        roles: ['*'],
      },
    },
    {
      path: 'tv',
      name: 'TV',
      component: () => import('@/views/dashboard/tv/index.vue'),
      meta: {
        locale: 'menu.dashboard.tv',
        requiresAuth: true,
        roles: ['*'],
      },
    },
    {
      path: 'data-analysis',
      name: 'DataAnalysis',
      component: () => import('@/views/dashboard/data-analysis/index.vue'),
      meta: {
        locale: 'menu.dataAnalysis',
        requiresAuth: true,
        roles: ['*'],
      },
    },
  ],
};

export default DASHBOARD;
