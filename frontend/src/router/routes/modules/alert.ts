import { DEFAULT_LAYOUT } from '../base';
import { AppRouteRecordRaw } from '../types';

const ALERT: AppRouteRecordRaw = {
  path: '/alert',
  name: 'alert',
  component: DEFAULT_LAYOUT,
  meta: {
    locale: 'menu.alert',
    requiresAuth: true,
    icon: 'icon-notification',
    order: 6,
  },
  children: [
    {
      path: 'workspace',
      name: 'alert/workspace',
      component: () => import('@/views/alert/workspace/index.vue'),
      meta: {
        locale: 'menu.alert.workspace',
        requiresAuth: true,
        roles: ['*'],
      },
    },
    {
      path: 'integration',
      name: 'alert/integration',
      component: () => import('@/views/integration/index.vue'),
      meta: {
        locale: 'menu.source.integration',
        requiresAuth: true,
        roles: ['*'],
      },
    },
    {
      path: 'integration/detail/:id',
      name: 'alert/integration/detail',
      component: () => import('@/views/integration/detail.vue'),
      meta: {
        locale: 'menu.source.integration',
        requiresAuth: true,
        hideInMenu: true,
        roles: ['*'],
      },
    },
    {
      path: 'workspace/detail/:id',
      name: 'alert/workspace/detail',
      component: () => import('@/views/alert/workspace/detail.vue'),
      meta: {
        locale: 'menu.alert.workspace',
        requiresAuth: true,
        roles: ['*'],
        hideInMenu: true,
        activeMenu: 'alert/workspace',
      },
    },
  ],
};

export default ALERT;
