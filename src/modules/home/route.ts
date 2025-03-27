import { defineRoute } from '@/router/define';

export const homeRoute = defineRoute({
  name: 'home',
  path: '/',
  component: () => import('./Route.vue'),
});
