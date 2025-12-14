import { defineViewRoute, type InferViewRouteInfo } from '@/router/define';

export const homeListRoute = defineViewRoute({
  name: 'home',
  path: '',
  component: () => import('./Route.vue'),
});

export type HomeListRouteInfo = InferViewRouteInfo<typeof homeListRoute>;
