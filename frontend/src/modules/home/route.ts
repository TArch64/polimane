import { defineViewRoute, type InferViewRouteInfo } from '@/router/define';

export const homeRoute = defineViewRoute({
  name: 'home',
  path: '/',
  component: () => import('./Route.vue'),
});

export type HomeRouteInfo = InferViewRouteInfo<typeof homeRoute>;
