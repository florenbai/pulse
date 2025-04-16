import { DEFAULT_LAYOUT } from '../base';
import { AppRouteRecordRaw } from '../types';

const TEMPLATE: AppRouteRecordRaw = {
  path: '/template',
  name: 'notify',
  component: DEFAULT_LAYOUT,
  meta: {
    locale: 'menu.notify',
    requiresAuth: true,
    icon: 'icon-bookmark',
    order: 6,
  },
  children: [
    {
      path: 'channels',
      name: 'channel/list',
      component: () => import('@/views/channel/index.vue'),
      meta: {
        locale: 'menu.channel',
        requiresAuth: true,
        roles: ['*'],
      },
    },
    {
      path: 'list',
      name: 'template/list',
      component: () => import('@/views/template/index.vue'),
      meta: {
        locale: 'menu.template.list',
        requiresAuth: true,
        roles: ['*'],
      },
    },
    {
      path: 'detail/:id',
      name: 'template/detail',
      component: () => import('@/views/template/detail.vue'),
      meta: {
        requiresAuth: true,
        roles: ['*'],
        hideInMenu: true,
      },
    },
  ],
};

export default TEMPLATE;
