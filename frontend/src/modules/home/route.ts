import { defineRoute, type InferViewRouteInfo } from '@/router/define';

export const homeRoute = defineRoute({
  name: 'home',
  path: '/',
  component: () => import('./Route.vue'),
});

export type HomeRouteInfo = InferViewRouteInfo<typeof homeRoute>;
