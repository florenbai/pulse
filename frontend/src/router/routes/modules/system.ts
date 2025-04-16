import { DEFAULT_LAYOUT } from '../base';
import { AppRouteRecordRaw } from '../types';

const SYSTEM: AppRouteRecordRaw = {
  path: '/system',
  name: 'system',
  component: DEFAULT_LAYOUT,
  meta: {
    locale: 'menu.system',
    requiresAuth: true,
    icon: 'icon-settings',
    order: 7,
  },
  children: [
    {
      path: 'user',
      name: 'system/user',
      component: () => import('@/views/system/user/index.vue'),
      meta: {
        locale: 'menu.system.user',
        requiresAuth: true,
        roles: ['*'],
      },
    },
    {
      path: 'team',
      name: 'system/team',
      component: () => import('@/views/system/team/index.vue'),
      meta: {
        locale: 'menu.system.team',
        requiresAuth: true,
        roles: ['*'],
      },
    },
    {
      path: 'strategy',
      name: 'system/strategy',
      component: () => import('@/views/system/strategy/index.vue'),
      meta: {
        locale: 'menu.system.strategy',
        requiresAuth: true,
        roles: ['*'],
      },
    },
    {
      path: 'level',
      name: 'system/level',
      component: () => import('@/views/system/level/index.vue'),
      meta: {
        locale: 'menu.system.level',
        requiresAuth: true,
        roles: ['*'],
      },
    },
    {
      path: 'menu',
      name: 'system/menu',
      component: () => import('@/views/system/menu/index.vue'),
      meta: {
        locale: 'menu.system.menu',
        requiresAuth: true,
        roles: ['*'],
      },
    },
  ],
};

export default SYSTEM;
