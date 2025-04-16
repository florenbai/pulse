import { DEFAULT_LAYOUT } from '../base';
import { AppRouteRecordRaw } from '../types';

const SCHEDULE: AppRouteRecordRaw = {
  path: '/schedule',
  name: 'schedule',
  component: DEFAULT_LAYOUT,
  meta: {
    locale: 'menu.schedule',
    requiresAuth: true,
    icon: 'icon-calendar-clock',
    order: 7,
  },
  children: [
    {
      path: 'index',
      name: 'schedule/index',
      component: () => import('@/views/schedule/index/index.vue'),
      meta: {
        locale: 'menu.schedule',
        requiresAuth: true,
        roles: ['*'],
      },
    },
    {
      path: 'detail/:id',
      name: 'schedule/detail',
      component: () => import('@/views/schedule/index/detail.vue'),
      meta: {
        locale: 'menu.schedule.detail',
        requiresAuth: true,
        roles: ['*'],
        hideInMenu: true,
      },
    },
  ],
};

export default SCHEDULE;
